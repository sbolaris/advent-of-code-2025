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

//main function
func main() {
	fmt.Println("Cephlamath Day 6")
	// open input file
	file, err := os.Open("test_input.txt")
	//file, err := os.Open("input.txt")
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
	file.Close()
	// call cephlamath solver function
	result := cephlamath_solver(math_problems)
	fmt.Println("Cephlamath Final Result: ", result)


}