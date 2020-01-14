package metrics

import (
	"github.com/33cn/chain33/queue"
	"github.com/33cn/chain33/types"
)

type Metrics struct {
	client          queue.Client
	store           *Store
	cfg             *types.Config
	listenAddr      string
	listenID        string
}


func NewMetrics(cfg *types.Config) *Metrics {
	m := &Metrics{}
	m.store = NewStore(cfg.P2P.MetricsDB)
	return m
}

func (m *Metrics) handleAddMetricsInfo(msg *queue.Message) {
	if info, ok := msg.GetData().(*types.MetricsInfo); ok {

		if m.store.connected {
			if m.listenAddr == "" {
				m.listenAddr = info.Dst
				m.listenID = info.DstID
			}
			m.store.insertMetrics(m.listenID,m.listenAddr,info)
		}
	}
}


func (m *Metrics) SetQueueClient(c queue.Client) {
	m.client = c

	go func() {
		m.client.Sub("metrics")
		for msg := range m.client.Recv() {
			switch msg.Ty {
			case types.EventAddMetricsInfo:
				m.handleAddMetricsInfo(msg)
			}
		}
	}()
}
