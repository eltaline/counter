package counter

import "sync/atomic"

type Cuint64 struct {
	Counter uint64
}

func NewUint64() *Cuint64 {
	return new(Cuint64)
}

func (c *Cuint64) Add(val uint64) uint64 {
	return atomic.AddUint64((&c.Counter), val)
}

func (c *Cuint64) Incr() uint64 {
	return atomic.AddUint64((&c.Counter), 1)
}

func (c *Cuint64) Decr() uint64 {

	for v := c.Get(); v > 0; v = c.Get() {
		if atomic.CompareAndSwapUint64((&c.Counter), v, v-1) {
			return v - 1
		}
	}
	return 0

}

func (c *Cuint64) Sub(val uint64) uint64 {

	for v := c.Get(); (v - val) >= 0; v = c.Get() {
		if atomic.CompareAndSwapUint64((&c.Counter), v, v-val) {
			return v - val
		}
	}
	return 0

}

func (c *Cuint64) Set(v uint64) {
	atomic.StoreUint64((&c.Counter), v)
}

func (c *Cuint64) Get() uint64 {
	return atomic.LoadUint64((&c.Counter))
}
