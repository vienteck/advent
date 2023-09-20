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
	pt := map[string]int{string(ds[0]): 0}
	for i := 0; i < len(ds)-1; i++ {
		//prep data
		if i < 3 {
			dt[string(ds[i])] = i
			continue
		}
		_, ok := dt[string(ds[i])]
		if ok {
			log.Printf("duplicated key: %v", string(ds[i]))
			p := pt[string(ds[i])] + 1
			delete(dt, string(ds[i]))
			pt = map[string]int{}
			dt[string(ds[i])] = i
			pt[string(ds[p])] = p + 1
			//If value already exists in our map that means we have a duplicate and our window needs to move
		} else {
			dt[string(ds[i])] = i
			if len(dt) == 4 {
				log.Printf("Marker found at position %v", i)
				break
			}
		}

	}
}
