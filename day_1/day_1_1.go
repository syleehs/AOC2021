package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func day_1_1() int {
	// Open the file.
	f, _ := os.Open("input")
	// Create a new Scanner for the file.
	scanner := bufio.NewScanner(f)
	// Loop over all lines & Count for each time the current line is greater than the previous
	prev_line := 0
	count := -1
	for scanner.Scan() {
		line := scanner.Text()
		curr_line, err := strconv.Atoi(line)
		if curr_line > prev_line {
			count++
		}
		if err != nil {
			return 0
		}
		prev_line = curr_line
	}
	return count
}
func main() {
	fmt.Println(day_1_1())
}
