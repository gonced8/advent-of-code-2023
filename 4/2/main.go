package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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
	// file, err := os.Open("../example.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a new scanner
	scanner := bufio.NewScanner(file)

	// Initialize variables
	var winning_numbers []int
	var my_numbers []int
	cache := make(map[int]int)
	sum := 0

	// Use scanner to read the file line by line
	for scanner.Scan() {
		// Extract winning and my numbers
		line := scanner.Text()

		// Parse the card number (e.g., "Card    1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53")
		card_str := regexp.MustCompile(`\d+`).FindString(line)
		card, err := strconv.Atoi(card_str)
		if err != nil {
			panic(err)
		}

		// Parse winning and my numbers
		numbers := strings.Split(strings.Split(line, ": ")[1], " | ")
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

		// Increment the future cards in the cache
		for i := card + 1; i <= card+matches; i++ {
			cache[i] += cache[card] + 1
		}

		// Update the sum and then delete the card from the cache
		sum += cache[card] + 1
		delete(cache, card)
	}

	// Check for errors during Scan
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Print the result
	fmt.Println("Sum:", sum)
}
