package main

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	pb "github.com/33cn/chain33/types"
	"google.golang.org/grpc"
)

const MaxCallRecvMsgSize = 100 * 1024 * 1024

type BroadcastSearcher struct {
}

func (bs *BroadcastSearcher) find(dstAddr string, hashs []string) (*pb.PeersBroadInfoReply, error) {
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
	resp, err := gcli.GetPeersBroadInfo(
		context.Background(),
		&pb.P2PPeersBroadInfoParams{hashs},
		grpc.FailFast(true))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}

func (bs *BroadcastSearcher) Search(startAddr string, hashs []string) map[string]*pb.PeersBroadInfoReply {
	usedPeers := make(map[string]string)
	replys := make(map[string]*pb.PeersBroadInfoReply)

	dstAddr := startAddr
	for {
		reply, err := bs.find(dstAddr, hashs)
		if err != nil {
			break
		}

		replys[dstAddr] = reply

		for _, peerInfo := range reply.Peers {
			ipPort := peerInfo.Ip + ":" + strconv.Itoa(int(peerInfo.Port))
			if usedPeers[ipPort] == "" {
				usedPeers[ipPort] = ""
			}
		}

		usedPeers[dstAddr] = dstAddr

		newDstAddr := ""
		for peer, usedPeer := range usedPeers {
			if usedPeer == "" {
				newDstAddr = peer
				break
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
