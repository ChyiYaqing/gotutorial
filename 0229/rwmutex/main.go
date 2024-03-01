package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var data []string
var rwMutex sync.RWMutex

func main() {

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		for i := 0; i < 10; i++ {
			writer(i)
		}
		wg.Done()
	}()

	for i := 0; i < 8; i++ {
		go func(id int) {
			for {
				reader(id)
			}
		}(i)
	}
	wg.Wait()
	fmt.Println("Program Complete")
}

func writer(i int) {
	rwMutex.Lock()
	{
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		fmt.Println("****> : Performing Write")
		data = append(data, fmt.Sprintf("String: %d", i))
	}
	rwMutex.Unlock()
}

func reader(id int) {
	rwMutex.RLock()
	{
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		fmt.Printf("%d : Performing Read : Length[%d]\n", id, len(data))
	}
	rwMutex.RUnlock()
}
