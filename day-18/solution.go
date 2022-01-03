package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// learn about pointers
// interface vs pointer to interface

func main() {
	// input := "[[[[[9,8],1],2],3],4]"
	// input := "[[6,[5,[4,[3,2]]]],1]"
	// input := "[[[[[10,8],1],2],3],4]"
	// input := "[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]"
	// input := "[[[[[1,1],[2,2]],[3,3]],[4,4]],[5,5]]"

	// arr := []int{1,2,3}
	// res := insertOn(arr, 1, 9)
	// fmt.Println(res)

	// tree := parseTree(input)
	// exploded, _ := explode(tree)
	// print(tree2list(exploded, 0))

	// var a []leaf = []leaf{{value: 1}, {value: 2}, {value: 3}}
	// res := deleteAround(a, 0)
	// fmt.Println(res)

	// list := tree2list(tree, 0)
	// print(list)
	// tree2 := list2tree(list, 0)
	// print(tree2list(tree2, 0))

	// list := tree2list(tree, 0)
	// traversed := traverseWithDepth(p, 0)
	// fmt.Print("result: ")
	// print(list)
	// mag := calcMagnitude(tree)
	// fmt.Println(mag)
	// tree = reduce(tree)
	// print(tree2list(tree, 0))

	numbers := parseInput()
	// acc := accMagnitude(numbers)
	// fmt.Println(acc)

	max := maxMagnitude(numbers)
	fmt.Println(max)
}

func maxMagnitude(rawNumbers []string) int {
	max := 0
	for a := 0; a < len(rawNumbers); a++ {
		for b := 0; b < len(rawNumbers); b++ {
			if a != b {
				ap := parseTree(rawNumbers[a])
				bp := parseTree(rawNumbers[b])

				sum := add(ap, bp)
				sum = reduce(sum)

				mag := calcMagnitude(sum)
				if mag > max {
					max = mag
				}
			}
		}
	}

	return max
}

func accMagnitude(rawNumbers []string) int {
	acc := parseTree(rawNumbers[0])
	for i := 1; i < len(rawNumbers); i++ {
		p := parseTree(rawNumbers[i])
		acc = add(acc, p)
		acc = reduce(acc)
	}

	return calcMagnitude(acc)
}

func calcMagnitude(tree interface{}) int {
	switch v := tree.(type) {
	case leaf:
		return v.value
	case branch:
		left := 3 * calcMagnitude(v.left)
		right := 2 * calcMagnitude(v.right)
		return left + right
	default:
		fmt.Println("wrong input in calcMagnitude()")
		panic(v)
	}
}

func add(p1 interface{}, p2 interface{}) interface{} {
	return branch{left: p1, right: p2}
}

func parseInput() []string {
	acc := make([]string, 0)
	data, _ := os.ReadFile("./input.txt")
	for _, row := range strings.Split(string(data), "\n") {
		acc = append(acc, row)
	}

	return acc
}

func reduce(tree interface{}) interface{} {
	toReduce := tree
	for true {
		reducedTree, reduced := doReduce(toReduce)
		if !reduced {
			return reducedTree
		}
		toReduce = reducedTree
	}

	fmt.Println("wrong input in reduce()")
	panic(tree)
}

func doReduce(tree interface{}) (interface{}, bool) {
	explodedTree, exploded := explode(tree)
	if exploded {
		return explodedTree, true
	}

	return split(tree)
}

func split(tree interface{}) (interface{}, bool) {
	list := tree2list(tree, 0)
	splitted := false
	for i, a := range list {
		v, isLeaf := a.(leaf)
		if isLeaf && v.value >= 10 {
			left := leaf{value: int(math.Floor(float64(v.value) / 2))}
			left.depth = v.depth + 1
			right := leaf{value: int(math.Ceil(float64(v.value) / 2))}
			right.depth = v.depth + 1
			b := branch{left: left, right: right}
			b.depth = v.depth
			list[i] = b
			list = insertOn(list, i-1, left)
			list = insertOn(list, i+1, right)
			splitted = true
			break
		}
	}

	return list2tree(list, 0), splitted
}

func print(list []interface{}) {
	for _, a := range list {
		switch v := a.(type) {
		case leaf:
			fmt.Print(" " + fmt.Sprint(v.value) + " ")
		case branch:
			fmt.Print(" # ")
		default:
			fmt.Println("wrong input in print()")
			panic(v)
		}
	}
	fmt.Println()
}

func printDepth(list []interface{}) {
	for _, a := range list {
		switch v := a.(type) {
		case leaf:
			fmt.Print(" " + fmt.Sprint(v.depth) + " ")
		case branch:
			fmt.Print(" (" + fmt.Sprint(v.depth) + ") ")
		default:
			fmt.Println("wrong input in print()")
			panic(v)
		}
	}
	fmt.Println()
}

func tree2list(tree interface{}, depth int) []interface{} {
	switch v := tree.(type) {
	case leaf:
		v.depth = depth
		return []interface{}{v}
	case branch:
		l := tree2list(v.left, depth+1)
		r := tree2list(v.right, depth+1)
		res := make([]interface{}, 0)
		v.depth = depth
		res = append(res, l...)
		res = append(res, v)
		res = append(res, r...)
		return res
	default:
		fmt.Println("wrong input in tree2list()")
		panic(v)
	}
}

func list2tree(list []interface{}, depth int) interface{} {
	validateList(list)
	if len(list) == 1 {
		return list[0]
	}

	for i, a := range list {
		if getDepth(a) == depth {
			left := list2tree(list[:i], depth+1)
			right := list2tree(list[i+1:], depth+1)
			return branch{left: left, right: right}
		}
	}

	fmt.Println("wrong input in list2tree")
	panic(fmt.Sprintf("%v, %v, %T", depth, list, list))
}

func validateList(list []interface{}) {
	for _, a := range list {
		switch v := a.(type) {
		case leaf:
			continue
		case branch:
			continue
		default:
			fmt.Println("not valid element in list")
			panic(fmt.Sprintf("%v, %T", v, v))
		}
	}
}

func getDepth(tree interface{}) int {
	switch v := tree.(type) {
	case leaf:
		return v.depth
	case branch:
		return v.depth
	default:
		fmt.Println("wrong input in getDepth")
		panic(v)
	}
}

func explode(tree interface{}) (interface{}, bool) {
	list := tree2list(tree, 0)

	exploded := false
	var explodedBranch *branch = nil
	var explodedIndex int = 0

	for i, a := range list {
		b, isBranch := a.(branch)
		if isBranch && b.depth >= 4 {
			_, isLeftLeaf := b.left.(leaf)
			_, isRightLeaf := b.right.(leaf)
			if isLeftLeaf && isRightLeaf {
				exploded = true
				explodedIndex = i
				explodedBranch = &b
				break
			}
		}
	}

	if exploded {
		l := leaf{value: 0}
		l.depth = (*explodedBranch).depth
		list[explodedIndex] = l

		leftFrom := explodedIndex - 2
		if leftFrom >= 0 {
			for i := leftFrom; i >= 0; i-- {
				l, isLeaf := list[i].(leaf)
				if isLeaf {
					l.value += explodedBranch.left.(leaf).value
					list[i] = l
					break
				}
			}
		}

		rightFrom := explodedIndex + 2
		if rightFrom <= len(list)-1 {
			for i := rightFrom; i <= len(list)-1; i++ {
				r, isLeaf := list[i].(leaf)
				if isLeaf {
					r.value += explodedBranch.right.(leaf).value
					list[i] = r
					break
				}
			}
		}

		list = deleteAround(list, explodedIndex)
	}

	return list2tree(list, 0), exploded
}

func deleteAround(list []interface{}, index int) []interface{} {
	list, removedLeft := deleteFrom(list, index-1)
	if removedLeft {
		list, _ := deleteFrom(list, index)
		return list
	} else {
		list, _ := deleteFrom(list, index+1)
		return list
	}
}

func insertOn(list []interface{}, on int, v interface{}) []interface{} {
	if on == -1 {
		return append([]interface{}{v}, list...)
	} else if on == len(list) {
		return append(list, v)
	} else {
		res := make([]interface{}, 0)
		res = append(res, list[:on+1]...)
		res = append(res, v)
		return append(res, list[on+1:]...)
	}
}

func deleteFrom(list []interface{}, index int) ([]interface{}, bool) {
	if index < 0 || index > len(list)-1 {
		return list, false
	}
	if index == 0 {
		return list[1:], true
	}
	if index == len(list)-1 {
		return list[:len(list)-1], true
	}
	return append(list[:index], list[index+1:]...), true
}

func parseTree(rawNumber string) interface{} {
	v, isInt := strconv.ParseInt(rawNumber, 10, 64)
	if isInt == nil {
		return leaf{value: int(v)}
	} else {
		l, r := splitPair(rawNumber)
		lp := parseTree(l)
		rp := parseTree(r)
		return branch{left: lp, right: rp}
	}
}

func splitPair(input string) (left string, right string) {
	middleIndex := -1
	trimmed := input[1 : len(input)-1]

	if string(trimmed[0]) == "[" {
		groupCounter := 1
		for i, ch := range trimmed[1:] {
			if string(ch) == "[" {
				groupCounter++
			} else if string(ch) == "]" {
				groupCounter--
			}
			if groupCounter == 0 {
				middleIndex = i + 2
				break
			}
		}
	} else {
		for i, ch := range trimmed {
			if string(ch) == "," {
				middleIndex = i
				break
			}
		}
	}

	left = trimmed[:middleIndex]
	right = trimmed[middleIndex+1:]
	return
}

type node struct {
	depth int
}

type branch struct {
	node
	left, right interface{}
}

type leaf struct {
	node
	value int
}
