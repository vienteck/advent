package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	//read file and map calories to elves
	//read file
	file, err := os.Open("calories.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	highestElfCalories := 0
	currentElfCalorie := 0
	topfour := [4]int{0, 0, 0, 0}
	for scanner.Scan() {
		row := scanner.Text()
		if row == "" {
			//log.Printf("Elf %d has %d calories", elf, elves[fmt.Sprintf("Elf%d", elf)])
			if currentElfCalorie > highestElfCalories {
				highestElfCalories = currentElfCalorie
				moveleadersdown(0, &topfour, highestElfCalories)
			} else if currentElfCalorie > topfour[1] {
				moveleadersdown(1, &topfour, currentElfCalorie)
			} else if currentElfCalorie > topfour[2] {
				moveleadersdown(2, &topfour, currentElfCalorie)
			} else if currentElfCalorie > topfour[3] {
				moveleadersdown(3, &topfour, currentElfCalorie)
			}
			currentElfCalorie = 0
			continue
		}
		calories, err := strconv.Atoi(row)

		if err != nil {
			continue
		}

		currentElfCalorie += calories
	}
	log.Printf("The highest calories an elf has is %d", highestElfCalories)
	log.Printf("The other highest 3 have %d calories", topfour[0]+topfour[2]+topfour[1])
}

func moveleadersdown(place int, array *[4]int, newvalue int) {

	newarray := [4]int{0, 0, 0, 0}
	for i := 0; i < 4; i++ {
		if i == place {
			newarray[i] = newvalue
		}
		if i < place {
			newarray[i] = (*array)[i]
		}
		if i > place {
			newarray[i] = (*array)[i-1]
		}
	}
	(*array) = newarray
}
