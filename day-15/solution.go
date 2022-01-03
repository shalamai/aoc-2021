package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	tile := parseInput()
	field := multiply5(tile)
	res := findWay3(field)
	fmt.Println(res)
}

func multiply5(seed [][]int) [][]int {
	for r := 0; r < len(seed); r++ {
		curr := seed[r]
		for i := 0; i < 4; i++ {
			curr = incrementArr(curr)
			seed[r] = append(seed[r], curr...)
		}
	}

	l := len(seed)
	for i := 0; i < l*4; i++ {
		seed = append(seed, make([]int, len(seed[0])))
		for c := 0; c < len(seed[0]); c++ {
			seed[l + i][c] = increment(seed[i][c])
		}
	}

	return seed
}

func increment(n int) int {
	if n < 9 {
		return n + 1
	} else {
		return 1
	}
}

func incrementArr(arr []int) []int {
	res := make([]int, len(arr))
	copy(res, arr)

	for i := 0; i < len(res); i++ {
		res[i] = increment(res[i])
	}

	return res
}

func findWay3(field [][]int) int {
	price := makeBlankPrice(len(field), len(field[0]))

	field[0][0] = 0
	price[0][0] = 0

	queue := make([]point, 0)
	queue = append(queue, point{0, 0})

	for true {
		if len(queue) == 0 {
			break
		}

		p := queue[0]
		queue = queue[1:]

		dirs := []point{
			{p.r - 1, p.c},
			{p.r + 1, p.c},
			{p.r, p.c - 1},
			{p.r, p.c + 1},
		}

		for _, d := range dirs {
			updated := updatePrice(field, price, d, price[p.r][p.c])
			if updated {
				queue = append(queue, d)
			}
		}
	}

	return price[len(field)-1][len(field[0])-1]
}

func updatePrice(field [][]int, price [][]int, p point, fromPrice int) bool {
	if p.r >= len(field) || p.r < 0 || p.c >= len(field[0]) || p.c < 0 {
		return false
	}

	potentialPrice := fromPrice + field[p.r][p.c]
	if price[p.r][p.c] == -1 || price[p.r][p.c] > potentialPrice {
		price[p.r][p.c] = potentialPrice
		return true
	}

	return false
}

func makeBlankPrice(r int, c int) [][]int {
	rows := make([][]int, 0)
	for i := 0; i < r; i++ {
		row := make([]int, 0)
		for j := 0; j < c; j++ {
			row = append(row, -1)
		}
		rows = append(rows, row)
	}

	return rows
}

type point struct {
	r, c int
}

var cache map[point]int = make(map[point]int)

func findWay2(field [][]int, current point, visited []point) (int, bool) {
	if val, ok := cache[current]; ok {
		return val, true
	}

	if contains(visited, current) || current.r < 0 || current.r >= len(field) || current.c < 0 || current.c >= len(field[0]) {
		return 0, false
	}

	if current.r == len(field)-1 && current.c == len(field[0])-1 {
		return field[current.r][current.c], true
	}

	results := make([]int, 0)
	dirs := []point{
		{current.r, current.c + 1},
		{current.r, current.c - 1},
		{current.r + 1, current.c},
		{current.r - 1, current.c},
	}

	visitedNew := visited
	visitedNew = append(visitedNew, current)
	for _, d := range dirs {
		res, ok := findWay2(field, d, visitedNew)
		if ok {
			cache[d] = res
			results = append(results, res)
		}
	}

	if len(results) == 0 {
		return 0, false
	}

	return minArr(results) + field[current.r][current.c], true
}

func findWay(field [][]int) int {
	for i := 2; i < len(field); i++ {
		field[i][0] += field[i-1][0]
	}

	for i := 2; i < len(field[0]); i++ {
		field[0][i] += field[0][i-1]
	}

	for r := 1; r < len(field); r++ {
		for c := 1; c < len(field[0]); c++ {
			field[r][c] += min(field[r-1][c], field[r][c-1])
		}
	}

	return field[len(field)-1][len(field[0])-1]
}

func contains(as []point, b point) bool {
	for _, a := range as {
		if a == b {
			return true
		}
	}

	return false
}

func minArr(as []int) int {
	min := as[0]
	for _, a := range as {
		if a < min {
			min = a
		}
	}

	return min
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func parseInput() [][]int {
	data, _ := os.ReadFile("./input.txt")

	arr := make([][]int, 0)
	for _, r := range strings.Split(string(data), "\n") {
		row := make([]int, 0)
		for _, c := range r {
			v, _ := strconv.ParseInt(string(c), 10, 64)
			row = append(row, int(v))
		}
		arr = append(arr, row)
	}

	return arr
}
