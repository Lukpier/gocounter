package gocounter

import "sync/atomic"

type Counter64 struct {
	val int64
}

func (c *Counter64) Inc() int64 {
	return atomic.AddInt64((*int64)(&c.val), 1)
}

func (c *Counter64) Get() int64 {
	return atomic.LoadInt64((*int64)(&c.val))
}
