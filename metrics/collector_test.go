package metrics

import (
	"net"
	"testing"
	"time"

	"github.com/33cn/chain33/types"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	bc := NewMetricsCollector()

	bc.Add("1111", &types.MetricsInfo{
		"1111",
		"aaaaa",
		"172.20.0.6:13802",
		100,
		types.Now().UnixNano(),
	})

	bc.Add("2222", &types.MetricsInfo{
		"2222",
		"bbbbb",
		"172.20.0.9:13802",
		100,
		types.Now().UnixNano(),
	})
	var keys []string
	keys = append(keys, "1111")
	keys = append(keys, "2222")
	reply := bc.Get(keys)

	assert.Equal(t, true, len(reply) == 2)

}
