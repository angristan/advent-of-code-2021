package main

import (
	"advent-of-code-2021/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/montanaflynn/stats"
)

type Point struct {
	x, y int
}

func part1(params utils.RunParams) {
	input := getInput(params)

	lowPoints := []float64{}

	for k1 := range input {
		for k2 := range input[k1] {
			neighbors := []float64{}
			if k1 >= 1 {
				neighbors = append(neighbors, input[k1-1][k2])
			}
			if k2 >= 1 {
				neighbors = append(neighbors, input[k1][k2-1])
			}
			if k2 < len(input[k1])-1 {
				neighbors = append(neighbors, input[k1][k2+1])
			}
			if k1 < len(input)-1 {
				neighbors = append(neighbors, input[k1+1][k2])
			}

			min, err := stats.Min(neighbors)
			if err != nil {
				panic(err)
			}

			if input[k1][k2] < min {
				lowPoints = append(lowPoints, input[k1][k2])
			}
		}
	}

	res := 0.0
	for _, point := range lowPoints {
		res += point + 1
	}

	fmt.Println(res)
}

func part2(params utils.RunParams) {
	input := getInput(params)

	lowPoints := []Point{}

	for k1 := range input {
		for k2 := range input[k1] {
			neighbors := []float64{}
			if k1 >= 1 {
				neighbors = append(neighbors, input[k1-1][k2])
			}
			if k2 >= 1 {
				neighbors = append(neighbors, input[k1][k2-1])
			}
			if k2 < len(input[k1])-1 {
				neighbors = append(neighbors, input[k1][k2+1])
			}
			if k1 < len(input)-1 {
				neighbors = append(neighbors, input[k1+1][k2])
			}

			min, err := stats.Min(neighbors)
			if err != nil {
				panic(err)
			}

			if input[k1][k2] < min {
				lowPoints = append(lowPoints, Point{x: k1, y: k2})
			}
		}
	}

	// key is lowpoint, value is locations
	basins := map[Point][]Point{}

	for _, point := range lowPoints {
		basins[point] = []Point{point}

		locationsToExplore := []Point{point}

		for len(locationsToExplore) > 0 {

			currentLocation := locationsToExplore[0]
			locationsToExplore = locationsToExplore[1:]

			neighbors := []Point{}

			if currentLocation.x >= 1 {
				neighbors = append(neighbors, Point{x: currentLocation.x - 1, y: currentLocation.y})
			}
			if currentLocation.y >= 1 {
				neighbors = append(neighbors, Point{x: currentLocation.x, y: currentLocation.y - 1})
			}
			if currentLocation.y < len(input[currentLocation.x])-1 {
				neighbors = append(neighbors, Point{x: currentLocation.x, y: currentLocation.y + 1})
			}
			if currentLocation.x < len(input)-1 {
				neighbors = append(neighbors, Point{x: currentLocation.x + 1, y: currentLocation.y})
			}

			for _, neighbor := range neighbors {
				if input[neighbor.x][neighbor.y] == 9 {
					continue
				}

				if contains(basins[point], neighbor) {
					continue
				}

				if input[neighbor.x][neighbor.y] > input[currentLocation.x][currentLocation.y] {
					basins[point] = append(basins[point], neighbor)
					locationsToExplore = append(locationsToExplore, neighbor)
				}
			}
		}
	}

	basinSizes := []int{}
	for _, basin := range basins {
		basinSizes = append(basinSizes, len(basin))
	}

	sort.Ints(basinSizes)

	res := basinSizes[len(basinSizes)-1] * basinSizes[len(basinSizes)-2] * basinSizes[len(basinSizes)-3]

	fmt.Println(res)
}

func contains(s []Point, e Point) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	part1(utils.RunParams{Sample: true})
	part1(utils.RunParams{Sample: false})
	part2(utils.RunParams{Sample: true})
	part2(utils.RunParams{Sample: false})
}

func getInput(params utils.RunParams) [][]float64 {
	lines, err := utils.ReadFileToString("09", params)
	if err != nil {
		panic(err)
	}

	result := [][]float64{}
	for i, line := range lines {
		result = append(result, []float64{})
		for _, number := range strings.Split(line, "") {
			num, err := strconv.Atoi(number)
			if err != nil {
				panic(err)
			}
			result[i] = append(result[i], float64(num))
		}
	}

	return result
}
