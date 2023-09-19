package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("datastream.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	datastream := scanner.Text()
	datatracker := map[string]bool{}
	for i := 0; i < len(datastream); i++ {
		_, ok := datatracker[string(datastream[i])]
		if ok {
			//If value already exists in our map that means we have a duplicate and our window needs to move
		}

	}
	log.Print(datastream)
}
