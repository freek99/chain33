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
	"google.golang.org/grpc"
)

// cmd parse
var startAddr string
var key string
var output string
var perfType string

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
	flag.StringVar(&startAddr, "startAddr", "", "start Address,for example:127.0.0.1:13802")
	flag.StringVar(&key, "key", "", "when perfType=broadcast key is block hash or tx hash")
	flag.StringVar(&output, "output", "",
		"when perfType=broadcast output to png image file,for example:./my.png,"+
			"when perfType=rollback output to a text file,for example :/rollback.txt")
	flag.StringVar(&perfType, "perfType", "", "perfType value is 'broadcast' or 'rollback' ")
	flag.Parse()
}

func main() {
	// collect and analyze data
	if perfType == "" {
		fmt.Println("err=", "perfType can't be empty")
		return
	}

	if startAddr == "" {
		fmt.Println("err=", "startAddr can't be empty")
		return
	}
	if key == "" {
		fmt.Println("err=", "key can't be empty")
		return
	}

	if perfType == "rollback" {
		if output != "" {
			f, err := os.OpenFile(output, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()

			log.SetOutput(f)
		}
	} else if perfType == "broadcast" {
		if output == "" {
			fmt.Println("err=", "output can't be empty")
			return
		}
	}


	var keys []string
	keys = append(keys, key)
	searcher := &Searcher{make(map[string]*grpc.ClientConn)}


	switch {
	case perfType == "rollback":
		monitor := &RollbackAnalyzer{}

		addrs := strings.Split(startAddr, ",")
		replyList := make(map[string]*pb.MetricsInfoReply)

		for _, addr := range addrs {
			replys := searcher.Search(addr, keys, false)
			for _, reply := range replys {
				replyList[addr] = reply
			}
		}
		monitor.Analyze(replyList)
	case perfType == "broadcast":
		replys := searcher.Search(startAddr, keys,true)
		analyzer := &BroadcastAnalyzer{}
		stat := analyzer.Analyze(replys)
		fmt.Println(
			"\nstartAddr=", startAddr,
			"\nhash=", keys,
			"\ntotalSize=", stat.TotalSize, "byte",
			"\nduration1=", stat.Duration1, "ms",
			"\nduration2=", stat.Duration2, "ms",
			"\ntimes=", stat.Times,
		)

		// output to GraphViz file
		gvPath := "/tmp/" + key + ".gv"

		viewer := &BroadcastViewer{}
		graphvizData := viewer.ExportToGraphVizData(replys)
		ioutil.WriteFile(gvPath, graphvizData, 0666)

		// output to png image file
		toPng := "dot " + gvPath + " -T png -o " + output
		execCmd(toPng)
		fmt.Println("graph=", output)
		//fmt.Println("Now you can run this command for png output:\n"+toPng)
	}
}
