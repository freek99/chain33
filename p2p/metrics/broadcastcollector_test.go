package metrics

import (
	"github.com/33cn/chain33/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	bc := &BroadcastCollector{}
	bc.Init()

	bc.Add(&types.PeersBroadInfo{
		"1111",
		"aaaaa",
		"172.20.0.6:13802",
		100,
		types.Now().UnixNano(),
	})

	bc.Add(&types.PeersBroadInfo{
		"2222",
		"bbbbb",
		"172.20.0.9:13802",
		100,
		types.Now().UnixNano(),
	})
	var hashs []string
	hashs = append(hashs,"1111")
	hashs = append(hashs,"2222")
	reply := bc.Get(hashs)

	assert.Equal(t,true, len(reply) == 2)
}
