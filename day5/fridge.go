//inputs 
package main

import( 
	"fmt"
 	"strconv"
 	"os"
	"bufio"
	"strings"
	"sync"
)

//part A 
// count how many fresh items are in the fridge based on the fresh range provided
func countFreshItems(fresh_ranges []string, item_ids []int) int {
	fresh_count := 0
	for _, item_id := range item_ids {
		is_fresh_chan := make(chan bool)
		for _, fresh_range := range fresh_ranges {
			go isFresh(item_id, fresh_range, is_fresh_chan)
			if <-is_fresh_chan {
				fresh_count += 1
				break // no need to check other ranges if already fresh
			}
		}
	}
	return fresh_count
}



//part B 

//subroutines
func isFresh(item_id int, fresh_range string, is_fresh chan<- bool) {
	//start and end of range
	ranges := strings.Split(fresh_range, "-")
	if len(ranges) != 2 {
		fmt.Println("Invalid range format:", fresh_range)
		is_fresh <- false
		return
	}
	start, err1 := strconv.Atoi(ranges[0])
	end, err2 := strconv.Atoi(ranges[1])
	if err1 != nil || err2 != nil {
		fmt.Println("Error converting range bounds to integers:", err1, err2)
		is_fresh <- false
		return
	}
	if item_id >= start && item_id <= end {
		is_fresh <- true
	} else {
		is_fresh <- false
	}
}

//database acesss
// take in file from input read each line and put the first set in to the fresh range
// once a empty line is found store the next set in the inventory range 
func access_database(file_path string) ([]string, []int) {
	fresh_range := []string{}
	item_ids := []int{}
	data_mode := "ranges"
	file, err := os.Open(file_path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return fresh_range, item_ids
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines := scanner.Text()
		fmt.Println("Processing line: ", string(lines))
		if lines == "" {
			data_mode = "items"
			continue
		}
		if data_mode == "ranges" {
			fresh_range = append(fresh_range, string(lines))
		} else if data_mode == "items" {
			id, err := strconv.Atoi(lines)
			if err != nil {
				fmt.Println("Error converting item ID to integer:", err)
				continue
			}
			item_ids = append(item_ids, id)

		}
	
	}
	file.Close()
	return fresh_range, item_ids
}

//main function
func main() {
	var wg sync.WaitGroup
	fmt.Println("Advent of Code 2025 Day 5 inventory management")
	file_path := "./input.txt"
	fresh_range, item_ids := access_database(file_path)
	fmt.Println("Fresh Range: ", fresh_range)
	fmt.Println("Item IDs: ", item_ids)
	//check each item id if fresh
	wg.Add(4)
	number_of_fresh_items := countFreshItems(fresh_range, item_ids)
	fmt.Println("Number of fresh items in inventory: ", number_of_fresh_items)
	wg.Done()
}