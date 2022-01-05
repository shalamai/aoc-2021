package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	step := 50

	res1 := part1(step)
	fmt.Println(res1)
}

func part1(step int) int {
	field, encoding := parseInput()

	res := expand(field, step)
	var background byte = 0
	for i := 0; i < step; i++ {
		res, background = doStep(res, encoding, background)
	}

	return calcLights(res)
}

func parseInput() ([][]byte, []byte) {
	data, _ := os.ReadFile("./input.txt")
	parts := strings.Split(string(data), "\n\n")

	encoding := toBinary(parts[0])

	rawField := strings.Split(parts[1], "\n")
	field := make([][]byte, len(rawField))
	for i, r := range rawField {
		field[i] = toBinary(r)
	}

	return field, encoding
}

func toBinary(input string) []byte {
	res := make([]byte, len(input))
	for i, ch := range input {
		if ch == '#' {
			res[i] = 1
		} else {
			res[i] = 0
		}
	}

	return res
}

func expand(field [][]byte) [][]byte {
	extra := 1
	res := make([][]byte, len(field)+extra*2)
	for i := 0; i < len(res); i++ {
		res[i] = make([]byte, len(field[0])+extra*2+2)
	}

	for r := 0; r < len(field); r++ {
		for c := 0; c < len(field[0]); c++ {
			res[r+extra][c+extra] = field[r][c]
		}
	}

	return res
}

func doStep(field [][]byte, encoding []byte, background byte) (field2 [][]byte, background2 byte) {
	if background == 0 {
		background2 = encoding[0]
	} else {
		background2 = encoding[511]
	}

	field2 = expand(field)
	for r2 := 0; r2 < len(field2); r2++ {
		for c2 := 0; c2 < len(field2[0]); c2++ {
			r := r2 - 1 // because new row is expaned
			c := c2 - 1 // because new column is expaned
			binaryIndex :=
				cell(field, r-1, c-1, background) +
					cell(field, r-1, c, background) +
					cell(field, r-1, c+1, background) +
					cell(field, r, c-1, background) +
					cell(field, r, c, background) +
					cell(field, r, c+1, background) +
					cell(field, r+1, c-1, background) +
					cell(field, r+1, c, background) +
					cell(field, r+1, c+1, background)

			index, _ := strconv.ParseInt(binaryIndex, 2, 64)
			field2[r2][c2] = encoding[index]
		}
	}

	return
}

func cell(field [][]byte, r int, c int, background byte) string {
	if field[r][c] == 2 {
		return strconv.Itoa(int(background))
	} else {
		return strconv.Itoa(int(field[r][c]))
	}
}

func calcLights(field [][]byte) int {
	acc := 0
	for r := 0; r < len(field); r++ {
		for c := 0; c < len(field[0]); c++ {
			if field[r][c] == 1 {
				acc++
			}
		}
	}

	return acc
}
