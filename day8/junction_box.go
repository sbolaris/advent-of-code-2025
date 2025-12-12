//imports 
package main

import(
	"fmt"
	"os"
	"bufio"
	"strings"
	"time"
	"math"
	"sort"
	"slices"
)

// function for part A
// given x,y,z coordinates, find groups of connected junction boxes
// by repeatedly connecting the closest unconnected pairs
func circuits_optimizer(junction_data [][]int, total_con int) int {
	// Create list of ALL possible connections with distances
	type connection struct {
		j1, j2 int
		dist   float64
	}
	
	var all_connections []connection
	for i := 0; i < len(junction_data); i++ {
		for j := i + 1; j < len(junction_data); j++ {
			dist := distance(junction_data[i], junction_data[j])
			all_connections = append(all_connections, connection{i, j, dist})
		}
	}
	
	// KEY INSIGHT: Sort ALL distances globally (shortest first)
	sort.Slice(all_connections, func(i, j int) bool {
		return all_connections[i].dist < all_connections[j].dist
	})
	
	// Union-Find setup
	parent := make([]int, len(junction_data))
	groupSize := make([]int, len(junction_data))
	for i := range parent {
		parent[i] = i
		groupSize[i] = 1
	}
	
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	
	union := func(x, y int) bool {
		rootX, rootY := find(x), find(y)
		if rootX == rootY {
			return false
		}
		if groupSize[rootX] < groupSize[rootY] {
			rootX, rootY = rootY, rootX
		}
		parent[rootY] = rootX
		groupSize[rootX] += groupSize[rootY]
		return true
	}
	
	// Connect the shortest total_con connections globally
	connections_made := 0
	actual_merges := 0
	for _, conn := range all_connections {
		if connections_made >= total_con {
			break
		}
		
		connections_made++ // Count every attempt
		if union(conn.j1, conn.j2) {
			actual_merges++ // Count only successful merges
		}
	}
	
	fmt.Printf("Attempted connections: %d, Actual merges: %d\n", connections_made, actual_merges)
	
	// Count final group sizes
	finalGroups := make(map[int]int)
	for i := 0; i < len(junction_data); i++ {
		root := find(i)
		finalGroups[root] = groupSize[root]
	}
	
	var sizes []int
	for _, size := range finalGroups {
		sizes = append(sizes, size)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))
	
	fmt.Printf("Group sizes: %v\n", sizes)
	
	if len(sizes) >= 3 {
		return sizes[0] * sizes[1] * sizes[2]
	}
	result := 1
	for _, size := range sizes {
		result *= size
	}
	return result
}

//part B
// need to determine the first connection that creates one big circuit
// take those two junction boxes x coordinates and return the multipled value
// example multiplies the x coodinates of the first two boxes bu answer want the last 2 
func find_complete_circuit_path(junction_data [][]int) int {
	// Try each junction box as starting point
	for start := 0; start < len(junction_data); start++ {
		// Union-Find setup for this attempt
		parent := make([]int, len(junction_data))
		groupSize := make([]int, len(junction_data))
		for i := range parent {
			parent[i] = i
			groupSize[i] = 1
		}
		
		var find func(int) int
		find = func(x int) int {
			if parent[x] != x {
				parent[x] = find(parent[x])
			}
			return parent[x]
		}
		
		union := func(x, y int) bool {
			rootX, rootY := find(x), find(y)
			if rootX == rootY {
				return false
			}
			if groupSize[rootX] < groupSize[rootY] {
				rootX, rootY = rootY, rootX
			}
			parent[rootY] = rootX
			groupSize[rootX] += groupSize[rootY]
			return true
		}
		
		// Build connections starting from this junction
		connected := make([]bool, len(junction_data))
		connected[start] = true
		connections_made := 0
		
		// Keep connecting to closest unconnected junction until all are connected
		for connections_made < len(junction_data)-1 {
			closest_dist := math.Inf(1)
			closest_pair := []int{-1, -1}
			
			// Find closest pair where one is connected and one isn't
			for i := 0; i < len(junction_data); i++ {
				for j := i + 1; j < len(junction_data); j++ {
					// Skip if both already connected or both unconnected
					if connected[i] == connected[j] {
						continue
					}
					
					dist := distance(junction_data[i], junction_data[j])
					if dist < closest_dist {
						closest_dist = dist
						closest_pair = []int{i, j}
					}
				}
			}
			
			// Make the connection
			if closest_pair[0] != -1 {
				i, j := closest_pair[0], closest_pair[1]
				if union(i, j) {
					connected[i] = true
					connected[j] = true
					connections_made++
					
					// Check if all are now in one group
					root := find(0)
					all_connected := true
					for k := 1; k < len(junction_data); k++ {
						if find(k) != root {
							all_connected = false
							break
						}
					}
					
					if all_connected {
						// Found complete connection! Calculate result
						x1 := junction_data[i][0]
						x2 := junction_data[j][0]
						result := x1 * x2
						
						fmt.Printf("Complete connection found starting from junction %d\n", start)
						fmt.Printf("Last connection: junction %d (x=%d) to junction %d (x=%d)\n", i, x1, j, x2)
						fmt.Printf("X-coordinates product: %d * %d = %d\n", x1, x2, result)
						
						return result
					}
				}
			} else {
				break // No more connections possible
			}
		}
	}
	
	return -1 // No complete circuit found
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

//calculate distance between two junctions
func distance(j1, j2 []int) float64 {
	dx := j1[0] - j2[0]
	dy := j1[1] - j2[1]
	dz := j1[2] - j2[2]
	return math.Sqrt(float64(dx*dx + dy*dy + dz*dz))
}

//check if connection already exists
func containsConnection(connections [][]int, j1 int) int {
	index := -1
	for i, conn := range connections {
		//if conn contains j1 and j2 in any order
		if slices.Contains(conn, j1) {
			index = i
		}
		
	}
	return index
}


//main function
func main() {
	fmt.Println("Advent of Code 2025 - Day 8: Junction Box")
	fileName := "test_input.txt"
	connections := 10  //default number of connections
	if len(os.Args) > 1{
		fileName = os.Args[1]
		connections = 1000 // Connect the 1000 closest pairs
	}
	// load junction data
	junction_data := loadJunctionData(fileName)

	// Part A
	startA := time.Now()
	resultA := circuits_optimizer(junction_data, connections)
	elapsedA := time.Since(startA)
	fmt.Printf("Part A: Optimal Circuit Path Length: %d (Time: %s)\n", resultA, elapsedA)//84968

	// Part B
	startB := time.Now()
	resultB := find_complete_circuit_path(junction_data)
	elapsedB := time.Since(startB)
	fmt.Printf("Part B: Complete Circuit X-Product: %d (Time: %s)\n", resultB, elapsedB)

}
