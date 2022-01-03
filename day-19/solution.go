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
			distance := calcDistance(a.coord, b.coord)
			if distance > max {
				max = distance
			}
		}
	}

	return max
}

func getOceanMap() ([]scanner, []beacon) {
	beaconsByScanners := parseInput()
	initialScanner := scanner{coord{x: 0, y: 0, z: 0}, scannerState{top: up, facing: xpos}}

	scanners := scannersMap(initialScanner, beaconsByScanners)
	beacons := beaconsMap(scanners, beaconsByScanners)

	return values(scanners), beacons
}

func parseInput() [][]beacon {
	scannersData := make([][]beacon, 0)
	rawScannersData, _ := os.ReadFile("./input.txt")
	for _, rawScannerData := range strings.Split(string(rawScannersData), "\n\n") {
		scannerData := make([]beacon, 0)
		for _, row := range strings.Split(rawScannerData, "\n")[1:] {
			rawNumbers := strings.Split(row, ",")
			x, _ := strconv.ParseInt(rawNumbers[0], 10, 64)
			y, _ := strconv.ParseInt(rawNumbers[1], 10, 64)
			z, _ := strconv.ParseInt(rawNumbers[2], 10, 64)

			scannerData = append(scannerData, beacon{coord{int(x), int(y), int(z)}})
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
				s, isDetected := detectScanner(scanners[pinnedScanner], beacons[pinnedScanner], beacons[i])
				if isDetected {
					scanners[i] = s
					queue = append(queue, i)
				}
			}
		}
	}

	return scanners
}

func detectScanner(scanner1 scanner, beacons1 []beacon, beacons2 []beacon) (scanner2 scanner, isOverlapped bool) {
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

	return scanner{coord{absoluteBeacon1.x - x, absoluteBeacon1.y - y, absoluteBeacon1.z - z}, scannerState}
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
	return beacon{coord{scanner.x + x, scanner.y + y, scanner.z + z}}
}

func calcAbsoluteBeaconDelta(scanner scannerState, b beacon) (x int, y int, z int) {
	afterFacing := coord{}
	switch scanner.facing {
	case xpos:
		afterFacing = coord{b.x, b.y, b.z}
	case xneg:
		afterFacing = coord{-b.x, b.y, -b.z}
	case zpos:
		afterFacing = coord{-b.z, b.y, b.x}
	case zneg:
		afterFacing = coord{b.z, b.y, -b.x}
	case ypos:
		afterFacing = coord{-b.y, b.x, b.z}
	case yneg:
		afterFacing = coord{b.y, -b.x, b.z}
	}

	afterTop := coord{}
	switch scanner.top {
	case up:
		afterTop = coord{afterFacing.x, afterFacing.y, afterFacing.z}
	case down:
		afterTop = coord{afterFacing.x, -afterFacing.y, -afterFacing.z}
	case left:
		afterTop = coord{afterFacing.x, afterFacing.z, -afterFacing.y}
	case right:
		afterTop = coord{afterFacing.x, -afterFacing.z, afterFacing.y}
	}

	return afterTop.x, afterTop.y, afterTop.z 
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

func calcDistance(p1 coord, p2 coord) int {
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

type scannerState struct {
	top    top
	facing facing
}

type scanner struct {
	coord
	scannerState
}

type beacon struct {
	coord
}
