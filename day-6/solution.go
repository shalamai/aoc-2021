package main

import "fmt"

var input []int = []int{
	3, 1, 5, 4, 4, 4, 5, 3, 4, 4, 1, 4, 2, 3, 1, 3, 3, 2, 3, 2, 5, 1, 1, 4, 4, 3, 2, 4, 2, 4, 1, 5, 3, 3, 2, 2, 2, 5, 5, 1, 3, 4, 5, 1, 5, 5, 1, 1, 1, 4, 3, 2, 3, 3, 3, 4, 4, 4, 5, 5, 1, 3, 3, 5, 4, 5, 5, 5, 1, 1, 2, 4, 3, 4, 5, 4, 5, 2, 2, 3, 5, 2, 1, 2, 4, 3, 5, 1, 3, 1, 4, 4, 1, 3, 2, 3, 2, 4, 5, 2, 4, 1, 4, 3, 1, 3, 1, 5, 1, 3, 5, 4, 3, 1, 5, 3, 3, 5, 4, 2, 3, 4, 1, 2, 1, 1, 4, 4, 4, 3, 1, 1, 1, 1, 1, 4, 2, 5, 1, 1, 2, 1, 5, 3, 4, 1, 5, 4, 1, 3, 3, 1, 4, 4, 5, 3, 1, 1, 3, 3, 3, 1, 1, 5, 4, 2, 5, 1, 1, 5, 5, 1, 4, 2, 2, 5, 3, 1, 1, 3, 3, 5, 3, 3, 2, 4, 3, 2, 5, 2, 5, 4, 5, 4, 3, 2, 4, 3, 5, 1, 2, 2, 4, 3, 1, 5, 5, 1, 3, 1, 3, 2, 2, 4, 5, 4, 2, 3, 2, 3, 4, 1, 3, 4, 2, 5, 4, 4, 2, 2, 1, 4, 1, 5, 1, 5, 4, 3, 3, 3, 3, 3, 5, 2, 1, 5, 5, 3, 5, 2, 1, 1, 4, 2, 2, 5, 1, 4, 3, 3, 4, 4, 2, 3, 2, 1, 3, 1, 5, 2, 1, 5, 1, 3, 1, 4, 2, 4, 5, 1, 4, 5, 5, 3, 5, 1, 5, 4, 1, 3, 4, 1, 1, 4, 5, 5, 2, 1, 3, 3,
}

func main() {
	
	// simulate([]int{3,4,3,1,2}, 80)
	fmt.Println(simulateFast(input, 256))
}

func print(initial int, days int) {
	fmt.Printf("%v - %v\n", days, 1+calc(days-initial))
}

func calc(days int) int {
	if days < 6 {
		return 1
	}

	return calc(days-8) + calc(days-6)
}

func simulateFast(fishes []int, days int) int {
	var population [9]int
	for _, f := range fishes {
		population[f]++
	}

	for i := 0; i < days; i++ {
		new := population[0]
		for i := 1; i < 9; i++ {
			population[i - 1] = population[i]
		}
		population[8] = new
		population[6] += new
	}

	acc := 0
	for i := 0; i < 9; i++ {
		acc += population[i]
	}

	return acc
}



func simulateForOne(days int) map[int]int {
	acc := make(map[int]int)
	population := []int{8}
	newFishes := make([]int, 0)
	for i := 0; i <= days; i++ {
		population = append(population, newFishes...)
		newFishes = newFishes[:0]
		for i, fish := range population {
			if fish == 0 {
				population[i] = 6
				newFishes = append(newFishes, 8)
			} else {
				population[i]--
			}
		}

		acc[i] = len(population)
	}

	return acc
}

func simulate(fishes []int, days int) {
	population := fishes
	newFishes := make([]int, 0)
	for i := 0; i <= days; i++ {
		population = append(population, newFishes...)
		newFishes = newFishes[:0]
		for i, fish := range population {
			if fish == 0 {
				population[i] = 6
				newFishes = append(newFishes, 8)
			} else {
				population[i]--
			}
		}
	}

	fmt.Println(len(population))
}
