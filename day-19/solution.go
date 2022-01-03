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
	_, beacons := getOceanMap()
	return len(beacons)
}

func part2() int {
	scanners, _ := getOceanMap()
	max := 0
	for _, a := range scanners {
		for _, b := range scanners {
			distance := calcDistance(a.pos, b.pos)
			if distance > max {
				max = distance
			}
		}
	}

	return max
}

func getOceanMap() ([]scanner, []beacon) {
	beaconsByScanners := parseInput()
	initialScanner := scanner{pos{x: 0, y: 0, z: 0}, scannerState{top: up, facing: xpos}}
	
	scanners := scannersMap(initialScanner, beaconsByScanners)
	beacons := beaconsMap(scanners, beaconsByScanners)

	return values(scanners), beacons
}

func parseInput() [][]beacon {
	scannersData := make([][]beacon, 0)
	rawData, _ := os.ReadFile("./input.txt")
	for _, group := range strings.Split(string(rawData), "\n\n") {
		scannerData := make([]beacon, 0)
		for _, group := range strings.Split(group, "\n")[1:] {
			rawNumbers := strings.Split(group, ",")
			x, _ := strconv.ParseInt(rawNumbers[0], 10, 64)
			y, _ := strconv.ParseInt(rawNumbers[1], 10, 64)
			z, _ := strconv.ParseInt(rawNumbers[2], 10, 64)

			scannerData = append(scannerData, beacon{pos{int(x), int(y), int(z)}})
		}
		scannersData = append(scannersData, scannerData)
	}

	return scannersData
}

func scannersMap(initialScanner scanner, beacons [][]beacon) map[int]scanner {
	scanners := make(map[int]scanner)
	scanners[0] = initialScanner
	
	queue := make([]int, 0)
	queue = append(queue, 0)
	
	for len(queue) > 0 {
		pinnedScanner := queue[0]
		queue = queue[1:]

		for i := 0; i < len(beacons); i++ {
			if _, found := scanners[i]; !found {
				s, isOverlapped := detectScanner(scanners[pinnedScanner], beacons[pinnedScanner], beacons[i])
				if isOverlapped {
					scanners[i] = s
					queue = append(queue, i)
				}
			}
		}
	}

	return scanners
}

func detectScanner(scanner1 scanner, beacons1 []beacon, beacons2 []beacon) (scanner2 scanner, isOverlapped bool) {
	facings := []facing{xpos, xneg, ypos, yneg, zpos, zneg}
	tops := []top{up, down, left, right}

	absoluteBeacons1 := calcAbsoluteBeacons(scanner1, beacons1)


	for _, b1 := range beacons1 {
		for _, b2 := range beacons2 {
			for _, facing := range facings {
				for _, top := range tops {
					maybeScanner2 := calcAbsoluteScanner(scanner1, b1, b2, facing, top)
					absoluteBeacons2 := calcAbsoluteBeacons(maybeScanner2, beacons2)
					sameBeacons := beaconIntersaction(absoluteBeacons1, absoluteBeacons2)
					if len(sameBeacons) >= 12 {
						return maybeScanner2, true
					}		
				}
			}
		}
	}

	return scanner{}, false
}

func beaconsMap(scanners map[int]scanner, beacons [][]beacon) []beacon {
	acc := make([]beacon, 0)
	for i, bs := range beacons {
		absoluteBeacons := calcAbsoluteBeacons(scanners[i], bs)
		acc = beaconUnion(acc, absoluteBeacons)
	}

	return acc
}

func calcAbsoluteScanner(scanner1 scanner, b1 beacon, b2 beacon, facing facing, top top) scanner {
	absoluteBeacon1 := calcAbsoluteBeacon(scanner1, b1)

	scannerState := scannerState{top, facing}
	x, y, z := calcAbsoluteBeaconDelta(scannerState, b2)

	return scanner{pos{absoluteBeacon1.x - x, absoluteBeacon1.y - y, absoluteBeacon1.z - z}, scannerState}
}

func calcAbsoluteBeacons(scanner scanner, beacons []beacon) []beacon {
	res := make([]beacon, 0)
	for _, b := range beacons {
		res = append(res, calcAbsoluteBeacon(scanner, b))
	}

	return res
}

func calcAbsoluteBeacon(scanner scanner, b beacon) beacon {
	x, y, z := calcAbsoluteBeaconDelta(scanner.scannerState, b)
	return beacon{pos{scanner.x + x, scanner.y + y, scanner.z + z}}
}

func calcAbsoluteBeaconDelta(scanner scannerState, b beacon) (x int, y int, z int) {
	switch scanner.facing {
	case xpos:
		switch scanner.top {
		case up:
			return b.x, b.y, b.z
		case down:
			return b.x, -b.y, -b.z
		case left:
			return b.x, b.z, -b.y
		case right:
			return b.x, -b.z, b.y
		}
	case xneg:
		switch scanner.top {
		case up:
			return -b.x, b.y, -b.z
		case down:
			return -b.x, -b.y, b.z
		case left:
			return -b.x, -b.z, -b.y
		case right:
			return -b.x, b.z, b.y
		}
	case zpos:
		switch scanner.top {
		case up:
			return -b.z, b.y, b.x
		case down:
			return b.z, -b.y, b.x
		case left:
			return -b.y, -b.z, b.x
		case right:
			return b.y, b.z, b.x
		}
	case zneg:
		switch scanner.top {
		case up:
			return b.z, b.y, -b.x
		case down:
			return -b.z, -b.y, -b.x
		case left:
			return -b.y, b.z, -b.x
		case right:
			return b.y, -b.z, -b.x
		}
	case ypos:
		switch scanner.top {
		case up:
			return -b.y, b.x, b.z
		case down:
			return b.y, b.x, -b.z
		case left:
			return -b.z, b.x, -b.y
		case right:
			return b.z, b.x, b.y
		}
	case yneg:
		switch scanner.top {
		case up:
			return b.y, -b.x, b.z
		case down:
			return -b.y, -b.x, -b.z
		case left:
			return b.z, -b.x, -b.y
		case right:
			return -b.z, -b.x, b.y
		}
	}

	panic("wrong input in calcAbsoluteBeaconDelta")
}

func beaconIntersaction(bs1 []beacon, bs2 []beacon) []beacon {
	set := make([]beacon, 0)
	hash := make(map[beacon]bool)

	for _, b := range bs1 {
		hash[b] = true
	}

	for _, b := range bs2 {
		if _, found := hash[b]; found {
			set = append(set, b)
		}
	}

	return set
}

func beaconUnion(bs1 []beacon, bs2 []beacon) []beacon {
	set := make([]beacon, 0)
	hash := make(map[beacon]bool)

	for _, b := range bs1 {
		hash[b] = true
		set = append(set, b)
	}

	for _, b := range bs2 {
		if _, found := hash[b]; !found {
			set = append(set, b)
		}
	}

	return set
}

func calcDistance(p1 pos, p2 pos) int {
	return int(math.Abs(float64(p1.x - p2.x)) + math.Abs(float64(p1.y - p2.y)) + math.Abs(float64(p1.z - p2.z)))
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

type top int

const (
	up    top = iota
	down      = iota
	left      = iota
	right     = iota
)

type pos struct {
	x, y, z int
}

type scannerState struct {
	top    top
	facing facing
}

type scanner struct {
	pos
	scannerState
}

type beacon struct {
	pos
}
