//imports 
package main

import "fmt"
import "bufio"
import "os"

// part A --- 1602
// the @ symbol are rolls of paper 
// the forklift can only access paper if there are fewer than 4 rolls of paper in the 8 adjacent positions
// identify which rolls of paper can be accessed by the forklift
// return the number of accessible rolls of paper
// change accessible rolls of paper by moving the forklift to x 
func forklift_navigation(paper_rolls [][]rune, remove bool) int {
	accessible_rolls := 0
	x := []int{}
	y := []int{}
	// at i, j positions to see if less than 4 adjacent @ symbols
	for i := 0; i < len(paper_rolls); i++ {
		for j := 0; j < len(paper_rolls[i]); j++ {
			if paper_rolls[i][j] == '@' {
				adjacent_count := 0
				// check 8 adjacent positions
				for x := -1; x <= 1; x++ {
					for y := -1; y <= 1; y++ {
						if x == 0 && y == 0 {
							continue
						}
						adj_i := i + x
						adj_j := j + y
						if adj_i >= 0 && adj_i < len(paper_rolls) && adj_j >= 0 && adj_j < len(paper_rolls[i]) {
							if paper_rolls[adj_i][adj_j] == '@' {
								adjacent_count++
							}
						}
					}
				}
				if adjacent_count < 4 {
					//fmt.Println("Accessible roll of paper at: ", i, j)
					accessible_rolls++
					if remove {
						x = append(x, i)
						y = append(y, j)
					}
				}
			}
		}
	}
	if remove {
		for idx := 0; idx < len(x); idx++ {
			paper_rolls[x[idx]][y[idx]] = '.'
		}
	}
	return accessible_rolls
}

// part B --- 9518
// now the elves need help moving the rolls of paper
// once a roll can be accessed it can be removed and the forklift can move to that position (x)
func forklift_navigation_remove(paper_rolls [][]rune) int {
	total_accessible := 0
	for {
		accessible := forklift_navigation(paper_rolls, true)
		if accessible == 0 {
			break
		}
		total_accessible += accessible
	}
	return total_accessible
}

//create function to read the map from the elfs (input file)
func readPaperRollsFromFile(filePath string) [][]rune {
	paper_rolls := [][]rune{}
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return paper_rolls
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines := scanner.Text()
		row := []rune{}
		for _, char := range lines {
			row = append(row, char)
		}
		paper_rolls = append(paper_rolls, row)
	}
	file.Close()
	return paper_rolls
}

func main() {
	fmt.Println("Day 4: Forklift Navigation")
	// read input file line by line
	// put each character a 2D grid
	// call forklift navigation function and print result
	file_path := "./input.txt"
	var paper_rolls = readPaperRollsFromFile(file_path)
	// call part A function and print result
	accessible_rolls := forklift_navigation(paper_rolls, false)
	fmt.Println("Accessible Rolls of Paper: ", accessible_rolls)
	//call part B function and print result
	accessible_rolls_remove := forklift_navigation_remove(paper_rolls)
	fmt.Println("Total Accessible Rolls of Paper after removal: ", accessible_rolls_remove)
}