package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("assignments.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	overlap := 0
	overlaptwo := 0
	for scanner.Scan() {
		log.Print(scanner.Text())
		split := strings.Split(scanner.Text(), ",")
		var a [2]int
		val, err := strconv.Atoi(strings.Split(split[0], "-")[0])
		if err != nil {
			log.Fatal(err)
		}
		a[0] = val
		val, err = strconv.Atoi(strings.Split(split[0], "-")[1])
		if err != nil {
			log.Fatal(err)
		}
		a[1] = val
		var b [2]int
		val, err = strconv.Atoi(strings.Split(split[1], "-")[0])
		if err != nil {
			log.Fatal(err)
		}
		b[0] = val
		val, err = strconv.Atoi(strings.Split(split[1], "-")[1])
		if err != nil {
			log.Fatal(err)
		}
		b[1] = val

		log.Print(a == b)
		if (a[0] >= b[0] && a[1] <= b[1]) || (b[0] >= a[0] && b[1] <= a[1]) {
			overlap++
		}
		if (a[0] >= b[0] && a[0] <= b[1]) || (a[1] >= b[0] && a[1] <= b[1]) {
			overlaptwo++
			continue
		}
		if (b[0] >= a[0] && b[0] <= a[1]) || (b[1] >= a[0] && b[1] <= a[1]) {
			overlaptwo++
			continue
		}
	}
	log.Print(overlaptwo)
}
