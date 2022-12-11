package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	readFile, err := os.Open("input")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		fmt.Println(line)
		for i := 0; i < len(line); i++ {
			if i >= 13 {
				piece := line[i-13 : i+1]
				fmt.Println(piece)
				if isThePiece(piece) {
					fmt.Printf("This ^^^ at %d\n", i+1)
					break
				}
			}
		}
	}

	readFile.Close()

}

func countChar(piece string, c string) int {
	result := 0
	for _, v := range piece {
		if string(v) == c {
			result += 1
		}
	}
	return result
}

func isThePiece(piece string) bool {
	for _, c := range piece {
		if countChar(piece, string(c)) > 1 {
			return false
		}
	}
	return true
}
