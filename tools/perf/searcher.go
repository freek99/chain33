package main

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	pb "github.com/33cn/chain33/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

const MaxCallRecvMsgSize = 100 * 1024 * 1024

type Searcher struct {
	conns map[string]*grpc.ClientConn
}

func (bs *Searcher) find(dstAddr string, hashs []string) (*pb.MetricsInfoReply, error) {
	if bs.conns[dstAddr] == nil {
		kp := keepalive.ClientParameters{
			Time:                time.Second * 5,
			Timeout:             time.Second * 20,
			PermitWithoutStream: true,
		}
		conn, err := grpc.Dial(
			dstAddr,
			grpc.WithInsecure(),
			grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(MaxCallRecvMsgSize)),
			grpc.WithKeepaliveParams(kp),
		)
		if err != nil {
			fmt.Println(err)
			return nil, errors.New("Dial error")
		}

		bs.conns[dstAddr] = conn

	}

	gcli := pb.NewP2PgserviceClient(bs.conns[dstAddr])
	resp, err := gcli.GetMetricsInfo(
		context.Background(),
		&pb.MetricsInfoParams{hashs},
		grpc.FailFast(true))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}

func (bs *Searcher) Search(startAddr string, keys []string, deepSearch bool) map[string]*pb.MetricsInfoReply {
	usedPeers := make(map[string]string)
	replys := make(map[string]*pb.MetricsInfoReply)

	dstAddr := startAddr

	for {
		reply, err := bs.find(dstAddr, keys)
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

		if newDstAddr == "" {
			break
		}

		dstAddr = newDstAddr

		if !deepSearch {
			break
		}
	}

	return replys
}
