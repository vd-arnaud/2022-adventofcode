package main

import (
	"bufio"
	"fmt"
	"os"
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

	var elf int = 0
	var highestElf int = 0
	var calories int = 0
	var highestCalories = 0

	for fileScanner.Scan() {
		line = fileScanner.Text()
		if line == "" {
			fmt.Println("vide")
			if calories > highestCalories {
				highestElf = elf
				highestCalories = calories
			}
			calories = 0
			elf += 1
		} else {
			food, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println("Error while converting line to int")
				continue
			}
			calories += food
		}
		fmt.Println(line)
	}

	readFile.Close()

	fmt.Printf("Elf with highest calories is %d with %d calories total", highestElf, highestCalories)
}
