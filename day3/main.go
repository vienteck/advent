package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Create an empty map.
	charPriority := make(map[string]int)

	// Populate the map with lowercase characters (a to z) with priority starting at 1.
	for i, char := 'a', 1; i <= 'z'; i, char = i+1, char+1 {
		charPriority[string(i)] = char
	}

	// Populate the map with uppercase characters (A to Z) with priority starting at 27.
	for i, char := 'A', 27; i <= 'Z'; i, char = i+1, char+1 {
		charPriority[string(i)] = char
	}

	file, err := os.Open("items.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	sumofprios := 0
	line := 0
	for scanner.Scan() {
		line++
		row := scanner.Text()
		mid := len(row) / 2
		lower := row[:mid]
		upper := row[mid:]
		mapper := make(map[string]int)
		for x := range lower {
			mapper[string(lower[x])] = 1
		}
		fmt.Printf("top %v len %v\nbot %v len %v\n", lower, len(lower), upper, len(upper))
		for x := range upper {
			if mapper[string(upper[x])] == 1 {
				fmt.Printf(
					"Line %v Character %v matches. Points: %v\n",
					line,
					string(upper[x]),
					charPriority[string(upper[x])],
				)

				mapper[string(upper[x])] += 1
				sumofprios += charPriority[string(upper[x])]
			}
		}

	}

	fmt.Print(sumofprios)
}
