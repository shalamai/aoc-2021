package main

import "fmt"

func main() {

	// input0 := [10][10]int{
	// 	{5, 4, 8, 3, 1, 4, 3, 2, 2, 3},
	// 	{2, 7, 4, 5, 8, 5, 4, 7, 1, 1},
	// 	{5, 2, 6, 4, 5, 5, 6, 1, 7, 3},
	// 	{6, 1, 4, 1, 3, 3, 6, 1, 4, 6},
	// 	{6, 3, 5, 7, 3, 8, 5, 4, 7, 8},
	// 	{4, 1, 6, 7, 5, 2, 4, 6, 4, 5},
	// 	{2, 1, 7, 6, 8, 4, 1, 7, 2, 1},
	// 	{6, 8, 8, 2, 8, 8, 1, 1, 3, 4},
	// 	{4, 8, 4, 6, 8, 4, 8, 5, 5, 4},
	// 	{5, 2, 8, 3, 7, 5, 1, 5, 2, 6},
	// }

	input1 := [10][10]int{
		{5, 7, 2, 3, 5, 7, 3, 1, 5, 8},
		{3, 1, 5, 4, 7, 4, 8, 5, 6, 3},
		{4, 7, 8, 3, 5, 1, 4, 8, 7, 8},
		{3, 8, 4, 8, 1, 4, 2, 3, 7, 5},
		{3, 6, 3, 7, 7, 2, 4, 1, 5, 1},
		{8, 5, 8, 3, 1, 7, 2, 4, 8, 4},
		{7, 7, 4, 7, 4, 4, 4, 1, 8, 4},
		{1, 6, 1, 3, 3, 6, 7, 8, 8, 2},
		{6, 2, 2, 8, 6, 1, 4, 2, 2, 7},
		{4, 7, 3, 2, 2, 2, 5, 3, 3, 4},
	}

	fmt.Println(simulate2(input1))
}

func simulate(input [10][10]int, steps int) int {
	flashes := 0

	for i := 0; i < steps; i++ {
		for r := 0; r < 10; r++ {
			for c := 0; c < 10; c++ {
				input[r][c]++
			}
		}

		flashed := make([]point, 0)
		for true {
			q := make([]point, 0)
			for r := 0; r < 10; r++ {
				for c := 0; c < 10; c++ {
					if input[r][c] > 9 {
						input[r][c] = 0
						flashed = append(flashed, point{r, c})
						flashes++
						q = append(q, point{r + 1, c})
						q = append(q, point{r - 1, c})
						q = append(q, point{r, c + 1})
						q = append(q, point{r, c - 1})
						q = append(q, point{r + 1, c + 1})
						q = append(q, point{r + 1, c - 1})
						q = append(q, point{r - 1, c + 1})
						q = append(q, point{r - 1, c - 1})
					}
				}
			}

			if len(q) == 0 {
				break
			}

			for _, o := range q {
				if o.r >= 0 && o.r < 10 && o.c >= 0 && o.c < 10 && !contains(flashed, o) {
					input[o.r][o.c]++
				}
			}
		}
	}

	return flashes
}

func simulate2(input [10][10]int) int {
	flashes := 0

	step := 0
	for true {
		step++

		for r := 0; r < 10; r++ {
			for c := 0; c < 10; c++ {
				input[r][c]++
			}
		}

		flashed := make([]point, 0)
		for true {
			q := make([]point, 0)
			for r := 0; r < 10; r++ {
				for c := 0; c < 10; c++ {
					if input[r][c] > 9 {
						input[r][c] = 0
						flashed = append(flashed, point{r, c})
						flashes++
						q = append(q, point{r + 1, c})
						q = append(q, point{r - 1, c})
						q = append(q, point{r, c + 1})
						q = append(q, point{r, c - 1})
						q = append(q, point{r + 1, c + 1})
						q = append(q, point{r + 1, c - 1})
						q = append(q, point{r - 1, c + 1})
						q = append(q, point{r - 1, c - 1})
					}
				}
			}

			if len(q) == 0 {
				break
			}

			for _, o := range q {
				if o.r >= 0 && o.r < 10 && o.c >= 0 && o.c < 10 && !contains(flashed, o) {
					input[o.r][o.c]++
				}
			}
		}

		if allFlashed(input) {
			return step
		}
	}

	return -1
}

func contains(as []point, b point) bool {
	for _, a := range as {
		if a == b {
			return true
		}
	}

	return false
}

type point struct {
	r, c int
}

func print(as [10][10]int) {
	for r := 0; r < 10; r++ {
		for c := 0; c < 10; c++ {
			fmt.Print(as[r][c])
		}
		fmt.Print("\n")
	}
}

func allFlashed(as [10][10]int) bool {
	for r := 0; r < 10; r++ {
		for c := 0; c < 10; c++ {
			if as[r][c] != 0 {
				return false
			}
		}
	}

	return true
}
