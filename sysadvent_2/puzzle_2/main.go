package main

import (
	"bufio"
	"fmt"
	"os"
)

// A adv plays rock
// B adv plays paper
// C adv plays scissors

// X = i need loose = 0 point
// Y = i need draw = 3
// Z = i need win = 6

var score int = 0

const LOOSE = 0
const DRAW = 3
const WIN = 6

var points = map[string]int{
	"rock":     1,
	"paper":    2,
	"scissors": 3,
}

var code_tools = map[int]string{
	1: "rock",
	2: "paper",
	3: "scissors",
}

var name_tools = map[string]string{
	"rock":     "A",
	"paper":    "B",
	"scissors": "C",
}

func getScore(adv string, result_needed string) int {
	//case all 9 possibilities
	switch result_needed {
	case "X":
		switch adv {
		case name_tools["rock"]:
			return points["scissors"] + LOOSE
		case name_tools["paper"]:
			return points["rock"] + LOOSE
		case name_tools["scissors"]:
			return points["paper"] + LOOSE
		}
	case "Y":
		switch adv {
		case name_tools["rock"]:
			return points["rock"] + DRAW
		case name_tools["paper"]:
			return points["paper"] + DRAW
		case name_tools["scissors"]:
			return points["scissors"] + DRAW
		}
	case "Z":
		switch adv {
		case name_tools["rock"]:
			return points["paper"] + WIN
		case name_tools["paper"]:
			return points["scissors"] + WIN
		case name_tools["scissors"]:
			return points["rock"] + WIN
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
