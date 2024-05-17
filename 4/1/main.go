package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse_numbes(numbers string) []int {
	var result []int
	for _, num := range strings.Split(numbers, " ") {
		n, err := strconv.Atoi(num)
		if err != nil {
			continue
		}
		result = append(result, n)
	}
	return result
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

	var winning_numbers []int
	var my_numbers []int

	// Initialize variables
	sum := 0

	// Use scanner to read the file line by line
	for scanner.Scan() {
		// Extract winning and my numbers
		line := scanner.Text()

		// Strip the begging string and split the line into winning and my numbers
		numbers := strings.Split(strings.Split(line, ": ")[1], " | ")

		// Parse winning and my numbers
		winning_numbers = parse_numbes(numbers[0])
		my_numbers = parse_numbes(numbers[1])

		// Count the number of matches
		matches := 0
		for _, n := range winning_numbers {
			for _, m := range my_numbers {
				if n == m {
					matches++
					break
				}
			}
		}

		// Add the number of matches to the sum
		if matches > 0 {
			sum += 1 << (matches - 1)
		}
	}

	// Check for errors during Scan
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Print the result
	fmt.Println("Sum:", sum)
}
