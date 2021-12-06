package main

import (
	"advent-of-code-2021/utils"
	"fmt"
	"strconv"
	"strings"
)

func reproduce(fishes []int, days int) int {
	for i := 0; i < days; i++ {
		// Keep count of fishes ready to reproduce
		last0 := fishes[0]

		// Decrease reproduce timer
		for k := 0; k <= 7; k++ {
			fishes[k] = fishes[k+1]
		}
		// Add new fishes
		fishes[8] = last0
		// Reset fishes that have reproduced
		fishes[6] += last0
	}

	totalFishes := 0
	for _, v := range fishes {
		totalFishes += v
	}

	return totalFishes
}

func part1(params utils.RunParams) {
	// array of days_remaining->fish_count
	fishes := getInput(params)

	// Problem:
	// How many lanternfish would there be after 80 days?
	numberOfDays := 80

	reproducedFishesCount := reproduce(fishes, numberOfDays)

	fmt.Println(reproducedFishesCount)
}

func part2(params utils.RunParams) {
	// array of days_remaining->fish_count
	fishes := getInput(params)

	// Problem:
	// How many lanternfish would there be after 256 days?
	numberOfDays := 256

	reproducedFishesCount := reproduce(fishes, numberOfDays)

	fmt.Println(reproducedFishesCount)
}

func main() {
	part1(utils.RunParams{Sample: true})
	part1(utils.RunParams{Sample: false})
	part2(utils.RunParams{Sample: true})
	part2(utils.RunParams{Sample: false})
}

func getInput(params utils.RunParams) []int {
	// each fish as a single number that represents
	// the number of days until it creates a new lanternfish.
	lines, err := utils.ReadFileToString("06", params)
	if err != nil {
		panic(err)
	}

	// Convert fish list into array of days_remaining->fish_count
	result := make([]int, 9)
	numbersStr := strings.Split(lines[0], ",")
	for _, numberStr := range numbersStr {
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			panic(err)
		}
		result[number]++
	}

	return result
}
