package main

import (
	"advent-of-code-2021/utils"
	"fmt"
	"sort"
	"strings"
)

type Status int

const (
	StatusOK Status = iota
	StatusUnfinished
	StatusCorrupted
)

func part1(params utils.RunParams) {
	input := getInput(params)

	illegalChars := []string{}

	points := map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}

	for _, v := range input {
		line := v

		status := getLineStatus(line)

		if status.status == StatusCorrupted {
			illegalChars = append(illegalChars, status.illegalChar)
		}
	}

	res := 0
	for _, char := range illegalChars {
		res += points[char]
	}

	fmt.Println("Part 1:", res)
}

func part2(params utils.RunParams) {
	input := getInput(params)

	scores := []int{}

	for _, v := range input {
		line := v

		status := getLineStatus(line)

		if status.status == StatusUnfinished {
			missingChars := getMissingChars(status.remainingChars)
			score := getScore(missingChars)
			scores = append(scores, score)
		}
	}

	sort.Ints(scores)

	res := scores[len(scores)/2]

	fmt.Println("Part 2:", res)
}

func getScore(line []string) int {

	points := map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}

	score := 0

	for _, char := range line {
		score *= 5
		score += points[char]
	}

	return score
}

func getMissingChars(line []string) []string {
	matchingChars := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
		"<": ">",
	}

	result := []string{}

	for _, char := range line {
		switch char {
		case "(", "[", "{", "<":
			result = append(result, matchingChars[char])
		}
	}

	result = reverse(result)

	return result

}

type LineStatusResult struct {
	status         Status
	illegalChar    string
	remainingChars []string
}

func getLineStatus(line []string) LineStatusResult {
	matchingChars := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
		">": "<",
	}

	encounteredClosingChar := true
	for encounteredClosingChar {
		encounteredClosingChar = false
	myloop:
		for i, char := range line {
			switch char {
			case ")", "]", "}", ">":
				if line[i-1] == matchingChars[char] {
					line = remove(line, i-1)
					line = remove(line, i-1)
					encounteredClosingChar = true
					break myloop
				} else {
					return LineStatusResult{
						status:         StatusCorrupted,
						illegalChar:    char,
						remainingChars: line,
					}
				}
			}
		}
	}

	if len(line) == 0 {
		return LineStatusResult{
			status:         StatusOK,
			illegalChar:    "",
			remainingChars: line,
		}
	}

	return LineStatusResult{
		status:         StatusUnfinished,
		illegalChar:    "",
		remainingChars: line,
	}
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func reverse(s []string) []string {
	a := make([]string, len(s))
	copy(a, s)

	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}

	return a
}

func main() {
	part1(utils.RunParams{Sample: true})
	part1(utils.RunParams{Sample: false})
	part2(utils.RunParams{Sample: true})
	part2(utils.RunParams{Sample: false})
}

func getInput(params utils.RunParams) [][]string {
	lines, err := utils.ReadFileToString("10", params)
	if err != nil {
		panic(err)
	}

	result := [][]string{}
	for i, line := range lines {
		result = append(result, []string{})
		result[i] = append(result[i], strings.Split(line, "")...)
	}

	return result
}
