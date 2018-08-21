package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

var (
	// Default buffered channel capacity of 5
	chanCapacity = 5

	// Jobs channel to recieve incoming jobs
	jobs = make(chan Job, chanCapacity)

	// Data channel to recieve processed jobs from workers
	data = make(chan Data, chanCapacity)
)

// Job type
type Job struct {
	Value int `json:"value"`
}

// Data type
type Data struct {
	GoroutineID int `json:"goroutine_id"`
	Job         Job `json:"job"`
	Result      int `json:"result"`
}

// Each worker watches the jobs channel, and processes
// a job, creates a Data object and places the result on the
// data channel
func worker(w *sync.WaitGroup, goroutineID int) {
	for j := range jobs {
		data <- Data{goroutineID, j, double(j.Value)}
	}

	// When no more jobs are on the jobs channel
	// resolve waitgroup.
	w.Done()
}

// Create a new job and send to jobs channel
// Once all jobs have been created, close jobs channel
func createJobs(n int) {
	for i := 1; i <= n; i++ {
		jobs <- Job{i}
	}
	close(jobs)
}

// Create a pool of workers
// Add one waitgroup reference for each worker
// passing each a reference to the waitgroup to signal when
// it's completed it's tasks.
func makeWP(n int) {
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go worker(&wg, i)
	}
	// block and wait for all waitgroup references to be resolved
	wg.Wait()

	// once all workers have completed their tasks close the data channel
	close(data)
}

func double(v int) int {
	return v * 2
}

func main() {
	jobsFlag := flag.Int("jobs", 100, "Number of jobs for program to run")
	workersFlag := flag.Int("workers", 5, "Number of concurrent workers to spawn")
	flag.Parse()

	start := time.Now()

	// Begin creation of jobs for processing
	// Run as it's own Goroutine to continue execution of program
	go createJobs(*jobsFlag)

	// Create a channel done for signalling when workers and data processing is complete
	done := make(chan interface{})

	// Once a worker has processed a job and sent the result to the data channel,
	// this Goroutine recieves the data and prints it out.
	// Run as it's own Goroutine to continue execution of program
	// Once all values from data channel have been revieved, place value on
	// done channel to signal completion
	go func() {
		for d := range data {
			fmt.Printf("%+v\n", d)
		}
		done <- true
	}()

	// Spawn workers to handle jobs
	makeWP(*workersFlag)

	// Blocks until value placed on done channel.
	// Once value is recieved, throw away value, print out
	// time elapsed and exit program.
	<-done

	fmt.Println("Time elapsed -> ", time.Since(start))
}
