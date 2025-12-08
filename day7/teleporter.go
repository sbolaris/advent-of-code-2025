//imports 
package main

import(
	"fmt"
	"strconv"
	"os"
	"bufio"
	"strings"
	"time"
	"runtime"
)

// function for part A
// in the first row there should be an S (start)
// go to th next row and follow the column, if its a ^ put | next to it 
// if its a . then change it to a | and move down 
// count 1 split for every ^
func tachyon_spliter(teleporter_data [][]rune) int {
	final_splits := 0
	// iterate over rows
	num_rows := len(teleporter_data)
	num_cols := len(teleporter_data[0])
	start_index := strings.Index(string(teleporter_data[0]), "S")
	tachyon_beam := []int{start_index}
	splits := 0
	// iterate over each row
	for row := 1; row < num_rows; row++ {
		tachyon_beam, splits = tachyon_emitter(teleporter_data[row], tachyon_beam)
		final_splits += splits
	}
	return final_splits
}

//function for part B


//subroutines if needed
func tachyon_emitter(row []rune, beam_positions []int) ([]int, int) {
	new_beam_positions := []int{}
	splits := 0
	// process the current row with the current beam positions
	for _, pos := range beam_positions {
		if pos >= 0 && pos < len(row) {
			if row[pos] == '^' {
				// split the beam
				new_beam_positions = append(new_beam_positions, pos-1)
				new_beam_positions = append(new_beam_positions, pos+1)
				splits++
			} else if row[pos] == '.' {
				// continue straight down
				new_beam_positions = append(new_beam_positions, pos)
			}
		}
	}

	return new_beam_positions, splits
}

// main function to read input solve issues with tachyonic teleporter
func main() {
	fmt.Println("Tachyonic Teleporter Calibration Program")
	
	// Check if benchmark mode is requested
	if len(os.Args) > 1 && os.Args[1] == "benchmark" {
		RunCustomBenchmarks()
		return
	}
	
	// Try test input first, then input.txt
	fileName := "input.txt"
	if _, err := os.Stat("test_input.txt"); err == nil {
		fileName = "test_input.txt"
		fmt.Println("Using test_input.txt for debugging")
	}
	
	// read input file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	teleporter_data := [][]rune{}
	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)
		teleporter_data = append(teleporter_data, runes)
	}

	// call tachyon_spliter function
	start := time.Now()
	result := tachyon_spliter(teleporter_data)
	duration := time.Since(start)
	fmt.Printf("Tachyonic Teleporter Calibration Result: %d\n", result)
	fmt.Printf("Time taken for tachyon_spliter: %v\n", duration)
}