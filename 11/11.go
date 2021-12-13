package main

import (
	"advent-of-code-2021/utils"
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func part1(params utils.RunParams) {
	input := getInput(params)

	flashed := 0
	// 100 steps
	for i := 0; i < 100; i++ {
		toVisit := []Point{}
		for y := 0; y < len(input); y++ {
			for x := 0; x < len(input[y]); x++ {
				input[y][x] += 1
				if input[y][x] == 10 {
					flashed++
					toVisit = append(toVisit, Point{x: x, y: y})
				}
			}
		}

		for len(toVisit) > 0 {
			currentLocation := toVisit[0]
			toVisit = toVisit[1:]

			neighbors := []Point{}

			// up
			if currentLocation.y >= 1 {
				neighbors = append(neighbors, Point{x: currentLocation.x, y: currentLocation.y - 1})
			}
			// down
			if currentLocation.y < len(input)-1 {
				neighbors = append(neighbors, Point{x: currentLocation.x, y: currentLocation.y + 1})
			}
			// left
			if currentLocation.x >= 1 {
				neighbors = append(neighbors, Point{x: currentLocation.x - 1, y: currentLocation.y})
			}
			// right
			if currentLocation.x < len(input[currentLocation.y])-1 {
				neighbors = append(neighbors, Point{x: currentLocation.x + 1, y: currentLocation.y})
			}
			// up left
			if currentLocation.y >= 1 && currentLocation.x >= 1 {
				neighbors = append(neighbors, Point{x: currentLocation.x - 1, y: currentLocation.y - 1})
			}
			// up right
			if currentLocation.y >= 1 && currentLocation.x < len(input[currentLocation.y])-1 {
				neighbors = append(neighbors, Point{x: currentLocation.x + 1, y: currentLocation.y - 1})
			}
			// down left
			if currentLocation.y < len(input)-1 && currentLocation.x >= 1 {
				neighbors = append(neighbors, Point{x: currentLocation.x - 1, y: currentLocation.y + 1})
			}
			// down right
			if currentLocation.y < len(input)-1 && currentLocation.x < len(input[currentLocation.y])-1 {
				neighbors = append(neighbors, Point{x: currentLocation.x + 1, y: currentLocation.y + 1})
			}

			for _, neighbor := range neighbors {
				input[neighbor.y][neighbor.x] += 1
				if input[neighbor.y][neighbor.x] == 10 {
					flashed++
					toVisit = append(toVisit, neighbor)
				}
			}
		}

		for y := 0; y < len(input); y++ {
			for x := 0; x < len(input[y]); x++ {
				if input[y][x] > 9 {
					input[y][x] = 0
				}
			}
		}

		syncFlashed := true
		for y := 0; y < len(input); y++ {
			for x := 0; x < len(input[y]); x++ {
				if input[y][x] != 0 {
					syncFlashed = false
				}
			}
		}

		if syncFlashed {
			fmt.Println("sync", i+1)
		}
	}

	fmt.Println(flashed)
}

func part2(params utils.RunParams) {
	input := getInput(params)

	flashed := 0
	step := 0
	for {
		step++
		toVisit := []Point{}
		for y := 0; y < len(input); y++ {
			for x := 0; x < len(input[y]); x++ {
				input[y][x] += 1
				if input[y][x] == 10 {
					flashed++
					toVisit = append(toVisit, Point{x: x, y: y})
				}
			}
		}

		for len(toVisit) > 0 {
			currentLocation := toVisit[0]
			toVisit = toVisit[1:]

			neighbors := []Point{}

			// up
			if currentLocation.y >= 1 {
				neighbors = append(neighbors, Point{x: currentLocation.x, y: currentLocation.y - 1})
			}
			// down
			if currentLocation.y < len(input)-1 {
				neighbors = append(neighbors, Point{x: currentLocation.x, y: currentLocation.y + 1})
			}
			// left
			if currentLocation.x >= 1 {
				neighbors = append(neighbors, Point{x: currentLocation.x - 1, y: currentLocation.y})
			}
			// right
			if currentLocation.x < len(input[currentLocation.y])-1 {
				neighbors = append(neighbors, Point{x: currentLocation.x + 1, y: currentLocation.y})
			}
			// up left
			if currentLocation.y >= 1 && currentLocation.x >= 1 {
				neighbors = append(neighbors, Point{x: currentLocation.x - 1, y: currentLocation.y - 1})
			}
			// up right
			if currentLocation.y >= 1 && currentLocation.x < len(input[currentLocation.y])-1 {
				neighbors = append(neighbors, Point{x: currentLocation.x + 1, y: currentLocation.y - 1})
			}
			// down left
			if currentLocation.y < len(input)-1 && currentLocation.x >= 1 {
				neighbors = append(neighbors, Point{x: currentLocation.x - 1, y: currentLocation.y + 1})
			}
			// down right
			if currentLocation.y < len(input)-1 && currentLocation.x < len(input[currentLocation.y])-1 {
				neighbors = append(neighbors, Point{x: currentLocation.x + 1, y: currentLocation.y + 1})
			}

			for _, neighbor := range neighbors {
				input[neighbor.y][neighbor.x] += 1
				if input[neighbor.y][neighbor.x] == 10 {
					flashed++
					toVisit = append(toVisit, neighbor)
				}
			}
		}

		for y := 0; y < len(input); y++ {
			for x := 0; x < len(input[y]); x++ {
				if input[y][x] > 9 {
					input[y][x] = 0
				}
			}
		}

		syncFlashed := true
		for y := 0; y < len(input); y++ {
			for x := 0; x < len(input[y]); x++ {
				if input[y][x] != 0 {
					syncFlashed = false
				}
			}
		}

		if syncFlashed {
			fmt.Println("synced", step)
			break
		}
	}
}

func main() {
	part1(utils.RunParams{Sample: true})
	part1(utils.RunParams{Sample: false})
	part2(utils.RunParams{Sample: true})
	part2(utils.RunParams{Sample: false})
}

func getInput(params utils.RunParams) [][]int {
	lines, err := utils.ReadFileToString("11", params)
	if err != nil {
		panic(err)
	}

	result := [][]int{}
	for i, line := range lines {
		result = append(result, []int{})
		for _, char := range strings.Split(line, "") {
			digit, err := strconv.Atoi(char)
			if err != nil {
				panic(err)
			}
			result[i] = append(result[i], digit)
		}
	}

	return result
}
