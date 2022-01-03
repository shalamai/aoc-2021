package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"regexp"
)

func main() {
	seq, boards := parseInput()
	// best := findBestBoard(seq, boards)
	// fmt.Println(best)

	worst := findWorstBoard(seq, boards)
	fmt.Println(worst)
	
}

func findWorstBoard(seq []string, boards [][][]string) int { 
	var winBoards []int 
	for _, num := range seq {
		markNumber(num, boards)
		
		for i, board := range boards {
			if !contains(winBoards, i) && isPeremoha(board) {
				winBoards = append(winBoards, i)
				if len(winBoards) == len(boards) {
					n, _ := strconv.ParseInt(num, 10, 64)
					return sumOfUnmarkedNums(board) * int(n)
				}
			} 
			
		}
	}

	return -1
}

func findBestBoard(seq []string, boards [][][]string) int {
	for _, num := range seq {
		markNumber(num, boards)
		for _, board := range boards {
			if isPeremoha(board) {
				n, _ := strconv.ParseInt(num, 10, 64)
				sum := sumOfUnmarkedNums(board)
				return sum * int(n)
			}
		}
	}

	return -1
}

func contains(ns []int, n int) bool {
	for _, v := range ns {
		if v == n {
			return true
		}
	}

	return false
}

func sumOfUnmarkedNums(board [][]string) int {
	sum := 0
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			if board[row][col] != "-1" {
				num, _ := strconv.ParseInt(board[row][col], 10, 64)
				sum += int(num)
			}
		}
	}

	return sum
}

func markNumber(num string, boards [][][]string) {
	for _, board := range boards {
		for row := 0; row < 5; row++ {
			for col := 0; col < 5; col++ {
				if board[row][col] == num {
					board[row][col] = "-1"
				}
			}
		}
	}
}

func isPeremoha(board [][]string) bool {
	for i := 0; i < 5; i++ {
		if isRowDone(board, i) || isColumnDone(board, i) {
			return true
		}
	}

	return false
}

func isRowDone(board [][]string, rowIndex int) bool {
	for i := 0; i < 5; i++ {
		if board[rowIndex][i] != "-1" {
			return false
		}
	}

	return true
}

func isColumnDone(board [][]string, colIndex int) bool {
	for i := 0; i < 5; i++ {
		if board[i][colIndex] != "-1" {
			return false
		}
	}

	return true
}

func parseInput() (seq []string, boards [][][]string) {
	data, _ := os.ReadFile("./input.txt")
	arr := strings.Split(string(data), "\n\n")

	seq = strings.Split(arr[0], ",")

	for _, rawBoard := range arr[1:] {
		var board [][]string
		for _, row := range strings.Split(rawBoard, "\n") {
			board = append(board, regexp.MustCompile("[ ]+").Split(row, -1))
		}
		boards = append(boards, board)
	}
	return
}
