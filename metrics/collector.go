package metrics

import (
	"fmt"
	"github.com/33cn/chain33/common"
	lru "github.com/hashicorp/golang-lru"
	"strings"
)

const (
	maxLRUCacheSize = 10 * 102400
)

type Collector struct {
	data   *lru.Cache
	enable bool
}

func NewMetricsCollector() *Collector {
	c := &Collector{}
	c.init()

	return c
}

func (c *Collector) init() {
	c.enable = true

	var err error
	c.data, err = lru.New(maxLRUCacheSize)
	if err != nil {
		fmt.Println(err)
		c.enable = false
	}
}

func (c *Collector) SetEnabled(enable bool) {
	c.enable = enable
}

func (c *Collector) IsEnabled() bool {
	return c.enable
}

func (c *Collector) Len() int {
	if !c.enable {
		return 0
	}

	return c.data.Len()
}

func (c *Collector) makeKey(key string, action string) string {

	return key + "_" + action + "_" + common.GetRandPrintString(10, 20)
}

func (c *Collector) Add(key string, action string, item interface{}) {
	if c.enable {
		c.data.Add(c.makeKey(key, action), item)
	}
}

func (c *Collector) Get(keys []string) []interface{} {

	var items []interface{}

	if !c.enable {
		return items
	}

	lruKeys := c.data.Keys()
	for _, key := range lruKeys {
		keyStr := key.(string)
		keyExist := false
		for _, hash := range keys {
			keyExist = strings.Contains(keyStr, hash)
			if keyExist {
				break
			}
		}

		if keyExist {
			if item, ok := c.data.Get(key); ok {
				items = append(items, item)
			}
		}

	}

	return items
}
