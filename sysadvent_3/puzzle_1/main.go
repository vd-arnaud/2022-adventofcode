package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	var priority = map[string]int{}
	var p1 int = 1
	var p2 int = 27
	for r := 'a'; r <= 'z'; r++ {
		R := unicode.ToUpper(r)
		priority[string(r)] = p1
		priority[string(R)] = p2
		p1++
		p2++
	}

	var totalPriority = 0

	readFile, err := os.Open("input")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		fmt.Println(line)
		len_line := len(line)
		comp1 := line[:len_line/2]
		comp2 := line[len_line/2:]
		fmt.Println(comp1)
		fmt.Println(comp2)
		var commonChar string
		for _, c1 := range comp1 {
			for _, c2 := range comp2 {
				if c1 == c2 {
					commonChar = string(c1)
					break
				}
			}
		}
		fmt.Println(commonChar)
		fmt.Println(priority[commonChar])
		totalPriority += priority[commonChar]
	}

	readFile.Close()

	fmt.Println(totalPriority)
	fmt.Println(priority)
}
