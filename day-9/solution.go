package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	// "time"
)

var input0 = [][]int{
	{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
	{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
	{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
	{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
	{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
}

func main() {
	fmt.Println(part2(parseInput()))
}

func parseInput() [][]int {
	data, _ := os.ReadFile("./input.txt")
	rows := make([][]int, 0)

	for _, r := range strings.Split(string(data), "\n") {
		row := make([]int, 0)
		for _, char := range []rune(r) {
			n, _ := strconv.ParseInt(string(char), 10, 64)
			row = append(row, int(n))
		}
		rows = append(rows, row)
	}

	return rows
}

func part1(input [][]int) int {
	acc := 0

	for r := 0; r < len(input); r++ {
		for c := 0; c < len(input[r]); c++ {
			if (r == 0 || input[r-1][c] > input[r][c]) &&
				(r == len(input)-1 || input[r+1][c] > input[r][c]) &&
				(c == 0 || input[r][c-1] > input[r][c]) &&
				(c == len(input[r])-1 || input[r][c+1] > input[r][c]) {
				acc += input[r][c] + 1
			}
		}
	}

	return acc
}

type point struct {
	r, c int
}

func part2(input [][]int) int {
	areas := make([]int, 0)
	visited := make([]point, 0)

	for r := 0; r < len(input); r++ {
		for c := 0; c < len(input[r]); c++ {
			if input[r][c] == 9 || contains(visited, point{r, c}) {
				continue
			}

			area := 0
			queue := make([]point, 0)
			queue = append(queue, point{r, c})

			fmt.Println(r)

			for len(queue) != 0 {
				p := queue[0]
				queue = queue[1:]

				if p.r < 0 || p.c < 0 || p.r >= len(input) || p.c >= len(input[r]) || input[p.r][p.c] == 9 || contains(visited, p) {
					continue
				}

				area++
				visited = append(visited, point{p.r, p.c})

				queue = append(queue, point{p.r - 1, p.c})
				queue = append(queue, point{p.r + 1, p.c})
				queue = append(queue, point{p.r, p.c - 1})
				queue = append(queue, point{p.r, p.c + 1})
			}

			areas = append(areas, area)
		}
	}

	sort.Ints(areas)

	fmt.Println(areas)

	return areas[len(areas)-1] * areas[len(areas)-2] * areas[len(areas)-3]
}

func contains(s []point, e point) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
