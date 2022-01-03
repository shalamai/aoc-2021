package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := parseInput()
	// filteredLines := filterHorisontalOrVertical(lines)

	fmt.Println(len(lines))
	field := makeField()
	applyLinesOnField(lines, field)
	res := calculateField(field)
	fmt.Println(res)
}

func makeField() [][]int {
	field := make([][]int, 1000)
	for i := range field {
		field[i] = make([]int, 1000)
	}

	return field
}

func calculateField(field [][]int) int {
	sum := 0

	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if field[i][j] >= 2 {
				sum++
			}
		}
	}

	return sum
}

func applyLinesOnField(lines []line, field [][]int) {
	for _, line := range lines {
		if line.start.x == line.end.x {
			lower, greater := sortByY(line.start, line.end)
			for i := lower.y; i <= greater.y; i++ {
				field[i][line.start.x]++
			}
		} else if line.start.y == line.end.y {
			lower, greater := sortByX(line.start, line.end)
			for i := lower.x; i <= greater.x; i++ {
				field[line.start.y][i]++
			}
		} else if isEncreasing(line) {
			lower, greater := sortByX(line.start, line.end)
			j := lower.y
			for i := lower.x; i <= greater.x; i++ {
				field[j][i]++
				j++
			}
		} else {
			lower, greater := sortByX(line.start, line.end)

			j := lower.y
			for i := lower.x; i <= greater.x; i++ {
				field[j][i]++
				j--
			}
		}
	}
}

func filterHorisontalOrVertical(lines []line) (filtered []line) {
	for _, line := range lines {
		if line.start.x == line.end.x || line.start.y == line.end.y {
			filtered = append(filtered, line)
		}
	}

	return
}

func isEncreasing(l line) bool {
	var lower, greater point
	if l.start.x < l.end.x {
		lower = l.start
		greater = l.end
	} else {
		lower = l.end
		greater = l.start
	}

	return lower.y < greater.y
}

func sortByX(p1 point, p2 point) (point, point) {
	if p1.x < p2.x {
		return p1, p2
	} else {
		return p2, p1
	}
}

func sortByY(p1 point, p2 point) (point, point) {
	if p1.y < p2.y {
		return p1, p2
	} else {
		return p2, p1
	}
}

func parseInput() (lines []line) {
	data, _ := os.ReadFile("./input.txt")
	for _, row := range strings.Split(string(data), "\n") {
		points := strings.Split(row, " -> ")
		lines = append(lines, line{start: parsePoint(points[0]), end: parsePoint(points[1])})
	}

	return
}

func parsePoint(data string) point {
	coords := strings.Split(data, ",")
	x, _ := strconv.ParseInt(coords[0], 10, 64)
	y, _ := strconv.ParseInt(coords[1], 10, 64)
	return point{x: int(x), y: int(y)}
}

type line struct {
	start, end point
}

type point struct {
	x, y int
}
