package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	state, rules := parseInput()

	res0 := run(state, rules, 40)
	fmt.Println(res0)
	res1 := sum(res0)
	fmt.Println(res1)
	res2 := res(res1)
	fmt.Println(res2)

}

func run(state map[string]int64, rules map[string]string, steps int) map[string]int64 {
	for i := 0; i < steps; i++ {
		stateNew := make(map[string]int64)
		for k, v := range state {
			middle := rules[k]
			stateNew[k[:1] + middle] += v
			stateNew[middle + k[1:]] += v
		}
		state = stateNew
	}

	return state
}

func sum(state map[string]int64) (acc map[string]int64) {
	acc = make(map[string]int64)

	starts := make(map[string]int64)
	ends := make(map[string]int64)

	for k, v := range state {
		starts[k[:1]] += v
		ends[k[1:]] += v
	}

	for i := byte('A'); i < byte('Z') + 1; i++ {
		start := starts[string(i)]
		end := ends[string(i)]
		if start > 0 || end > 0 {
			if start > end {
				acc[string(i)] = start
			} else {
				acc[string(i)] = end
			}
		}
	}
	
	return
}

func res(counts map[string]int64) int64 {
	max := int64(0)
	min := int64(9223372036854775807)

	for _, v := range counts {
		if v > max {
			max = v
		}

		if v < min {
			min = v
		}
	}

	return max - min
}

func parseInput() (state map[string]int64, rules map[string]string) {
	state = make(map[string]int64)
	rules = make(map[string]string)

	data, _ := os.ReadFile("./input.txt")

	d := strings.Split(string(data), "\n\n")
	d1 := d[0]
	d2 := d[1]
	
	for i := 0; i < len(d1) - 1; i++ {
		state[d1[i:i+2]]++
	}

	for _, r := range strings.Split(d2, "\n") {
		kv := strings.Split(r, " -> ")
		rules[kv[0]] = kv[1]
	}

	return
}