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

func part2() int {
	players := parseInput()
	playersScores := make([]int, len(players))
	universe0 := universe{players, playersScores}

	winners := make([]int, len(players))
	rolls := quantumRollX3(3)
	
	q := make([]universe, 0)
	q = append(q, universe0)

	for len(q) > 0 {
		u := q[0]
		q = q[1:]

		us, ws := playRound(u, rolls)
		for i, w := range ws {
			winners[i] += w
		}
		q = append(q, us...)
	}

	return max(winners)
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

func playRound(u universe, rolls []int) (us []universe, winners []int) {
	us = make([]universe, 0)
	winners = make([]int, len(u.players))
	for i, p := range u.players {
		for _, r := range rolls {
			p2 := movePlayer(p, r)
			if u.playersScores[i] + p2 >= 21 {
				winners[i]++
			} else {
				u2 := cpUniverse(u)
				u2.players[i] = p2
				u2.playersScores[i] = u.playersScores[i] + p2
				us = append(us, u2)
			}
		}
	}

	return
}

func cpUniverse(u universe) universe {
	players2 := make([]int, len(u.players))
	playersScores2 := make([]int, len(u.playersScores))

	copy(players2, u.players)
	copy(playersScores2, u.playersScores)

	return universe{players2, playersScores2}
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

func max(as []int) int {
	max := as[0]
	for _, a := range as {
		if a > max {
			max = a
		}
	}

	return max
}

type universe struct {
	players       []int
	playersScores []int
}
