package main

import (
	"advent-of-code-2021/utils"
	"fmt"
	"strconv"
	"strings"
)

type Command struct {
	direction string
	distance  int
}

func part1(params utils.RunParams) {
	// series of commands with direction (up, down, forward) and distance
	input := getInput(params)

	// Problem
	// Calculate the horizontal position and depth you would have after following the planned course.
	//  What do you get if you multiply your final horizontal position by your final depth?

	horizontalPosition := 0
	depth := 0

	for i := range input {
		command := input[i]
		if command.direction == "up" {
			depth -= command.distance
		} else if command.direction == "down" {
			depth += command.distance
		} else if command.direction == "forward" {
			horizontalPosition += command.distance
		}
	}

	fmt.Println(horizontalPosition * depth)
}

func part2(params utils.RunParams) {
	// series of commands with direction (up, down, forward) and distance BUT with aim
	input := getInput(params)

	// Problem
	// Calculate the horizontal position and depth you would have after following the planned course.
	//  What do you get if you multiply your final horizontal position by your final depth?

	horizontalPosition := 0
	depth := 0
	aim := 0

	for i := range input {
		command := input[i]
		if command.direction == "up" {
			aim -= command.distance
		} else if command.direction == "down" {
			aim += command.distance
		} else if command.direction == "forward" {
			horizontalPosition += command.distance
			depth += aim * command.distance
		}
	}

	fmt.Println(horizontalPosition * depth)
}

func main() {
	part1(utils.RunParams{Sample: true})
	part1(utils.RunParams{Sample: false})
	part2(utils.RunParams{Sample: true})
	part2(utils.RunParams{Sample: false})
}

func getInput(params utils.RunParams) []Command {
	lines, err := utils.ReadFileToString("02", params)
	if err != nil {
		panic(err)
	}

	commands := make([]Command, len(lines))
	for i := range lines {
		command := strings.Split(lines[i], " ")
		direction := command[0]
		distance, err := strconv.Atoi(command[1])
		if err != nil {
			panic(err)
		}

		commands[i] = Command{direction, distance}
	}

	return commands
}
