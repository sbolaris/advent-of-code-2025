//imports 
package main

import(
	"fmt"
	"strconv"
	"os"
	"bufio"
	"strings"
	"time"
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
		tachyon_beam, splits =tachyon_emitter(teleporter_data[row], tachyon_beam)
		final_splits += splits
	}

}

//function for part B


//subroutines if needed
func tachyon_emitter([]rune, []int) []int, int {
	new_beam_positions := []int{}
	splits := 0
	// process the current row with the current beam positions
	for _, pos := range beam_positions {
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

	// placeholder
	return new_beam_positions, splits
}

// main function to read input solve issues with tachyonic teleporter
func main() {
	fmt.Println("Tachyonic Teleporter Calibration Program")
	// read input file
	file, err := os.Open("input.txt")
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