package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	pb "github.com/33cn/chain33/types"
)

// cmd parse
var host string
var key string
var output string
var perfType string
var dbConfig string

func execCmd(s string) {
	cmd := exec.Command(`/bin/sh`, `-c`, s)
	var out bytes.Buffer

	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", out.String())
}

func init() {
	flag.StringVar(&host, "host", "", "host,for example:127.0.0.1:13802")
	flag.StringVar(&key, "key", "", "when perfType=broadcast key is block hash or tx hash")
	flag.StringVar(&output, "output", "",
		"when perfType=broadcast output to png image file,for example:./my.png,"+
			"when perfType=rollback output to a text file,for example :/rollback.txt")
	flag.StringVar(&perfType, "perfType", "", "perfType value is 'broadcast' or 'rollback' ")
	flag.StringVar(&dbConfig, "dbConfig", "", "db config  ")
	flag.Parse()
}

func main() {
	// collect and analyze data
	if perfType == "" {
		fmt.Println("err=", "perfType can't be empty")
		return
	}

	if dbConfig == "" {
		fmt.Println("err=", "dbConfig can't be empty")
		return
	}

	if perfType == "rollback" {
		if host == "" {
			fmt.Println("err=", "host can't be empty")
			return
		}
		if output != "" {
			f, err := os.OpenFile(output, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()

			log.SetOutput(f)
		}
	} else if perfType == "broadcast" {
		if key == "" {
			fmt.Println("err=", "key can't be empty")
			return
		}
		if output == "" {
			fmt.Println("err=", "output can't be empty")
			return
		}
	}

	db := &MetricsDB{}
	if db.connect(dbConfig) {
		fmt.Println("connect db success")
	} else {
		return
	}

	switch {
	case perfType == "rollback":
		addrs := strings.Split(host, ",")
		replyList := make(map[string][]*pb.MetricsInfo)

		for _, addr := range addrs {
			replys := db.searchRewardAction(addr)
			for _, reply := range replys {
				replyList[addr] = reply
			}
		}

		analyzer := &RollbackAnalyzer{}
		analyzer.Analyze(replyList)

	case perfType == "broadcast":
		reply := db.searrchBroadcastAction(key)

		//replys := searcher.Search(startAddr, keys, true)
		analyzer := &BroadcastAnalyzer{}
		stat := analyzer.Analyze(reply)
		fmt.Println(
			"\nhash=", key,
			"\ntotalSize=", stat.TotalSize, "byte",
			"\nduration1=", stat.Duration1, "ms",
			"\nduration2=", stat.Duration2, "ms",
			"\ntimes=", stat.Times,
		)

		// output to GraphViz file
		gvPath := "/tmp/" + key + ".gv"

		viewer := &BroadcastViewer{}
		graphvizData := viewer.ExportToGraphVizData(reply)
		ioutil.WriteFile(gvPath, graphvizData, 0666)

		// output to png image file
		toPng := "dot " + gvPath + " -T png -o " + output
		execCmd(toPng)
		fmt.Println("graph=", output)
		//fmt.Println("Now you can run this command for png output:\n"+toPng)
	}
}
