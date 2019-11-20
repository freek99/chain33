package metrics

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os/exec"
	"testing"
)

func System(s string) {
	cmd := exec.Command(`/bin/sh`, `-c`, s)
	var out bytes.Buffer

	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", out.String())
}

func Test(t *testing.T)  {
	startAddr := "172.21.0.5:13802"
	hash := "0x7b1e9b87170cb83084c26208a504d0eba15cf0032f88aa537a627e7685909cd8"

	searcher := &BroadcastSearcher{}
	replys := searcher.Search(startAddr, hash)

	analyzer := &BroadcastAnalyzer{}
	stat := analyzer.Analyze(replys)

	fmt.Println(
		"totalSize=", stat.TotalSize, "byte",
		"duration1=", stat.Duration1, "ms",
		"duration2=", stat.Duration2, "ms",
		"times=", stat.Times,
	)

	viewer := &BroadcastViewer{}
	graphvizData := viewer.ExportToGraphVizData(replys)
	filename := "data/" + string("img")
	ioutil.WriteFile(filename+".gv", graphvizData, 0666)
	System("dot " + filename + ".gv" + " -T png -o " + filename + ".png")

}