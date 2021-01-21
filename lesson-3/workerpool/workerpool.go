package main

import (
	"log"
	"sync"
)

type Job struct {
	payload []byte
}

type Worker struct {
	wg      *sync.WaitGroup
	num     int // only for example
	jobChan <-chan *Job
}

func main() {
	wg := &sync.WaitGroup{}
	jobChan := make(chan *Job)
	for i := 0; i < 5; i++ {
		worker := NewWorker(i+1, wg, jobChan)
		wg.Add(1)
		go worker.Handle()
	}

	jobChan <- &Job{
		payload: []byte("Some message 1"),
	}
	jobChan <- &Job{
		payload: []byte("Some message 2"),
	}
	jobChan <- &Job{
		payload: []byte("Some message 3"),
	}
	close(jobChan)
	wg.Wait()
}

func (w *Worker) Handle() {
	defer w.wg.Done()
	for job := range w.jobChan {
		log.Printf("worker %d processing job with payload %s", w.num, string(job.payload))
	}
}

func NewWorker(num int, wg *sync.WaitGroup, jobChan <-chan *Job) *Worker {
	return &Worker{
		wg:      wg,
		num:     num,
		jobChan: jobChan,
	}
}
