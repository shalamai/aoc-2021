package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	res1 := part1()
	fmt.Println(res1)
}

func part1() int {
	instructions := parseInput()
	for number := 99999999999999; number >= 11111111111111; number-- {
		if !containsZero(number) {
			fmt.Println(number)
			a := &alu{0, 0, 0, 0}
			a.exec(instructions, number2stream(number))
			if a.z == 0 {
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
	x, y, z, w int
}

func (al *alu) exec(instructions []instruction, inputStream []int) {
	for _, inst := range instructions {
		switch inst.operator {
		case "inp":
			a := al.varPointer(inst.operand1)
			n := inputStream[0]
			inputStream = inputStream[1:]
			*a = n
		case "add":
			a := al.varPointer(inst.operand1)
			b := al.varPointer(inst.operand2)
			*a = *a + *b
		case "mul":
			a := al.varPointer(inst.operand1)
			b := al.varPointer(inst.operand2)
			*a = *a * *b
		case "div":
			a := al.varPointer(inst.operand1)
			b := al.varPointer(inst.operand2)
			*a = *a / *b
		case "mod":
			a := al.varPointer(inst.operand1)
			b := al.varPointer(inst.operand2)
			*a = *a % *b
		case "eql":
			a := al.varPointer(inst.operand1)
			b := al.varPointer(inst.operand2)
			if *a == *b {
				*a = 1	
			} else {
				*a = 0	
			}
		default:
			panic("aaaaa!!! wrong operator")
		}
	}
}

func (a *alu) varPointer(v string) *int {
	switch v {
	case "x":
		return &a.x
	case "y":
		return &a.y
	case "z":
		return &a.z
	case "w":
		return &a.w
	default:
		n, _ := strconv.ParseInt(v, 10, 64)
		n2 := int(n)
		return &n2
	}
}