package main

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

// Helper function to load test data
func loadTestData() ([][]string, [][]rune) {
	file, err := os.Open("test_input.txt")
	if err != nil {
		panic("Error opening test file: " + err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	math_problems := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		math_problems = append(math_problems, fields)
	}

	// Reset file pointer and read as runes
	file.Seek(0, 0)
	scanner2 := bufio.NewScanner(file)
	problems_math := [][]rune{}
	for scanner2.Scan() {
		line := scanner2.Text()
		runes := []rune(line)
		problems_math = append(problems_math, runes)
	}

	return math_problems, problems_math
}

func BenchmarkCephlaMathSolver(b *testing.B) {
	math_problems, _ := loadTestData()
	
	// Reset timer to exclude setup time
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		cephlamath_solver(math_problems)
	}
}

func BenchmarkVerticalMath(b *testing.B) {
	_, problems_math := loadTestData()
	
	// Reset timer to exclude setup time
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		vertical_math(problems_math)
	}
}

// Benchmark both functions together
func BenchmarkBothFunctions(b *testing.B) {
	math_problems, problems_math := loadTestData()
	
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		cephlamath_solver(math_problems)
		vertical_math(problems_math)
	}
}

// Memory allocation benchmarks
func BenchmarkCephlaMathSolverAllocs(b *testing.B) {
	math_problems, _ := loadTestData()
	
	b.ResetTimer()
	b.ReportAllocs() // Report memory allocations
	
	for i := 0; i < b.N; i++ {
		cephlamath_solver(math_problems)
	}
}

func BenchmarkVerticalMathAllocs(b *testing.B) {
	_, problems_math := loadTestData()
	
	b.ResetTimer()
	b.ReportAllocs() // Report memory allocations
	
	for i := 0; i < b.N; i++ {
		vertical_math(problems_math)
	}
}