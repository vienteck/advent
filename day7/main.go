package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("data.txt")

	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	var pwd []string
	var curdir string
	cdsize := 0
	sizes := make(map[string]int)
	sizes["/"] = 0
	for scanner.Scan() {
		line := scanner.Text()
		sep := strings.Split(line, " ")
		if line == "" {
			continue
		}
		//these are commands
		if string(line[0]) == "$" {
			if strings.Contains(line, "cd") {
				if cdsize != 0 {
					val, pass := sizes[strings.Join(pwd, "/")]
					if pass {
						fmt.Printf("Value already exists in map %v", val)
					}
					sizes[strings.Join(pwd, "/")] = cdsize
					cdsize = 0
				}

				if strings.Contains(line, "..") {
					pwd = pwd[:len(pwd)-1]
				} else {
					if sep[2] == "/" {
						pwd = []string{}
					} else {
						pwd = append(pwd, sep[2])
					}
				}
				fmt.Println(pwd)
				curdir = sep[2]
				log.Print(curdir)
			}

		} else {
			if sep[0] != "dir" {
				val, err := strconv.Atoi(sep[0])
				if err != nil {
					log.Fatal(err)
				}
				cdsize += val
			}
			fmt.Println("     " + line)
		}

	}
	if cdsize != 0 {
		sizes[strings.Join(pwd, "/")] = cdsize
		cdsize = 0
	}

	tracker := make(map[string]int)
	asdf := 0
	for key, val := range sizes {
		d := strings.Split(key, "/")
		fmt.Printf("%v (%v)\n", key, val)
		for i := 0; i < len(d); i++ {
			asdf++
			log.Print(asdf)
			fmt.Println(strings.Join(d[:i], "/"))
			tracker[strings.Join(d[:i], "/")] += val
		}
		// fmt.Println(strings.Join(d, "/"))
		// tracker[strings.Join(d, "/")] += val

	}
	solution := 0
	for key, val := range tracker {
		fmt.Printf("%v %v\n", key, val)
		if val < 100000 {
			solution += val
		}
	}
	fmt.Print(solution)
}
