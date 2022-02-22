package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Aim() int {
	// Open the file.
	f, _ := os.Open("input")
	// Create a new Scanner for the file.
	scanner := bufio.NewScanner(f)
	// Loop over all lines & Count for each time the current line is greater than the previous
	coordinates := []int{0, 0}
	for scanner.Scan() {
		line := scanner.Text()
		coordinates = parse(coordinates, line)
	}
	return coordinates[0] * coordinates[1]
}

func parse(coordinates []int, line string) []int {
	commands := strings.Split(line, " ")
	var direction string = commands[0]
	distance, err := strconv.Atoi(commands[1])
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	switch direction {
	case "forward":
		coordinates[0] += distance
	case "up":
		coordinates[1] -= distance
	case "down":
		coordinates[1] += distance
	}
	return coordinates
}

func main() {
	fmt.Println(Aim())
}
