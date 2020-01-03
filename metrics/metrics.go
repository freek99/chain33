package metrics

import (
	"errors"

	"github.com/33cn/chain33/queue"
	"github.com/33cn/chain33/types"
)

type Metrics struct {
	collector *Collector
	client    queue.Client
}

func (m *Metrics) GetCollector() *Collector {
	return m.collector
}

func NewMetrics() *Metrics {
	m := &Metrics{}
	m.collector = NewMetricsCollector()
	return m
}

func (m *Metrics) handleAddMetricsInfo(msg *queue.Message) {
	if info, ok := msg.GetData().(*types.MetricsInfo); ok {
		m.collector.Add(info.Key,info.Action, info)
	}
}

func (m *Metrics) handleGetMetricsInfo(msg *queue.Message) {
	if params, ok := msg.GetData().(*types.MetricsInfoParams); ok {
		infos := m.collector.Get(params.Keys)
		var replys []*types.MetricsInfo
		for _, intf := range infos {
			replys = append(replys, intf.(*types.MetricsInfo))
		}
		msg.Reply(m.client.NewMessage("", types.EventGetMetricsInfoReply, replys))
	} else {
		err := errors.New("EventGetMetricsInfo error")
		msg.Reply(m.client.NewMessage("", types.EventGetMetricsInfoReply, err))
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
			case types.EventGetMetricsInfo:
				m.handleGetMetricsInfo(msg)
			}

		}
	}()
}
