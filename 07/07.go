package main

import (
	"advent-of-code-2021/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func part1(params utils.RunParams) {
	input := getInput(params)

	_, max := findMinAndMax(input)

	tmp := make([]int, max+1)

	for goal := range tmp {
		for _, value := range input {
			tmp[goal] += int(math.Abs(float64(value - goal)))
		}
	}

	res, _ := findMinAndMax(tmp)

	fmt.Println("Part 1:", res)

}

func part2(params utils.RunParams) {
	input := getInput(params)

	_, max := findMinAndMax(input)

	tmp := make([]int, max+1)

	for goal := range tmp {
		for _, value := range input {
			tmp[goal] += getFuel(value, goal)
		}
	}

	res, _ := findMinAndMax(tmp)

	fmt.Println("Part 2:", res)

}

func getFuel(start int, end int) int {
	fuel := 0
	for i := 0; i <= int(math.Abs(float64(start-end))); i++ {
		fuel += i
	}
	return fuel
}

func findMinAndMax(a []int) (min int, max int) {
	min = a[0]
	max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}

func main() {
	part1(utils.RunParams{Sample: true})
	part1(utils.RunParams{Sample: false})
	part2(utils.RunParams{Sample: true})
	part2(utils.RunParams{Sample: false})
}

func getInput(params utils.RunParams) []int {
	lines, err := utils.ReadFileToString("07", params)
	if err != nil {
		panic(err)
	}

	result := []int{}
	numbersStr := strings.Split(lines[0], ",")
	for _, numberStr := range numbersStr {
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			panic(err)
		}
		result = append(result, number)
	}

	return result
}
