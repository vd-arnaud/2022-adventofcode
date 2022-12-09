package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func findCommonChar(lines []string) string {
	var compteurs map[string]int
	compteurs = make(map[string]int)
	for _, line := range lines {
		alreadyFound := []string{}
		for _, char := range line {
			if !contains(alreadyFound, string(char)) {
				alreadyFound = append(alreadyFound, string(char))
				compteurs[string(char)]++
				if compteurs[string(char)] == 3 {
					return string(char)
				}
			}
		}
	}
	return ""
}

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
	var i int = 0
	var lines []string
	for fileScanner.Scan() {
		fmt.Println(i)
		line := fileScanner.Text()
		fmt.Println(line)
		lines = append(lines, line)
		if (i+1)%3 == 0 {
			fmt.Printf("Stop at index %d", i)
			fmt.Println(lines)
			commonChar := findCommonChar(lines)
			fmt.Printf("Common char : %s\n", commonChar)
			totalPriority += priority[commonChar]
			//reset
			lines = []string{}
		}
		i++
	}

	readFile.Close()

	fmt.Println(totalPriority)
}
