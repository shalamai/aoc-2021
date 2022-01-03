package main

import "fmt"

func main() {
	t := target{287, 309, -76, -48}
	// t := target{20, 30, -10, -5}

	res := possibleSpeeds(t)
	fmt.Println(res)
}

func possibleSpeeds(t target) int {
	minY := t.y0
	maxY := -t.y0 - 1
	
	acc := 0
	for x := speedToReach(t.x0); x <= t.x1; x++ {
		for y := minY; y <= maxY; y++ {
			if isSpeedValid(x, y, t) {
				acc++
			}
		}
	}

	return acc
}

func isSpeedValid(x int, y int, t target) bool {
	currX := 0
	currY := 0
	xSpeed := x
	ySpeed := y

	for true {
		currX += xSpeed
		currY += ySpeed
		xSpeed = adjustX(xSpeed)
		ySpeed = adjustY(ySpeed)

		if currX >= t.x0 && currX <= t.x1 && currY >= t.y0 && currY <= t.y1 {
			return true
		}
		if currX > t.x1 || currY < t.y0 {
			break
		}
	}

	return false
}

func adjustX(x int) int {
	if x > 0 {
		return x - 1
	} else {
		return 0
	}
}

func adjustY(y int) int {
	return y - 1
}

func speedToReach(x int) int {
	distance := 0
	i := 0
	for ; distance < x; i++ {
		distance += i
	}
	return i - 1
}

func f(n int) int {
	acc := 0
	for i := 1; i < n+1; i++ {
		acc += i
	}
	return acc
}

type target struct {
	x0, x1, y0, y1 int
}
