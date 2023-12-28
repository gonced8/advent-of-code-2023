package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func IntMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func sum_part_numbers(lines []string, idx int) (sum int) {
	// Initialize sum
	sum = 0

	// Find all numbers in the line
	numbers := regexp.MustCompile(`\d+`).FindAllStringIndex(lines[idx], -1)

	// Loop through all numbers
	for _, num := range numbers {
	AdjacentPositionsLoop:
		// Loop through all adjacent positions
		for i := 0; i < IntMin(len(lines), 3); i++ {
			for j := IntMax(num[0]-1, 0); j < IntMin(num[1]+1, len(lines[idx])); j++ {
				// Skip the current number
				if i == idx && j >= num[0] && j < num[1] {
					continue
				}

				// Check if it is not a number nor a dot
				if lines[i][j] != '.' && (lines[i][j] < '0' || lines[i][j] > '9') {
					// Parse number and add it to the sum
					n, err := strconv.Atoi(lines[idx][num[0]:num[1]])
					if err != nil {
						panic(err)
					}
					sum += n
					break AdjacentPositionsLoop
				}
			}
		}
	}

	return
}

func main() {
	// Open input file
	file, err := os.Open("../input.txt")
	//file, err := os.Open("../example.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a new scanner
	scanner := bufio.NewScanner(file)

	// Create buffers
	var lines []string // At most 3 lines

	// Initialize variables
	var idx int
	sum := 0

	// Use scanner to read the file line by line
	for scanner.Scan() {
		// Append the line to the buffer
		lines = append(lines, scanner.Text())

		switch len(lines) {
		case 1:
			// Skip the first iteration (no adjacent lines to compare)
			continue
		case 2:
			idx = 0
		default:
			// If buffer size exceeds 3, remove the oldest line
			lines = lines[1:]
			fallthrough
		case 3:
			idx = 1
		}

		sum += sum_part_numbers(lines, idx)
	}

	// Process the last line
	sum += sum_part_numbers(lines[1:], 1)

	// Print the result
	fmt.Println("Sum:", sum)

	// Check for errors during Scan
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
