package worker

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type Job struct {
	ID  int
	URL string
}

type Result struct {
	JobID      int
	URL        string
	StatusCode int
	Err        error
}

func RunPool() {
	jobs := make(chan Job)
	results := make(chan Result)

	var wg sync.WaitGroup

	for range 3 {
		wg.Add(1)
		go poolWorker(jobs, results, &wg)
	}

	go func() {
		urls := []string{
			"https://httpbin.org/status/200",
			"https://httpbin.org/status/404",
			"https://httpbin.org/status/500",
			"https://httpbin.org/status/500",
			"https://httpbin.org/delay/1",
			"https://httpbin.org/delay/1",
			"https://httpbin.org/status/200",
			"https://httpbin.org/status/200",
			"https://httpbin.org/status/404",
			"https://httpbin.org/status/500",
			"https://httpbin.org/delay/1",
			"https://httpbin.org/status/200",
			"https://httpbin.org/status/200",
			"https://httpbin.org/status/404",
			"https://httpbin.org/status/404",
		}

		for i, url := range urls {
			jobs <- Job{ID: i + 1, URL: url}
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Printf("job-%d %s → %d\n", result.JobID, result.URL, result.StatusCode)
	}
}

func poolWorker(jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		parts := strings.Split(job.URL, "/")
		statusCode, _ := strconv.Atoi(parts[len(parts)-1])
		results <- Result{
			JobID:      job.ID,
			URL:        job.URL,
			StatusCode: statusCode,
		}
	}
}
