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
	lastitem := [9]int{}
	line := 1
	row := 7
	var instructions [][3]int
	for scanner.Scan() {
		if line < 9 {
			prepdata(scanner.Text(), row, &grid, &lastitem)
			row--
		}
		if line > 10 {

			instructions = append(instructions, getinstructions(scanner.Text()))
		}
		line++
	}
	//we have our base grid and out instructions now
	for i := 0; i < len(instructions); i++ {
		processStep(instructions[i], &grid, &lastitem)
	}
	results := []string{}
	for i := 0; i < 9; i++ {
		results = append(results, string(grid[i][lastitem[i]-1]))
	}
	log.Print(results)
}
func processStep(instructions [3]int, grid *[9][80]string, lastemptycell *[9]int) {
	//[0] amount to move from column  [1] to columns [2]
	//remove cells fom grid and put into a temporary slice
	movefrom := instructions[1] - 1
	moveto := instructions[2] - 1
	temp := make([]string, instructions[0])
	copy(temp, (*grid)[movefrom][lastemptycell[movefrom]-instructions[0]:lastemptycell[movefrom]])
	for i := lastemptycell[movefrom] - instructions[0]; i < lastemptycell[movefrom]; i++ {
		(*grid)[movefrom][i] = ""
	}
	log.Print(temp)
	//update lastemptycell array
	lastemptycell[movefrom] = lastemptycell[movefrom] - instructions[0]
	//move cells to requested column
	for i := len(temp) - 1; i >= 0; i-- {
		(*grid)[moveto][lastemptycell[moveto]] = temp[i]
		lastemptycell[moveto]++
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

func prepdata(data string, row int, grid *[9][80]string, lastitem *[9]int) {

	for i, column := 1, 0; i < len(data); i++ {

		if data[i-1] == '[' {
			(*grid)[column][row] = string(data[i])
			if lastitem[column] == 0 {
				lastitem[column] = row + 1
			}
		}
		if i%4 == 0 {
			column++
		}
	}
}
