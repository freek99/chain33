package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os/exec"
)

// cmd parse
var startAddr string
var hash string
var output string

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
	flag.StringVar(&startAddr, "startAddr", "", "start Address,")
	flag.StringVar(&hash, "hash", "", "block hash or tx hash")
	flag.StringVar(&output, "output", "", "output to png image file")
	flag.Parse()
}
func main() {
	// collect and analyze data
	if startAddr == ""  {
		fmt.Println("startAddr can't be empty")
		return
	}
	if hash == ""  {
		fmt.Println("hash can't be empty")
		return
	}
	var hashs []string
	hashs = append(hashs, hash)

	searcher := &BroadcastSearcher{}
	replys := searcher.Search(startAddr, hashs)
	analyzer := &BroadcastAnalyzer{}
	stat := analyzer.Analyze(replys)
	fmt.Println(
		"\nstartAddr=", startAddr,
		"\nhash=", hash,
		"\ntotalSize=", stat.TotalSize, "byte",
		"\nduration1=", stat.Duration1, "ms",
		"\nduration2=", stat.Duration2, "ms",
		"\ntimes=", stat.Times,
		"\nstartNode=", stat.StartNode,
		"\nendNode=", stat.EndNode,
	)

	// output to GraphViz file
	if output == ""  {
		fmt.Println("output can't be empty")
		return
	}

	gvPath := "/tmp/"+hash+".gv"

	viewer := &BroadcastViewer{}
	graphvizData := viewer.ExportToGraphVizData(replys)
	ioutil.WriteFile(gvPath, graphvizData, 0666)

	// output to png image file
	toPng := "dot " + gvPath  + " -T png -o " + output
	execCmd(toPng)
	//fmt.Println("Now you can run this command for png output:\n"+toPng)
}
