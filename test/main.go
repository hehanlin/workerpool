package main

import (
	"fmt"
	"time"

	"github.com/hehanlin/workerpool"
)

func main() {
	pool := workerpool.NewWorkerPool(3)
	pool.Debug = true
	for i := 0; i < 100; i++ {
		i := i
		pool.Add(func() {
			// 业务代码
			time.Sleep(1 * time.Second)
			fmt.Printf(",task num %d\n", i)
		})
	}
	pool.Wait()
}
