package main

import (
	"bufio"
	"fmt"
	"os"
)

// A ou X rock (1 point)
// B ou Y  paper (2 points)
// C ou Z scissors (3 points)

// loose = 0 point
// draw = 3
// win = 6

var score int = 0

const LOOSE = 0
const DRAW = 3
const WIN = 6

var points = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

func getScore(adv string, me string) int {
	//case all 9 possibilities
	switch me {
	case "X":
		switch adv {
		case "A":
			return points[me] + DRAW
		case "B":
			return points[me] + LOOSE
		case "C":
			return points[me] + WIN
		}
	case "Y":
		switch adv {
		case "A":
			return points[me] + WIN
		case "B":
			return points[me] + DRAW
		case "C":
			return points[me] + LOOSE
		}
	case "Z":
		switch adv {
		case "A":
			return points[me] + LOOSE
		case "B":
			return points[me] + WIN
		case "C":
			return points[me] + DRAW
		}
	}

	//add what he played
	return 0
}

func main() {
	readFile, err := os.Open("input")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		//fmt.Printf("var1 = %T\n", line[0])
		score += getScore(string(line[0]), string(line[2]))
	}

	readFile.Close()
	fmt.Printf("Total score is %d", score)
}
