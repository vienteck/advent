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
	pwd = append(pwd, "/")
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
			fmt.Println(line)
			if strings.Contains(line, "cd") {
				if cdsize != 0 {
					// v := strings.Join(pwd, "/")
					val, pass := sizes[strings.Join(pwd, "/")]
					if pass {
						fmt.Printf("Value already exists in map %v\n", val)
					}

					for i := 1; i < len(pwd)+1; i++ {
						fmt.Printf("Adding %v to directory %v\n", cdsize, strings.Join(pwd[:i], "/"))
						sizes[strings.Join(pwd[:i], "/")] += cdsize
					}

					// sizes[strings.Join(pwd, "/")] = cdsize
					cdsize = 0
				}

				if strings.Contains(line, "..") {
					pwd = pwd[:len(pwd)-1]
				} else {
					if sep[2] == "/" {
						pwd = []string{}
						pwd = append([]string{}, "")
					} else {
						pwd = append(pwd, sep[2])
					}
				}
				fmt.Println("dir - /" + strings.Join(pwd, "/"))
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

	solution := 0
	target := 30000000 - (70000000 - sizes[""])
	fmt.Println(target)
	difference := -99999999999
	var currcandidate string
	for key, val := range sizes {
		if 0 > (target-val) && (target-val) > difference {
			difference = target - val
			currcandidate = key
		}
		// fmt.Printf("%v %v\n", key, val)
		if val < 100000 {
			solution += val
		}
	}
	fmt.Print(difference)
	fmt.Print(currcandidate)
	fmt.Print(solution)
}
