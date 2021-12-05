package main

import (
	"advent-of-code-2021/utils"
	"fmt"
	"strconv"
	"strings"
)

func part1(params utils.RunParams) {
	// input is a array of bit arrays
	diagnosticReport := getInput(params)

	// Problem:
	// Use the binary numbers in your diagnostic report to calculate the gamma rate and epsilon rate,
	// then multiply them together. What is the power consumption of the submarine?
	// (Be sure to represent your answer in decimal, not binary.)

	// Each bit in the gamma rate can be determined by finding the most common bit
	// in the corresponding position of all numbers in the diagnostic report.
	// The epsilon rate is calculated in a similar way; rather than use the most common bit, the least common bit from each position is used.

	gammaRateBinary := ""
	epsilonRateBinary := ""

	diagnosticCount := len(diagnosticReport)
	bitCountPerDiagnostic := len(diagnosticReport[0])

	// Find the most common bit in each column

	// For each bit position
	for i := 0; i < bitCountPerDiagnostic; i++ {
		countOf0 := 0
		countOf1 := 0

		// For each diagnostic report
		for j := 0; j < diagnosticCount; j++ {
			// Get the bit at the current position
			bit := diagnosticReport[j][i]

			if bit == "0" {
				countOf0++
			} else {
				countOf1++
			}

		}
		if countOf0 > countOf1 {
			gammaRateBinary += "0"
			epsilonRateBinary += "1"
		} else {
			gammaRateBinary += "1"
			epsilonRateBinary += "0"
		}
	}

	gammaRateDecimal, err := strconv.ParseInt(gammaRateBinary, 2, 64)
	if err != nil {
		panic(err)
	}
	epsilonRateDecimal, err := strconv.ParseInt(epsilonRateBinary, 2, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println("Power consumption:", gammaRateDecimal*epsilonRateDecimal)
}

func part2(params utils.RunParams) {
	// input is a array of bit arrays
	diagnosticReports := getInput(params)

	// Problem:
	// verify the life support rating, which can be determined by multiplying the oxygen generator rating by the CO2 scrubber rating.
	// Keep only numbers selected by the bit criteria for the type of rating value for which you are searching. Discard numbers which do not match the bit criteria.
	// If you only have one number left, stop; this is the rating value for which you are searching.
	// Otherwise, repeat the process, considering the next bit to the right.

	// diagnosticCount := len(diagnosticReport)
	bitCountPerDiagnostic := len(diagnosticReports[0])

	// To find oxygen generator rating, determine the most common value (0 or 1) in the current bit position, and keep only numbers with that bit in that position. If 0 and 1 are equally common, keep values with a 1 in the position being considered.

	// Determine matching diagnostic report to get oxygen generator rating
	diagnosticReportsForOxygenGenerator := diagnosticReports
	// For each bit position
	for i := 0; i < bitCountPerDiagnostic; i++ {
		countOf0 := 0
		countOf1 := 0
		chosenBitForOxgygenGenerator := ""

		// For each diagnostic report
		for j := 0; j < len(diagnosticReportsForOxygenGenerator); j++ {
			// Get the bit at the current position
			bit := diagnosticReportsForOxygenGenerator[j][i]

			if bit == "0" {
				countOf0++
			} else {
				countOf1++
			}
		}

		if countOf0 > countOf1 {
			chosenBitForOxgygenGenerator = "0"
		} else {
			chosenBitForOxgygenGenerator = "1"
		}

		newdiagnosticReportsForOxygenGenerator := make([][]string, 0)
		for j := 0; j < len(diagnosticReportsForOxygenGenerator); j++ {
			// Get the bit at the current position
			bit := diagnosticReportsForOxygenGenerator[j][i]
			if bit == chosenBitForOxgygenGenerator {
				// Add the diagnostic report
				newdiagnosticReportsForOxygenGenerator = append(newdiagnosticReportsForOxygenGenerator, diagnosticReportsForOxygenGenerator[j])
			}
		}
		diagnosticReportsForOxygenGenerator = newdiagnosticReportsForOxygenGenerator
		if len(diagnosticReportsForOxygenGenerator) == 1 {
			break
		}
	}

	// To find CO2 scrubber rating, determine the least common value (0 or 1) in the current bit position, and keep only numbers with that bit in that position. If 0 and 1 are equally common, keep values with a 0 in the position being considered.

	// Determine matching diagnostic report to get co2 scrubber rating
	diagnosticReportsForCO2Scrubber := diagnosticReports
	// For each bit position
	for i := 0; i < bitCountPerDiagnostic; i++ {
		countOf0 := 0
		countOf1 := 0
		chosenBitForCO2Scrubber := ""

		// For each diagnostic report
		for j := 0; j < len(diagnosticReportsForCO2Scrubber); j++ {
			// Get the bit at the current position
			bit := diagnosticReportsForCO2Scrubber[j][i]

			if bit == "0" {
				countOf0++
			} else {
				countOf1++
			}
		}

		if countOf0 < countOf1 {
			chosenBitForCO2Scrubber = "0"
		}
		if countOf1 < countOf0 {
			chosenBitForCO2Scrubber = "1"
		}
		if countOf0 == countOf1 {
			chosenBitForCO2Scrubber = "0"
		}

		newdiagnosticReportsForCO2Scrubber := make([][]string, 0)
		for j := 0; j < len(diagnosticReportsForCO2Scrubber); j++ {
			// Get the bit at the current position
			bit := diagnosticReportsForCO2Scrubber[j][i]
			if bit == chosenBitForCO2Scrubber {
				// Add the diagnostic report
				newdiagnosticReportsForCO2Scrubber = append(newdiagnosticReportsForCO2Scrubber, diagnosticReportsForCO2Scrubber[j])
			}
		}
		diagnosticReportsForCO2Scrubber = newdiagnosticReportsForCO2Scrubber
		if len(diagnosticReportsForCO2Scrubber) == 1 {
			break
		}
	}

	binaryOxygenGeneratorRating := strings.Join(diagnosticReportsForOxygenGenerator[0], "")
	binaryCO2ScrubberRating := strings.Join(diagnosticReportsForCO2Scrubber[0], "")

	oxygenGeneratorRating, err := strconv.ParseInt(binaryOxygenGeneratorRating, 2, 64)
	if err != nil {
		panic(err)
	}
	CO2ScrubberRating, err := strconv.ParseInt(binaryCO2ScrubberRating, 2, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println("life support rating of the submarine:", oxygenGeneratorRating*CO2ScrubberRating)

}

func main() {
	part1(utils.RunParams{Sample: true})
	part1(utils.RunParams{Sample: false})
	part2(utils.RunParams{Sample: true})
	part2(utils.RunParams{Sample: false})
}

func getInput(params utils.RunParams) [][]string {
	lines, err := utils.ReadFileToString("03", params)
	if err != nil {
		panic(err)
	}

	diagnosticReport := make([][]string, len(lines))
	for i := range lines {
		bits := strings.Split(lines[i], "")

		diagnosticReport[i] = append(diagnosticReport[i], bits...)
	}

	return diagnosticReport
}
