// Code generated by protoc-gen-go. DO NOT EDIT.
// source: p2pnext.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MessageComm struct {
	// shared between all requests
	Version    string `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	Timestamp  int64  `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
	Id         string `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
	Gossip     bool   `protobuf:"varint,4,opt,name=gossip" json:"gossip,omitempty"`
	NodeId     string `protobuf:"bytes,5,opt,name=nodeId" json:"nodeId,omitempty"`
	NodePubKey []byte `protobuf:"bytes,6,opt,name=nodePubKey,proto3" json:"nodePubKey,omitempty"`
	Sign       []byte `protobuf:"bytes,7,opt,name=sign,proto3" json:"sign,omitempty"`
}

func (m *MessageComm) Reset()                    { *m = MessageComm{} }
func (m *MessageComm) String() string            { return proto.CompactTextString(m) }
func (*MessageComm) ProtoMessage()               {}
func (*MessageComm) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *MessageComm) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *MessageComm) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *MessageComm) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MessageComm) GetGossip() bool {
	if m != nil {
		return m.Gossip
	}
	return false
}

func (m *MessageComm) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *MessageComm) GetNodePubKey() []byte {
	if m != nil {
		return m.NodePubKey
	}
	return nil
}

func (m *MessageComm) GetSign() []byte {
	if m != nil {
		return m.Sign
	}
	return nil
}

type MessageUtil struct {
	Common *MessageComm `protobuf:"bytes,1,opt,name=common" json:"common,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*MessageUtil_PeerInfo
	//	*MessageUtil_Version
	//	*MessageUtil_VersionAck
	//	*MessageUtil_External
	//	*MessageUtil_Getblocks
	//	*MessageUtil_Invdatas
	Value isMessageUtil_Value `protobuf_oneof:"value"`
}

func (m *MessageUtil) Reset()                    { *m = MessageUtil{} }
func (m *MessageUtil) String() string            { return proto.CompactTextString(m) }
func (*MessageUtil) ProtoMessage()               {}
func (*MessageUtil) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

type isMessageUtil_Value interface {
	isMessageUtil_Value()
}

type MessageUtil_PeerInfo struct {
	PeerInfo *P2PPeerInfo `protobuf:"bytes,2,opt,name=peerInfo,oneof"`
}
type MessageUtil_Version struct {
	Version *P2PVersion `protobuf:"bytes,3,opt,name=version,oneof"`
}
type MessageUtil_VersionAck struct {
	VersionAck *P2PVerAck `protobuf:"bytes,4,opt,name=versionAck,oneof"`
}
type MessageUtil_External struct {
	External *P2PExternalInfo `protobuf:"bytes,5,opt,name=external,oneof"`
}
type MessageUtil_Getblocks struct {
	Getblocks *P2PGetBlocks `protobuf:"bytes,6,opt,name=getblocks,oneof"`
}
type MessageUtil_Invdatas struct {
	Invdatas *InvDatas `protobuf:"bytes,7,opt,name=invdatas,oneof"`
}

func (*MessageUtil_PeerInfo) isMessageUtil_Value()   {}
func (*MessageUtil_Version) isMessageUtil_Value()    {}
func (*MessageUtil_VersionAck) isMessageUtil_Value() {}
func (*MessageUtil_External) isMessageUtil_Value()   {}
func (*MessageUtil_Getblocks) isMessageUtil_Value()  {}
func (*MessageUtil_Invdatas) isMessageUtil_Value()   {}

func (m *MessageUtil) GetValue() isMessageUtil_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *MessageUtil) GetCommon() *MessageComm {
	if m != nil {
		return m.Common
	}
	return nil
}

func (m *MessageUtil) GetPeerInfo() *P2PPeerInfo {
	if x, ok := m.GetValue().(*MessageUtil_PeerInfo); ok {
		return x.PeerInfo
	}
	return nil
}

func (m *MessageUtil) GetVersion() *P2PVersion {
	if x, ok := m.GetValue().(*MessageUtil_Version); ok {
		return x.Version
	}
	return nil
}

func (m *MessageUtil) GetVersionAck() *P2PVerAck {
	if x, ok := m.GetValue().(*MessageUtil_VersionAck); ok {
		return x.VersionAck
	}
	return nil
}

func (m *MessageUtil) GetExternal() *P2PExternalInfo {
	if x, ok := m.GetValue().(*MessageUtil_External); ok {
		return x.External
	}
	return nil
}

func (m *MessageUtil) GetGetblocks() *P2PGetBlocks {
	if x, ok := m.GetValue().(*MessageUtil_Getblocks); ok {
		return x.Getblocks
	}
	return nil
}

func (m *MessageUtil) GetInvdatas() *InvDatas {
	if x, ok := m.GetValue().(*MessageUtil_Invdatas); ok {
		return x.Invdatas
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*MessageUtil) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _MessageUtil_OneofMarshaler, _MessageUtil_OneofUnmarshaler, _MessageUtil_OneofSizer, []interface{}{
		(*MessageUtil_PeerInfo)(nil),
		(*MessageUtil_Version)(nil),
		(*MessageUtil_VersionAck)(nil),
		(*MessageUtil_External)(nil),
		(*MessageUtil_Getblocks)(nil),
		(*MessageUtil_Invdatas)(nil),
	}
}

func _MessageUtil_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*MessageUtil)
	// value
	switch x := m.Value.(type) {
	case *MessageUtil_PeerInfo:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PeerInfo); err != nil {
			return err
		}
	case *MessageUtil_Version:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Version); err != nil {
			return err
		}
	case *MessageUtil_VersionAck:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.VersionAck); err != nil {
			return err
		}
	case *MessageUtil_External:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.External); err != nil {
			return err
		}
	case *MessageUtil_Getblocks:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Getblocks); err != nil {
			return err
		}
	case *MessageUtil_Invdatas:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Invdatas); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("MessageUtil.Value has unexpected type %T", x)
	}
	return nil
}

func _MessageUtil_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*MessageUtil)
	switch tag {
	case 2: // value.peerInfo
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(P2PPeerInfo)
		err := b.DecodeMessage(msg)
		m.Value = &MessageUtil_PeerInfo{msg}
		return true, err
	case 3: // value.version
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(P2PVersion)
		err := b.DecodeMessage(msg)
		m.Value = &MessageUtil_Version{msg}
		return true, err
	case 4: // value.versionAck
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(P2PVerAck)
		err := b.DecodeMessage(msg)
		m.Value = &MessageUtil_VersionAck{msg}
		return true, err
	case 5: // value.external
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(P2PExternalInfo)
		err := b.DecodeMessage(msg)
		m.Value = &MessageUtil_External{msg}
		return true, err
	case 6: // value.getblocks
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(P2PGetBlocks)
		err := b.DecodeMessage(msg)
		m.Value = &MessageUtil_Getblocks{msg}
		return true, err
	case 7: // value.invdatas
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(InvDatas)
		err := b.DecodeMessage(msg)
		m.Value = &MessageUtil_Invdatas{msg}
		return true, err
	default:
		return false, nil
	}
}

func _MessageUtil_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*MessageUtil)
	// value
	switch x := m.Value.(type) {
	case *MessageUtil_PeerInfo:
		s := proto.Size(x.PeerInfo)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *MessageUtil_Version:
		s := proto.Size(x.Version)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *MessageUtil_VersionAck:
		s := proto.Size(x.VersionAck)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *MessageUtil_External:
		s := proto.Size(x.External)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *MessageUtil_Getblocks:
		s := proto.Size(x.Getblocks)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *MessageUtil_Invdatas:
		s := proto.Size(x.Invdatas)
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// *
// 请求获取远程节点的节点信息
type MessagePeerInfoReq struct {
	// / p2p版本
	MessageData *MessageComm `protobuf:"bytes,1,opt,name=messageData" json:"messageData,omitempty"`
}

func (m *MessagePeerInfoReq) Reset()                    { *m = MessagePeerInfoReq{} }
func (m *MessagePeerInfoReq) String() string            { return proto.CompactTextString(m) }
func (*MessagePeerInfoReq) ProtoMessage()               {}
func (*MessagePeerInfoReq) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *MessagePeerInfoReq) GetMessageData() *MessageComm {
	if m != nil {
		return m.MessageData
	}
	return nil
}

type MessagePeerInfoResp struct {
	MessageData *MessageComm `protobuf:"bytes,1,opt,name=messageData" json:"messageData,omitempty"`
	Message     *P2PPeerInfo `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *MessagePeerInfoResp) Reset()                    { *m = MessagePeerInfoResp{} }
func (m *MessagePeerInfoResp) String() string            { return proto.CompactTextString(m) }
func (*MessagePeerInfoResp) ProtoMessage()               {}
func (*MessagePeerInfoResp) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

func (m *MessagePeerInfoResp) GetMessageData() *MessageComm {
	if m != nil {
		return m.MessageData
	}
	return nil
}

func (m *MessagePeerInfoResp) GetMessage() *P2PPeerInfo {
	if m != nil {
		return m.Message
	}
	return nil
}

type MessageP2PVersionReq struct {
	MessageData *MessageComm `protobuf:"bytes,1,opt,name=messageData" json:"messageData,omitempty"`
	Message     *P2PVersion  `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *MessageP2PVersionReq) Reset()                    { *m = MessageP2PVersionReq{} }
func (m *MessageP2PVersionReq) String() string            { return proto.CompactTextString(m) }
func (*MessageP2PVersionReq) ProtoMessage()               {}
func (*MessageP2PVersionReq) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{4} }

func (m *MessageP2PVersionReq) GetMessageData() *MessageComm {
	if m != nil {
		return m.MessageData
	}
	return nil
}

func (m *MessageP2PVersionReq) GetMessage() *P2PVersion {
	if m != nil {
		return m.Message
	}
	return nil
}

type MessageP2PVersionResp struct {
	MessageData *MessageComm `protobuf:"bytes,1,opt,name=messageData" json:"messageData,omitempty"`
	Message     *P2PVersion  `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *MessageP2PVersionResp) Reset()                    { *m = MessageP2PVersionResp{} }
func (m *MessageP2PVersionResp) String() string            { return proto.CompactTextString(m) }
func (*MessageP2PVersionResp) ProtoMessage()               {}
func (*MessageP2PVersionResp) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{5} }

func (m *MessageP2PVersionResp) GetMessageData() *MessageComm {
	if m != nil {
		return m.MessageData
	}
	return nil
}

func (m *MessageP2PVersionResp) GetMessage() *P2PVersion {
	if m != nil {
		return m.Message
	}
	return nil
}

type MessagePingReq struct {
	MessageData *MessageComm `protobuf:"bytes,1,opt,name=messageData" json:"messageData,omitempty"`
	Message     *P2PPing     `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *MessagePingReq) Reset()                    { *m = MessagePingReq{} }
func (m *MessagePingReq) String() string            { return proto.CompactTextString(m) }
func (*MessagePingReq) ProtoMessage()               {}
func (*MessagePingReq) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{6} }

func (m *MessagePingReq) GetMessageData() *MessageComm {
	if m != nil {
		return m.MessageData
	}
	return nil
}

func (m *MessagePingReq) GetMessage() *P2PPing {
	if m != nil {
		return m.Message
	}
	return nil
}

type MessagePingResp struct {
	MessageData *MessageComm `protobuf:"bytes,1,opt,name=messageData" json:"messageData,omitempty"`
	Message     *P2PPong     `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *MessagePingResp) Reset()                    { *m = MessagePingResp{} }
func (m *MessagePingResp) String() string            { return proto.CompactTextString(m) }
func (*MessagePingResp) ProtoMessage()               {}
func (*MessagePingResp) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{7} }

func (m *MessagePingResp) GetMessageData() *MessageComm {
	if m != nil {
		return m.MessageData
	}
	return nil
}

func (m *MessagePingResp) GetMessage() *P2PPong {
	if m != nil {
		return m.Message
	}
	return nil
}

type MessageAddrReq struct {
	MessageData *MessageComm `protobuf:"bytes,1,opt,name=messageData" json:"messageData,omitempty"`
	Message     *P2PGetAddr  `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *MessageAddrReq) Reset()                    { *m = MessageAddrReq{} }
func (m *MessageAddrReq) String() string            { return proto.CompactTextString(m) }
func (*MessageAddrReq) ProtoMessage()               {}
func (*MessageAddrReq) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{8} }

func (m *MessageAddrReq) GetMessageData() *MessageComm {
	if m != nil {
		return m.MessageData
	}
	return nil
}

func (m *MessageAddrReq) GetMessage() *P2PGetAddr {
	if m != nil {
		return m.Message
	}
	return nil
}

type MessageAddrResp struct {
	MessageData *MessageComm `protobuf:"bytes,1,opt,name=messageData" json:"messageData,omitempty"`
	Message     *P2PAddr     `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *MessageAddrResp) Reset()                    { *m = MessageAddrResp{} }
func (m *MessageAddrResp) String() string            { return proto.CompactTextString(m) }
func (*MessageAddrResp) ProtoMessage()               {}
func (*MessageAddrResp) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{9} }

func (m *MessageAddrResp) GetMessageData() *MessageComm {
	if m != nil {
		return m.MessageData
	}
	return nil
}

func (m *MessageAddrResp) GetMessage() *P2PAddr {
	if m != nil {
		return m.Message
	}
	return nil
}

type MessageAddrList struct {
	MessageData *MessageComm `protobuf:"bytes,1,opt,name=messageData" json:"messageData,omitempty"`
	Message     *P2PAddrList `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *MessageAddrList) Reset()                    { *m = MessageAddrList{} }
func (m *MessageAddrList) String() string            { return proto.CompactTextString(m) }
func (*MessageAddrList) ProtoMessage()               {}
func (*MessageAddrList) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{10} }

func (m *MessageAddrList) GetMessageData() *MessageComm {
	if m != nil {
		return m.MessageData
	}
	return nil
}

func (m *MessageAddrList) GetMessage() *P2PAddrList {
	if m != nil {
		return m.Message
	}
	return nil
}

type MessageExternalNetReq struct {
	MessageData *MessageComm `protobuf:"bytes,1,opt,name=messageData" json:"messageData,omitempty"`
}

func (m *MessageExternalNetReq) Reset()                    { *m = MessageExternalNetReq{} }
func (m *MessageExternalNetReq) String() string            { return proto.CompactTextString(m) }
func (*MessageExternalNetReq) ProtoMessage()               {}
func (*MessageExternalNetReq) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{11} }

func (m *MessageExternalNetReq) GetMessageData() *MessageComm {
	if m != nil {
		return m.MessageData
	}
	return nil
}

type MessageExternalNetResp struct {
	MessageData *MessageComm     `protobuf:"bytes,1,opt,name=messageData" json:"messageData,omitempty"`
	Message     *P2PExternalInfo `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *MessageExternalNetResp) Reset()                    { *m = MessageExternalNetResp{} }
func (m *MessageExternalNetResp) String() string            { return proto.CompactTextString(m) }
func (*MessageExternalNetResp) ProtoMessage()               {}
func (*MessageExternalNetResp) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{12} }

func (m *MessageExternalNetResp) GetMessageData() *MessageComm {
	if m != nil {
		return m.MessageData
	}
	return nil
}

func (m *MessageExternalNetResp) GetMessage() *P2PExternalInfo {
	if m != nil {
		return m.Message
	}
	return nil
}

type MessageGetBlocksReq struct {
	MessageData *MessageComm  `protobuf:"bytes,1,opt,name=messageData" json:"messageData,omitempty"`
	Message     *P2PGetBlocks `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *MessageGetBlocksReq) Reset()                    { *m = MessageGetBlocksReq{} }
func (m *MessageGetBlocksReq) String() string            { return proto.CompactTextString(m) }
func (*MessageGetBlocksReq) ProtoMessage()               {}
func (*MessageGetBlocksReq) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{13} }

func (m *MessageGetBlocksReq) GetMessageData() *MessageComm {
	if m != nil {
		return m.MessageData
	}
	return nil
}

func (m *MessageGetBlocksReq) GetMessage() *P2PGetBlocks {
	if m != nil {
		return m.Message
	}
	return nil
}

type MessageGetBlocksResp struct {
	MessageData *MessageComm `protobuf:"bytes,1,opt,name=messageData" json:"messageData,omitempty"`
	Message     *InvDatas    `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *MessageGetBlocksResp) Reset()                    { *m = MessageGetBlocksResp{} }
func (m *MessageGetBlocksResp) String() string            { return proto.CompactTextString(m) }
func (*MessageGetBlocksResp) ProtoMessage()               {}
func (*MessageGetBlocksResp) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{14} }

func (m *MessageGetBlocksResp) GetMessageData() *MessageComm {
	if m != nil {
		return m.MessageData
	}
	return nil
}

func (m *MessageGetBlocksResp) GetMessage() *InvDatas {
	if m != nil {
		return m.Message
	}
	return nil
}

type MessageGetMempoolReq struct {
	MessageData *MessageComm   `protobuf:"bytes,1,opt,name=messageData" json:"messageData,omitempty"`
	Message     *P2PGetMempool `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *MessageGetMempoolReq) Reset()                    { *m = MessageGetMempoolReq{} }
func (m *MessageGetMempoolReq) String() string            { return proto.CompactTextString(m) }
func (*MessageGetMempoolReq) ProtoMessage()               {}
func (*MessageGetMempoolReq) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{15} }

func (m *MessageGetMempoolReq) GetMessageData() *MessageComm {
	if m != nil {
		return m.MessageData
	}
	return nil
}

func (m *MessageGetMempoolReq) GetMessage() *P2PGetMempool {
	if m != nil {
		return m.Message
	}
	return nil
}

type MessageVersion struct {
	MessageData *MessageComm `protobuf:"bytes,1,opt,name=messageData" json:"messageData,omitempty"`
	Message     *Versions    `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *MessageVersion) Reset()                    { *m = MessageVersion{} }
func (m *MessageVersion) String() string            { return proto.CompactTextString(m) }
func (*MessageVersion) ProtoMessage()               {}
func (*MessageVersion) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{16} }

func (m *MessageVersion) GetMessageData() *MessageComm {
	if m != nil {
		return m.MessageData
	}
	return nil
}

func (m *MessageVersion) GetMessage() *Versions {
	if m != nil {
		return m.Message
	}
	return nil
}

type MessageHeaderReq struct {
	MessageData *MessageComm   `protobuf:"bytes,1,opt,name=messageData" json:"messageData,omitempty"`
	Message     *P2PGetHeaders `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *MessageHeaderReq) Reset()                    { *m = MessageHeaderReq{} }
func (m *MessageHeaderReq) String() string            { return proto.CompactTextString(m) }
func (*MessageHeaderReq) ProtoMessage()               {}
func (*MessageHeaderReq) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{17} }

func (m *MessageHeaderReq) GetMessageData() *MessageComm {
	if m != nil {
		return m.MessageData
	}
	return nil
}

func (m *MessageHeaderReq) GetMessage() *P2PGetHeaders {
	if m != nil {
		return m.Message
	}
	return nil
}

type MessageHeaderResp struct {
	MessageData *MessageComm `protobuf:"bytes,1,opt,name=messageData" json:"messageData,omitempty"`
	Message     *P2PHeaders  `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *MessageHeaderResp) Reset()                    { *m = MessageHeaderResp{} }
func (m *MessageHeaderResp) String() string            { return proto.CompactTextString(m) }
func (*MessageHeaderResp) ProtoMessage()               {}
func (*MessageHeaderResp) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{18} }

func (m *MessageHeaderResp) GetMessageData() *MessageComm {
	if m != nil {
		return m.MessageData
	}
	return nil
}

func (m *MessageHeaderResp) GetMessage() *P2PHeaders {
	if m != nil {
		return m.Message
	}
	return nil
}

type MessageInvDataReq struct {
	MessageData *MessageComm `protobuf:"bytes,1,opt,name=messageData" json:"messageData,omitempty"`
	Message     *InvData     `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *MessageInvDataReq) Reset()                    { *m = MessageInvDataReq{} }
func (m *MessageInvDataReq) String() string            { return proto.CompactTextString(m) }
func (*MessageInvDataReq) ProtoMessage()               {}
func (*MessageInvDataReq) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{19} }

func (m *MessageInvDataReq) GetMessageData() *MessageComm {
	if m != nil {
		return m.MessageData
	}
	return nil
}

func (m *MessageInvDataReq) GetMessage() *InvData {
	if m != nil {
		return m.Message
	}
	return nil
}

type MessagePeerList struct {
	MessageData *MessageComm `protobuf:"bytes,1,opt,name=messageData" json:"messageData,omitempty"`
	Message     *PeerList    `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *MessagePeerList) Reset()                    { *m = MessagePeerList{} }
func (m *MessagePeerList) String() string            { return proto.CompactTextString(m) }
func (*MessagePeerList) ProtoMessage()               {}
func (*MessagePeerList) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{20} }

func (m *MessagePeerList) GetMessageData() *MessageComm {
	if m != nil {
		return m.MessageData
	}
	return nil
}

func (m *MessagePeerList) GetMessage() *PeerList {
	if m != nil {
		return m.Message
	}
	return nil
}

type MessageNetInfo struct {
	MessageData *MessageComm `protobuf:"bytes,1,opt,name=messageData" json:"messageData,omitempty"`
	Message     *NodeNetInfo `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *MessageNetInfo) Reset()                    { *m = MessageNetInfo{} }
func (m *MessageNetInfo) String() string            { return proto.CompactTextString(m) }
func (*MessageNetInfo) ProtoMessage()               {}
func (*MessageNetInfo) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{21} }

func (m *MessageNetInfo) GetMessageData() *MessageComm {
	if m != nil {
		return m.MessageData
	}
	return nil
}

func (m *MessageNetInfo) GetMessage() *NodeNetInfo {
	if m != nil {
		return m.Message
	}
	return nil
}

type MessagePeersReply struct {
	Common     *MessageComm `protobuf:"bytes,1,opt,name=common" json:"common,omitempty"`
	PeersReply *PeersReply  `protobuf:"bytes,2,opt,name=peersReply" json:"peersReply,omitempty"`
}

func (m *MessagePeersReply) Reset()                    { *m = MessagePeersReply{} }
func (m *MessagePeersReply) String() string            { return proto.CompactTextString(m) }
func (*MessagePeersReply) ProtoMessage()               {}
func (*MessagePeersReply) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{22} }

func (m *MessagePeersReply) GetCommon() *MessageComm {
	if m != nil {
		return m.Common
	}
	return nil
}

func (m *MessagePeersReply) GetPeersReply() *PeersReply {
	if m != nil {
		return m.PeersReply
	}
	return nil
}

type MessageBroadCast struct {
	Common  *MessageComm   `protobuf:"bytes,1,opt,name=common" json:"common,omitempty"`
	Message *BroadCastData `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *MessageBroadCast) Reset()                    { *m = MessageBroadCast{} }
func (m *MessageBroadCast) String() string            { return proto.CompactTextString(m) }
func (*MessageBroadCast) ProtoMessage()               {}
func (*MessageBroadCast) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{23} }

func (m *MessageBroadCast) GetCommon() *MessageComm {
	if m != nil {
		return m.Common
	}
	return nil
}

func (m *MessageBroadCast) GetMessage() *BroadCastData {
	if m != nil {
		return m.Message
	}
	return nil
}

func init() {
	proto.RegisterType((*MessageComm)(nil), "types.MessageComm")
	proto.RegisterType((*MessageUtil)(nil), "types.MessageUtil")
	proto.RegisterType((*MessagePeerInfoReq)(nil), "types.MessagePeerInfoReq")
	proto.RegisterType((*MessagePeerInfoResp)(nil), "types.MessagePeerInfoResp")
	proto.RegisterType((*MessageP2PVersionReq)(nil), "types.MessageP2PVersionReq")
	proto.RegisterType((*MessageP2PVersionResp)(nil), "types.MessageP2PVersionResp")
	proto.RegisterType((*MessagePingReq)(nil), "types.MessagePingReq")
	proto.RegisterType((*MessagePingResp)(nil), "types.MessagePingResp")
	proto.RegisterType((*MessageAddrReq)(nil), "types.MessageAddrReq")
	proto.RegisterType((*MessageAddrResp)(nil), "types.MessageAddrResp")
	proto.RegisterType((*MessageAddrList)(nil), "types.MessageAddrList")
	proto.RegisterType((*MessageExternalNetReq)(nil), "types.MessageExternalNetReq")
	proto.RegisterType((*MessageExternalNetResp)(nil), "types.MessageExternalNetResp")
	proto.RegisterType((*MessageGetBlocksReq)(nil), "types.MessageGetBlocksReq")
	proto.RegisterType((*MessageGetBlocksResp)(nil), "types.MessageGetBlocksResp")
	proto.RegisterType((*MessageGetMempoolReq)(nil), "types.MessageGetMempoolReq")
	proto.RegisterType((*MessageVersion)(nil), "types.MessageVersion")
	proto.RegisterType((*MessageHeaderReq)(nil), "types.MessageHeaderReq")
	proto.RegisterType((*MessageHeaderResp)(nil), "types.MessageHeaderResp")
	proto.RegisterType((*MessageInvDataReq)(nil), "types.MessageInvDataReq")
	proto.RegisterType((*MessagePeerList)(nil), "types.MessagePeerList")
	proto.RegisterType((*MessageNetInfo)(nil), "types.MessageNetInfo")
	proto.RegisterType((*MessagePeersReply)(nil), "types.MessagePeersReply")
	proto.RegisterType((*MessageBroadCast)(nil), "types.MessageBroadCast")
}

func init() { proto.RegisterFile("p2pnext.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 739 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x97, 0xdb, 0x6e, 0xda, 0x4a,
	0x14, 0x86, 0x81, 0x24, 0x1c, 0x16, 0x7b, 0xe7, 0x30, 0xc9, 0x8e, 0xac, 0xad, 0x1e, 0x10, 0x57,
	0xf4, 0x10, 0x92, 0x9a, 0xbc, 0x40, 0x48, 0xab, 0x90, 0xb6, 0x89, 0x90, 0xa5, 0xf6, 0xa2, 0x77,
	0x06, 0x4f, 0x89, 0x15, 0xec, 0x99, 0x78, 0x06, 0x1a, 0xaa, 0x5e, 0xf4, 0xb5, 0xfa, 0x58, 0x7d,
	0x83, 0x8a, 0x61, 0x4e, 0x38, 0x10, 0x29, 0x6e, 0xb8, 0xb3, 0x67, 0xd6, 0x5a, 0x9f, 0xff, 0xb5,
	0x66, 0x7e, 0x01, 0xfc, 0x4b, 0x5d, 0x1a, 0xe3, 0x5b, 0xde, 0xa4, 0x09, 0xe1, 0x04, 0x6d, 0xf0,
	0x09, 0xc5, 0xec, 0xff, 0x0a, 0x75, 0xe9, 0x6c, 0xa5, 0xfe, 0x2b, 0x0f, 0xd5, 0x0b, 0xcc, 0x98,
	0x3f, 0xc0, 0xa7, 0x24, 0x8a, 0x90, 0x03, 0xa5, 0x31, 0x4e, 0x58, 0x48, 0x62, 0x27, 0x5f, 0xcb,
	0x37, 0x2a, 0x9e, 0x7a, 0x45, 0x4f, 0xa0, 0xc2, 0xc3, 0x08, 0x33, 0xee, 0x47, 0xd4, 0x29, 0xd4,
	0xf2, 0x8d, 0x35, 0xcf, 0x2c, 0xa0, 0x4d, 0x28, 0x84, 0x81, 0xb3, 0x26, 0x52, 0x0a, 0x61, 0x80,
	0xf6, 0xa1, 0x38, 0x20, 0x8c, 0x85, 0xd4, 0x59, 0xaf, 0xe5, 0x1b, 0x65, 0x4f, 0xbe, 0x4d, 0xd7,
	0x63, 0x12, 0xe0, 0xf3, 0xc0, 0xd9, 0x10, 0xb1, 0xf2, 0x0d, 0x3d, 0x03, 0x98, 0x3e, 0x75, 0x47,
	0xbd, 0x0f, 0x78, 0xe2, 0x14, 0x6b, 0xf9, 0xc6, 0x3f, 0x9e, 0xb5, 0x82, 0x10, 0xac, 0xb3, 0x70,
	0x10, 0x3b, 0x25, 0xb1, 0x23, 0x9e, 0xeb, 0xbf, 0x0b, 0xfa, 0xdb, 0x3f, 0xf1, 0x70, 0x88, 0x5e,
	0x42, 0xb1, 0x4f, 0xa2, 0x48, 0x7e, 0x7a, 0xd5, 0x45, 0x4d, 0x21, 0xb7, 0x69, 0xe9, 0xf3, 0x64,
	0x04, 0x3a, 0x82, 0x32, 0xc5, 0x38, 0x39, 0x8f, 0xbf, 0x12, 0x21, 0xc6, 0x44, 0x77, 0xdd, 0x6e,
	0x57, 0xee, 0x74, 0x72, 0x9e, 0x8e, 0x42, 0x07, 0xa6, 0x33, 0x6b, 0x22, 0x61, 0xc7, 0x24, 0x7c,
	0x9e, 0x6d, 0x74, 0x72, 0xa6, 0x5d, 0x2e, 0x80, 0x7c, 0x3c, 0xe9, 0x5f, 0x8b, 0x26, 0x54, 0xdd,
	0xed, 0xb9, 0x8c, 0x93, 0xfe, 0x75, 0x27, 0xe7, 0x59, 0x51, 0xe8, 0x18, 0xca, 0xf8, 0x96, 0xe3,
	0x24, 0xf6, 0x87, 0xa2, 0x3d, 0x55, 0x77, 0xdf, 0x64, 0xbc, 0x93, 0x3b, 0xea, 0xc3, 0x54, 0x24,
	0x6a, 0x41, 0x65, 0x80, 0x79, 0x6f, 0x48, 0xfa, 0xd7, 0x4c, 0x74, 0xae, 0xea, 0xee, 0x9a, 0xb4,
	0x33, 0xcc, 0xdb, 0x62, 0xab, 0x93, 0xf3, 0x4c, 0x1c, 0x3a, 0x80, 0x72, 0x18, 0x8f, 0x03, 0x9f,
	0xfb, 0x4c, 0xf4, 0xb4, 0xea, 0x6e, 0xc9, 0x9c, 0xf3, 0x78, 0xfc, 0x76, 0xba, 0x3c, 0x65, 0xa8,
	0x90, 0x76, 0x09, 0x36, 0xc6, 0xfe, 0x70, 0x84, 0xeb, 0xef, 0x01, 0xc9, 0x76, 0xaa, 0x26, 0x79,
	0xf8, 0x06, 0x1d, 0x43, 0x35, 0x9a, 0xad, 0x4e, 0x53, 0xef, 0x69, 0xbf, 0x1d, 0x56, 0x9f, 0xc0,
	0xee, 0x9d, 0x5a, 0x8c, 0x66, 0x2b, 0x86, 0x5e, 0x43, 0x49, 0xbe, 0x2e, 0x9f, 0xa7, 0xa7, 0x42,
	0xea, 0x13, 0xd8, 0x53, 0x68, 0x3d, 0xbd, 0xcc, 0x42, 0xd0, 0xab, 0x34, 0xfb, 0xee, 0xd1, 0x30,
	0xe8, 0xef, 0xf0, 0xdf, 0x02, 0x74, 0x66, 0xdd, 0x0f, 0x62, 0x53, 0xd8, 0x54, 0xec, 0x30, 0x1e,
	0x64, 0x17, 0xdc, 0x48, 0x43, 0x37, 0xad, 0x66, 0x4f, 0x2b, 0x6b, 0xe2, 0x0d, 0x6c, 0xcd, 0x11,
	0x33, 0xeb, 0xbc, 0x17, 0x49, 0x6c, 0x24, 0xd3, 0x22, 0x4f, 0x82, 0x20, 0x59, 0xcd, 0x54, 0xcf,
	0x30, 0x17, 0xc5, 0x17, 0xe8, 0x9c, 0x41, 0x57, 0xa2, 0x73, 0x1e, 0x39, 0x9a, 0x43, 0x7e, 0x0c,
	0x19, 0x5f, 0xc1, 0xd5, 0x51, 0xa5, 0x0d, 0xf6, 0x42, 0x9f, 0x5f, 0xe5, 0x48, 0x97, 0x98, 0x67,
	0x37, 0x81, 0x9f, 0x79, 0xd8, 0x5f, 0x54, 0x2f, 0x73, 0x03, 0x8f, 0xd2, 0x6a, 0x96, 0x78, 0xa8,
	0x7d, 0x23, 0x95, 0x0f, 0x69, 0xb3, 0xcc, 0x7e, 0x6a, 0x0e, 0xd2, 0xf8, 0x45, 0x5e, 0x6c, 0xd8,
	0xdf, 0xb4, 0x11, 0x59, 0xec, 0xcc, 0xda, 0x5f, 0xa4, 0xe1, 0x69, 0x53, 0x37, 0xe0, 0x1f, 0x36,
	0xf8, 0x02, 0x47, 0x94, 0x90, 0x61, 0x76, 0xd5, 0xcd, 0x34, 0x78, 0x6f, 0x4e, 0xb5, 0xaa, 0x6f,
	0x5d, 0x17, 0x75, 0x47, 0xa5, 0x47, 0x3d, 0xb6, 0x60, 0x59, 0xd6, 0x12, 0x7c, 0x0b, 0xdb, 0xb2,
	0x4c, 0x07, 0xfb, 0x01, 0x4e, 0x56, 0x26, 0x76, 0x56, 0xde, 0x22, 0x8f, 0x61, 0x27, 0x45, 0x5e,
	0x89, 0xdb, 0xdf, 0xe1, 0x32, 0xcd, 0x95, 0xe3, 0x5f, 0x81, 0xe1, 0xab, 0xca, 0x1a, 0x9a, 0x18,
	0xc3, 0xc7, 0xf8, 0x6f, 0x5c, 0x69, 0xe9, 0x68, 0x55, 0x5d, 0xc3, 0xe4, 0xfa, 0x34, 0x5d, 0x62,
	0x2e, 0x7e, 0xac, 0x3d, 0xb2, 0x11, 0x5e, 0x92, 0x40, 0x95, 0xb6, 0x95, 0xee, 0x58, 0x4a, 0x99,
	0x87, 0xe9, 0x70, 0xf2, 0xa0, 0xdf, 0xa0, 0x6f, 0x00, 0xa8, 0xce, 0x4c, 0xcf, 0x53, 0x6f, 0x78,
	0x56, 0x50, 0x3d, 0xd6, 0x87, 0xb8, 0x9d, 0x10, 0x3f, 0x38, 0xf5, 0x19, 0x7f, 0x10, 0x72, 0xe9,
	0xd1, 0xd5, 0xe5, 0xe6, 0xa6, 0xd9, 0x7e, 0xfe, 0xe5, 0xe9, 0x20, 0xe4, 0x57, 0xa3, 0x5e, 0xb3,
	0x4f, 0xa2, 0xc3, 0x56, 0xab, 0x1f, 0x1f, 0xf6, 0xaf, 0xfc, 0x30, 0x6e, 0xb5, 0x0e, 0x45, 0x5e,
	0xaf, 0x28, 0xfe, 0x46, 0xb4, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0xbb, 0x63, 0x1e, 0x5a, 0x69,
	0x0c, 0x00, 0x00,
}
