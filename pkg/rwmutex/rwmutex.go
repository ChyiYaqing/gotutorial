package rwmutex

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu    sync.RWMutex
	count uint64
}

func (c *Counter) Incr() {
	// Lock/Unlock: 写操作调用的方法
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c *Counter) Count() uint64 {
	// RLock/RUnlock 读操作调用的方法
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.count
}

func TestRWMutex() {
	var counter Counter
	for i := 0; i < 10; i++ {
		go func() {
			for {
				fmt.Println(counter.Count())
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for {
		counter.Incr()
		time.Sleep(time.Second)
	}
}
