package ticker

import "time"

type Ticker interface {
	C() <-chan time.Time
	Stop()
	Reset(d time.Duration)
}

type ticker struct {
	t *time.Ticker
}

func (t ticker) C() <-chan time.Time {
	return t.t.C
}

func (t ticker) Stop() {
	t.t.Stop()
}

func (t ticker) Reset(d time.Duration) {
	t.t.Reset(d)
}

func NewTicker(d time.Duration) Ticker {
	return &ticker{time.NewTicker(d)}
}
