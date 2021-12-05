package main

import (
	"advent-of-code-2021/utils"
	"fmt"
	"strconv"
)

func part1(params utils.RunParams) {
	// depth measureement, one depth per line
	input := getInput(params)

	// Problem:
	// How many measurements are larger than the previous measurement?
	result := 0

	for i := range input {
		if i == 0 {
			continue
		}

		if input[i] > input[i-1] {
			result++
		}
	}

	fmt.Println(result)
}

func part2(params utils.RunParams) {
	// depth measureement, one depth per line
	input := getInput(params)

	// Problem:
	// Consider sums of a three-measurement sliding window.
	// How many sums are larger than the previous sum?
	result := 0

	for i := range input {
		if i == 0 {
			continue
		}

		prevSum := input[i-1] + input[i]
		sum := input[i]
		if i <= len(input)-2 {
			prevSum = prevSum + input[i+1]
			sum = sum + input[i+1]
		}
		if i <= len(input)-3 {
			sum = sum + input[i+2]
		}

		if sum > prevSum {
			result++
		}
	}

	fmt.Println(result)
}

func main() {
	part1(utils.RunParams{Sample: true})
	part1(utils.RunParams{Sample: false})
	part2(utils.RunParams{Sample: true})
	part2(utils.RunParams{Sample: false})
}

func getInput(params utils.RunParams) []int {

	lines, err := utils.ReadFileToString("01", params)
	if err != nil {
		panic(err)
	}

	numbers := make([]int, len(lines))
	for i := range lines {
		y, err := strconv.Atoi(lines[i])
		numbers[i] = y

		if err != nil {
			panic(err)
		}
	}

	return numbers
}
