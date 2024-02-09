package collector

import (
	"context"
	"sync"
	"time"
)

type historyItem struct {
	prev    int64
	current int64
}

type history map[time.Duration]*historyItem

type QueryCollector struct {
	sync.Mutex
	ctx context.Context

	history history
	calls   int64
}

func NewQueryCollector(ctx context.Context) *QueryCollector {
	qc := &QueryCollector{
		ctx: ctx,
		history: history{
			time.Second: {},
			time.Minute: {},
		},
	}
	qc.run()
	return qc
}

func (qc *QueryCollector) run() {
	for d, h := range qc.history {
		d := d
		h := h
		go func() {
			ticker := time.NewTicker(d)
			defer ticker.Stop()
			for {
				select {
				case <-qc.ctx.Done():
					return
				case <-ticker.C:
					h.prev = h.current
					h.current = qc.calls
				}
			}
		}()
	}
}

func (qc *QueryCollector) IncCalls() {
	qc.Lock()
	defer qc.Unlock()
	qc.calls++
}

func (qc *QueryCollector) CallsFrequency(d time.Duration) int64 {
	h, ok := qc.history[d]
	if !ok {
		return 0
	}
	return h.current - h.prev
}

func (qc *QueryCollector) TotalCalls() int64 {
	return qc.calls
}
