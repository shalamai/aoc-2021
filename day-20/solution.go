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

func expand(field [][]byte, extra int) [][]byte {
	res := make([][]byte, len(field)+extra*2+2)
	for i := 0; i < len(res); i++ {
		res[i] = make([]byte, len(field[0])+extra*2+2)
	}

	for r := 0; r < len(res); r++ {
		res[r][0] = 2
		res[r][len(res[0])-1] = 2
	}

	for c := 0; c < len(res[0]); c++ {
		res[0][c] = 2
		res[len(res)-1][c] = 2
	}

	for r := 0; r < len(field); r++ {
		for c := 0; c < len(field[0]); c++ {
			res[extra+1+r][extra+1+c] = field[r][c]
		}
	}

	return res
}

func cp(field [][]byte) [][]byte {
	res := make([][]byte, len(field))
	for i := 0; i < len(field); i++ {
		row := make([]byte, len(field[0]))
		copy(row, field[i])
		res[i] = row
	}

	return res
}

func doStep(field [][]byte, encoding []byte, background byte) (field2 [][]byte, background2 byte) {
	if background == 0 {
		background2 = encoding[0]
	} else {
		background2 = encoding[511]
	}

	field2 = cp(field)
	for r := 1; r < len(field2)-1; r++ {
		for c := 1; c < len(field2[0])-1; c++ {
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

			field2[r][c] = encoding[index]
		}
	}

	return field2, background2
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
