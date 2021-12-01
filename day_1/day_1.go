package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile() []int {
	var result []int
	
	file, err := os.Open("input.txt")
	checkError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		depth, err := strconv.Atoi(scanner.Text())
		checkError(err)

		result = append(result, depth)
	}

	checkError(scanner.Err())
	
	return result
}

func Puzzle1(data []int) int {
	count := 0
	prevDepth := -1
	
	for _, depth := range data {
		if ( depth > prevDepth) && ( prevDepth != -1) {
			count += 1
		}
		
		prevDepth = depth
	}
	
	return count
}

func Puzzle2(data []int) int {
	count := 0
	prevDepth := -1
	windowSize := 3
	var windows []int
	
	for _, depth := range data {
		windows = append(windows, 0)
		
		for k := 0; k < windowSize; k++ {
			index := len(windows) - k - 1
			
			if ( index >= 0) {
				windows[index] += depth
			}
		}
	}
	
	for _, depth := range windows {
		if ( depth > prevDepth) && ( prevDepth != -1) {
			count += 1
		}
		
		prevDepth = depth
	}
	
	return count
}

func main() {
	inputData := ReadFile()
	fmt.Println("Answer 1 : ", Puzzle1(inputData))
	fmt.Println("Answer 2 : ", Puzzle2(inputData))
}
