package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	data, _ := os.ReadFile("./input0.txt")
	paths := buildPaths(string(data))
	// fmt.Println(paths)

	// res := explore("start", paths, []string{"start"}, make([]string, 0))
	res := explore2("start", paths, []string{"start"}, make([]string, 0), false)
	fmt.Println(res)
}

func buildPaths(input string) map[string][]string {
	paths := make(map[string][]string)

	for _, link := range strings.Split(input, "\n") {
		parsed := strings.Split(link, "-")
		paths[parsed[0]] = append(paths[parsed[0]], parsed[1]) 
		paths[parsed[1]] = append(paths[parsed[1]], parsed[0]) 
	}

	return paths
}

func explore(from string, pathes map[string][]string, visited []string, ways []string) int {
	if from == "end" {
		fmt.Println(ways)
		return 1
	}

	acc := 0
	for _, path := range pathes[from] {
		if contains(visited, path) {
			continue
		}

		visitedNew := visited
		if unicode.IsLower(rune(path[0])) {
			visitedNew = append(visitedNew, path)
		}

		waysNew := ways
		waysNew = append(waysNew, path)
		acc += explore(path, pathes, visitedNew, waysNew)
	}

	return acc
}

func explore2(from string, pathes map[string][]string, visited []string, ways []string, isSmallVisited bool) int {
	if from == "end" {
		fmt.Println(ways)
		return 1
	}

	acc := 0
	for _, path := range pathes[from] {
		if path == "start" {
			continue
		}

		if isSmallVisited {
			if contains(visited, path) {
				continue
			}
	
			visitedNew := visited
			if unicode.IsLower(rune(path[0])) {
				visitedNew = append(visitedNew, path)
			}
	
			waysNew := ways
			waysNew = append(waysNew, path)
			acc += explore2(path, pathes, visitedNew, waysNew, true)
		} else {
			if contains(visited, path) {
				waysNew := ways
				waysNew = append(waysNew, path)
				acc += explore2(path, pathes, visited, waysNew, true)
				continue
			}
	
			visitedNew := visited
			if unicode.IsLower(rune(path[0])) {
				visitedNew = append(visitedNew, path)
			}
	
			waysNew := ways
			waysNew = append(waysNew, path)
			acc += explore2(path, pathes, visitedNew, waysNew, false)
		}
	}

	return acc
}

func contains(as []string, b string) bool {
	for _, a := range as {
		if a == b {
			return true
		}
	}

	return false
}