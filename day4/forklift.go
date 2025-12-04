//imports 
package main

import "fmt"

// part A
// the @ symbol are rolls of paper 
// the forklift can only access paper if there are fewer than 4 rolls of paper in the 8 adjacent positions
// identify which rolls of paper can be accessed by the forklift
// return the number of accessible rolls of paper
// change accessible rolls of paper by moving the forklift to x 
func forklift_navigation(paper_rolls string) int {
	accessible_rolls := 0
	// parse input into 2D grid
	grid := [][]rune{}
	// at i, j positions to see if less than 4 adjacent @ symbols
	for i, row := range paper_rolls {
		grid_row := []rune{}
		for _, char := range row {
			grid_row = append(grid_row, char)
		}
		grid = append(grid, grid_row)
	}
}

// part B

func main() {
	fmt.Println("Day 4: Forklift Navigation")
}

//test and benchmark part A fucntions
func benchmarkForklift_A() {
	fmt.Println("Benchmarking Forklift A")
	input  := "input.txt"

}
//test and benchmark part B fucntions
func benchmarkForklift_B() {
	fmt.Println("Benchmarking Forklift B")
}