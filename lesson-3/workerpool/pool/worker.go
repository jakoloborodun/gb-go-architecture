package pool

import (
	"fmt"
	"sync"
	"workerpool/job"
)

type Worker struct {
	Id      int
	jobChan chan *job.Job
}

func NewWorker(channel chan *job.Job, ID int) *Worker {
	return &Worker{
		Id:      ID,
		jobChan: channel,
	}
}

func (w *Worker) Start(wg *sync.WaitGroup) {
	fmt.Printf("Starting worker %d\n", w.Id)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for task := range w.jobChan {
			job.DoJob(task, w.Id)
		}
	}()
}
