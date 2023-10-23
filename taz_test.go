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
	"fmt"
	"testing"
	"time"
)

// Define a mock load test function for testing.
func mockLoadTest() error {
	// Simulate a workload or behavior for testing.
	time.Sleep(100 * time.Millisecond)
	return nil
}

func TestRunLoadTest(t *testing.T) {
	// Create a test configuration.
	config := LoadTestConfig{
		WorkerCount:       2,
		RequestsPerSecond: 5,
		TotalRequests:     10,
		LoadTestFunc:      mockLoadTest,
	}

	// Run the load test with a timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := RunLoadTest(ctx, config)
	if err != nil {
		t.Errorf("Error running load test: %v", err)
		return
	}
	fmt.Println(result.ResponseTimes)
	// Add assertions for the test results.
	if len(result.ResponseTimes) != 10 {
		t.Errorf("Expected 10 response times, got %d", len(result.ResponseTimes))
	}

	if len(result.Errors) != 0 {
		t.Errorf("Expected 0 errors, got %d", len(result.Errors))
	}
}
