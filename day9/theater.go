//imports
package main

import(
	"fmt"
	"os"
	"bufio"
	"strings"
	"math"
)

//part A
func largest_theater_area(tile_data [][]int) int {
	max_area := 0
	for i := 0; i < len(tile_data); i++ {
		for j := i+1; j < len(tile_data); j++ {
			// Compute area between row i and row j
			area := 0
			point1 := tile_data[i]
			point2 := tile_data[j]
			width := int(math.Abs(float64(point2[0] - point1[0])))+1
			height := int(math.Abs(float64(point2[1] - point1[1])))+1
			area = width * height
			if area > max_area {
				max_area = area
			}
			//fmt.Printf("Area between rows %d and %d: %d\n", i, j, area)
		}
	}
			
	return max_area
}

//part B
// can only do red or green tiles 
func largest_theater_area_with_green(tile_data [][]int) int {
	green_floor := floorWithGreenTiles(tile_data)
	max_area := 0
	// find largest rectangle of red and green tiles where the corner tiles are red
	// corner tiles are 1
	// all the inside tiles are 2
	for _, row1 := range green_floor {
		fmt.Println(row1)
	}
	return max_area
}



//subroutines
func loadSeatData(fileName string) [][]int {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	var seat_data [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		var row []int
		for _, part := range parts {
			var val int
			fmt.Sscanf(part, "%d", &val)
			row = append(row, val)
		}
		seat_data = append(seat_data, row)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	return seat_data
}

//floormap with green tiles
// linerally connect the red tiles with green tiles and fill in the floor
func floorWithGreenTiles(tile_data [][]int) [][]int {
	green_floor := [][]int{}
	max_x := 0
	max_y := 0
	for i := 0; i < len(tile_data); i++ {
		index_row := tile_data[i][0]
		index_col := tile_data[i][1]
		if index_row > max_x {
			max_x = index_row
		}
		if index_col > max_y {
			max_y = index_col
		}
	}
	// create 2D slice
	green_floor = make([][]int, max_x+1)
	for i := range green_floor {
		green_floor[i] = make([]int, max_y+1)
	}
	// fill in red tiles
	for i := 0; i < len(tile_data); i++ {
		index_row := tile_data[i][0]
		index_col := tile_data[i][1]
		green_floor[index_row][index_col] = 1 // red tile
	}
	// fill in green tiles between red tiles, horizontal and vertical only and inside of the bounding box
	// fill in green tiles in side of polygon made by 1 and 2's
	for i := 0; i < len(tile_data); i++ {
		for j := i+1; j < len(tile_data); j++ {
			point1 := tile_data[i]
			point2 := tile_data[j]
			if point1[0] == point2[0] { // same row
				row := point1[0]
				col_start := int(math.Min(float64(point1[1]), float64(point2[1])))
				col_end := int(math.Max(float64(point1[1]), float64(point2[1])))
				for c := col_start; c <= col_end; c++ {
					if green_floor[row][c] == 0 {
						green_floor[row][c] = 2 // green tile
					}
				}
			} else if point1[1] == point2[1] { // same column
				col := point1[1]
				row_start := int(math.Min(float64(point1[0]), float64(point2[0])))
				row_end := int(math.Max(float64(point1[0]), float64(point2[0])))
				for r := row_start; r <= row_end; r++ {
					if green_floor[r][col] == 0 {
						green_floor[r][col] = 2 // green tile
					}
				}
			}
		}
	}
	for row := 0; row < len(green_floor); row++ {
		// Find leftmost and rightmost boundary tiles in this row
		leftBoundary := -1
		rightBoundary := -1
		
		for col := 0; col < len(green_floor[row]); col++ {
			if green_floor[row][col] == 1 || green_floor[row][col] == 2 {
				if leftBoundary == -1 {
					leftBoundary = col
				}
				rightBoundary = col
			}
		}
		
		// Fill everything between boundaries with 2's
		if leftBoundary != -1 && rightBoundary != -1 {
			for col := leftBoundary; col <= rightBoundary; col++ {
				if green_floor[row][col] == 0 {
					green_floor[row][col] = 2
				}
			}
		}
	}
	return green_floor
}



//main function
func main() {
	fmt.Println("Advent of Code 2025 - Day 8: Junction Box")
	fileName := "test_input.txt"
	if len(os.Args) > 1{
		fileName = os.Args[1]
	}
	// load red tile data from file
	tile_data := loadSeatData(fileName)
	// Part A: Find largest theater area
	largest_area := largest_theater_area(tile_data)
	fmt.Printf("Largest theater area: %d\n", largest_area)
	// Part B: Find largest theater area with green tiles
	largest_area_with_green := largest_theater_area_with_green(tile_data)
	fmt.Printf("Largest theater area with green tiles: %d\n", largest_area_with_green)


}