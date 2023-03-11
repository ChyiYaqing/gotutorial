package mutex

import (
	"fmt"
	"sync"
)

func AdditionCount() {
	var count = 0
	// 使用WaitGroup等待10个goroutine完成
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		go func() {
			wg.Add(1)
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				// count++ 不是原子操作
				count++
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
