package job

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

type Job struct {
	ID      int
	Address string
	Method  string
	Body    string
}

func NewJob(id int, addr, method, body string) *Job {
	return &Job{
		ID:      id,
		Address: addr,
		Method:  method,
		Body:    body,
	}
}

func DoJob(job *Job, workerId int) {
	start := time.Now()
	client := &http.Client{}
	req, err := http.NewRequest(
		job.Method,
		job.Address,
		strings.NewReader(job.Body))
	if err != nil {
		log.Println(err)
	}

	res, err := client.Do(req)
	//_, err = client.Do(req)
	if err != nil {
		log.Println(err)
	}
	//data, _ := ioutil.ReadAll(res.Body)
	//defer res.Body.Close()
	status := res.StatusCode
	elapsed := time.Since(start)

	fmt.Printf("worker [%d] - status code [%d] and request time [%s]\n", workerId, status, elapsed)
}
