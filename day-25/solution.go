package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	res1 := part1()
	fmt.Println(res1)
}

func part1() int {
	field := parseInput()
	step := 0

	for doStep(field) {
		step++
		// print(field, step)
	}

	return step + 1
}

func parseInput() [][]string {
	acc := make([][]string, 0)
	data, _ := os.ReadFile("./input.txt")
	for _, r := range strings.Split(string(data), "\n") {
		row := make([]string, 0)
		for _, c := range r {
			row = append(row, string(c))
		}
		acc = append(acc, row)
	}
	return acc
}

func doStep(field [][]string) bool {
	movedEast := moveEast(field)
	cleanTrace(field)
	movedSouth := moveSouth(field)
	cleanTrace(field)
	return movedEast > 0 || movedSouth > 0
}

func moveEast(field [][]string) int {
	moved := 0
	for r := 0; r < len(field); r++ {
		for c := 0; c < len(field[0]); c++ {
			c2 := next(c, len(field[0]))
			if field[r][c] == ">" && field[r][c2] == "." {
				moved++
				field[r][c] = "#"
				field[r][c2] = ">"
				c++
			}
		}
	}

	return moved
}

func moveSouth(field [][]string) int {
	moved := 0
	for c := 0; c < len(field[0]); c++ {
		for r := 0; r < len(field); r++ {
			r2 := next(r, len(field))
			if field[r][c] == "v" && field[r2][c] == "." {
				moved++
				field[r][c] = "#"
				field[r2][c] = "v"
				r++
			}
		}
	}

	return moved
}

func cleanTrace(field [][]string) {
	for r := 0; r < len(field); r++ {
		for c := 0; c < len(field[0]); c++ {
			if field[r][c] == "#" {
				field[r][c] = "."
			}
		}
	}
}

func next(current, limit int) int {
	if current + 1 >= limit {
		return 0
	} else {
		return current + 1
	}
}

func print(field [][]string, step int) {
	fmt.Printf("\n\n\n step %v\n", step)
	for _, r := range field {
		for _, c := range r {
			fmt.Print(c)
		}
		fmt.Println()
	}
}
