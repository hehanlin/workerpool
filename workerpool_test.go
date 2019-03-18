package workerpool

import (
	"fmt"
	"testing"
	"time"
)

func TestWorkerPool(t *testing.T) {
	pool := NewWorkerPool(3)
	pool.Debug = true
	for i := 0; i < 100; i++ {
		i := i
		pool.Add(func() {
			// 业务代码
			time.Sleep(1 * time.Second)
			fmt.Printf("task num %d\n", i)
		})
	}
	pool.Wait()
}
