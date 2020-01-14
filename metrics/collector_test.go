package metrics

import (
	"github.com/33cn/chain33/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	bc := NewMetricsCollector()

	bc.Add("1111","RECV", &types.MetricsInfo{
		"",
		"",
		"1111",
		"aaaaa",
		"",
		"172.20.0.6:13802",
		100,
		types.Now().UnixNano(),
		"",
	})

	bc.Add("2222","RECV", &types.MetricsInfo{
		"",
		"",
		"2222",
		"aaaaa",
		"",
		"172.20.0.7:13802",
		100,
		types.Now().UnixNano(),
		"",
	})

	var keys []string
	keys = append(keys, "1111")
	keys = append(keys, "2222")
	reply := bc.Get(keys)

	assert.Equal(t, true, len(reply) == 2)

}
