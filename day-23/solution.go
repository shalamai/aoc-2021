package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	// start := time.Now()
	// res1 := part1()
	// fmt.Println(time.Since(start))
	// fmt.Println(res1)
	
	start := time.Now()
	res2 := part2()
	fmt.Println(time.Since(start))
	fmt.Println(res2)
}

func part1() int {
	initial := state{
		players: []player{
			{"A", coord{2, 3}}, {"B", coord{3, 3}}, {"C", coord{2, 5}}, {"A", coord{3, 5}},
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

	return findMin(initial).price
}

func part2() int {
	initial := state{
		players: []player{
			{"A", coord{2, 3}}, {"D", coord{3, 3}}, {"D", coord{4, 3}}, {"B", coord{5, 3}}, 
			{"C", coord{2, 5}}, {"C", coord{3, 5}}, {"B", coord{4, 5}}, {"A", coord{5, 5}},
			{"B", coord{2, 7}}, {"B", coord{3, 7}}, {"A", coord{4, 7}}, {"D", coord{5, 7}}, 
			{"D", coord{2, 9}}, {"A", coord{3, 9}}, {"C", coord{4, 9}}, {"C", coord{5, 9}},
		},
		targets: map[string][]coord{
			"A": {coord{5, 3}, coord{4, 3}, coord{3, 3}, coord{2, 3}},
			"B": {coord{5, 5}, coord{4, 5}, coord{3, 5}, coord{2, 5}},
			"C": {coord{5, 7}, coord{4, 7}, coord{3, 7}, coord{2, 7}},
			"D": {coord{5, 9}, coord{4, 9}, coord{3, 9}, coord{2, 9}},
		},
		history: make([]history, 0),
	}

	return findMin(initial).price
}

var cache = map[string]findMinRes{}

var cacheHit = 0

func findMin(s state) findMinRes {
	fmt.Printf("%v, %v\n", len(cache), cacheHit)
	if v, ok := cache[s.hash()]; ok {
		cacheHit++
		return v
	}

	if len(s.players) == 0 {
		return findMinRes{0, true}
	}

	min := 0
	solved := false
	for i := range s.players {
		for _, move := range moves(s, i) {
			if res := findMin(move.state); res.solved {
				total := move.price + res.price
				if !solved || total < min {
					min = total
					solved = true
				}
			}
		}
	}

	res := findMinRes{min, solved}
	cache[s.hash()] = res
	return res
}

func moves(s state, i int) []move {
	p := s.players[i]
	if p.r == 1 {
		return movesIntoTheRoom(s, i)
	} else {
		return movesIntoTheHall(s, i)
	}
}

func movesIntoTheHall(s state, pi int) []move {
	res := make([]move, 0)
	p := s.players[pi]

	for _, spot := range accessibleHallSpots(s, pi) {
		s2 := s.copy()
		s2.players[pi].coord = spot
		s2.history = append(s2.history, history{p.name, p.coord, spot})

		res = append(res, move{calcPrice(p, spot), s2})
	}

	return res
}

func movesIntoTheRoom(s state, pi int) []move {
	res := make([]move, 0)

	if isRoomAccessible(s, pi) && isRoomFree(s, pi) {
		p := s.players[pi]
		t := s.targets[p.name][0]

		s2 := s.copy()
		s2.targets[p.name] = s2.targets[p.name][1:]
		s2.players = append(s2.players[:pi], s2.players[pi+1:]...)
		s2.history = append(s2.history, history{p.name, p.coord, t})
		res = append(res, move{calcPrice(p, t), s2})
	}

	return res
}

func accessibleHallRange(s state, pi int) (int, int) {
	p := s.players[pi]
	from := 1
	to := 11
	for i2, p2 := range s.players {
		if pi == i2 || p2.r != 1 {
			continue
		}

		if p2.c < p.c && p2.c >= from {
			from = p2.c + 1
		}

		if p2.c > p.c && p2.c <= to {
			to = p2.c - 1
		}
	}

	return from, to
}

func accessibleHallSpots(s state, pi int) []coord {
	res := make([]coord, 0)
	from, to := accessibleHallRange(s, pi)

	for c := from; c <= to; c++ {
		if !containsInt(entrances, c) {
			res = append(res, coord{1, c})
		}
	}

	return res
}

func isRoomAccessible(s state, pi int) bool {
	from, to := accessibleHallRange(s, pi)
	roomC := s.targets[s.players[pi].name][0].c

	return roomC >= from && roomC <= to
}

func isRoomFree(s state, pi int) bool {
	for _, spot := range s.targets[s.players[pi].name] {
		for _, p2 := range s.players {
			if p2.coord == spot {
				return false
			}
		}
	}

	return true
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

var entrances = []int{3, 5, 7, 9}
var playerTypes = []string{"A", "B", "C", "D"}
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
	players []player
	targets map[string][]coord
	history []history
}

type move struct {
	price int
	state
}

type findMinRes struct {
	price int
	solved bool
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
		players: players2,
		targets: targets2,
		history: history2,
	}
}

func (s state) hash() string {
	h := ""

	for _, p := range s.players {
		h += fmt.Sprintf("%v", p)
	}
	
	for _, t := range playerTypes {
		for _, spot := range s.targets[t] {
			h += fmt.Sprintf("%v,%v", t, spot)
		}
	}

	return h
}
