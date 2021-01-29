package pool

import (
	"fmt"
	"math"
	"sync"
	"time"
	"workerpool/job"
)

type Pool struct {
	Jobs    []*job.Job
	wg      sync.WaitGroup
	num     int // only for example
	jobChan chan *job.Job
}

//func (w *Worker) Handle() {
//	defer w.wg.Done()
//	for job := range w.jobChan {
//		log.Printf("worker %d processing job with payload %s", w.num, string(job.payload))
//	}
//}

// Run starts all work in Pool and lock it until finish.
func (p *Pool) Run() {
	count := 0

	start := time.Now()
	for i := 1; i <= p.num; i++ {
		worker := NewWorker(p.jobChan, i)
		worker.Start(&p.wg)
	}

	for i := range p.Jobs {
		p.jobChan <- p.Jobs[i]
		count++
	}
	close(p.jobChan)

	p.wg.Wait()
	elapsed := time.Since(start)
	avgRPS := float64(count) / elapsed.Seconds()

	fmt.Printf("Execution time [%s]\n", elapsed)
	fmt.Printf("RPS [%v]\n", math.RoundToEven(avgRPS))
}

func NewPool(jobs []*job.Job, num int, wg sync.WaitGroup) *Pool {
	return &Pool{
		Jobs:    jobs,
		wg:      wg,
		num:     num,
		jobChan: make(chan *job.Job, 1000),
	}
}
