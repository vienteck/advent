package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	trees := []string{}
	for scanner.Scan() {
		trees = append(trees, scanner.Text())
	}
	counter := 0

	for i := 0; i < len(trees); i++ {
		line := trees[i]
		for j := 0; j < len(line); j++ {
			var height int
			height, err = strconv.Atoi(string(line[j]))
			if i == 0 || i == len(trees)-1 || j == 0 || j == len(line)-1 {
				counter += 1
				continue
			}

			l, err := strconv.Atoi(string(line[j-1]))
			if err != nil {
				log.Fatal(err)
			}

			r, err := strconv.Atoi(string(line[j+1]))
			if err != nil {
				log.Fatal(err)
			}

			u, err := strconv.Atoi(string(trees[i-1][j]))
			if err != nil {
				log.Fatal(err)
			}

			d, err := strconv.Atoi(string(trees[i+1][j]))
			if err != nil {
				log.Fatal(err)
			}
			if l < height && r < height && d < height && u < height {
				counter++
			}
		}
	}
	fmt.Print(counter)
}
