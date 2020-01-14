package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"

	pb "github.com/33cn/chain33/types"
)

type RollbackAnalyzer struct {
}

func parseOtherInfo(other string) (int,string,int)  {
	infos := strings.Split(other, ",")
	h, _ := strconv.Atoi(infos[0])
	addr := infos[1]
	depth, _ := strconv.Atoi(infos[2])

	return h,addr,depth
}

func minInt(left, right int) int {
	if left > right {
		return right
	}
	return left
}

func maxInt(left, right int) int {
	if left > right {
		return left
	}
	return right
}


func (ra *RollbackAnalyzer) Analyze(replys map[string][]*pb.MetricsInfo)  {

	var dates [1]time.Time
	for i := 0; i < 1; i++ {
		dates[i] = time.Date(2019, 12, 16+i, 0, 0, 0, 0, time.Local)
	}

	multiAddrs := make(map[string]map[string]map[time.Time][3]int)
	for addr, infos := range replys {

		if len(infos) <= 0 {
			continue
		}


		mineAddrs := make(map[string]map[time.Time][3]int)
		minHeight := int(9999999)
		maxHeight := int(0)
		maxDepth := int(0)
		allHeight := make(map[int]int)
		for _, info := range infos {
			h,addr,depth := parseOtherInfo(info.Other)
			allHeight[h] = 0

			maxDepth = maxInt(maxDepth,depth)
			minHeight = minInt(minHeight,h)
			maxHeight = maxInt(maxHeight,h)

			if mineAddrs[addr] == nil {
				mineDates := make(map[time.Time][3]int)
				for _, date := range dates {
					var values [3]int
					mineDates[date] = values
				}
				mineAddrs[addr] = mineDates
			}

			values := mineAddrs[addr][dates[0]]
			switch info.Action {
			case "ROLLBACK":
				values[0] = values[0] - 1
				values[1] = values[1] + 1
			case "ATTACH":
				values[0] = values[0] + 1
				values[2] = values[2] + 1
			}
			mineAddrs[addr][dates[0]] = values
			//switch  {
			//case info.Time > dates[0].Unix() && info.Time < dates[1].Unix():
			//	mineAddrs[info.Src][dates[0]] = mineAddrs[info.Src][dates[0]]+1
			//case info.Time > dates[1].Unix() && info.Time < dates[2].Unix():
			//	mineAddrs[info.Src][dates[1]] = mineAddrs[info.Src][dates[1]]+1
			//case info.Time > dates[2].Unix() && info.Time < dates[3].Unix():
			//	mineAddrs[info.Src][dates[2]] = mineAddrs[info.Src][dates[2]]+1
			//case info.Time > dates[3].Unix() && info.Time < dates[4].Unix():
			//	mineAddrs[info.Src][dates[3]] = mineAddrs[info.Src][dates[3]]+1
			//case info.Time > dates[4].Unix() && info.Time < dates[5].Unix():
			//	mineAddrs[info.Src][dates[4]] = mineAddrs[info.Src][dates[4]]+1
			//case info.Time > dates[5].Unix() && info.Time < dates[6].Unix():
			//	mineAddrs[info.Src][dates[5]] = mineAddrs[info.Src][dates[5]]+1
			//case info.Time > dates[6].Unix() && info.Time < dates[7].Unix():
			//	mineAddrs[info.Src][dates[6]] = mineAddrs[info.Src][dates[6]]+1
			//case info.Time > dates[7].Unix():
			//	mineAddrs[info.Src][dates[7]] = mineAddrs[info.Src][dates[7]]+1
			//}

			//log.Println("from", addr, "height", h, "depth", depth, "Interval", h-oldH, "pid", info.SrcID,
			//	"time", time.Unix(info.Time/1000000000, int64(info.Time/1000000000)*1000000000).Format("2006-01-02 15:04:05.999 "))

			//oldH = h
		}

		multiAddrs[addr] = mineAddrs

		log.Println("---", "addr", addr, "avgInterval", (maxHeight-minHeight)/len(allHeight), "maxRollDepth", maxDepth)

	}

	multiMineAddrs := make(map[string]map[string]map[time.Time][3]int)
	for netAddr, mineAddrs := range multiAddrs {
		for mineAddr, dateAddr := range mineAddrs {
			if multiMineAddrs[mineAddr] == nil {
				netAddrs := make(map[string]map[time.Time][3]int)
				for netAddr1, _ := range multiAddrs {
					dateAddrs := make(map[time.Time][3]int)
					for _, date := range dates {
						var values [3]int
						dateAddrs[date] = values
					}
					netAddrs[netAddr1] = dateAddrs
				}
				multiMineAddrs[mineAddr] = netAddrs
			}
			multiMineAddrs[mineAddr][netAddr] = dateAddr
		}
	}

	//for mineAddr,netAddrs := range multiMineAddrs {
	//
	//	for _,date := range dates {
	//		avgCount := 0
	//		hasValue := 0
	//		for _, dateAddrs := range netAddrs {
	//			if dateAddrs[date] != 0 {
	//				hasValue++
	//				avgCount = avgCount + dateAddrs[date]
	//			}
	//		}
	//		if hasValue>0 {
	//			avgCount = avgCount / hasValue
	//			fmt.Println("cccc",avgCount)
	//		}
	//
	//		for netAddr, dateAddrs := range netAddrs {
	//			if dateAddrs[date] == 0 {
	//				multiMineAddrs[mineAddr][netAddr][date] = avgCount
	//			}
	//		}
	//
	//	}
	//}

	netAddrStr := ""
	for addr, _ := range replys {
		for _, date := range dates {
			netAddrStr = netAddrStr + "," + addr + "_" + date.Format("20060102") + ",roll,add"
		}
	}
	netAddrStr = strings.Replace(netAddrStr, ":13802", "", -1)
	fmt.Printf("minerAddr,returnAddr,order" + netAddrStr + ",ticket\n")

	ScaleFactor := 10000
	for mineAddr, netAddrs := range multiMineAddrs {
		netAddrString := ""
		for _, dateAddrs := range netAddrs {
			for _, count := range dateAddrs {
				netAddrString = netAddrString + strconv.Itoa(count[0]*ScaleFactor) + ","
				netAddrString = netAddrString + strconv.Itoa(count[1]*ScaleFactor) + ","
				netAddrString = netAddrString + strconv.Itoa(count[2]*ScaleFactor) + ","
			}
		}

		cmd := exec.Command("chain33-cli", "ticket", "cold", "-m", mineAddr)
		strBuf, err := cmd.Output()
		if err == nil {
			var returnAddrs []string
			returnAddr := string(strBuf)
			if returnAddr != "ErrNotFound" && returnAddr != "" {
				comma := strings.Index(returnAddr, "[")
				commaEnd := strings.Index(returnAddr, "]")
				addrs := returnAddr[comma+1 : commaEnd-1]
				tmpReturnAddrs := strings.Split(addrs, ",")
				for _, addr := range tmpReturnAddrs {
					comma := strings.Index(addr, "1")
					addr1 := addr[comma : comma+34]
					returnAddrs = append(returnAddrs, addr1)
				}

			}
			returnAddrs = append(returnAddrs, mineAddr)

			returnAddrStr := ""
			ticket := 0.00
			for _, returnAddr := range returnAddrs {
				cmd = exec.Command("chain33-cli", "account", "balance", "-a", returnAddr, "-e", "ticket")
				strBuf, err = cmd.Output()
				if  err != nil {
					fmt.Println(err)
					return
				}

				ticketStr := string(strBuf)

				comma1 := strings.Index(ticketStr, "frozen")
				comma2 := strings.Index(ticketStr, "addr")
				if comma1 < 0 || comma2 > len(ticketStr)-1 {
					//todo deal error output
					continue
				}

				tmpTicket, err1 := strconv.ParseFloat(ticketStr[comma1+10:comma2-8], 64)
				if err1 == nil {
					ticket = ticket + tmpTicket
				}


				returnAddrStr = returnAddrStr + "_" + returnAddr
			}

			fmt.Printf("%s,%s,%d,%s%f\n", mineAddr, returnAddrStr, 0, netAddrString, ticket)
		} else {
			//log.Println(err)
		}
	}

}
