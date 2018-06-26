package counter

import (
	"time"
)

type Counter struct {
	C chan uint64
	d time.Duration
	t uint64
}

func startTimer(c *Counter) {
	tickChan := time.NewTicker(c.d).C
	go func() {
		for range tickChan {
			c.t++
			c.C <- c.t
		}
	}()
}

func NewCounter(d time.Duration) *Counter {
	c := make(chan uint64, 1)
	counter := &Counter{
		d: d,
		C: c,
	}
	startTimer(counter)
	return counter
}
