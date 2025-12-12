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
	// Create a structure to hold all distances
	connections_count := 0
	junction_distance := make([][]int, len(junction_data))
	//create distance matrix
	for i := 0; i < len(junction_data); i++ {
		junction_distance[i] = make([]int, len(junction_data))
		for j := 0; j < len(junction_data); j++ {
			if i == j {
				junction_distance[i][j] = 0
			} else {
				dist := distance(junction_data[i], junction_data[j])
				junction_distance[i][j] = int(dist * 1000)  //scale to avoid float issues
			}
		}
	}
	//print out junction distance matrix
	// for i := 0; i < len(junction_distance); i++ {
	// 	fmt.Println(junction_distance[i])
	// }
	// 0-19-
	//

	connections := [][]int{}
	//go through distance matrix and find closest pairs using index to join int to connections
	for i := 0; i < len(junction_distance); i++ {
		//print out what is in connections
		for _, conn := range connections {
			fmt.Println("Current connection:", conn)
		}

		if connections_count == total_con {
			break
		}
		if containsConnection(connections, i) >=0 {
			continue
		}
		// find min distance in row i
		//zero := slices.Index(junction_distance[i], slices.Min(junction_distance[i]))
		min_dist := slices.Min(slices.Delete(junction_distance[i], i, i+1))  //remove self-distance
		min_index := slices.Index(junction_distance[i], min_dist) //+1 to account for deleted self-distance
		if min_index >= i {
			min_index +=1  //adjust index after deletion
		}
		fmt.Println("Junction ", i, " closest to Junction ", min_index, " with distance ", min_dist)
		// if min_index == i {
		// 	continue
		// }
		// if nieghbor not already connected, connect them in new entry
		if containsConnection(connections, i,) <0 && containsConnection(connections,min_index)<0 {
			connections = append(connections, []int{i, min_index})
			connections_count +=1
			continue
		}
		//if one junction is already connected, connect the other to it
		if containsConnection(connections, i) >=0 && containsConnection(connections,min_index)<0 {
			index := containsConnection(connections, i)
			connections[index] = append(connections[index], min_index)
			connections_count +=1
			continue
		}
		if containsConnection(connections, min_index) >=0 && containsConnection(connections,i)<0 {
			index := containsConnection(connections, min_index)
			connections[index] = append(connections[index], i)
			connections_count +=1
			continue
		}
		if containsConnection(connections, i) >=0 && containsConnection(connections,min_index)>=0 {
			index1 := containsConnection(connections, i)
			index2 := containsConnection(connections, min_index)
			// combine the two connections to one array and delete the second
			if index1 != index2 {
				connections[index1] = append(connections[index1], connections[index2]...)
				//delete index2
				connections = append(connections[:index2], connections[index2+1:]...)
				connections_count +=1
			}
		}
	}
	
	//sort connections by size of the list and multiply the size of the top 3 connections
	sort.Slice(connections, func(i, j int) bool {
		return len(connections[i]) > len(connections[j])
	})
	//print out connections
	for _, conn := range connections {
		fmt.Println("Junction conns", conn)
	}
	
	total_length := len(connections[0])*len(connections[1])*len(connections[2])
	return total_length
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
		connections = 1000  // Connect 1000 pairs for input.txt
	}
	// load junction data
	junction_data := loadJunctionData(fileName)

	// Part A
	startA := time.Now()
	resultA := circuits_optimizer(junction_data, connections)
	elapsedA := time.Since(startA)
	fmt.Printf("Part A: Optimal Circuit Path Length: %d (Time: %s)\n", resultA, elapsedA)

	// Part B
	// startB := time.Now()
	// resultB := another_function(junction_data)
	// elapsedB := time.Since(startB)
	// fmt.Printf("Part B: Result: %d (Time: %s)\n", resultB, elapsedB)

}