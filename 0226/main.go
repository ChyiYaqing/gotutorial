package main

import "github.com/pingcap/failpoint"

func main() {
	failpoint.Inject("testPanic", func() {
		panic("failpoint triggerd")
	})
}
