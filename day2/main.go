package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type legend map[string]map[string]int

func main() {
	//A for Rock, B for Paper, and C for Scissors
	//X for Rock, Y for Paper, and Z for Scissors
	file, err := os.Open("strategy.txt")
	l := legend{
		"X": {"A": 4, "B": 1, "C": 7},
		"Y": {"A": 8, "B": 5, "C": 2},
		"Z": {"A": 3, "B": 9, "C": 6},
	}
	//X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win.
	p := legend{
		"A": {"Y": 4, "X": 3, "Z": 8},
		"B": {"Z": 9, "Y": 5, "X": 1},
		"C": {"X": 2, "Z": 7, "Y": 6},
	}

	if err != nil {

		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	currentscore := 0
	currentscoretwo := 0
	for scanner.Scan() {
		signs := strings.Split(scanner.Text(), " ")
		currentscore += l[signs[1]][signs[0]]
		currentscoretwo += p[signs[0]][signs[1]]
	}
	log.Print(currentscore)
	log.Print(currentscoretwo)

}
