package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {

	readFile, err := os.Open("input")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var line string

	var calories int = 0
	var elfs []int = []int{}

	for fileScanner.Scan() {
		line = fileScanner.Text()
		if line == "" {
			// empty line, elf switch
			elfs = append(elfs, calories)
			calories = 0
		} else {
			food, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println("Error while converting line to int")
				continue
			}
			calories += food
		}
	}
	//last line
	elfs = append(elfs, calories)

	readFile.Close()

	fmt.Println(len(elfs))
	sort.Sort(sort.Reverse(sort.IntSlice(elfs)))
	totalThreeBest := 0
	for _, calories := range elfs[0:3] {
		fmt.Println(calories)
		totalThreeBest += calories
	}
	fmt.Println(totalThreeBest)
}
