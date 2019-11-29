package main

import (
	"strconv"
	"strings"

	pb "github.com/33cn/chain33/types"
	"github.com/awalterschulze/gographviz"
)

type BroadcastViewer struct {
}

func (bv *BroadcastViewer) ExportToGraphVizData(replys map[string]*pb.PeersBroadInfoReply) []byte {

	graph := gographviz.NewGraph()
	graphAst, _ := gographviz.Parse([]byte(`digraph G{}`))
	gographviz.Analyse(graphAst, graph)

	for fromIPPort, reply := range replys {
		for _,info := range reply.Infos {
			tmpDstAddr := strings.Replace(fromIPPort, ":", ".", -1)
			tmpDstAddr = strings.Replace(tmpDstAddr, ".", "", -1)
			tmpSrcAddr := strings.Replace(info.SrcIPPort, ":", ".", -1)
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

	}

	return []byte(graph.String())
}
