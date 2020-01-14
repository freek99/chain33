package main

import (
	"strconv"
	"strings"

	pb "github.com/33cn/chain33/types"
	"github.com/awalterschulze/gographviz"
)

type BroadcastViewer struct {
}

func (bv *BroadcastViewer) ExportToGraphVizData(infos []*pb.MetricsInfo) []byte {
	graph := gographviz.NewGraph()
	graphAst, _ := gographviz.Parse([]byte(`digraph G{}`))
	gographviz.Analyse(graphAst, graph)

	for _, info := range infos {
		tmpDstAddr := strings.Replace(info.Dst, ":", ".", -1)
		tmpDstAddr = strings.Replace(tmpDstAddr, ".", "", -1)
		tmpSrcAddr := strings.Replace(info.Src, ":", ".", -1)
		tmpSrcAddr = strings.Replace(tmpSrcAddr, ".", "", -1)
		attrs := make(map[string]string)
		attrs["color"] = "blue"
		//attrs["label"] = strings.Replace(tmpSrcAddr,".","x",-1)
		graph.AddNode("G", tmpSrcAddr, attrs)
		//attrs["label"] = strings.Replace(tmpDstAddr,".","x",-1)
		graph.AddNode("G", tmpDstAddr, attrs)
		attrs["color"] = "green"
		attrs["label"] = strconv.Itoa(int(info.Size))
		graph.AddEdge(tmpSrcAddr, tmpDstAddr, true, attrs)
	}

	return []byte(graph.String())
}
