package worker

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type job struct {
	ID int
}

type jobResult struct {
	JobID     int
	IsSuccess bool
}

func RunPoolWithContext() {
	const (
		numJobs    = 100
		numWorkers = 5
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	jobs := make(chan job, numJobs)
	results := make(chan jobResult, numJobs)

	var wg sync.WaitGroup

	for range numWorkers {
		wg.Add(1)
		go contextWorker(ctx, &wg, jobs, results)
	}

	go func() {
		defer close(jobs)

		for i := 1; i <= numJobs; i++ {
			select {
			case <-ctx.Done():
				return
			case jobs <- job{ID: i}:
			}
		}
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		if !r.IsSuccess {
			fmt.Printf("job %d failed\n", r.JobID)
			continue
		}
		fmt.Printf("job %d success\n", r.JobID)
	}
}

func contextWorker(ctx context.Context, wg *sync.WaitGroup, jobs <-chan job, results chan<- jobResult) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case j, ok := <-jobs:
			if !ok {
				return
			}
			results <- executeJob(ctx, j)
		}
	}
}

func executeJob(ctx context.Context, j job) jobResult {
	jobCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	select {
	case <-jobCtx.Done():
		return jobResult{JobID: j.ID, IsSuccess: false}
	case <-time.After(10 * time.Millisecond):
		return jobResult{JobID: j.ID, IsSuccess: true}
	}
}
