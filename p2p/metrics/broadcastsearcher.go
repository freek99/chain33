package metrics

import (
	"context"
	"errors"
	"fmt"
	"math/rand"

	pb "github.com/33cn/chain33/types"
	"google.golang.org/grpc"
)

const MaxCallRecvMsgSize = 100 * 1024 * 1024

type BroadcastSearcher struct {
}

func (bs *BroadcastSearcher) find(dstAddr string, hash string) (*pb.PeersBroadInfoReply, error) {
	conn, err := grpc.Dial(
		dstAddr,
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(MaxCallRecvMsgSize)),
	)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Dial error")
	}

	gcli := pb.NewP2PgserviceClient(conn)
	resp, err := gcli.GetBroadcastData(
		context.Background(),
		&pb.P2PPing{Nonce: int64(rand.Int31n(102040)), Addr: hash, Port: 13802},
		grpc.FailFast(true))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}

func (bs *BroadcastSearcher) Search(startAddr string, hash string) []*pb.PeersBroadInfoReply {
	usedPeers := make(map[string]string)
	var replys []*pb.PeersBroadInfoReply

	dstAddr := startAddr
	for {
		reply, err := bs.find(dstAddr, hash)
		if err != nil {
			break
		}
		replys = append(replys, reply)

		for _, info := range reply.Infos {
			if usedPeers[info.SrcIPPort] == "" {
				usedPeers[info.SrcIPPort] = ""
			}
		}

		usedPeers[dstAddr] = dstAddr

		newDstAddr := ""
		for peer, usedPeer := range usedPeers {
			if usedPeer == "" {
				newDstAddr = peer
				break
			} else {
			}
		}

		if newDstAddr != "" {
			dstAddr = newDstAddr
		} else {
			break
		}
	}

	return replys
}
