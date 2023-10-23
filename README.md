<h1 align="center">Taz - A Go Load Testing Framework</h1>

<p align="center">
  A simple yet powerful load testing framework for Go, designed to help you simulate and measure the performance of your applications under various loads.
</p>

<p align="center">
  <a href="https://pkg.go.dev/github.com/itpey/taz">
    <img src="https://pkg.go.dev/badge/github.com/itpey/taz.svg" alt="Go Reference">
  </a>
  <a href="https://github.com/itpey/taz/blob/main/LICENSE">
    <img src="https://img.shields.io/github/license/itpey/taz" alt="license">
  </a>
</p>

## Table of Contents

- [Features](#features)
- [Getting Started](#getting-started)
  - [Installation](#installation)
  - [Usage](#usage)
- [Configuration](#configuration)
- [Results](#results)
- [Error Handling](#error-handling)
- [Examples](#examples)
- [Running Tests](#running-tests)
- [License](#license)

## Features

- **Configurability:** Taz allows you to fine-tune your load tests by specifying the number of worker goroutines, desired requests per second, and the total number of requests to send.

- **Response Times:** Measure the response times for each request to gain insights into your application's performance.

- **Error Handling:** Track and report errors encountered during load testing to help identify potential issues.

- **Context Support:** Utilize Go's context package to manage the load test's lifecycle and control its execution.

## Getting Started

### Installation

To start using Taz in your Go project, follow these simple steps:

```bash
go get github.com/itpey/taz
```

### Usage

Import the Taz package in your code:

```go
import "github.com/itpey/taz"
```

Create a LoadTestConfig to configure your load test:

```go
config := taz.LoadTestConfig{
    WorkerCount:       10,
    RequestsPerSecond: 100,
    TotalRequests:     1000,
    LoadTestFunc: func() error {
        // Your workload simulation logic here
        return nil
    },
}
```
Execute the load test using the RunLoadTest function:

```go
result, err := taz.RunLoadTest(context.Background(), config)
if err is not nil {
    // Handle errors
}
```

## Configuration

Taz offers several configuration options to customize your load tests:

- **WorkerCount:** The number of worker goroutines to use in the load test.

- **RequestsPerSecond:** The desired rate of requests to be sent per second.

- **TotalRequests:** The total number of requests to send during the load test.

- **LoadTestFunc:** A user-defined function that simulates the workload for each request.

## Results

The RunLoadTest function returns a LoadTestResult structure that contains important information:

- **ResponseTimes:** An array of response times for each request, allowing you to analyze performance.

- **Errors:** An array of errors encountered during the load test, helping you identify potential issues.

## Error Handling

Proper error handling is crucial when working with Taz. Check the Errors array in the LoadTestResult to identify and troubleshoot any issues that may have occurred during the load test.

## Examples

Here's an example of a simple load test using Taz:

```go
package main

import (
	"context"
	"fmt"
	"math/rand"
	"github.com/itpey/taz"
	"time"
)

func main() {
	// Define the load test configuration
	config := taz.LoadTestConfig{
		WorkerCount:       10,
		RequestsPerSecond: 100,
		TotalRequests:     1000,
		LoadTestFunc:      simulateWorkload,
	}

	// Create a context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	
	// Run
	result, err := taz.RunLoadTest(ctx, config)

	// Check for errors
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Output load test results
	fmt.Printf("Load Test Results:\n")
	fmt.Printf("Total Requests: %d\n", config.TotalRequests)
	fmt.Printf("Successful Requests: %d\n", config.TotalRequests-uint64(len(result.Errors)))
	fmt.Printf("Failed Requests: %d\n", len(result.Errors))
	fmt.Printf("Average Response Time: %s\n", calculateAverageResponseTime(result.ResponseTimes))
}

// Simulate a workload by sleeping for a random duration (representing the work being done).
func simulateWorkload() error {
	// Simulate a workload by sleeping for a random duration between 100ms and 500ms.
	sleepDuration := time.Duration(rand.Intn(400)+100) * time.Millisecond
	time.Sleep(sleepDuration)
	return nil
}

// Calculate the average response time from the list of response times.
func calculateAverageResponseTime(responseTimes []time.Duration) time.Duration {
	if len(responseTimes) == 0 {
		return 0
	}

	var total time.Duration
	for _, rt := range responseTimes {
		total += rt
	}
	return total / time.Duration(len(responseTimes))
}
```

## Running Tests

To run tests for Taz, use the following command:

```bash
go test github.com/itpey/taz
```

## License

Taz is open-source software released under the Apache License, Version 2.0. You can find a copy of the license in the [LICENSE](https://github.com/itpey/taz/blob/main/LICENSE) file.
