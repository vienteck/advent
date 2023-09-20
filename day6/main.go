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
	ds := scanner.Text()
	dt := map[string]int{}
	pt := 1
	for i := 0; i < len(ds); i++ {
		log.Printf("Current Pointer: %v", pt)
		log.Printf("Current character: %v", string(ds[i]))
		//prep data
		if i < 3 {
			dt[string(ds[i])] = 1
			continue
		}
		_, ok := dt[string(ds[i])]
		if len(dt) > 3 && !ok {
			dt[string(ds[i])] = 1
			log.Print(dt)
			break
			//this should meant that we have a unique 4 characters in a row
		}
		if ok {
			//If value already exists in our map that means we have a duplicate and our window needs to move
			dt = map[string]int{}
			pt = i + 2
		} else {
			dt[string(ds[i])] = 1
		}

	}
	log.Print(ds)
}
