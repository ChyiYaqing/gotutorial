package mutex

import (
	"fmt"
	"sync"
)

type Counter struct {
	// 嵌入字段
	sync.Mutex
	count uint64
}

// 计数器加1
func (c *Counter) Incr() {
	c.Lock()
	c.count++
	c.Unlock()
}

// 读取计数器值
func (c *Counter) Count() uint64 {
	c.Lock()
	defer c.Unlock()
	return c.count
}

func AdditionCount() {
	var counter Counter
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				counter.Incr()
			}
		}()
	}
	wg.Wait()
	fmt.Println(counter.Count())
}
