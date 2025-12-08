// imports
package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
	"testing"
	"time"
)

// Test data loader
func loadTestData() [][]rune {
	file, err := os.Open("test_input.txt")
	if err != nil {
		panic("Error opening test file: " + err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	teleporter_data := [][]rune{}
	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)
		teleporter_data = append(teleporter_data, runes)
	}
	return teleporter_data
}

func loadInputData() [][]rune {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("Error opening input file: " + err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	teleporter_data := [][]rune{}
	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)
		teleporter_data = append(teleporter_data, runes)
	}
	return teleporter_data
}

// Unit Tests
func TestTachyonSpliterBasic(t *testing.T) {
	// Test with minimal data
	testData := [][]rune{
		[]rune(".S."),
		[]rune("..."),
		[]rune(".^."),
	}
	
	result := tachyon_spliter(testData)
	if result != 1 {
		t.Errorf("Expected 1 split, got %d", result)
	}
}

func TestTachyonSpliterWithTestInput(t *testing.T) {
	testData := loadTestData()
	
	result := tachyon_spliter(testData)
	// Add your expected result here based on manual calculation
	// For now, just ensure it doesn't crash and returns a reasonable number
	if result < 0 {
		t.Errorf("Expected non-negative result, got %d", result)
	}
	
	t.Logf("Test input result: %d splits", result)
}

func TestTachyonEmitterFunction(t *testing.T) {
	// Test the helper function directly
	testCases := []struct {
		name           string
		row            []rune
		beamPositions  []int
		expectedSplits int
		description    string
	}{
		{
			name:           "NoSplits",
			row:            []rune("..."),
			beamPositions:  []int{1},
			expectedSplits: 0,
			description:    "Beam passes through empty space",
		},
		{
			name:           "OneSplit",
			row:            []rune(".^."),
			beamPositions:  []int{1},
			expectedSplits: 1,
			description:    "Beam hits one splitter",
		},
		{
			name:           "MultipleSplits",
			row:            []rune("^.^"),
			beamPositions:  []int{0, 2},
			expectedSplits: 2,
			description:    "Multiple beams hit splitters",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			newPositions, splits := tachyon_emitter(tc.row, tc.beamPositions)
			
			if splits != tc.expectedSplits {
				t.Errorf("%s: expected %d splits, got %d", tc.description, tc.expectedSplits, splits)
			}
			
			t.Logf("%s: got %d splits, new positions: %v", tc.description, splits, newPositions)
		})
	}
}

// Benchmark Tests
func BenchmarkTachyonSpliterTestInput(b *testing.B) {
	testData := loadTestData()
	
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		tachyon_spliter(testData)
	}
}

func BenchmarkTachyonSpliterFullInput(b *testing.B) {
	// Only run if input.txt exists
	if _, err := os.Stat("input.txt"); os.IsNotExist(err) {
		b.Skip("input.txt not found, skipping full input benchmark")
	}
	
	inputData := loadInputData()
	
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		tachyon_spliter(inputData)
	}
}

func BenchmarkTachyonSpliterWithAllocs(b *testing.B) {
	testData := loadTestData()
	
	b.ResetTimer()
	b.ReportAllocs()
	
	for i := 0; i < b.N; i++ {
		tachyon_spliter(testData)
	}
}

func BenchmarkTachyonEmitter(b *testing.B) {
	row := []rune(".^.^.^.")
	positions := []int{1, 3, 5}
	
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		tachyon_emitter(row, positions)
	}
}

// Performance analysis with custom benchmarking
type BenchmarkResult struct {
	FunctionName      string
	Duration          time.Duration
	MemoryAllocated   uint64
	MemoryAllocations uint64
	Result            int
}

func RunCustomBenchmarks() {
	fmt.Printf("\n%s\n", strings.Repeat("=", 80))
	fmt.Printf("TACHYON TELEPORTER BENCHMARK RESULTS\n")
	fmt.Printf("%s\n", strings.Repeat("=", 80))

	testData := loadTestData()
	var inputData [][]rune
	
	// Try to load input data
	if _, err := os.Stat("input.txt"); err == nil {
		inputData = loadInputData()
	}

	iterations := 1000
	results := []BenchmarkResult{}

	// Benchmark with test data
	fmt.Printf("Running %d iterations on test data...\n", iterations)
	result := benchmarkFunction("TachyonSpliter_TestData", iterations, func() int {
		return tachyon_spliter(testData)
	})
	results = append(results, result)

	// Benchmark with full input if available
	if inputData != nil {
		fmt.Printf("Running %d iterations on full input...\n", iterations)
		result := benchmarkFunction("TachyonSpliter_FullInput", iterations, func() int {
			return tachyon_spliter(inputData)
		})
		results = append(results, result)
	}

	// Print results table
	fmt.Printf("\n%-25s %-15s %-15s %-10s %-10s\n", 
		"Function", "Avg Time", "Memory (bytes)", "Allocs", "Result")
	fmt.Printf("%s\n", strings.Repeat("-", 80))
	
	for _, r := range results {
		fmt.Printf("%-25s %-15s %-15d %-10d %-10d\n",
			r.FunctionName,
			r.Duration,
			r.MemoryAllocated,
			r.MemoryAllocations,
			r.Result,
		)
	}
	
	fmt.Printf("%s\n", strings.Repeat("=", 80))
}

func benchmarkFunction(name string, iterations int, fn func() int) BenchmarkResult {
	runtime.GC()
	
	var totalDuration time.Duration
	var totalMemAlloc uint64
	var totalMemAllocations uint64
	var result int
	
	for i := 0; i < iterations; i++ {
		var m1 runtime.MemStats
		runtime.ReadMemStats(&m1)
		
		start := time.Now()
		result = fn()
		duration := time.Since(start)
		
		var m2 runtime.MemStats
		runtime.ReadMemStats(&m2)
		
		totalDuration += duration
		totalMemAlloc += m2.TotalAlloc - m1.TotalAlloc
		totalMemAllocations += m2.Mallocs - m1.Mallocs
	}
	
	return BenchmarkResult{
		FunctionName:      name,
		Duration:          totalDuration / time.Duration(iterations),
		MemoryAllocated:   totalMemAlloc / uint64(iterations),
		MemoryAllocations: totalMemAllocations / uint64(iterations),
		Result:            result,
	}
}
