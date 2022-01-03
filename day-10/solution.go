package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	

	fmt.Println(part2())
}

func part1() int {
	acc := 0
	pointsMap := map[string]int {
		")": 3,
		"]": 57,
		"}": 1197,
		">":25137,
	}

	for _, r := range parseInput() {
		ch := findIllegalCharacter(r)
		if ch != "" {
			acc += pointsMap[ch]
		}
	}

	return acc
}

func part2() int {
	scores := make([]int, 0)
	pointsMap := map[string]int {
		"(": 1,
		"[": 2,
		"{": 3,
		"<": 4,
	}

	for _, r := range parseInput() {
		missing := findMissingPart(r)
		if missing != "" {
			acc := 0
			for _, ch := range missing {
				acc *= 5
				acc += pointsMap[string(ch)]
			}
			scores = append(scores, acc)
		}
	}

	sort.Ints(scores)
	return scores[(len(scores) - 1) / 2]
}

func parseInput() []string {
	data, _ := os.ReadFile("./input.txt")
	return strings.Split(string(data), "\n")
}


func findIllegalCharacter(s string) string {
	stack := make(Stack, 0)
	stack.Push(s[0])
	for i := 1; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' || s[i] == '<' {
			stack.Push(s[i])
		} else {
			ch, exists := stack.Pop()
			if !exists {
				return ""
			}
			if (s[i] == ')' && ch != '(') || (s[i] == ']' && ch != '[') || (s[i] == '}' && ch != '{') || (s[i] == '>' && ch != '<') {
				return string(s[i])
			} 
		}
	}

	return ""
}

func findMissingPart(s string) string {
	stack := make(Stack, 0)
	stack.Push(s[0])
	for i := 1; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' || s[i] == '<' {
			stack.Push(s[i])
		} else {
			ch, exists := stack.Pop()
			if !exists {
				return ""
			}
			if (s[i] == ')' && ch != '(') || (s[i] == ']' && ch != '[') || (s[i] == '}' && ch != '{') || (s[i] == '>' && ch != '<') {
				return ""
			} 
		}
	}

	return reverse(string(stack))
}

type Stack []byte

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(ch byte) {
	*s = append(*s, ch) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (byte, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1 // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index] // Remove it from the stack by slicing it off.
		return element, true
	}
}

func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

			// swap the letters of the string,
			// like first with last and so on.
			rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}