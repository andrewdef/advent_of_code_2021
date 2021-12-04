package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

type bingoboard struct {
	numbers [5][5]int
	isWinning bool
}

type bingo struct {
	extractedNumbers []int
	boards []bingoboard
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func MarkNumber(numberToMark int, board *bingoboard) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if board.numbers[i][j] == numberToMark {
				board.numbers[i][j] = -1
			}
		}
	}
}

func IsWinning(board *bingoboard) bool {
	var count int
	for i := 0; i < 5; i++ {
		count = 0
		for j := 0; j < 5; j++ {
			if board.numbers[i][j] == -1 {
				count++
			}
		}
		if count == 5 {
			return true
		}
	}
	
	for i := 0; i < 5; i++ {
		count = 0
		for j := 0; j < 5; j++ {
			if board.numbers[j][i] == -1 {
				count++
			}
		}
		if count == 5 {
			return true
		}
	}
	
	return false
}

func GetBoardScore(board *bingoboard) int {
	score := 0
	
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if board.numbers[i][j] != -1 {
				score += board.numbers[i][j]
			}
		}
	}
	
	return score
}

func ReadFile() bingo {
	result := bingo{}
	
	var extractedNumbers []int
	var boards []bingoboard
	
	var i,k int
	
	file, err := os.Open("input.txt")
	checkError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	scanner.Scan() 
	for _, numString := range strings.Split(scanner.Text(), ",") {
		num, err := strconv.Atoi(numString)
		checkError(err)
		
		extractedNumbers = append(extractedNumbers, num)
	}
	
	scanner.Scan()

	var board bingoboard
	board = bingoboard{}
	k = 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			boards = append(boards, board)
			board = bingoboard{}
			k = 0
		} else {
			i = 0
			for _, numString := range strings.Split(scanner.Text(), " ") {
				if numString == "" {
					continue
				}
				
				num, err := strconv.Atoi(numString)
				checkError(err)
				
				board.numbers[k][i] = num
				i++
			}
			k++
		}
	}
	boards = append(boards, board)

	checkError(scanner.Err())
	
	result.extractedNumbers = extractedNumbers
	result.boards = boards
	
	return result
}

func Puzzle1(input bingo) int {
	for _, num := range input.extractedNumbers {
		for i := 0; i < len(input.boards); i++ {
			board := &input.boards[i]
			
			MarkNumber(num, board)

			if IsWinning(board) {
				return GetBoardScore(board) * num
			}
		}
	}

	return -1
}

func Puzzle2(input bingo) int {
	notWonBoards := len(input.boards)
	
	for _, num := range input.extractedNumbers {
		for i := 0; i < len(input.boards); i++ {
			board := &input.boards[i]
			
			if board.isWinning {
				continue
			}
			
			MarkNumber(num, board)

			if IsWinning(board) {
				board.isWinning = true
				notWonBoards -= 1
				
				if notWonBoards == 0 {
					return GetBoardScore(board) * num
				}
			}
		}
	}

	return -1
}

func main() {
	inputData := ReadFile()
	fmt.Println("Answer 1 : ", Puzzle1(inputData))
	fmt.Println("Answer 2 : ", Puzzle2(inputData))
}
