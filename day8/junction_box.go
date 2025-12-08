//imports 
package main

import(
	"fmt"
	"os"
	"bufio"
	"strings"
	"time"
	"math"
	"slices"
)

// function for part A
// given x,y,z coordinates, find groups of connected junction boxes
// based on closest neighbor distances and multiply their group sizes together
func circuits_optimizer(junction_data [][]int) int {

}

//subroutines
//load data input
func loadJunctionData(filename string) [][]int {
	file, err := os.Open(filename)
	if (err != nil) {
		panic("Error opening file: " + err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	junction_data := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		ints := []int{}
		for _, part := range parts {
			var num int
			fmt.Sscanf(part, "%d", &num)
			ints = append(ints, num)
		}
		junction_data = append(junction_data, ints)
	}
	return junction_data
}



//main function
func main() {
	fmt.Println("Advent of Code 2025 - Day 8: Junction Box")
	fileName := "test_input.txt"
	if len(os.Args) > 1{
		fileName = os.Args[1]
	}
	// load junction data
	junction_data := loadJunctionData(fileName)

	// Part A
	startA := time.Now()
	resultA := circuits_optimizer(junction_data)
	elapsedA := time.Since(startA)
	fmt.Printf("Part A: Optimal Circuit Path Length: %d (Time: %s)\n", resultA, elapsedA)

	// Part B
	// startB := time.Now()
	// resultB := another_function(junction_data)
	// elapsedB := time.Since(startB)
	// fmt.Printf("Part B: Result: %d (Time: %s)\n", resultB, elapsedB)

}