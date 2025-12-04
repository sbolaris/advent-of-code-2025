//imports 
package main
import (
	"fmt"
)

//test and benchmark part A fucntions
func benchmarkForklift_A() {
	fmt.Println("Benchmarking Forklift A")
	input  := "test_input.txt"
	var elfmap = readPaperRollsFromFile(input)
	total_rolls := forklift_navigation(elfmap, false)
	fmt.Println("Total Accessible Paper Rolls: ", total_rolls)
}
//test and benchmark part B fucntions
func benchmarkForklift_B() {
	fmt.Println("Benchmarking Forklift B")
	input  := "test_input.txt"
	var elfmap = readPaperRollsFromFile(input)
	total_rolls := forklift_navigation_remove(elfmap)
	fmt.Println("Total Accessible Paper Rolls after removal: ", total_rolls)
}