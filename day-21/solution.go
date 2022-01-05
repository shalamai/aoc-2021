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

	res2 := part2()
	fmt.Println(res2)
}

func part1() int {
	players := parseInput()
	playersScores := make([]int, len(players))
	diceRoll := 0
	diceState := 0

	for true {
		var steps int
		for i, p := range players {
			steps, diceState = rollDiceX3(diceState, 100)
			diceRoll += 3
			players[i] = movePlayer(p, steps)
			playersScores[i] += players[i]
			if playersScores[i] >= 1000 {
				return min(playersScores) * diceRoll
			}
		}
	}

	return -1
}

func part2() int64 {
	players := parseInput()
	universe0 := universe{players[0], players[1], 0, 0, 0}

	var p1Wons, p2Wons int64
	rolls := quantumRollX3(3)

	q := make([]universe, 0)
	q = append(q, universe0)

	for len(q) > 0 {
		u := q[0]
		q = q[1:]

		us, wons := play(u, rolls)
		if u.turn == 0 {
			p1Wons += int64(wons)
		} else {
			p2Wons += int64(wons)
		}
		q = append(q, us...)
	}

	if p1Wons > p2Wons {
		return p1Wons
	} else {
		return p2Wons
	}
}

func parseInput() []int {
	acc := make([]int, 0)
	data, _ := os.ReadFile("./input.txt")
	for _, row := range strings.Split(string(data), "\n") {
		parts := strings.Split(row, ": ")
		position, _ := strconv.ParseInt(parts[1], 10, 64)
		acc = append(acc, int(position))
	}
	return acc
}

func rollDiceX3(state int, max int) (res int, state2 int) {
	acc := 0
	for i := 0; i < 3; i++ {
		if state < max {
			state++
		} else {
			state = 1
		}
		acc += state
	}

	return acc, state
}

func quantumRollX3(max int) []int {
	acc := make([]int, 0)
	for i := 1; i <= max; i++ {
		acc = append(acc, i)
	}

	for i := 0; i < 2; i++ {
		acc2 := make([]int, 0)
		for j := 1; j <= max; j++ {
			for _, v := range acc {
				acc2 = append(acc2, v+j)
			}
		}
		acc = acc2
	}

	return acc
}

func movePlayer(from int, steps int) int {
	return 1 + (from-1+steps)%10
}

func play(u universe, rolls []int) (us []universe, wons int) {
	us = make([]universe, 0)
	wons = 0

	for _, r := range rolls {
		if u.turn == 0 {
			pNext := movePlayer(u.p1, r)
			if u.p1Score+pNext >= 10 {
				wons++
			} else {
				us = append(us, universe{pNext, u.p2, u.p1Score + pNext, u.p2Score, 1})
			}
		} else {
			pNext := movePlayer(u.p2, r)
			if u.p2Score+pNext >= 10 {
				wons++
			} else {
				us = append(us, universe{u.p1, pNext, u.p1Score, u.p2Score + pNext, 0})
			}
		}
	}

	return
}

func min(as []int) int {
	min := as[0]
	for _, a := range as {
		if a < min {
			min = a
		}
	}

	return min
}

type universe struct {
	p1, p2, p1Score, p2Score, turn int
}
