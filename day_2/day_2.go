package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

type command struct {
	direction string
	amount int
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile() []command {
	var result []command
	
	file, err := os.Open("input.txt")
	checkError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tokens := strings.Fields(scanner.Text())
		amount, err := strconv.Atoi(tokens[1])
		checkError(err)

		result = append(result, command{direction: tokens[0], amount:amount})
	}

	checkError(scanner.Err())
	
	return result
}

func Puzzle1(data []command) int {
	pos := 0
	depth := 0
	
	for _, cmd := range data {
		if ( cmd.direction == "up" ) {
			depth -= cmd.amount
		} else if ( cmd.direction == "down" ) {
			depth += cmd.amount
		} else if ( cmd.direction == "forward" ) {
			pos += cmd.amount
		}
	}
	
	return pos * depth
}

func Puzzle2(data []command) int {
	pos := 0
	depth := 0
	aim := 0
	
	for _, cmd := range data {
		if ( cmd.direction == "up" ) {
			aim -= cmd.amount
		} else if ( cmd.direction == "down" ) {
			aim += cmd.amount
		} else if ( cmd.direction == "forward" ) {
			pos += cmd.amount
			depth += aim * cmd.amount
		}
	}
	
	return pos * depth
}

func main() {
	inputData := ReadFile()
	fmt.Println("Answer 1 : ", Puzzle1(inputData))
	fmt.Println("Answer 2 : ", Puzzle2(inputData))
}
