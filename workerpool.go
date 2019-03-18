package workerpool

import (
	"context"
	"fmt"
	"sync"
)

type Pool struct {
	Queue     chan func() // 任务队列
	once      *sync.Once
	waitGroup *sync.WaitGroup
	WorkerNum int // goroutine数量
	Debug     bool
	ctx       context.Context
}

func NewWorkerPool(num int) *Pool {
	if num <= 0 {
		num = 1
	}
	return &Pool{
		WorkerNum: num,
		Queue:     make(chan func(), num),
		waitGroup: &sync.WaitGroup{},
		once:      &sync.Once{},
		ctx:       context.Background(),
	}
}

func (p *Pool) Add(fn func()) {
	p.once.Do(func() {
		p.waitGroup.Add(p.WorkerNum)
		go p.loop()
	})

	p.Queue <- fn
}

func (p *Pool) loop() {
	for i := 0; i < p.WorkerNum; i++ {
		go func(id int) {
			defer p.waitGroup.Done()
			for fn := range p.Queue {
				if fn == nil {
					return
				}
				if p.Debug {
					fmt.Printf("worker num %d\n", id)
				}
				fn()
			}
		}(i)
	}
}

func (p *Pool) Wait() {
}
