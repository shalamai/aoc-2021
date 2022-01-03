package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// acc := 0
	// data, _ := os.ReadFile("./input.txt")
	// for _, r := range strings.Split(string(data), "\n") {
	// 	parts := strings.Split(r, "|")
	// 	for _, d := range strings.Split(parts[1], " ") {
	// 		if len(d) == 2 || len(d) == 3 || len(d) == 4 || len(d) == 7 {
	// 			acc++
	// 		}
	// 	}
	// }

	codes, digits := parseInput()

	acc := 0
	for i := 0; i < len(codes); i++ {
		// fmt.Println(resolveEncoding(codes[i]))
			acc += decrypt(resolveEncoding(codes[i]), digits[i])
	} 

	fmt.Println(acc)
}

func parseInput() (codes [][]string, digits [][]string) {
	data, _ := os.ReadFile("./input.txt")
	for _, r := range strings.Split(string(data), "\n") {
		parts := strings.Split(r, "|")
		codes = append(codes, strings.Split(parts[0], " "))
		digits = append(digits, strings.Split(parts[1], " "))
	}

	return
}


func decrypt(encoding map[string]string, digits []string) int {
	res := ""
	for _, d := range digits {
		res += decryptDigit(encoding, d)
	}

	i, _ := strconv.ParseInt(res, 10, 64)
	return int(i)
}

func decryptDigit(encoding map[string]string, digit string) string {
	
	codeToDigit := map[string]string {
		"abcefg": "0",
		"cf": "1",
		"acdeg": "2",
		"acdfg" : "3",
		"bcdf": "4",
		"abdfg": "5",
		"abdefg": "6",
		"acf": "7",
		"abcdefg": "8",
		"abcdfg": "9",
	}

	original := make([]string, 0)
	for _, c := range digit {
		original = append(original, encoding[string(c)]) 
	}

	sort.Strings(original)
	fmt.Println(original)
	return codeToDigit[strings.Join(original, "")]
}

func resolveEncoding(encoded []string) map[string]string {
	encoding := make(map[string]string)

	encoding["a"] = decryptA(encoded)
	charactersStats := countCharacters(strings.Join(encoded, ""))
	encoding["b"] = getKeyByValue(charactersStats, 6)[0]
	encoding["e"] = getKeyByValue(charactersStats, 4)[0]
	encoding["f"] = getKeyByValue(charactersStats, 9)[0]
	encoding["c"] = strings.Replace(strings.Join(getKeyByValue(charactersStats, 8), ""), encoding["a"], "", 1)
	encoding["g"] = decryptG(encoding, encoded)
	encoding["d"] = diff("abcedfg", strings.Join(values(encoding), ""))[0]

	return reverse(encoding)
}

func decryptA(encoded []string) string {
	one := findByLength(encoded, 2)[0]
	seven := findByLength(encoded, 3)[0]
	res := diff(seven, one)
	return res[0]
}

func decryptG(encoding map[string]string, encoded []string) string {
	for _, w := range findByLength(encoded, 6) {
		if strings.Contains(w, encoding["a"]) &&
			strings.Contains(w, encoding["b"]) &&
			strings.Contains(w, encoding["c"]) &&
			strings.Contains(w, encoding["e"]) &&
			strings.Contains(w, encoding["f"]) {
				return diff(w, encoding["a"] + encoding["b"] + encoding["c"] + encoding["e"] + encoding["f"])[0]
		}
	}

	return ""
}

func reverse(m map[string]string) map[string]string {
	res := make(map[string]string)
	for k, v := range m {
		res[v] = k
	}

	return res
}

func values(m map[string]string) []string {
	res := make([]string, 0)
	for _, v := range m {
		res = append(res, v)
	}

	return res
}

func diff(s1 string, s2 string) []string {
	diff := make([]string, 0)
	for _, c := range s1 {
		if !strings.Contains(s2, string(c)) {
			diff = append(diff, string(c))
		}
	}

	return diff
}

func countCharacters(s string) map[string]int {
	m := make(map[string]int)
	for _, c := range s {
		m[string(c)]++
	}

	return m
}

func findByLength(arr []string, n int) []string {
	res := make([]string, 0)
	for _, e := range arr {
		if len(e) == n {
			res = append(res, e)
		}
	}

	return res
}

func getKeyByValue(m map[string]int, target int) []string {
	res := make([]string, 0)
	for k, v := range m {
		if v == target {
			res = append(res, k)
		}
	}

	return res
}
