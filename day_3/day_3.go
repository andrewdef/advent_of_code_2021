package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile() []string {
	var result []string
	
	file, err := os.Open("input.txt")
	checkError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	checkError(scanner.Err())
	
	return result
}

func FindMatchingNumbers(data []string, posToSearch int, criteria string) []string {
	counter := 0
	var results []string
	var toSearch byte
	
	for _, num := range data {
		r := num[posToSearch]

		if ( r == '0' ) {
			counter -= 1
		} else {
			counter += 1
		}
	}
	
	if ( criteria == "mostCommon" ) {
		if ( counter < 0) {
			toSearch = '0'
		} else {
			toSearch = '1'
		}
	} else {
		if ( counter < 0) {
			toSearch = '1'
		} else {
			toSearch = '0'
		}
	}
	
	for _, num := range data {
		r := num[posToSearch]

		if ( r == toSearch ) {
			results = append(results, num)
		}
	}
	
	return results
}

func Puzzle1(data []string) int64 {
	var counters []int
	var epsilon []string
	var gamma []string
	
	for _, num := range data {
		for i, r := range num {
			if ( len(counters) < (i + 1) ) {
				counters = append(counters, 0)
			}

			if ( r == '0' ) {
				counters[i] -= 1
			} else {
				counters[i] += 1
			}
		}
	}
	
	for _, count := range counters {
		if ( count < 0 ) {
			gamma = append(gamma, "0")
			epsilon = append(epsilon, "1")
		} else {
			gamma = append(gamma, "1")
			epsilon = append(epsilon, "0")
		}
	}
	
	g, err := strconv.ParseInt(strings.Join(gamma, ""), 2, 64)
	checkError(err)
	
	e, err := strconv.ParseInt(strings.Join(epsilon, ""), 2, 64)
	checkError(err)	
	
	return g * e
}

func Puzzle2(data []string) int64 {
	oxygen_gen_rating := ""
	co2_scrubber_rating := ""
	var pos int
	var inputData []string
	
	pos = 0
	inputData = data
	for oxygen_gen_rating == "" {
		inputData = FindMatchingNumbers(inputData, pos, "mostCommon")
		
		if len(inputData) == 1 {
			oxygen_gen_rating = inputData[0]
		}
		
		pos++
	}
	
	pos = 0
	inputData = data
	for co2_scrubber_rating == "" {
		inputData = FindMatchingNumbers(inputData, pos, "leastCommon")
		
		if len(inputData) == 1 {
			co2_scrubber_rating = inputData[0]
		}
		
		pos++
	}
	
	ogr, err := strconv.ParseInt(oxygen_gen_rating, 2, 64)
	checkError(err)
	
	co2sr, err := strconv.ParseInt(co2_scrubber_rating, 2, 64)
	checkError(err)	
	
	return ogr * co2sr
}

func main() {
	inputData := ReadFile()
	fmt.Println("Answer 1 : ", Puzzle1(inputData))
	fmt.Println("Answer 2 : ", Puzzle2(inputData))
}
