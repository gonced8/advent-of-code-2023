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

type Number struct {
	bounds []int
	value  string
}

func prod_gear_rations(lines []string, idx int) (prod int) {
	prod = 0

	// Find all '*' in the line
	gears := regexp.MustCompile(`\*`).FindAllStringIndex(lines[idx], -1)

	// Find all numbers in all lines
	numbers := []Number{}
	for _, line := range lines {
		for _, bounds := range regexp.MustCompile(`\d+`).FindAllStringIndex(line, -1) {
			numbers = append(numbers, Number{bounds, line[bounds[0]:bounds[1]]})
		}
	}

	// Loop through all numbers
	for _, gear := range gears {
		adjacent_numbers := 0
		gear_prod := 1

		// Loop through all numbers
		for _, num := range numbers {
			if gear[1] >= num.bounds[0] && gear[0] <= num.bounds[1] {

				// There can be only two adjacent numbers
				if adjacent_numbers++; adjacent_numbers > 2 {
					break
				}

				// Parse number and multiply it to the product
				n, err := strconv.Atoi(num.value)
				if err != nil {
					panic(err)
				}
				gear_prod *= n
			}
		}

		if adjacent_numbers == 2 {
			prod += gear_prod
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
			idx = 0 // The "center" line is the first line
		default:
			// If buffer size exceeds 3, remove the oldest line
			lines = lines[1:]
			fallthrough
		case 3:
			idx = 1 // The "center" line is the middle line
		}

		sum += prod_gear_rations(lines, idx)
	}

	// Process the last line
	sum += prod_gear_rations(lines[1:], 1)

	// Print the result
	fmt.Println("Sum:", sum)

	// Check for errors during Scan
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
