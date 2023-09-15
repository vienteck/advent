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
	sumofgroups := 0
	line := 0
	rs := 0
	groupcheck := make(map[string]int)
	for scanner.Scan() {
		if rs < 3 {
			rs++
		} else {
			groupcheck = make(map[string]int)
			rs = 1
		}
		line++
		row := scanner.Text()
		mapper := make(map[string]int)
		//O(n)
		for i := 0; i < len(row); i++ {
			if rs == 1 {
				groupcheck[string(row[i])] = rs
			} else {
				val, ok := groupcheck[string(row[i])]
				if ok && val == 2 && rs == 3 {
					sumofgroups += charPriority[string(row[i])]
				}
				if ok {
					groupcheck[string(row[i])] = rs
				}
			}

			if i < len(row)/2 {
				mapper[string(row[i])] = 1
			} else {
				if mapper[string(row[i])] == 1 {
					fmt.Printf(
						"Line %v Character %v matches. Points: %v\n",
						line,
						string(row[i]),
						charPriority[string(row[i])],
					)

					mapper[string(row[i])] += 1
					sumofprios += charPriority[string(row[i])]
				}
			}
		}
	}

	fmt.Print(sumofprios)
}
