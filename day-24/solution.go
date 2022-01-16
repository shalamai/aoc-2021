package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// res1 := part1()
	// fmt.Println(res1)

	instructions := parseInput()
	a := &alu{map[string]int{"x": 0, "y": 0, "z": 0, "w": 0}}

	start := time.Now()
	a.exec(instructions, number2stream(11111111111111))
	fmt.Println(time.Since(start))
	fmt.Println(a)
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
