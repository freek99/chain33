package download

import (
	"context"
	"sort"
	"sync"
	"time"

	"github.com/33cn/chain33/common/log/log15"

	core "github.com/libp2p/go-libp2p-core"

	prototypes "github.com/33cn/chain33/p2pnext/protocol/types"
	uuid "github.com/google/uuid"

	"github.com/33cn/chain33/queue"
	"github.com/33cn/chain33/types"
	net "github.com/libp2p/go-libp2p-core/network"
)

var (
	log = log15.New("module", "p2p.peer")
)

func init() {
	prototypes.RegisterProtocolType(protoTypeID, &DownloadProtol{})
	var downloadHandler = new(DownloadHander)
	prototypes.RegisterStreamHandlerType(protoTypeID, DownloadBlockReq, downloadHandler)
	//prototypes.RegisterStreamHandlerType(protoTypeID, DownloadBlockResp, downloadHandler)
}

const (
	protoTypeID      = "DownloadProtocolType"
	DownloadBlockReq = "/chain33/downloadBlockReq/1.0.0"
	//DownloadBlockResp = "/chain33/downloadBlockResp/1.0.0"
)

//type Istream
type DownloadProtol struct {
	*prototypes.BaseProtocol
	*prototypes.BaseStreamHandler

	requests map[string]*types.MessageGetBlocksReq // used to access request data from response handlers
}

func (d *DownloadProtol) InitProtocol(data *prototypes.GlobalData) {
	d.BaseProtocol = new(prototypes.BaseProtocol)
	d.requests = make(map[string]*types.MessageGetBlocksReq)
	d.GlobalData = data

	//注册事件处理函数
	prototypes.RegisterEventHandler(types.EventFetchBlocks, d.handleEvent)
}

type DownloadHander struct {
	*prototypes.BaseStreamHandler
}

func (d *DownloadHander) SetProtocol(protocol prototypes.IProtocol) {
	d.BaseStreamHandler = new(prototypes.BaseStreamHandler)
	d.Protocol = protocol
}

//接收Response消息
func (d *DownloadProtol) OnResp(datas *types.InvDatas, s core.Stream) {
	client := d.QueueClient

	for _, invdata := range datas.GetItems() {
		if invdata.Ty != 2 {
			continue
		}

		block := invdata.GetBlock()
		newmsg := client.NewMessage("blockchain", types.EventSyncBlock,
			&types.BlockPid{Pid: s.Conn().LocalPeer().String(), Block: block})
		client.SendTimeout(newmsg, false, 20*time.Second)

	}

}
func (h *DownloadHander) VerifyRequest(data []byte) bool {

	return true
}

//Handle 处理请求
func (d *DownloadHander) Handle(stream core.Stream) {
	protocol := d.GetProtocol().(*DownloadProtol)

	//解析处理
	if stream.Protocol() == DownloadBlockReq {
		var data types.MessageGetBlocksReq
		err := d.BaseStreamHandler.ReadProtoMessage(&data, stream)
		if err != nil {
			log.Error("Handle", "err", err)
			return
		}
		recvData := data.Message
		protocol.OnReq(data.MessageData.Id, recvData, stream)
		return
	}

	return

}

func (d *DownloadProtol) OnReq(id string, message *types.P2PGetBlocks, s core.Stream) {

	//获取headers 信息
	//允许下载的最大高度区间为256
	if message.GetEndHeight()-message.GetStartHeight() > 256 || message.GetEndHeight() < message.GetStartHeight() {
		return

	}
	//开始下载指定高度
	reqblock := &types.ReqBlocks{Start: message.GetStartHeight(), End: message.GetEndHeight()}
	client := d.GetQueueClient()
	msg := client.NewMessage("blockchain", types.EventGetBlocks, reqblock)
	err := client.Send(msg, true)
	if err != nil {
		log.Error("GetBlocks", "Error", err.Error())
		return
	}
	resp, err := client.WaitTimeout(msg, time.Second*20)
	if err != nil {
		log.Error("GetBlocks Err", "Err", err.Error())
		return
	}

	blockDetails := resp.Data.(*types.BlockDetails)
	var p2pInvData = make([]*types.InvData, 0)
	var invdata types.InvData
	for _, item := range blockDetails.Items {
		invdata.Ty = 2 //2 block,1 tx
		invdata.Value = &types.InvData_Block{Block: item.Block}
		p2pInvData = append(p2pInvData, &invdata)
	}
	peerID := d.GetHost().ID()
	pubkey, _ := d.GetHost().Peerstore().PubKey(peerID).Bytes()
	blocksResp := &types.MessageGetBlocksResp{MessageData: d.NewMessageCommon(id, peerID.Pretty(), pubkey, false),
		Message: &types.InvDatas{p2pInvData}}

	err = d.BaseStreamHandler.SendProtoMessage(blocksResp, s)
	//wlen, err := rw.WriteString(fmt.Sprintf("%v\n", string(types.Encode(blocksResp))))
	if err != nil {
		log.Error("SendProtoMessage", "err", err)
		d.GetConnsManager().Delete(s.Conn().RemotePeer().Pretty())
		return
	}
	log.Info("%s: Ping response to %s sent.", s.Conn().LocalPeer().String(), s.Conn().RemotePeer().String())

}

//GetBlocks 接收来自chain33 blockchain模块发来的请求
func (d *DownloadProtol) handleEvent(msg *queue.Message) {

	req := msg.GetData().(*types.ReqBlocks)
	pids := req.GetPid()
	if len(pids) == 0 { //根据指定的pidlist 获取对应的block header
		log.Debug("GetBlocks:pid is nil")
		msg.Reply(d.GetQueueClient().NewMessage("blockchain", types.EventReply, types.Reply{Msg: []byte("no pid")}))
		return
	}

	msg.Reply(d.GetQueueClient().NewMessage("blockchain", types.EventReply, types.Reply{IsOk: true, Msg: []byte("ok")}))

	for _, pid := range pids {

		data := d.PeerInfoManager.Load(pid)
		if data == nil {
			continue
		}
		peerinfo := data.(*types.P2PPeerInfo)
		//去指定的peer上获取对应的blockHeader
		peerId := peerinfo.GetName()
		Pconn := d.ConnManager.Get(peerId)
		if Pconn == nil {
			continue
		}
		//具体的下载逻辑
		jobS := d.initStreamJob()

		for height := req.GetStart(); height <= req.GetEnd(); height++ {
			jStream := d.getFreeJobStream(jobS)
			getblocks := &types.P2PGetBlocks{StartHeight: req.GetStart(), EndHeight: req.GetEnd(),
				Version: 0}

			peerID := d.GetHost().ID()
			pubkey, _ := d.GetHost().Peerstore().PubKey(peerID).Bytes()
			blockReq := &types.MessageGetBlocksReq{MessageData: d.NewMessageCommon(uuid.New().String(), peerID.Pretty(), pubkey, false),
				Message: getblocks}

			if err := d.SendProtoMessage(blockReq, jStream.Stream); err == nil {
				d.requests[blockReq.MessageData.Id] = blockReq
			}

		}
	}

}

type JobStream struct {
	Limit  int
	Stream net.Stream
	mtx    sync.Mutex
}

// jobs datastruct
type jobs []*JobStream

//Len size of the Invs data
func (i jobs) Len() int {
	return len(i)
}

//Less Sort from low to high
func (i jobs) Less(a, b int) bool {
	return i[a].Limit < i[b].Limit
}

//Swap  the param
func (i jobs) Swap(a, b int) {
	i[a], i[b] = i[b], i[a]
}

func (d *DownloadProtol) initStreamJob() jobs {
	var jobStreams jobs
	conns := d.ConnManager.Fetch()
	for _, conn := range conns {
		var jstream JobStream

		stream, err := d.Host.NewStream(context.Background(), conn.RemotePeer(), DownloadBlockReq)
		if err != nil {
			log.Error("NewStream", "err", err)
			continue
		}
		jstream.Stream = stream
		jstream.Limit = 0
		jobStreams = append(jobStreams, &jstream)

	}

	return jobStreams
}

func (d *DownloadProtol) getFreeJobStream(js jobs) *JobStream {
	var MaxJobLimit int = 10
	sort.Sort(js)
	for _, job := range js {
		if job.Limit < MaxJobLimit {
			job.mtx.Lock()
			job.Limit++
			job.mtx.Unlock()
			return job
		}
	}

	return nil

}

func (d *DownloadProtol) releaseJobStream(js *JobStream) {
	js.mtx.Lock()
	defer js.mtx.Unlock()
	js.Limit--
	if js.Limit < 0 {
		js.Limit = 0
	}
}