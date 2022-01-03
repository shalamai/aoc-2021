package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	field, folds := parseInput()
	// doFold(field, folds[0])

	// fmt.Println(len(field))
	// fmt.Println(field)
	// // res := calcUnique(field)
	// fmt.Println(res)

	for _, f := range folds {
		doFold(field, f)
	}

	printField(field)
}

func printField(field []coord) {
	arr := make([][]int, getMaxY(field))

	maxX := getMaxX(field)
	for i := range arr {
		arr[i] = make([]int, maxX)
	}

	for _, c := range field {
		arr[c.y][c.x] = 1
	}

	for _, r := range arr {
		for _, l := range r {
			if l == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}

func getMaxX(field []coord) int {
	max := 0
	for _, c := range field {
		if c.x > max {
			max = c.x
		}
	}

	return max + 1
}

func getMaxY(field []coord) int {
	max := 0
	for _, c := range field {
		if c.y > max {
			max = c.y
		}
	}

	return max + 1
}

func calcUnique(field []coord) int {
	set := make(map[coord]bool)
	for _, c := range field {
		if !set[c] {
			set[c] = true
		}
	}

	return len(set)
}

func doFold(field []coord, f fold) {
	if f.axis == "x" {
		for i, c := range field {
			if c.x > f.coord {
				c.x = f.coord - (c.x - f.coord) 
			}
			field[i] = c
		}
	} else {
		for i, c := range field {
			if c.y > f.coord {
				c.y = f.coord - (c.y - f.coord) 
			}
			field[i] = c
		}
	}

	fmt.Println(field)
}

func parseInput() (field []coord, folds []fold) {
	data, _ := os.ReadFile("./input2.txt")
	res0 := strings.Split(string(data), "\n\n")

	for _, c := range strings.Split(res0[0], "\n") {
		xy := strings.Split(c, ",")
		x, _ := strconv.ParseInt(xy[0], 10, 64)
		y, _ := strconv.ParseInt(xy[1], 10, 64)

		field = append(field, coord{int(x), int(y)})
	}

	for _, f := range strings.Split(res0[1], "\n") {
		a := strings.Split(f, " ")
		f := strings.Split(a[2], "=")
		c, _ := strconv.ParseInt(f[1], 10, 64)

		folds = append(folds, fold{f[0], int(c)})
	}

	return
}

type fold struct {
	axis  string
	coord int
}

type coord struct {
	x, y int
}
