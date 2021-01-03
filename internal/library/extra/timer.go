package extra

import (
	"sync/atomic"
)

type Counter struct {
	Err uint32
}

//生成原子计数器
func NewCounter() *Counter {

	return &Counter{0}
}

func (c *Counter) AddErr() {
	atomic.AddUint32(&c.Err, 1)
}

func (c *Counter) ClearErr() {
	atomic.StoreUint32(&c.Err, 0)
}

func (c *Counter) CountErr() int {
	return int(atomic.LoadUint32(&c.Err))
}
