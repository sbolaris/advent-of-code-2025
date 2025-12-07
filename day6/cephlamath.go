//imports 
package main
import(
	"fmt"
	"strconv"
	"os"
	"bufio"
	"strings"
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
	// iterate over columns
	num_columns := len(worksheet[0])
	num_rows := len(worksheet)
	// Can you do this recusrively?
	vertical_nums := []int{}
	for col := num_columns; 0 < num_columns; col-- {
		// get the operand from the last row
		operand := string(worksheet[num_rows-1][col])
		// perform operation on all numbers above it
		column_result := 0
		current_number, err := strconv.Atoi(string(worksheet[0:num_rows-1][col]))
		if err != nil {
			fmt.Println("Error converting to integer:", err)
			continue
		}
		fmt.Println("current number is: ", current_number)
		vertical_nums = append(vertical_nums, current_number) 
		if operand != ""{
			for _, num := range(vertical_nums){
				switch operand {
				case "+":
					column_result += num
				case "-":
					column_result -= num
				case "*":
					//if row == 0 {
						//column_result = num
					//} else {
						column_result *= num
					//}
				case "/":
					// if row == 0 {
					// 	column_result = num
					// } else {
						if num != 0 {
							column_result /= num
						} else {
							fmt.Println("Division by zero encountered.")
						}
					//}
				default:
					fmt.Println("Unknown operand:", operand)
				}
			}
			final_result += column_result
			vertical_nums = []int{}
		}
	}
	return final_result

}


//main function
func main() {
	fmt.Println("Cephlamath Day 6")
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
	//file.Close()
	// call cephlamath solver function
	result := cephlamath_solver(math_problems)
	fmt.Println("Cephlamath Final Result: ", result)
	// part b requires the file be read in differently and not split into fields
	scanner = bufio.NewScanner(file)
	problems_math := [][]rune{}
	for scanner.Scan() {
		lines := scanner.Text()
		lineOfRunes := []rune(lines)
		problems_math = append(problems_math, lineOfRunes)
	}
	file.Close()
	//call part B where it read the numbers as the cols instead of human math
	resultB := vertical_math(problems_math)
	fmt.Println("Cephlamath Final Result: ", resultB)


}