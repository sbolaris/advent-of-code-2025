//imports
package main

import(
	"fmt"
	"os"
	"bufio"
	"strings"
	"math"
	"time"
	"runtime"
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
	// ...existing code...
	// for _, row1 := range green_floor {
	// 	fmt.Println(row1)
	// }
	
	// Precompute valid tile pairs to avoid redundant checks
	validPairs := [][]int{}
	for i := 0; i < len(tile_data); i++ {
		for j := i + 1; j < len(tile_data); j++ {
			r1, c1 := tile_data[i][0], tile_data[i][1]
			r2, c2 := tile_data[j][0], tile_data[j][1]
			
			// Only store pairs that form proper rectangles
			if r1 != r2 && c1 != c2 {
				// Normalize coordinates
				if r1 > r2 {
					r1, r2 = r2, r1
				}
				if c1 > c2 {
					c1, c2 = c2, c1
				}
				
				// Quick area check before expensive validation
				area := (r2 - r1 + 1) * (c2 - c1 + 1)
				if area > max_area {
					validPairs = append(validPairs, []int{r1, c1, r2, c2, area})
				}
			}
		}
	}
	
	// Sort by area descending to check largest rectangles first
	for i := 0; i < len(validPairs)-1; i++ {
		for j := i + 1; j < len(validPairs); j++ {
			if validPairs[i][4] < validPairs[j][4] {
				validPairs[i], validPairs[j] = validPairs[j], validPairs[i]
			}
		}
	}
	
	// Check rectangles starting from largest
	for _, pair := range validPairs {
		r1, c1, r2, c2, area := pair[0], pair[1], pair[2], pair[3], pair[4]
		
		if area <= max_area {
			break // All remaining rectangles are smaller
		}
		
		// Quick validation with early exit
		valid := true
		outer:
		for r := r1; r <= r2; r++ {
			for c := c1; c <= c2; c++ {
				if green_floor[r][c] == 0 {
					valid = false
					break outer
				}
			}
		}
		
		if valid {
			max_area = area
		}
	}
	
	return max_area
}

// Fast O(n) algorithm for largest rectangle in histogram
func largestRectInHistogram(heights []int) int {
	stack := []int{}
	maxArea := 0
	
	for i := 0; i <= len(heights); i++ {
		h := 0
		if i < len(heights) {
			h = heights[i]
		}
		
		for len(stack) > 0 && h < heights[stack[len(stack)-1]] {
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			
			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}
			
			area := height * width
			if area > maxArea {
				maxArea = area
			}
		}
		stack = append(stack, i)
	}
	
	return maxArea
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

// Check if the rectangle formed by (r1, c1) and (r2, c2) has red corners and green interior
func isValidRectangle(green_floor [][]int, r1, c1, r2, c2 int) bool {
	// Check that corners are boundary tiles (1's or 2's)
	if (green_floor[r1][c1] == 0) || (green_floor[r1][c2] == 0) || 
	   (green_floor[r2][c1] == 0) || (green_floor[r2][c2] == 0) {
		return false
	}
	
	// Check that all tiles in rectangle are filled (1's or 2's, no 0's)
	for r := r1; r <= r2; r++ {
		for c := c1; c <= c2; c++ {
			if green_floor[r][c] == 0 {
				return false
			}
		}
	}
	return true
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
	startA := time.Now()
	largest_area := largest_theater_area(tile_data)
	elapsedA := time.Since(startA)
	var memStatsA runtime.MemStats
	runtime.ReadMemStats(&memStatsA)
	fmt.Printf("Largest theater area: %d\n", largest_area) //4759420470
	fmt.Printf("Part A took %s and used %d bytes of memory\n", elapsedA, memStatsA.Alloc)

	// Part B: Find largest theater area with green tiles
	startB := time.Now()
	largest_area_with_green := largest_theater_area_with_green(tile_data)
	elapsedB := time.Since(startB)
	var memStatsB runtime.MemStats
	runtime.ReadMemStats(&memStatsB)
	fmt.Printf("Largest theater area with green tiles: %d\n", largest_area_with_green)
	fmt.Printf("Part B took %s and used %d bytes of memory\n", elapsedB, memStatsB.Alloc)
}