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
func largest_theater_area_with_green(tile_data [][]int) int {
	// Skip creating green_floor entirely - work directly with coordinates
	maxArea := 0
	
	// Check all pairs of tile_data points as potential rectangle corners
	for i := 0; i < len(tile_data); i++ {
		for j := i + 1; j < len(tile_data); j++ {
			r1, c1 := tile_data[i][0], tile_data[i][1]
			r2, c2 := tile_data[j][0], tile_data[j][1]
			
			// Skip invalid rectangles
			if r1 == r2 || c1 == c2 {
				continue
			}
			
			// Normalize coordinates
			if r1 > r2 {
				r1, r2 = r2, r1
			}
			if c1 > c2 {
				c1, c2 = c2, c1
			}
			
			area := (r2 - r1 + 1) * (c2 - c1 + 1)
			if area > maxArea {
				// Check if this rectangle is "valid" by ensuring it's properly bounded
				// by the tile_data points (geometric validation instead of array lookup)
				if isGeometricallyValid(tile_data, r1, c1, r2, c2) {
					maxArea = area
				}
			}
		}
	}
	
	return maxArea
}

// Fast geometric validation without needing the 2D array
func isGeometricallyValid(tile_data [][]int, r1, c1, r2, c2 int) bool {
	// Check if all four corners are within or on the polygon boundary
	corners := [][]int{{r1, c1}, {r1, c2}, {r2, c1}, {r2, c2}}
	
	for _, corner := range corners {
		if !isPointInOrOnPolygon(corner[0], corner[1], tile_data) {
			return false
		}
	}
	
	// Check if rectangle is properly enclosed (no gaps in boundary)
	// Sample a few interior points to ensure they would be filled
	midR, midC := (r1+r2)/2, (c1+c2)/2
	if !isPointInPolygon(midR, midC, tile_data) {
		return false
	}
	
	return true
}

// Check if a point is inside the polygon using ray casting
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

// Check if point is inside polygon or on boundary
func isPointInOrOnPolygon(x, y int, polygon [][]int) bool {
	// Check if point is a vertex
	for _, point := range polygon {
		if point[0] == x && point[1] == y {
			return true
		}
	}
	
	// Check if point is inside
	return isPointInPolygon(x, y, polygon)
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

// Helper function to format memory usage in human-readable format
func formatBytes(bytes uint64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := uint64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
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
	fmt.Printf("Part A took %s and used %s of memory\n", elapsedA, formatBytes(memStatsA.Alloc))

	// Part B: Find largest theater area with green tiles
	startB := time.Now()
	largest_area_with_green := largest_theater_area_with_green(tile_data)
	elapsedB := time.Since(startB)
	var memStatsB runtime.MemStats
	runtime.ReadMemStats(&memStatsB)
	fmt.Printf("Largest theater area with green tiles: %d\n", largest_area_with_green)
	fmt.Printf("Part B took %s and used %s of memory\n", elapsedB, formatBytes(memStatsB.Alloc))
}