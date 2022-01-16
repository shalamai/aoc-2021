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


// == largest	
//as0 -> 6
//as1 -> 5
//as2 -> 9
//as3 -> 8
//as4 -> 4
//as5 -> 9
//as6 -> 1
//as7 -> 9
//as8 -> 9
//as9 -> 9
//as10 -> 7
//as11 -> 9
//as12 -> 3
//as13 -> 9

// == smallest
//as0 -> 1
//as1 -> 1
//as2 -> 2
//as3 -> 1
//as4 -> 1
//as5 -> 6
//as6 -> 1
//as7 -> 9
//as8 -> 5
//as9 -> 4
//as10 -> 1
//as11 -> 7
//as12 -> 1
//as13 -> 3


	// largest := 65984919997939
	smallest := 11211619541713
	instructions := parseInput()
	a := &alu{map[string]int{"x": 0, "y": 0, "z": 0, "w": 0}}
	a.exec(instructions, number2stream(smallest))
	fmt.Println(a.vars["z"])
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
