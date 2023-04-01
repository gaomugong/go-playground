package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("worker", id, "started job", job)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", job)

		// send result to channel results
		results <- job * 2
	}
}

func main() {
	const jobNum = 5
	jobChan := make(chan int, jobNum)
	resultChan := make(chan int, jobNum)

	// start 3 workers and blocked with channel
	for w := 1; w <= 3; w++ {
		go worker(w, jobChan, resultChan)
	}

	// 	send 5 job to jobChan and close jobChan to indicate all job is pushed
	for i := 1; i <= jobNum; i++ {
		jobChan <- i
	}
	close(jobChan)

	// collect all the result of the workers
	for i := 1; i <= jobNum; i++ {
		fmt.Println(<-resultChan)
	}
}

// worker 3 started job 1
// worker 1 started job 2
// worker 2 started job 3
// worker 2 finished job 3
// worker 2 started job 4
// worker 1 finished job 2
// worker 1 started job 5
// worker 3 finished job 1
// 6
// 4
// 2
// worker 1 finished job 5
// 10
// worker 2 finished job 4
// 8
