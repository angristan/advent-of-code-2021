package main

import (
	"advent-of-code-2021/utils"
	"fmt"
	"strconv"
	"strings"

	"github.com/juliangruber/go-intersect"
)

type Input struct {
	patterns []string
	toDecode []string
}

func part1(params utils.RunParams) {
	input := getInput(params)

	// Problem:
	// In the output values, how many times do digits 1, 4, 7, or 8 appear?
	// (digits with a unique number of segments)

	result := 0
	for _, line := range input {
		for _, digit := range line.toDecode {
			switch len(digit) {
			case 2, 3, 4, 7: // 1, 7, 4, 8
				result++
			}
		}
	}

	fmt.Println(result)

}

func part2(params utils.RunParams) {
	input := getInput(params)

	// Problem:
	// For each line, there are 10 patterns. Based on these patterns, decode 4 digits
	// Then add up the resulting 4-digits numbers for each line to get the result

	result := 0

	//      0:      1:      2:      3:      4:
	//     aaaa    ....    aaaa    aaaa    ....
	//    b    c  .    c  .    c  .    c  b    c
	//    b    c  .    c  .    c  .    c  b    c
	//     ....    ....    dddd    dddd    dddd
	//    e    f  .    f  e    .  .    f  .    f
	//    e    f  .    f  e    .  .    f  .    f
	//     gggg    ....    gggg    gggg    ....

	//      5:      6:      7:      8:      9:
	//     aaaa    aaaa    aaaa    aaaa    aaaa
	//    b    .  b    .  .    c  b    c  b    c
	//    b    .  b    .  .    c  b    c  b    c
	//     dddd    dddd    ....    dddd    dddd
	//    .    f  e    f  .    f  e    f  .    f
	//    .    f  e    f  .    f  e    f  .    f
	//     gggg    gggg    ....    gggg    gggg

	// Unique digits based on their number of segments:  1, 4, 7, and 8
	// We can deduce the other digits based on how many segments they share
	// with the unique digits (excluding 8 bc it has all segments)

	// +-------+-----+----+----+----+
	// | digit | len | ∩1 | ∩7 | ∩4 |
	// +-------+-----+----+----+----+
	// |     0 |   6 |  2 |  3 |  3 |
	// |     2 |   5 |  1 |  2 |  2 |
	// |     3 |   5 |  2 |  3 |  3 |
	// |     5 |   5 |  1 |  2 |  3 |
	// |     6 |   6 |  1 |  2 |  3 |
	// |     9 |   6 |  2 |  3 |  4 |
	// +-------+-----+----+----+----+

	for _, line := range input {

		mapping := map[int]string{}
		// Unique digits
		for _, digit := range line.patterns {
			switch len(digit) {
			case 2:
				mapping[1] = digit
			case 3:
				mapping[7] = digit
			case 4:
				mapping[4] = digit
			case 7:
				mapping[8] = digit
			}
		}

		// The intersection table above, as code
		for _, digit := range line.patterns {
			switch len(digit) {
			case 5:
				if len(intersect.Simple(digit, mapping[4])) == 2 {
					mapping[2] = digit
					continue
				}
				if len(intersect.Simple(digit, mapping[1])) == 2 {
					mapping[3] = digit
					continue
				}
				mapping[5] = digit
			case 6:
				if len(intersect.Simple(digit, mapping[1])) == 1 {
					mapping[6] = digit
					continue
				}
				if len(intersect.Simple(digit, mapping[4])) == 4 {
					mapping[9] = digit
					continue
				}
				mapping[0] = digit
			}
		}

		res := ""

		for _, digit := range line.toDecode {
			for k, v := range mapping {
				// Letters are not ordered, so we have to check for each
				commonLetters := intersect.Simple(strings.Split(v, ""), strings.Split(digit, ""))
				if len(commonLetters) == len(v) && len(commonLetters) == len(digit) {
					res += fmt.Sprintf("%d", k)
				}
			}
		}

		num, _ := strconv.Atoi(res)
		result += num
	}

	fmt.Println(result)
}

func main() {
	part1(utils.RunParams{Sample: true})
	part1(utils.RunParams{Sample: false})
	part2(utils.RunParams{Sample: true})
	part2(utils.RunParams{Sample: false})
}

func getInput(params utils.RunParams) []Input {
	lines, err := utils.ReadFileToString("08", params)
	if err != nil {
		panic(err)
	}

	result := []Input{}
	for _, line := range lines {
		input := Input{}
		patternsStr := strings.Split(line, "|")[0]
		patternsStr = strings.TrimSpace(patternsStr)
		patternsDigit := strings.Split(patternsStr, " ")
		input.patterns = append(input.patterns, patternsDigit...)

		toDecodeStr := strings.Split(line, "|")[1]
		toDecodeStr = strings.TrimSpace(toDecodeStr)
		toDecodeDigits := strings.Split(toDecodeStr, " ")
		input.toDecode = append(input.toDecode, toDecodeDigits...)

		result = append(result, input)
	}

	return result
}
