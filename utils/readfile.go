package utils

import (
	"bufio"
	"fmt"
	"os"
)

type RunParams struct {
	Sample bool
}

func ReadFileToString(day string, params RunParams) ([]string, error) {
	var filename string

	if params.Sample {
		filename = fmt.Sprintf("%s/%s_sample.txt", day, day)
	} else {
		filename = fmt.Sprintf("%s/%s.txt", day, day)
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines := []string{}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}
