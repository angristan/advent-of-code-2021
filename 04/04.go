package main

import (
	"advent-of-code-2021/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type BingoNumber struct {
	Number int
	Marked bool
}

type BingoBoard struct {
	ID      int
	Numbers [][]BingoNumber
	Won     bool // For part 2
}

type BingoDrawnNumbers struct {
	Numbers []int
}

type BingoInput struct {
	DrawnNumbers BingoDrawnNumbers
	Boards       []BingoBoard
}

func part1(params utils.RunParams) {
	input := getInput(params)

	for _, number := range input.DrawnNumbers.Numbers {
		for _, board := range input.Boards {
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					if board.Numbers[i][j].Number == number {
						board.Numbers[i][j].Marked = true
					}
				}
			}
		}

		winningBoards := getWinningBoards(input)
		if len(winningBoards) == 0 {
			continue
		} else {
			fmt.Printf("score: %d\n", getScoreOfWinningBoard(number, winningBoards[0]))
			return
		}

	}
}

func part2(params utils.RunParams) {
	input := getInput(params)

	wonBoardsCount := 0

	for _, number := range input.DrawnNumbers.Numbers {
		for boardIndex, board := range input.Boards {
			if board.Won {
				continue
			}
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					if board.Numbers[i][j].Number == number {
						input.Boards[boardIndex].Numbers[i][j].Marked = true
					}
				}
			}
		}

		winningBoards := getWinningBoards(input)
		if len(winningBoards) == 0 {
			continue
		} else {
			for _, wonBoard := range winningBoards {
				if input.Boards[wonBoard.ID].Won {
					// Board has alread won
					// probably because a row AND a column became full at the same time
					continue
				}
				wonBoardsCount++
				input.Boards[wonBoard.ID].Won = true

				// Is this the last winning board?
				if wonBoardsCount == len(input.Boards) || number == input.DrawnNumbers.Numbers[len(input.DrawnNumbers.Numbers)-1] {
					fmt.Printf("score of last winning board: %d\n", getScoreOfWinningBoard(number, wonBoard))
					return
				}
			}
		}
	}
}

// sum of all unmarked numbers * winning number
func getScoreOfWinningBoard(winningNumber int, winningBoard BingoBoard) int {
	score := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !winningBoard.Numbers[i][j].Marked {
				score += winningBoard.Numbers[i][j].Number
			}
		}
	}

	return score * winningNumber
}

func getWinningBoards(input BingoInput) []BingoBoard {
	// Check if any row is all marked
	winningBoards := []BingoBoard{}
	for _, board := range input.Boards {
		if board.Won {
			continue
		}
		for i := 0; i < 5; i++ {
			allMarkedInRow := true
			for j := 0; j < 5; j++ {
				if !board.Numbers[i][j].Marked {
					allMarkedInRow = false
					break
				}
			}
			if allMarkedInRow {
				winningBoards = append(winningBoards, board)
				break
			}
		}
	}

	// Check if any column is all marked
	for _, board := range input.Boards {
		if board.Won {
			continue
		}
		for i := 0; i < 5; i++ {
			allMarkedInColumn := true
			for j := 0; j < 5; j++ {
				if !board.Numbers[j][i].Marked {
					allMarkedInColumn = false
					break
				}
			}
			if allMarkedInColumn {
				winningBoards = append(winningBoards, board)
				break
			}
		}
	}

	return winningBoards
}

func main() {
	part1(utils.RunParams{Sample: true})
	part1(utils.RunParams{Sample: false})
	part2(utils.RunParams{Sample: true})
	part2(utils.RunParams{Sample: false})
}

func getInput(params utils.RunParams) BingoInput {
	lines, err := utils.ReadFileToString("04", params)
	if err != nil {
		panic(err)
	}

	input := BingoInput{}

	boardsString := []string{}

	numbersString := lines[0]

	numbersStringSlice := strings.Split(numbersString, ",")
	for _, numberString := range numbersStringSlice {
		number, err := strconv.Atoi(numberString)
		if err != nil {
			panic(err)
		}
		input.DrawnNumbers.Numbers = append(input.DrawnNumbers.Numbers, number)
	}

	// remove first line
	boardsString = lines[1:]

	board := newBoard()
	board.ID = 0
	space := regexp.MustCompile(`\s+`)
	currentBoardRow := 0

	for _, line := range boardsString {
		if line == "" {
			continue
		}
		// Remove all duplicate whitespace since a row can be like this:
		// ` 3 15  0  2 22`
		line = strings.TrimSpace(line)
		line = space.ReplaceAllString(line, " ")
		numbers := strings.Split(line, " ")
		for col, numberStr := range numbers {
			number, err := strconv.Atoi(numberStr)
			if err != nil {
				panic(err)
			}
			board.Numbers[currentBoardRow][col].Number = number
		}

		if currentBoardRow == 4 {
			input.Boards = append(input.Boards, board)
			board = newBoard()
			board.ID = len(input.Boards)
			currentBoardRow = 0
		} else {
			currentBoardRow++
		}
	}

	return input
}

func newBoard() BingoBoard {
	board := BingoBoard{}
	board.Numbers = make([][]BingoNumber, 5)
	for j := range board.Numbers {
		board.Numbers[j] = make([]BingoNumber, 5)
		for k := range board.Numbers[j] {
			board.Numbers[j][k] = BingoNumber{}
		}
	}

	return board
}
