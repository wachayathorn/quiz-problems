package design

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type (
	job struct {
		ID int
	}

	result struct {
		JobID     int
		IsSuccess bool
	}
)

func run() {
	const (
		// numJobs is the total number of jobs to dispatch.
		numJobs = 100

		// numWorkers is the size of the worker pool.
		numWorkers = 5
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	jobs := make(chan job)
	results := make(chan result, numJobs)

	var wg sync.WaitGroup

	// start a fixed number of long-lived workers.
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go ctxWorker(ctx, &wg, jobs, results)
	}

	go func() {
		// close jobs after all jobs are dispatched or when context is cancelled
		defer close(jobs)

		// dispatch jobs while respecting context cancellation,
		for i := 1; i <= numJobs; i++ {
			select {
			case <-ctx.Done():
				return
			case jobs <- job{ID: i}:
			}
		}
	}()

	// once all workers have exited (jobs drained or context cancelled),
	// close the results channel so the main loop can finish.
	go func() {
		wg.Wait()
		close(results)
	}()

	// consume results until the channel is closed.
	for r := range results {
		if !r.IsSuccess {
			fmt.Printf("job %d failed\n", r.JobID)
			continue
		}
		fmt.Printf("job %d success\n", r.JobID)
	}
}

func ctxWorker(ctx context.Context, wg *sync.WaitGroup, jobs <-chan job, results chan<- result) {
	defer wg.Done()

	// run worker until ctx timeout or job channel closed
	for {
		select {
		case <-ctx.Done():
			return
		case j, ok := <-jobs:
			if !ok {
				return
			}
			results <- doJob(ctx, j)
		}
	}
}

func doJob(ctx context.Context, j job) result {
	jobCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	select {
	case <-jobCtx.Done():
		return result{JobID: j.ID, IsSuccess: false}
	case <-time.After(10 * time.Millisecond):
		return result{JobID: j.ID, IsSuccess: true}
	}
}
