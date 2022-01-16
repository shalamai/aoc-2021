package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// res1 := part1()
	// fmt.Println(res1)

	n := 11893121161114
	instructions := parseInput()
	a := &alu{map[string]int{"x": 0, "y": 0, "z": 0, "w": 0}}
	a.exec(instructions, number2stream(n))
	fmt.Println(a.vars["z"])
	r := execOptimized(number2stream(n))
	fmt.Println(r)
}

func part1() int {
	instructions := parseInput()
	for number := 99999999999999; number >= 11111111111111; number-- {
		if !containsZero(number) {
			fmt.Println(number)
			a := &alu{map[string]int{"x": 0, "y": 0, "z": 0, "w": 0}}
			a.exec(instructions, number2stream(number))
			if a.vars["z"] == 0 {
				return number
			}
		}
	}

	return -1
}

func part1Optimized() bool {
	for number := 99999999999999; number >= 11111111111111; number-- {
		if !containsZero(number) {
			fmt.Println(number)
			isOk := execOptimized(number2stream(number))
			if isOk {
				return true
			}
		}
	}

	return false
}

func parseInput() []instruction {
	acc := make([]instruction, 0)
	data, _ := os.ReadFile("./input.txt")
	for _, row := range strings.Split(string(data), "\n") {
		parts := strings.Split(row, " ")
		inst := instruction{operator: parts[0], operand1: parts[1]}
		if len(parts) > 2 {
			inst.operand2 = parts[2]
		}
		acc = append(acc, inst)
	}

	return acc
}

func containsZero(n int) bool {
	for _, c := range strconv.Itoa(n) {
		if c == '0' {
			return true
		}
	}
	return false
}

func number2stream(n int) []int {
	acc := make([]int, 0)

	for _, c := range strconv.Itoa(n) {
		d, _ := strconv.ParseInt(string(c), 10, 64)
		acc = append(acc, int(d))
	}

	return acc
}

func execOptimized(as []int) bool {
	z0 := as[0] + 7
	z1 := z0*26 + (as[1] + 15)
	z2 := z1*26 + (as[2] + 2)
	var z3 int
	if z2%26-3 == as[3] {
		z3 = z2 / 26
	} else {
		z3 = (z2/26)*26 + (15 + as[3])
	}
	z4 := z3*26 + (as[4] + 14)
		
	var z5 int
	if ((z4 % 26) - 9) == as[5] {
		z5 = z4 / 26
	} else {
		z5 = (z4 / 26) * 26 + (as[5] + 2)
	}

	z6 := z5*26 + (as[6] + 15)

	var z7 int
	if z6%26-7 == as[7] {
		z7 = z6 / 26
	} else {
		z7 = (z6/26)*26 + (as[7] + 1)
	}

	var z8 int
	if z7%26-11 == as[8] {
		z8 = z7 / 26
	} else {
		z8 = (z7/26)*26 + (as[8] + 15)
	}

	var z9 int
	if z8%26-4 == as[9] {
		z9 = z8 / 26
	} else {
		z9 = (z8/26)*26 + (15 + as[9])
	}

	z10 := z9*26 + (as[10] + 12)
	z11 := z10*26 + (as[11] + 2)

	var z12 int
	if z11%26-8 == as[12] {
		z12 = z11 / 26
	} else {
		z12 = (z11/26)*26 + (as[12] + 13)
	}

	var z13 int
	if z12%26-10 == as[13] {
		z13 = z12 / 26
	} else {
		z13 = (z12/26)*26 + (as[13] + 13)
	}

	fmt.Println(z13)
	return z13 == 0
}

type instruction struct {
	operator string
	operand1 string
	operand2 string
}

type alu struct {
	vars map[string]int
}

func (al *alu) exec(instructions []instruction, inputStream []int) {
	for _, inst := range instructions {
		if _, found := al.vars[inst.operand1]; !found {
			continue
		}

		switch inst.operator {
		case "inp":
			al.vars[inst.operand1] = inputStream[0]
			inputStream = inputStream[1:]
		case "add":
			al.vars[inst.operand1] = al.operand2value(inst.operand1) + al.operand2value(inst.operand2)
		case "mul":
			al.vars[inst.operand1] = al.operand2value(inst.operand1) * al.operand2value(inst.operand2)
		case "div":
			al.vars[inst.operand1] = al.operand2value(inst.operand1) / al.operand2value(inst.operand2)
		case "mod":
			al.vars[inst.operand1] = al.operand2value(inst.operand1) % al.operand2value(inst.operand2)
		case "eql":
			a := al.operand2value(inst.operand1)
			b := al.operand2value(inst.operand2)
			if a == b {
				al.vars[inst.operand1] = 1
			} else {
				al.vars[inst.operand1] = 0
			}
		default:
			panic("aaaaa!!! wrong operator")
		}
	}
}

func (al *alu) operand2value(operand string) int {
	if _, found := al.vars[operand]; found {
		return al.vars[operand]
	} else {
		n, _ := strconv.ParseInt(operand, 10, 64)
		return int(n)
	}
}
