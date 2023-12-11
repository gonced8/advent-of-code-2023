package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

	// Initialize variables
	sum := 0

	// Use scanner to read the file line by line
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		// Find the idx of ':' in "Game i: "
		idx := strings.Index(line, ":")

		// Loop games
		games := strings.Split(line[idx+1:], ";")
		colors := map[string]int{"red": 0, "green": 0, "blue": 0}

		for _, game := range games {
			// Loop colors
			counts := strings.Split(game, ",")

			for _, count := range counts {
				// Get count and color values
				var value int
				var color string
				fmt.Sscanf(count, " %d %s", &value, &color)

				if value > colors[color] {
					colors[color] = value
				}
			}
		}

		// Calculate power
		power := 1
		for _, value := range colors {
			power *= value
		}

		// Add to sum
		sum += power
	}

	// Check for errors during Scan
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Print the sum
	fmt.Println("Sum:", sum)
}
