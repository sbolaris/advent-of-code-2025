//imports 
package main
import "testing"
import "fmt"

//test and benchmark part A fucntions
func benchmarkForklift_A() {
	fmt.Println("Benchmarking Forklift A")
	input  := "test_input.txt"
	map := readPaperRollsFromFile(input)
	total_rolls = forklift_navigation(map)
	fmt.Println("Total Accessible Paper Rolls: ", total_rolls)
}
//test and benchmark part B fucntions
func benchmarkForklift_B() {
	fmt.Println("Benchmarking Forklift B")
	input  := "input.txt"
}