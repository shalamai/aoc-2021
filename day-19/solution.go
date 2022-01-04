package main

import (
	"fmt"
	"math"
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
	scanners := getScannerMap()
	
	acc := make([]coord, 0)
	for _, s := range scanners {
		acc = union(acc, s.beacons)
	}

	return len(acc)
}

func part2() int {
	scanners := getScannerMap()
	max := 0
	for _, a := range scanners {
		for _, b := range scanners {
			distance := distance(a.coord, b.coord)
			if distance > max {
				max = distance
			}
		}
	}

	return max
}

func getScannerMap() []scanner {
	relativeBeacons := parseInput()
	initialScanner := scanner{coord{x: 0, y: 0, z: 0}, relativeBeacons[0]}

	scanners := detectScanners(initialScanner, relativeBeacons)

	return values(scanners)
}

func parseInput() [][]coord {
	scannersData := make([][]coord, 0)
	rawScannersData, _ := os.ReadFile("./input.txt")
	for _, rawScannerData := range strings.Split(string(rawScannersData), "\n\n") {
		scannerData := make([]coord, 0)
		for _, row := range strings.Split(rawScannerData, "\n")[1:] {
			rawNumbers := strings.Split(row, ",")
			x, _ := strconv.ParseInt(rawNumbers[0], 10, 64)
			y, _ := strconv.ParseInt(rawNumbers[1], 10, 64)
			z, _ := strconv.ParseInt(rawNumbers[2], 10, 64)

			scannerData = append(scannerData, coord{int(x), int(y), int(z)})
		}
		scannersData = append(scannersData, scannerData)
	}

	return scannersData
}

func detectScanners(initialScanner scanner, beacons [][]coord) map[int]scanner {
	scanners := make(map[int]scanner)
	scanners[0] = initialScanner

	queue := make([]int, 0)
	queue = append(queue, 0)

	for len(queue) > 0 {
		pinnedScanner := queue[0]
		queue = queue[1:]

		for i := 0; i < len(beacons); i++ {
			if _, found := scanners[i]; !found {
				s, isDetected := detectScanner(scanners[pinnedScanner].beacons, beacons[i])
				if isDetected {
					scanners[i] = s
					queue = append(queue, i)
				}
			}
		}
	}

	return scanners
}

func detectScanner(absoluteBeacons1 []coord, relativeBeacons2 []coord) (scanner, bool) {
	for _, absolute1 := range absoluteBeacons1 {
		for _, relative2 := range relativeBeacons2 {
			for _, facing := range facings {
				for _, top := range tops {
					p := position{top, facing}
					coord := calcScannerCoord(p, absolute1, relative2)
					absoluteBeacons2 := calcAbsoluteBeacons(coord, p, relativeBeacons2)
					sameBeacons := intersection(absoluteBeacons1, absoluteBeacons2)
					if len(sameBeacons) >= 12 {
						return scanner{coord, absoluteBeacons2}, true
					}
				}
			}
		}
	}

	return scanner{}, false
}

func calcScannerCoord(p position, absoluteBeacon coord, relativeBeacon coord) coord {
	delta := rotate(p, relativeBeacon)
	return coord{absoluteBeacon.x - delta.x, absoluteBeacon.y - delta.y, absoluteBeacon.z - delta.z}
}

func calcAbsoluteBeacons(scannerCoord coord, scannerPosition position, beacons []coord) []coord {
	acc := make([]coord, 0)
	for _, b := range beacons {
		delta := rotate(scannerPosition, b)
		absoluteB := coord{scannerCoord.x + delta.x, scannerCoord.y + delta.y, scannerCoord.z + delta.z}
		acc = append(acc, absoluteB)
	}

	return acc
}

func rotate(p position, c coord) coord {
	afterFacing := coord{}
	switch p.facing {
	case xpos:
		afterFacing = coord{c.x, c.y, c.z}
	case xneg:
		afterFacing = coord{-c.x, c.y, -c.z}
	case zpos:
		afterFacing = coord{-c.z, c.y, c.x}
	case zneg:
		afterFacing = coord{c.z, c.y, -c.x}
	case ypos:
		afterFacing = coord{-c.y, c.x, c.z}
	case yneg:
		afterFacing = coord{c.y, -c.x, c.z}
	}

	afterTop := coord{}
	switch p.top {
	case up:
		afterTop = coord{afterFacing.x, afterFacing.y, afterFacing.z}
	case down:
		afterTop = coord{afterFacing.x, -afterFacing.y, -afterFacing.z}
	case left:
		afterTop = coord{afterFacing.x, afterFacing.z, -afterFacing.y}
	case right:
		afterTop = coord{afterFacing.x, -afterFacing.z, afterFacing.y}
	}

	return afterTop
}

func intersection(c1 []coord, c2 []coord) []coord {
	set := make([]coord, 0)
	hash := make(map[coord]bool)

	for _, b := range c1 {
		hash[b] = true
	}

	for _, b := range c2 {
		if _, found := hash[b]; found {
			set = append(set, b)
		}
	}

	return set
}

func union(c1 []coord, c2 []coord) []coord {
	set := make([]coord, 0)
	hash := make(map[coord]bool)

	for _, b := range c1 {
		hash[b] = true
		set = append(set, b)
	}

	for _, b := range c2 {
		if _, found := hash[b]; !found {
			set = append(set, b)
		}
	}

	return set
}

func distance(p1 coord, p2 coord) int {
	return int(math.Abs(float64(p1.x-p2.x)) + math.Abs(float64(p1.y-p2.y)) + math.Abs(float64(p1.z-p2.z)))
}

func values(m map[int]scanner) []scanner {
	acc := make([]scanner, 0)
	for _, s := range m {
		acc = append(acc, s)
	}

	return acc
}

type facing int

const (
	xpos facing = iota
	xneg        = iota
	ypos        = iota
	yneg        = iota
	zpos        = iota
	zneg        = iota
)

var facings = []facing{xpos, xneg, ypos, yneg, zpos, zneg}

type top int

const (
	up    top = iota
	down      = iota
	left      = iota
	right     = iota
)

var tops = []top{up, down, left, right}

type coord struct {
	x, y, z int
}

type position struct {
	top    top
	facing facing
}

type scanner struct {
	coord
	beacons []coord
}
