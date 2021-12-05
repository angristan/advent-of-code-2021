package main

import (
	"advent-of-code-2021/utils"
	"fmt"
)

type Line struct {
	x1, y1, x2, y2 int
}

type Point struct {
	x, y int
}

func part1(params utils.RunParams) {
	// Each line of vents is given as a line segment
	// in the format x1,y1 -> x2,y2
	// where x1,y1 are the coordinates of one end the line segment
	// and x2,y2 are the coordinates of the other end.
	input := getInput(params)

	// For now, only consider horizontal and vertical lines:
	// lines where either x1 = x2 or y1 = y2.
	horAndVerLines := make([]Line, 0)
	for _, line := range input {
		if line.x1 == line.x2 {
			horAndVerLines = append(horAndVerLines, line)
		} else if line.y1 == line.y2 {
			horAndVerLines = append(horAndVerLines, line)
		}
	}

	// you need to determine the number of points where at least two lines overlap.
	pointOccurence := make(map[Point]int)
	overalpingPointCount := 0
	for _, line := range horAndVerLines {
		// vertical line
		if line.x1 == line.x2 {
			if line.y1 < line.y2 { // forwards
				for y := line.y1; y <= line.y2; y++ {
					pointOccurence[Point{line.x1, y}]++
					if pointOccurence[Point{line.x1, y}] == 2 {
						overalpingPointCount++
					}
				}
			} else { // backwards
				for y := line.y2; y <= line.y1; y++ {
					pointOccurence[Point{line.x1, y}]++
					if pointOccurence[Point{line.x1, y}] == 2 {
						overalpingPointCount++
					}
				}
			}
		} else {
			// horizontal line
			if line.x1 < line.x2 { // forwards
				for x := line.x1; x <= line.x2; x++ {
					pointOccurence[Point{x, line.y1}]++
					if pointOccurence[Point{x, line.y1}] == 2 {
						overalpingPointCount++
					}
				}
			} else { // backwards
				for x := line.x2; x <= line.x1; x++ {
					pointOccurence[Point{x, line.y1}]++
					if pointOccurence[Point{x, line.y1}] == 2 {
						overalpingPointCount++
					}
				}
			}
		}
	}

	fmt.Println("Part 1:", overalpingPointCount)
}

func part2(params utils.RunParams) {
	// Each line of vents is given as a line segment
	// in the format x1,y1 -> x2,y2
	// where x1,y1 are the coordinates of one end the line segment
	// and x2,y2 are the coordinates of the other end.
	input := getInput(params)

	// you need to determine the number of points where at least two lines overlap.
	pointOccurence := make(map[Point]int)
	overalpingPointCount := 0
	for _, line := range input {
		// vertical line
		if line.x1 == line.x2 {
			if line.y1 < line.y2 { // forwards
				for y := line.y1; y <= line.y2; y++ {
					pointOccurence[Point{line.x1, y}]++
					if pointOccurence[Point{line.x1, y}] == 2 {
						overalpingPointCount++
					}
				}
			} else { // backwards
				for y := line.y2; y <= line.y1; y++ {
					pointOccurence[Point{line.x1, y}]++
					if pointOccurence[Point{line.x1, y}] == 2 {
						overalpingPointCount++
					}
				}
			}
			continue
		}
		// horizontal line
		if line.y1 == line.y2 {
			if line.x1 < line.x2 { // forwards
				for x := line.x1; x <= line.x2; x++ {
					pointOccurence[Point{x, line.y1}]++

					if pointOccurence[Point{x, line.y1}] == 2 {
						overalpingPointCount++
					}
				}
			} else { // backwards
				for x := line.x2; x <= line.x1; x++ {
					pointOccurence[Point{x, line.y1}]++

					if pointOccurence[Point{x, line.y1}] == 2 {
						overalpingPointCount++
					}
				}
			}
			continue
		}
		// diagonal line
		if line.x1 < line.x2 { // forwards
			y := line.y1
			for x := line.x1; x <= line.x2; x++ {
				pointOccurence[Point{x, y}]++
				if pointOccurence[Point{x, y}] == 2 {
					overalpingPointCount++
				}
				if line.y1 < line.y2 { // forwards
					y++
				} else { // backwards
					y--
				}
			}
		}
		if line.x1 > line.x2 { // backwards
			y := line.y2
			for x := line.x2; x <= line.x1; x++ {
				pointOccurence[Point{x, y}]++

				if pointOccurence[Point{x, y}] == 2 {
					overalpingPointCount++
				}
				if line.y1 > line.y2 { // forwards
					y++
				} else { // backwards
					y--
				}
			}
		}
	}

	fmt.Println("Part 2:", overalpingPointCount)
}

func main() {
	part1(utils.RunParams{Sample: true})
	part1(utils.RunParams{Sample: false})
	part2(utils.RunParams{Sample: true})
	part2(utils.RunParams{Sample: false})
}

func getInput(params utils.RunParams) []Line {
	lines, err := utils.ReadFileToString("05", params)
	if err != nil {
		panic(err)
	}

	var result []Line
	for _, line := range lines {
		var x1, y1, x2, y2 int
		_, err := fmt.Sscanf(line, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		if err != nil {
			panic(err)
		}
		result = append(result, Line{x1, y1, x2, y2})
	}

	return result
}
