package metrics

import (
	"strings"

	pb "github.com/33cn/chain33/types"
	lru "github.com/hashicorp/golang-lru"
)

const (
	MaxPerfCacheSize = 10240
)

type BroadcastCollector struct {
	data   *lru.Cache
	enable bool
}

func (bc *BroadcastCollector) Init() {
	var err error
	bc.data, err = lru.New(MaxPerfCacheSize)
	bc.enable = true
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
		bc.data.Add(bc.getKey(item), item)
	}
}


func (bc *BroadcastCollector) Get(itemID string) []*pb.PeersBroadInfo {

	var items []*pb.PeersBroadInfo

	if !bc.enable {
		return items
	}

	Keys := bc.data.Keys()
	for _, key := range Keys {
		keyStr := key.(string)
		itemExists := strings.Contains(keyStr, itemID)
		if itemExists {
			item, ok := bc.data.Get(key)
			if ok {
				items = append(items, item.(*pb.PeersBroadInfo))
			}
		}
	}
	return items
}

func (bc *BroadcastCollector) getKey(item *pb.PeersBroadInfo) string {
	return item.Hash + "_" + item.SrcIPPort
}
