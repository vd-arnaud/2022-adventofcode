package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {

	col1 := []string{"Z", "J", "N", "W", "P", "S"}
	col2 := []string{"G", "S", "T"}
	col3 := []string{"V", "Q", "R", "L", "H"}
	col4 := []string{"V", "S", "T", "D"}
	col5 := []string{"Q", "Z", "T", "D", "B", "M", "J"}
	col6 := []string{"M", "W", "T", "J", "D", "C", "Z", "L"}
	col7 := []string{"L", "P", "M", "W", "G", "T", "J"}
	col8 := []string{"N", "G", "M", "T", "B", "F", "Q", "H"}
	col9 := []string{"R", "D", "G", "C", "P", "B", "Q", "W"}

	cols := [][]string{col1, col2, col3, col4, col5, col6, col7, col8, col9}

	readFile, err := os.Open("input2")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		fmt.Println(line)
		r := regexp.MustCompile(`move\s(?P<count>\d+)\sfrom\s(?P<from>\d+)\sto\s(?P<from>\d+)`)
		fmt.Printf("%#v\n", r.FindStringSubmatch(line))
		matches := r.FindStringSubmatch(line)
		count, _ := strconv.Atoi(matches[1])
		from, _ := strconv.Atoi(matches[2])
		to, _ := strconv.Atoi(matches[3])
		fmt.Println(count, from, to)
		moveCountFromTo(count, &(cols[from-1]), &(cols[to-1]))
	}

	readFile.Close()
	fmt.Println(cols)
}

func moveCountFromTo(count int, col1, col2 *[]string) {
	for i := 0; i < count; i++ {
		elementToMove := (*col1)[len(*col1)-1]
		removeLastElement(col1)
		*col2 = append(*col2, elementToMove)
	}
}

func removeLastElement(slice *[]string) {
	*slice = (*slice)[:len(*slice)-1]
}
