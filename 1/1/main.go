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
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		number := ""

		// Find the first character that is a digit
		for _, c := range line {
			if c >= '0' && c <= '9' {
				number += string(c)
				break
			}
		}

		// Find the last character that is a digit
		for i := len(line) - 1; i >= 0; i-- {
			if line[i] >= '0' && line[i] <= '9' {
				number += string(line[i])
				break
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
