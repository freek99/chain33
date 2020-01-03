package main

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"

	pb "github.com/33cn/chain33/types"
)

type RollbackAnalyzer struct {
}

func (rm *RollbackAnalyzer) Analyze(replys map[string]*pb.MetricsInfoReply) int {

	var dates [1]time.Time
	for i := 0; i < 1; i++ {
		dates[i] = time.Date(2019, 12, 16+i, 0, 0, 0, 0, time.Local)
	}

	multiAddrs := make(map[string]map[string]map[time.Time]int)
	for addr, reply := range replys {

		if len(reply.Infos) <= 0 {
			continue
		}

		mineAddrs := make(map[string]map[time.Time]int)
		times := int32(0)
		minHeight := int(9999999)
		maxHeight := int(0)
		maxDepth := int(0)
		oldH := 0
		allHeight := make(map[int]int)
		for _, info := range reply.Infos {
			times++

			otherInfo := info.Other
			infos := strings.Split(otherInfo, ",")
			h, _ := strconv.Atoi(infos[0])
			addr := infos[1]
			depth,_ := strconv.Atoi(infos[2])
			allHeight[h] = 0

			if depth > maxDepth {
				maxDepth = depth
			}


			if h < minHeight {
				minHeight = h
			}
			if h > maxHeight {
				maxHeight = h
			}

			if mineAddrs[addr] == nil {
				mineDates := make(map[time.Time]int)
				for _, date := range dates {
					mineDates[date] = 0
				}
				mineAddrs[addr] = mineDates
			}

			mineAddrs[addr][dates[0]] = mineAddrs[addr][dates[0]] + 1

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

			log.Println("from", addr, "height", h, "depth", depth, "Interval", h-oldH,"hash",info.Key,"time",
				time.Unix(info.Time, 0).Format("2006-01-02 15:04:05 "))

			oldH = h
		}

		multiAddrs[addr] = mineAddrs

		log.Println("---","addr", addr, "avgInterval", (maxHeight-minHeight)/len(allHeight), "maxRollDepth", maxDepth)

		//return 0 //maxHeight

	}

	multiMineAddrs := make(map[string]map[string]map[time.Time]int)
	for netAddr, mineAddrs := range multiAddrs {
		for mineAddr, dateAddr := range mineAddrs {
			if multiMineAddrs[mineAddr] == nil {
				netAddrs := make(map[string]map[time.Time]int)
				for netAddr1, _ := range multiAddrs {
					dateAddrs := make(map[time.Time]int)
					for _, date := range dates {
						dateAddrs[date] = 0
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
			netAddrStr = netAddrStr + "," + addr + "_" + date.Format("20060102")
		}
	}
	netAddrStr = strings.Replace(netAddrStr, ":13802", "", -1)
	log.Printf("minerAddr,returnAddr,order" + netAddrStr + ",ticket\n")

	for mineAddr, netAddrs := range multiMineAddrs {
		netAddrString := ""
		for _, dateAddrs := range netAddrs {
			for _, count := range dateAddrs {
				netAddrString = netAddrString + strconv.Itoa(count*100000) + ","
			}
		}

		cmd := exec.Command("chain33-cli", "ticket", "cold", "-m", mineAddr)
		strBuf, err := cmd.Output()
		if err == nil {
			returnAddr := string(strBuf)
			if returnAddr != "ErrNotFound" && returnAddr != "" {
				comma := strings.Index(returnAddr, "1")
				returnAddr = returnAddr[comma : comma+34]
			} else {
				returnAddr = mineAddr
			}

			cmd = exec.Command("chain33-cli", "account", "balance", "-a", returnAddr, "-e", "ticket")
			strBuf, err = cmd.Output()

			ticket := "0"
			if err == nil {
				ticketStr := string(strBuf)

				comma1 := strings.Index(ticketStr, "frozen")
				comma2 := strings.Index(ticketStr, "addr")
				if comma1 < 0 || comma2 > len(ticketStr)-1 {
					//todo deal error output
					continue
				}
				ticket = ticketStr[comma1+10 : comma2-8]
			}

			log.Printf("%s,%s,%d,%s%s\n", mineAddr, returnAddr, 0, netAddrString, ticket)
		} else {
			//log.Println(err)
		}
	}

	return 0
}
