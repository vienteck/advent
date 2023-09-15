package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("procedures.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	grid := [9][80]string{}
	line := 1
	row := 7
	for scanner.Scan() {
		if line < 9 {
			prepdata(scanner.Text(), row, &grid)
			row--
		}
		if line > 11 {

			instructions := getinstructions(scanner.Text())
			log.Print(instructions)
		}
		line++
	}
}
func getinstructions(row string) [3]int {

	instructions := [3]int{}
	split := strings.Split(row, " ")
	parse, err := strconv.Atoi(split[1])
	if err != nil {
		log.Fatal(err)
	}
	instructions[0] = parse
	parse, err = strconv.Atoi(split[3])
	if err != nil {
		log.Fatal(err)
	}
	instructions[1] = parse
	parse, err = strconv.Atoi(split[5])
	if err != nil {
		log.Fatal(err)
	}
	instructions[2] = parse
	return instructions
}

func prepdata(data string, row int, grid *[9][80]string) {
	for i, column := 1, 0; i < len(data); i++ {
		if data[i-1] == '[' {
			(*grid)[column][row] = string(data[i])
		}
		if i%4 == 0 {
			column++
		}
	}
}
