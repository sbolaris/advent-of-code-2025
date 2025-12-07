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

//fucntion for part 1
// need to look at the columns of the math problems and perform the operand in the last line
// on all numbers above it to get the final answer which is the sum of all the column results
func cephlamath_solver(worksheet [][]string) int {
	final_result := 0
	// iterate over columns
	num_columns := len(worksheet[0])
	num_rows := len(worksheet)
	for col := 0; col < num_columns; col++ {
		// get the operand from the last row
		operand := worksheet[num_rows-1][col]
		// perform operation on all numbers above it
		column_result := 0
		for row := 0; row < num_rows-1; row++ {
			num, err := strconv.Atoi(worksheet[row][col])
			if err != nil {
				fmt.Println("Error converting to integer:", err)
				continue
			}
			switch operand {
			case "+":
				column_result += num
			case "-":
				column_result -= num
			case "*":
				if row == 0 {
					column_result = num
				} else {
					column_result *= num
				}
			case "/":
				if row == 0 {
					column_result = num
				} else {
					if num != 0 {
						column_result /= num
					} else {
						fmt.Println("Division by zero encountered.")
					}
				}
			default:
				fmt.Println("Unknown operand:", operand)
			}
		}
		final_result += column_result
	}
	return final_result
	
}

//fucntion for part 2
// need to read the file from right to left all the rows in the column form that number 
// do operations on all the numbers in that problem 
func vertical_math (worksheet [][]rune) int{
	// go through runes the check that that the last one by be a math operator
	final_result := 0
	// iterate over columns from right to left
	num_columns := len(worksheet[0])
	num_rows := len(worksheet)
	
	vertical_nums := []int{}
	
	for col := num_columns - 1; col >= 0; col-- {
		// get the operand from the last row
		operand := string(worksheet[num_rows-1][col])
		
		// Build the number from all rows above the operand row for this column
		number_str := ""
		for row := 0; row < num_rows-1; row++ {
			char := worksheet[row][col]
			// Skip spaces when building the number
			if char != ' ' {
				number_str += string(char)
			}
		}
		
		// Convert the built string to integer if it's not empty
		if number_str != "" {
			current_number, err := strconv.Atoi(number_str)
			if err != nil {
				fmt.Println("Error converting to integer:", err, "for string:", number_str)
				continue
			}
			//fmt.Println("Current number is:", current_number, "from column", col)
			vertical_nums = append(vertical_nums, current_number)
		}
		
		// If we have an operand and numbers to operate on
		if operand != " " && operand != "" && len(vertical_nums) > 0 {
			column_result := 0
			for i, num := range vertical_nums {
				switch operand {
				case "+":
					column_result += num
				case "-":
					if i == 0 {
						column_result = num
					} else {
						column_result -= num
					}
				case "*":
					if i == 0 {
						column_result = num
					} else {
						column_result *= num
					}
				case "/":
					if i == 0 {
						column_result = num
					} else {
						if num != 0 {
							column_result /= num
						} else {
							fmt.Println("Division by zero encountered.")
						}
					}
				default:
					fmt.Println("Unknown operand:", operand)
				}
			}
			//fmt.Println("Column result for operand", operand, ":", column_result)
			final_result += column_result
			vertical_nums = []int{} // Reset for next operation
		}
	}
	return final_result
}

//main function
func main() {
	fmt.Println("Cephlamath Day 6")
	
	// Check if benchmark mode is requested
	if len(os.Args) > 1 && os.Args[1] == "benchmark" {
		RunBenchmarks()
		return
	}
	
	// open input file
	//file, err := os.Open("test_input.txt")
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	scanner := bufio.NewScanner(file)
	math_problems := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		feilds := strings.Fields(line)
		// process line
		math_problems = append(math_problems, feilds)
	}
	
	// call cephlamath solver function
	start := time.Now()
	result := cephlamath_solver(math_problems)
	duration := time.Since(start)
	fmt.Printf("Cephlamath Final Result: %d\n", result)
	fmt.Printf("Time taken for cephlamath_solver: %v\n", duration)
	
	// part b requires the file be read in differently and not split into fields
	_, err = file.Seek(0, 0) // reset file pointer to beginning
	if err != nil {
		fmt.Println("Error resetting file pointer:", err)
		return
	}
	scanner2 := bufio.NewScanner(file)
	problems_math := [][]rune{}
	for scanner2.Scan() {
		line := scanner2.Text()
		runes := []rune(line)
		// process line
		problems_math = append(problems_math, runes)
	}
	file.Close()
	
	//call part B where it read the numbers as the cols instead of human math
	start = time.Now()
	resultB := vertical_math(problems_math)
	duration = time.Since(start)
	fmt.Printf("Cephlamath Final Result B: %d\n", resultB)
	fmt.Printf("Time taken for vertical_math: %v\n", duration)
}