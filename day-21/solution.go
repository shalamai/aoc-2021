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

	rolls := quantumRollX3(3)
	moves := quantumMoves(rolls)
	
	res := play(universe0, moves)
	
	if res.p1Wons > res.p2Wons {
		return res.p1Wons
	} else {
		return res.p2Wons
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

func quantumRollX3(max int) map[int]int {
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


	m := make(map[int]int)
	for _, v := range acc {
		m[v]++
	}

	return m
}

func quantumMoves(rolls map[int]int) map[int]map[int]int {
	acc := make(map[int]map[int]int)
	for from := 0; from < 10; from++ {
		acc2 := make(map[int]int)
		for roll, n := range rolls {
			move := movePlayer(from + 1, roll)
			acc2[move] = n
		}
		acc[from] = acc2
	}

	return acc
}

func movePlayer(from int, steps int) int {
	return 1 + (from-1+steps)%10
}

var cash map[universe]res = make(map[universe]res)
func play(u universe, moves map[int]map[int]int) res {
	if r, ok := cash[u]; ok {
		return r
	}
	
	var p1Wons, p2Wons int64

	if u.turn == 0 {
		for move, n := range moves[u.p1 - 1] {
			if u.p1Score+move >= 25 {
				p1Wons += int64(n)
			} else {
				res := play(universe{move, u.p2, u.p1Score + move, u.p2Score, 1}, moves)
				p1Wons += int64(n) * res.p1Wons
				p2Wons += int64(n) * res.p2Wons
			}
		}	
	} else {
		for move, n := range moves[u.p2 - 1] {
			if u.p2Score+move >= 25 {
				p2Wons += int64(n)
			} else {
				res := play(universe{u.p1, move, u.p1Score, u.p2Score + move, 0}, moves)
				p1Wons += int64(n) * res.p1Wons
				p2Wons += int64(n) * res.p2Wons
			}
		}	
	}

	r := res{p1Wons, p2Wons}
	cash[u] = r

	return r
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

type res struct {
	p1Wons, p2Wons int64
}
