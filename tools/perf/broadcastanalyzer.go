package main

import (
	"fmt"
	"sort"
	"time"

	pb "github.com/33cn/chain33/types"
)

type BroadcastStat struct {
	TotalSize int32
	Duration1 int64
	Duration2 int64
	Times     int32
}

type BroadcastAnalyzer struct {
}

func minInt64(left, right int64) int64 {
	if left > right {
		return right
	}
	return left
}

func maxInt64(left, right int64) int64 {
	if left > right {
		return left
	}
	return right
}

func (ba *BroadcastAnalyzer) Analyze(infos []*pb.MetricsInfo) *BroadcastStat {
	startTime := int64(^uint64(0) >> 1)
	endTime1 := int64(0)
	endTime2 := int64(0)
	size := int32(0)
	times := int32(0)
	endTime1List := make(map[string]int64)

	for _, info := range infos{

		startTime = minInt64(info.Time,startTime)

		if endTime1List[info.Dst] == 0 {
			endTime1List[info.Dst] = int64(^uint64(0) >> 1)
		}
		endTime1List[info.Dst] = minInt64(info.Time,endTime1List[info.Dst])
		endTime2 = maxInt64(info.Time,endTime2)

		size = size + info.Size
		times++
	}


	if len(endTime1List)>0 {
		type kv struct {
			Key   string
			Value int64
		}
		var sorts []kv
		for k, v := range endTime1List {
			sorts = append(sorts, kv{k, v})
		}
		sort.Slice(sorts, func(i, j int) bool {
			return sorts[i].Value > sorts[j].Value  // 升序
		})
		endTime1 = sorts[0].Value
	}

	scale := int64(1000000000)
	fmt.Println(
		time.Unix(startTime/scale, startTime-int64(startTime/scale)*scale).Format("2006-01-02 15:04:05.999 "),
		time.Unix(endTime1/scale, endTime1-int64(endTime1/scale)*scale).Format("2006-01-02 15:04:05.999 "),
		time.Unix(endTime2/scale, endTime2-int64(endTime2/scale)*scale).Format("2006-01-02 15:04:05.999 "),
		)

	duration1 := (endTime1 - startTime) / int64(time.Millisecond)
	duration2 := (endTime2 - startTime) / int64(time.Millisecond)

	return &BroadcastStat{size, duration1, duration2, times}
}
