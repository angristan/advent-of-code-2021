package main

import (
	"advent-of-code-2021/utils"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/montanaflynn/stats"
)

func part1(params utils.RunParams) {
	input := getInput(params)

	median, _ := stats.Median(input)

	res := 0
	for _, value := range input {
		res += int(math.Abs(float64(value - median)))
	}

	fmt.Println("Part 1:", res)

}

func part2(params utils.RunParams) {
	input := getInput(params)

	max, _ := stats.Max(input)

	tmp := make([]float64, int(max)+1)

	for goal := range tmp {
		for _, value := range input {
			tmp[goal] += getFuel(value, float64(goal))
		}
	}

	res, _ := stats.Min(tmp)

	fmt.Println("Part 2:", int(res))

}

func getFuel(start float64, end float64) float64 {
	n := int(math.Abs(float64(start - end)))

	// https://math.stackexchange.com/a/593320
	return float64(((n * n) + n) / 2)
}

func main() {
	part1(utils.RunParams{Sample: true})
	part1(utils.RunParams{Sample: false})
	part2(utils.RunParams{Sample: true})
	part2(utils.RunParams{Sample: false})
}

func getInput(params utils.RunParams) []float64 {
	lines, err := utils.ReadFileToString("07", params)
	if err != nil {
		panic(err)
	}

	result := []float64{}
	numbersStr := strings.Split(lines[0], ",")
	for _, numberStr := range numbersStr {
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			panic(err)
		}
		result = append(result, float64(number))
	}

	return result
}
