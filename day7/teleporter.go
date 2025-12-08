//imports 
package main

import(
	"fmt"
	"os"
	"bufio"
	"strings"
	"time"
	"slices"

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
	start_index := strings.Index(string(teleporter_data[0]), "S")
	tachyon_beam := []int{start_index}
	splits := 0
	// iterate over each row
	for row := 1; row < num_rows; row++ {
		tachyon_beam, splits = tachyon_emitter(teleporter_data[row], tachyon_beam)
		//debug statement
		//fmt.Println("Number of splits: ", splits, "in row ", row)
		final_splits += splits
	}
	return final_splits
}

//function for part B
func quantam_entanglement(teleporter_data [][]rune) int {
	time_lines := 0

	return time_lines
}

//subroutines if needed
// have to check that you dont add duplicate positions to the beam positions
// this causes the beam to be counted wrong and could lead exploded counts
func tachyon_emitter(row []rune, beam_positions []int) ([]int, int) {
	new_beam_positions := []int{}
	splits := 0
	// process the current row with the current beam positions
	for _, pos := range beam_positions {
		if pos >= 0 && pos < len(row) {
			if row[pos] == '^' {
				// split the beam
				//check if new possitions already exist in new_beam_positions
				if slices.Contains(new_beam_positions, pos-1) == false {
					new_beam_positions = append(new_beam_positions, pos-1)
				}
				if slices.Contains(new_beam_positions, pos+1) == false {
					new_beam_positions = append(new_beam_positions, pos+1)
				}
				splits++
			} else if row[pos] == '.' {
				// continue straight down
				if slices.Contains(new_beam_positions, pos) == false {
					new_beam_positions = append(new_beam_positions, pos)
				}
			}
		}
	}
	//debug statement
	//fmt.Println("Beam positions after row:", new_beam_positions)
	return new_beam_positions, splits
}

// main function to read input solve issues with tachyonic teleporter
func main() {
	fmt.Println("Tachyonic Teleporter Calibration Program")
	fileName := "test_input.txt"
	if len(os.Args) > 1{
		fileName = os.Args[1]
		// read input file
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

	//
}