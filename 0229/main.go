package main

import (
	"crypto/sha1"
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

func init() {
	g := runtime.GOMAXPROCS(1)
	fmt.Println(g)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		printHashes("A")
		wg.Done()
	}()

	go func() {
		printHashes("B")
		wg.Done()
	}()

	fmt.Println("Waiting To Finish")
	// The Wait call will block until the WaitGroup is set back to 0.
	wg.Wait()

	fmt.Println("\nTerminating Program")
}

func printHashes(prefix string) {
	for i := 1; i <= 50000; i++ {
		num := strconv.Itoa(i)
		sum := sha1.Sum([]byte(num))
		fmt.Printf("%s: %05d: %x\n", prefix, i, sum)
	}
	fmt.Println("Completed", prefix)
}
