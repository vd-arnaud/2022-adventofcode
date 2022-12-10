package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseLine(line string) ([]int, []int) {
	s := strings.Split(line, ",")
	range1 := strings.Split(s[0], "-")
	range2 := strings.Split(s[1], "-")
	r1e1, _ := strconv.Atoi(range1[0])
	r1e2, _ := strconv.Atoi(range1[1])
	r2e1, _ := strconv.Atoi(range2[0])
	r2e2, _ := strconv.Atoi(range2[1])

	var range1int []int = []int{r1e1, r1e2}
	var range2int []int = []int{r2e1, r2e2}
	return range1int, range2int

}

func isOneRangeOverlapingTheOther(range1 []int, range2 []int) bool {
	// 1-0    2-0   1-1             2-1
	// |--------[----|---------------]
	if range2[0] >= range1[0] && range1[1] >= range2[0] {
		return true
	}
	if range1[0] >= range2[0] && range2[1] >= range1[0] {
		return true
	}
	return false
}

func isOneRangeEatingTheOther(range1 []int, range2 []int) bool {
	//how many element in range1
	r1l := range1[1] - range1[0] + 1
	r2l := range2[1] - range2[0] + 1
	//fmt.Println(r1l)
	//fmt.Println(r2l)
	//Cas particulier les ranges font la meme longueur
	if r1l == r2l {
		if range1[0] == range2[0] && range1[1] == range2[1] {
			return true
		}
		return false
	}

	//Si le range1 est plus grand que le range2
	// [--------[----]---------------]
	if r1l > r2l {
		if range1[0] <= range2[0] && range1[1] >= range2[1] {
			return true
		}
		return false
	}

	//Si le range2 est plus grand que le range1
	// [--------[----]---------------]
	if r2l > r1l {
		if range2[0] <= range1[0] && range2[1] >= range1[1] {
			return true
		}
		return false
	}
	return false
}

func main() {
	readFile, err := os.Open("input")

	var total int = 0

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		range1, range2 := parseLine(line)

		if isOneRangeOverlapingTheOther(range1, range2) {
			total += 1
		}
	}

	readFile.Close()
	fmt.Println(total)
}
