//imports 
package main

import(
	"fmt"
	"os"
	"bufio"
	"strings"
	"time"
	"slices"
	"runtime"
	"sync"
)

// function for part A
// given x,y,z from input, find the shortest path to the junction box
func circuits_optimizer(junction_data [][]int) int {
	//create max workers based on cpu cores
	numCores := runtime.NumCPU()
	fmt.Printf("Using %d CPU cores for optimization\n", numCores)
	var wg sync.WaitGroup
	wg.Add(numCores)
	closest_box := make(chan int, numCores)
	//create worker pool
	// iterate over junction data and find nearest junction
	
	for _, junction := range junction_data {
		fmt.Println(junction)
		go findNearestJunction(junction, junction_data, closest_box, &wg)
	}
	wg.Wait()
	close(circuits)
	//collect results
	return 

}

// function for part B


//subroutines
//load data input
func loadJunctionData(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
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

//find nearsest junction
func findNearestJunction(current_position []int, junctions [][]int, closest_box chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	min_distance := -1
	nearest_index := -1
	for i, junction := range junctions {
		if junction != current_position {
			distance := abs(current_position[0]-junction[0]) + abs(current_position[1]-junction[1]) + abs(current_position[2]-junction[2])
			if min_distance == -1 || distance < min_distance {
				min_distance = distance
				nearest_index = i
			}
		}
	}
	closest_box <- nearest_index
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