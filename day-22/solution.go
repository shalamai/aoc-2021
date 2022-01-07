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
	steps := parseInput()
	area := cuboid{-50, 50,-50, 50,-50, 50}
	validSteps := make([]rebootStep, 0)
	for _, s := range steps {
		intersection, intersected := area.intersection(s.cuboid)
		if intersected{
			s.cuboid = intersection
			validSteps = append(validSteps, s)
		}
	}
	return calcReactorState(validSteps)
}

func part2() int {
	steps := parseInput()
	return calcReactorState(steps)
}

func parseInput() []rebootStep {
	data, _ := os.ReadFile("./input.txt")
	acc := make([]rebootStep, 0)
	for _, r := range strings.Split(string(data), "\n") {
		parts := strings.Split(r, " ")

		isOn := false
		if parts[0] == "on" {
			isOn = true
		}

		parseRange := func (input string) (int, int) {
			input = input[2:]
			parts := strings.Split(input, "..")
			from, _ := strconv.ParseInt(parts[0], 10, 64)
			to, _ := strconv.ParseInt(parts[1], 10, 64)
			return int(from), int(to)
		}
		
		coords := strings.Split(parts[1], ",")
		x1, x2 := parseRange(coords[0])
		y1, y2 := parseRange(coords[1])
		z1, z2 := parseRange(coords[2])

		acc = append(acc, rebootStep{cuboid{x1, x2, y1, y2, z1, z2}, isOn})
	}
	return acc
}

func calcReactorState(steps []rebootStep) int {
	acc := 0
	for i, s := range steps {
		if s.turnOn {
			acc += s.cuboid.volumeWithout(map2cuboids(steps[i+1:]))
		}
	}

	return acc
}

func map2cuboids(ss []rebootStep) []cuboid {
	acc := make([]cuboid, 0)
	for _, s := range ss {
		acc = append(acc, s.cuboid)
	}
	return acc
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

type cuboid struct {
	x1, x2, y1, y2, z1, z2 int
}

func (c cuboid) intersected(c2 cuboid) bool {
	
	intervalsIntersected := func(a1, a2, b1, b2 int) bool {
		return max(a1, b1) <= min(a2, b2)
	}

	return intervalsIntersected(c.x1, c.x2, c2.x1, c2.x2) &&
		intervalsIntersected(c.y1, c.y2, c2.y1, c2.y2) &&
		intervalsIntersected(c.z1, c.z2, c2.z1, c2.z2)
}

func (c cuboid) intersection(c2 cuboid) (cuboid, bool) {
	isIntersected := c.intersected(c2)

	if !isIntersected {
		return cuboid{}, false
	}

	intervalsIntersection := func(a1, a2, b1, b2 int) (int, int) {
		return max(a1, b1), min(a2, b2)
	}

	x1, x2 := intervalsIntersection(c.x1, c.x2, c2.x1, c2.x2)
	y1, y2 := intervalsIntersection(c.y1, c.y2, c2.y1, c2.y2)
	z1, z2 := intervalsIntersection(c.z1, c.z2, c2.z1, c2.z2)

	return cuboid{x1, x2, y1, y2, z1, z2}, true
}

func (c cuboid) volume() int {
	return (c.x2 - c.x1 + 1) * (c.y2 - c.y1 + 1) * (c.z2 - c.z1 + 1)
}

func (c cuboid) volumeWithout(toExclude []cuboid) int {
	intersections := make([]cuboid, 0)
	for _, e := range toExclude {
		intersection, intersected := c.intersection(e)
		if intersected {
			intersections = append(intersections, intersection)
		}
	}

	return c.volume() - volumeN(intersections)
}

func volumeN(cs []cuboid) int {
	acc := 0

	for len(cs) > 0 {
		c := cs[0]
		cs = cs[1:]

		acc += c.volumeWithout(cs)
	}

	return acc
}

type rebootStep struct {
	cuboid
	turnOn bool
}
