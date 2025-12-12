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
	// Find dimensions needed
	maxRow, maxCol := 0, 0
	for _, tile := range tile_data {
		if tile[0] > maxRow { maxRow = tile[0] }
		if tile[1] > maxCol { maxCol = tile[1] }
	}
	
	// Create 2D slice
	green_floor := make([][]int, maxRow+1)
	for i := range green_floor {
		green_floor[i] = make([]int, maxCol+1)
	}
	
	// Place red tiles at polygon vertices
	for _, tile := range tile_data {
		green_floor[tile[0]][tile[1]] = 1
	}

	// Connect polygon edges between consecutive vertices
	for i := 0; i < len(tile_data); i++ {
		j := (i + 1) % len(tile_data)
		r1, c1 := tile_data[i][0], tile_data[i][1]
		r2, c2 := tile_data[j][0], tile_data[j][1]
		
		// Draw line between consecutive vertices using Bresenham's algorithm
		dr := int(math.Abs(float64(r2 - r1)))
		dc := int(math.Abs(float64(c2 - c1)))
		sr := 1
		if r1 > r2 { sr = -1 }
		sc := 1
		if c1 > c2 { sc = -1 }
		err := dr - dc
		
		r, c := r1, c1
		for {
			if green_floor[r][c] == 0 {
				green_floor[r][c] = 2 // Mark edge as green if not already red
			}
			
			if r == r2 && c == c2 { break }
			
			e2 := 2 * err
			if e2 > -dc {
				err -= dc
				r += sr
			}
			if e2 < dr {
				err += dr
				c += sc
			}
		}
	}

	// Fill polygon interior using ray casting
	for row := 0; row < len(green_floor); row++ {
		for col := 0; col < len(green_floor[row]); col++ {
			if green_floor[row][col] == 0 && isPointInPolygon(row, col, tile_data) {
				green_floor[row][col] = 2
			}
		}
	}
	
	return green_floor
}

// Helper function for polygon fill - ray casting algorithm
func isPointInPolygon(x, y int, polygon [][]int) bool {
	intersections := 0
	n := len(polygon)
	
	for i := 0; i < n; i++ {
		j := (i + 1) % n
		x1, y1 := float64(polygon[i][0]), float64(polygon[i][1])
		x2, y2 := float64(polygon[j][0]), float64(polygon[j][1])
		
		if ((y1 > float64(y)) != (y2 > float64(y))) {
			xIntersection := x1 + (float64(y)-y1)*(x2-x1)/(y2-y1)
			if xIntersection > float64(x) {
				intersections++
			}
		}
	}
	return intersections%2 == 1
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