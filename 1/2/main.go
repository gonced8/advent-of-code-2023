package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Open input file
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a new scanner
	scanner := bufio.NewScanner(file)

	// Use scanner to read the file line by line
	sum := 0
	numbers := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		number := ""

		// Find the first character that is a number
	FindFirstLoop:
		for i := 0; i < len(line); i++ {
			if line[i] >= '0' && line[i] <= '9' {
				number += string(line[i])
				break FindFirstLoop
			} else {
				// Loop possible numbers
				for j := 0; j < len(numbers); j++ {
					end := i + len(numbers[j])
					if end < len(line) && line[i:end] == numbers[j] {
						number += strconv.Itoa(j + 1)
						break FindFirstLoop
					}
				}
			}
		}

		// Find the last character that is a digit
	FindLastLoop:
		for i := len(line) - 1; i >= 0; i-- {
			if line[i] >= '0' && line[i] <= '9' {
				number += string(line[i])
				break FindLastLoop
			} else {
				// Loop possible numbers
				for j := 0; j < len(numbers); j++ {
					start := i - len(numbers[j]) + 1
					if start >= 0 && line[start:i+1] == numbers[j] {
						number += strconv.Itoa(j + 1)
						break FindLastLoop
					}
				}
			}
		}

		// Add the number to the sum
		num, err := strconv.Atoi(number)
		if err != nil {
			panic(err)
		}
		sum += num
	}

	// Check for errors during Scan
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Print the sum
	fmt.Println("Sum:", sum)
}
