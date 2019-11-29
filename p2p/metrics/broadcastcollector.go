package metrics

import (
	"strings"

	pb "github.com/33cn/chain33/types"
	lru "github.com/hashicorp/golang-lru"
)

const (
	MaxPerfCacheSize = 102400
)

type BroadcastCollector struct {
	data   *lru.Cache
	enable bool
}

func (bc *BroadcastCollector) Init() {
	bc.enable = true

	var err error
	bc.data, err = lru.New(MaxPerfCacheSize)
	if err != nil {
		bc.enable = false
	}
}

func (bc *BroadcastCollector) SetEnabled(enable bool) {
	bc.enable = enable
}

func (bc *BroadcastCollector) IsEnabled() bool {
	return bc.enable
}

func (bc *BroadcastCollector) Len() int {
	return bc.data.Len()
}

func (bc *BroadcastCollector) Add(item *pb.PeersBroadInfo) {
	if bc.enable {
		bc.data.Add(bc.makeKey(item), item)
	}
}

func (bc *BroadcastCollector) Get(hashs []string) []*pb.PeersBroadInfo {

	var items []*pb.PeersBroadInfo

	if !bc.enable {
		return items
	}

	keys := bc.data.Keys()
	for _, key := range keys {
		keyStr := key.(string)
		hashExist := false
		for _, hash := range hashs {
			hashExist = strings.Contains(keyStr, hash)
			if hashExist {
				break
			}
		}

		if hashExist {
			if item, ok := bc.data.Get(key); ok {
				item := item.(*pb.PeersBroadInfo)
				items = append(items, item)
			}
		}

	}

	return items
}

func (bc *BroadcastCollector) makeKey(item *pb.PeersBroadInfo) string {
	return item.Hash + "_" + item.SrcIPPort
}
