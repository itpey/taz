// Copyright 2023 itpey
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package taz

import (
	"context"
	"errors"
	"sync"
	"time"
)

// LoadTestConfig represents the configuration for a load test.
type LoadTestConfig struct {
	WorkerCount       uint64       // Number of worker goroutines.
	RequestsPerSecond uint64       // Desired requests per second rate.
	TotalRequests     uint64       // Total number of requests to send.
	LoadTestFunc      func() error // The function to simulate a workload.
}

// LoadTestResult represents the results of a load test.
type LoadTestResult struct {
	ResponseTimes []time.Duration // Response times for each request.
	Errors        []error         // Errors encountered during the test.
}

// RunLoadTest executes a load test based on the provided configuration and context.
func RunLoadTest(ctx context.Context, cfg LoadTestConfig) (LoadTestResult, error) {
	if cfg.RequestsPerSecond == 0 {
		cfg.RequestsPerSecond = 10
	}
	if cfg.TotalRequests == 0 {
		cfg.TotalRequests = 100
	}
	if cfg.WorkerCount == 0 {
		cfg.WorkerCount = 10
	}
	if cfg.LoadTestFunc == nil {
		return LoadTestResult{}, errors.New("must provide a LoadTestFunc in config")
	}

	if ctx == nil {
		ctx = context.Background()
	}

	var wg sync.WaitGroup
	responseTimes := make([]time.Duration, 0, cfg.TotalRequests)
	errors := make([]error, 0)

	requestInterval := time.Second / time.Duration(cfg.RequestsPerSecond)

	requestsPerWorker := cfg.TotalRequests / cfg.WorkerCount
	remainderRequests := cfg.TotalRequests % cfg.WorkerCount

	workChannel := make(chan struct{})

	for i := uint64(0); i < cfg.WorkerCount; i++ {
		numRequests := requestsPerWorker
		if i < remainderRequests {
			numRequests++
		}

		wg.Add(1)
		go func(workerID uint64, numRequests uint64) {
			defer wg.Done()
			for j := uint64(0); j < numRequests; j++ {
				select {
				case <-workChannel:
					startTime := time.Now()
					select {
					case <-ctx.Done():
						return
					default:
					}
					if err := cfg.LoadTestFunc(); err != nil {
						errors = append(errors, err)
					}
					responseTimes = append(responseTimes, time.Since(startTime))

					time.Sleep(requestInterval)
				case <-ctx.Done():
					return
				}
			}

		}(i, numRequests)

	}
	close(workChannel)
	wg.Wait()

	return LoadTestResult{
		ResponseTimes: responseTimes,
		Errors:        errors,
	}, nil
}
