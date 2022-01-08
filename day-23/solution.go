package main

import (
	"fmt"
	"math"
)

// add memoization ?
func main() {
	res1 := part1()
	fmt.Println(res1)
}

func part1() int {
	initial := state{
		price: 0,
		players: []player{
			{"A", coord{2, 3}}, {"B", coord{3, 3}}, {"C", coord{2, 5}}, {"A", coord{3, 3}},
			{"B", coord{2, 7}}, {"D", coord{3, 7}}, {"D", coord{2, 9}}, {"C", coord{3, 9}},
		},
		targets: map[string][]coord{
			"A": {coord{3, 3}, coord{2, 3}},
			"B": {coord{3, 5}, coord{2, 5}},
			"C": {coord{3, 7}, coord{2, 7}},
			"D": {coord{3, 9}, coord{2, 9}},
		},
		history: make([]history, 0),
	}

	q := make([]state, 0)
	q = append(q, initial)

	min := math.MaxInt
	for len(q) > 0 {
		fmt.Println(len(q))
		s := q[0]
		q = q[1:]

		if len(s.players) == 0 {
			if s.price < min {
				min = s.price
				panic(s) // rm
			}
			continue
		}

		for i := range s.players {
			for _, s2 := range moves(s, i) {
				q = append(q, s2)
			}
		}
	}

	return min
}

func moves(s state, i int) []state {
	p := s.players[i]
	if p.r == 1 {
		return movesIntoTheRoom(s, i)
	} else {
		return movesIntoTheHall(s, i)
	}
}

func movesIntoTheHall(s state, i int) []state {
	res := make([]state, 0)

	from, to := accessibleHallRange(s, i)
	p := s.players[i]

	for j := from; j <= to; j++ {
		if !containsInt(entranceC, j) {
			c := coord{1, j}
			s2 := s.copy()
			s2.price += calcPrice(p, c)
			s2.players[i].coord = c
			s2.history = append(s2.history, history{p.name, p.coord, c})

			res = append(res, s2)
		}
	}

	return res
}

// change i to player
func movesIntoTheRoom(s state, i int) []state {
	res := make([]state, 0)

	from, to := accessibleHallRange(s, i)
	p := s.players[i]
	roomC := s.targets[p.name][0].c
	if roomC >= from && roomC <= to && accessibleRoom(s, i) {
		t := s.targets[p.name][0]

		s2 := s.copy()
		s2.price += calcPrice(p, t)
		s2.targets[p.name] = s2.targets[p.name][1:]
		s2.players = append(s2.players[:i], s2.players[i+1:]...)
		s2.history = append(s2.history, history{p.name, p.coord, t})
		res = append(res, s)
	}

	return res
}

func accessibleHallRange(s state, i int) (int, int) {
	p := s.players[i]
	from := 1
	to := 11
	for i2, p2 := range s.players {
		if i == i2 || p2.r != 1 {
			continue
		}

		if p2.c < p.c && p2.c > from {
			from = p2.c + 1
		}

		if p2.c > p.c && p2.c < to {
			to = p2.c - 1
		}
	}

	return from, to
}

func accessibleRoom(s state, i int) bool {
	p := s.players[i]
	room := s.targets[p.name]
	for i2, p2 := range s.players {
		if i == i2 {
			continue
		}

		if (p2.r == room[0].r && p2.c == room[0].c) || (p2.r == room[1].r && p2.c == room[1].c) {
			return false
		}
	}

	return true
}

func find(as []player, b player) (int, bool) {
	for i, a := range as {
		if a == b {
			return i, true
		}
	}

	return -1, false
}

func calcPrice(p player, to coord) int {
	distance := int(math.Abs(float64(p.coord.r)-float64(to.r)) + math.Abs(float64(p.coord.c)-float64(to.c)))
	return distance * price[p.name]
}

func containsInt(as []int, b int) bool {
	for _, a := range as {
		if a == b {
			return true
		}
	}

	return false
}

var entranceC = []int{3, 5, 7, 9}

var price map[string]int = map[string]int{
	"A": 1,
	"B": 10,
	"C": 100,
	"D": 1000,
}

type coord struct {
	r, c int
}

type player struct {
	name string
	coord
}

type history struct {
	name     string
	from, to coord
}

type state struct {
	price   int
	players []player
	targets map[string][]coord
	history []history
}

func (s state) copy() state {
	players2 := make([]player, len(s.players))
	copy(players2, s.players)
	
	targets2 := make(map[string][]coord)
	for k, v := range s.targets {
		targets2[k] = make([]coord, len(v))
		copy(targets2[k], v)
	}

	history2 := make([]history, len(s.history))
	copy(history2, s.history)

	return state{
		price: s.price,
		players: players2,
		targets: targets2,
		history: history2,
	}
}
