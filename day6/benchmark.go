package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

// BenchmarkResult holds the results of a benchmark run
type BenchmarkResult struct {
	FunctionName     string
	Duration         time.Duration
	MemoryAllocated  uint64
	MemoryAllocations uint64
	Result           int
}

// BenchmarkRunner provides utilities for benchmarking functions
type BenchmarkRunner struct {
	iterations int
	results    []BenchmarkResult
}

func NewBenchmarkRunner(iterations int) *BenchmarkRunner {
	return &BenchmarkRunner{
		iterations: iterations,
		results:    make([]BenchmarkResult, 0),
	}
}

// BenchmarkFunction runs a function multiple times and measures performance
func (br *BenchmarkRunner) BenchmarkFunction(name string, fn func() int) {
	// Force garbage collection before benchmark
	runtime.GC()
	
	var totalDuration time.Duration
	var totalMemAlloc uint64
	var totalMemAllocations uint64
	var result int
	
	for i := 0; i < br.iterations; i++ {
		// Get memory stats before
		var m1 runtime.MemStats
		runtime.ReadMemStats(&m1)
		
		start := time.Now()
		result = fn()
		duration := time.Since(start)
		
		// Get memory stats after
		var m2 runtime.MemStats
		runtime.ReadMemStats(&m2)
		
		totalDuration += duration
		totalMemAlloc += m2.TotalAlloc - m1.TotalAlloc
		totalMemAllocations += m2.Mallocs - m1.Mallocs
	}
	
	// Calculate averages
	avgDuration := totalDuration / time.Duration(br.iterations)
	avgMemAlloc := totalMemAlloc / uint64(br.iterations)
	avgMemAllocations := totalMemAllocations / uint64(br.iterations)
	
	benchResult := BenchmarkResult{
		FunctionName:      name,
		Duration:          avgDuration,
		MemoryAllocated:   avgMemAlloc,
		MemoryAllocations: avgMemAllocations,
		Result:            result,
	}
	
	br.results = append(br.results, benchResult)
}

// PrintResults prints the benchmark results in a formatted table
func (br *BenchmarkRunner) PrintResults() {
	fmt.Printf("\n%s\n", strings.Repeat("=", 80))
	fmt.Printf("BENCHMARK RESULTS (%d iterations each)\n", br.iterations)
	fmt.Printf("%s\n", strings.Repeat("=", 80))
	fmt.Printf("%-25s %-15s %-15s %-10s %-10s\n", 
		"Function", "Avg Time", "Memory (bytes)", "Allocs", "Result")
	fmt.Printf("%s\n", strings.Repeat("-", 80))
	
	for _, result := range br.results {
		fmt.Printf("%-25s %-15s %-15d %-10d %-10d\n",
			result.FunctionName,
			result.Duration,
			result.MemoryAllocated,
			result.MemoryAllocations,
			result.Result,
		)
	}
	
	fmt.Printf("%s\n", strings.Repeat("=", 80))
	
	// Find fastest and slowest
	if len(br.results) > 1 {
		fastest := br.results[0]
		slowest := br.results[0]
		
		for _, result := range br.results[1:] {
			if result.Duration < fastest.Duration {
				fastest = result
			}
			if result.Duration > slowest.Duration {
				slowest = result
			}
		}
		
		fmt.Printf("Fastest: %s (%s)\n", fastest.FunctionName, fastest.Duration)
		fmt.Printf("Slowest: %s (%s)\n", slowest.FunctionName, slowest.Duration)
		
		if slowest.Duration > 0 {
			speedup := float64(slowest.Duration) / float64(fastest.Duration)
			fmt.Printf("Speed difference: %.2fx\n", speedup)
		}
	}
}

// RunBenchmarks runs all benchmarks with the loaded data
func RunBenchmarks() {
	fmt.Println("Loading test data...")
	
	// Load data for both functions
	file, err := os.Open("test_input.txt")
	if err != nil {
		fmt.Println("Error opening test file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	math_problems := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		math_problems = append(math_problems, fields)
	}

	// Reset and read as runes
	file.Seek(0, 0)
	scanner2 := bufio.NewScanner(file)
	problems_math := [][]rune{}
	for scanner2.Scan() {
		line := scanner2.Text()
		runes := []rune(line)
		problems_math = append(problems_math, runes)
	}

	// Create benchmark runner
	runner := NewBenchmarkRunner(1000) // Run each function 1000 times
	
	fmt.Println("Running benchmarks...")
	
	// Benchmark cephlamath_solver
	runner.BenchmarkFunction("cephlamath_solver", func() int {
		return cephlamath_solver(math_problems)
	})
	
	// Benchmark vertical_math
	runner.BenchmarkFunction("vertical_math", func() int {
		return vertical_math(problems_math)
	})
	
	// Print results
	runner.PrintResults()
}