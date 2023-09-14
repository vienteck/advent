package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	//read file and map calories to elves
	elves := make(map[string]int)
	//read file
	file, err := os.Open("calories.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	elf := 1
	highestcalories := 1
	for scanner.Scan() {
		row := scanner.Text()
		if row == "" {
			//log.Printf("Elf %d has %d calories", elf, elves[fmt.Sprintf("Elf%d", elf)])
			if elves[fmt.Sprintf("Elf%d", elf)] > elves[fmt.Sprintf("Elf%d", highestcalories)] {
				highestcalories = elf
			}
			elf++
			continue
		}
		calories, err := strconv.Atoi(row)

		if err != nil {
			continue
		}

		elves[fmt.Sprintf("Elf%d", elf)] = elves[fmt.Sprintf("Elf%d", elf)] + calories

	}
	log.Printf("Elf number %d has the highest calories: %d", highestcalories, elves[fmt.Sprintf("Elf%d", highestcalories)])
}
