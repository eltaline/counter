package counter

import "sync/atomic"

type Cint64 struct {
	Counter int64
}

func NewInt64() *Cint64 {
	return new(Cint64)
}

func (c *Cint64) Add(val int64) int64 {
	return atomic.AddInt64((&c.Counter), val)
}

func (c *Cint64) Incr() int64 {
	return atomic.AddInt64((&c.Counter), 1)
}

func (c *Cint64) Decr() int64 {

	for v := c.Get(); v > 0; v = c.Get() {
		if atomic.CompareAndSwapInt64((&c.Counter), v, v-1) {
			return v - 1
		}
	}
	return 0

}

func (c *Cint64) Sub(val int64) int64 {

	for v := c.Get(); (v - val) >= 0; v = c.Get() {
		if atomic.CompareAndSwapInt64((&c.Counter), v, v-val) {
			return v - val
		}
	}
	return 0

}

func (c *Cint64) Set(v int64) {
	atomic.StoreInt64((&c.Counter), v)
}

func (c *Cint64) Get() int64 {
	return atomic.LoadInt64((&c.Counter))
}
