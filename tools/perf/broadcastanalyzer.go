package main

import (
	"time"

	pb "github.com/33cn/chain33/types"
)

type BroadcastStat struct {
	TotalSize int32
	Duration1 int64
	Duration2 int64
	Times     int32
	StartNode string
	EndNode   string
}

type BroadcastAnalyzer struct {
}

func (ba *BroadcastAnalyzer) Analyze(replys []*pb.PeersBroadInfoReply) *BroadcastStat {
	startTime := int64(^uint64(0) >> 1)
	endTime1 := int64(0)
	endTime2 := int64(0)
	size := int32(0)
	times := int32(0)
	startNode := ""
	endNode   := ""
	for _, reply := range replys {
		singleStartTime := int64(^uint64(0) >> 1)
		for _, info := range reply.Infos {
			if info.RecvTime < startTime {
				startNode = info.DstIPPort
				startTime = info.RecvTime
			}

			if info.RecvTime < singleStartTime {
				endNode = info.DstIPPort
				singleStartTime = info.RecvTime
			}

			if info.RecvTime > endTime2 {
				endTime2 = info.RecvTime
			}

			size = size + info.Size
			times++
		}

		if singleStartTime > endTime1 {
			endTime1 = singleStartTime
		}

	}

	duration1 := (endTime1 - startTime) / int64(time.Millisecond)
	duration2 := (endTime2 - startTime) / int64(time.Millisecond)

	return &BroadcastStat{size, duration1, duration2, times,startNode,endNode}
}
