// imports 
package main

import "fmt"
import "strconv"
import "os"
import "bufio"

// function for part 1
// each line of input is the single bank of batteries
// turn on exactly two bateries that two are the max jotlage
// this is the two max digits in the list
// return the two numbeers as a 2 digit int
func jotlage_meter(batt_bank string) int {
	max_joltage := 0
	// parse input into list of integers
	for i, char := range batt_bank {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			fmt.Println("Error converting character to integer:", err)
			continue
		}
		for _, char2 := range batt_bank[i+1:] {
			digit2, err := strconv.Atoi(string(char2))
			if err != nil {
				fmt.Println("Error converting character to integer:", err)
				continue
			}
			combined_joltage := strconv.Itoa(digit) + strconv.Itoa(digit2)
			combined_int, err := strconv.Atoi(combined_joltage)
			if err != nil {
				fmt.Println("Error converting combined joltage to integer:", err)
				continue
			}
			if combined_int > max_joltage {
				max_joltage = combined_int
				//fmt.Println("Max digits are:", digit, digit2)
			}
		}
		
	}
	// return their final number combined as int
	return max_joltage
}
func sumIntSlice(ints []int) int {
	sum := 0
	for _, v := range ints {
		sum += v
	}
	return sum
}


// function for part 2
// now joltages instead of two batteries, find 12 batteries that give max joltage

// main function
func main() {
	fmt.Println("Advent of Code 2025 - Day 3")
	//read input file to pass to functions 
	//file_path := "./test_input.txt"
	file_path := "./input.txt"
	file, err := os.Open(file_path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	total_joltage := []int{}
	total_joltage12 := []int{}
	scanner := bufio.NewScanner(file)
	//go workers := 4
	for scanner.Scan() {
		lines := scanner.Text()
		// call joltage meter function and print result
		fmt.Println("Processing line: ", string(lines))
		// call part 1 function and print result
		joltage := jotlage_meter(string(lines))
		//fmt.Println("Joltage Meter Reading: ", joltage) 
		total_joltage = append(total_joltage, joltage)
		//total_joltage = append(total_joltage, go jotlage_meter(string(lines)))
		// call part 2 function and print result
	}
	file.Close()
	sum_joltage := sumIntSlice(total_joltage) //16854
	fmt.Println("Total Joltage Meter Reading with 2 batteries: ", sum_joltage)
	
	
	

	
}