package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func readFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, scanner.Err()
}

func day1p1() (int, int) {
	inputFile := "d2p1.input"
	lines, err := readFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return 0, 0
	}

	var safe int = 0

	for _, line := range lines {
		splitLine := strings.Split(line, " ")
		isSafe := true
		isIncreasing := true
		isDecreasing := true

		for i := 1; i < len(splitLine); i++ {
			num, _ := strconv.Atoi(splitLine[i])
			numBefore, _ := strconv.Atoi(splitLine[i-1])

			if num <= numBefore {
				isIncreasing = false
			}
			if num >= numBefore {
				isDecreasing = false
			}

			if math.Abs(float64(num-numBefore)) < 1 || math.Abs(float64(num-numBefore)) > 3 {
				isSafe = false
				break
			}
		}
		if isSafe && (isIncreasing || isDecreasing) {
			safe++
		}
	}

	fmt.Println(safe)

	return 0, 0
}

func day1p2() (int, int) {
	inputFile := "d2p2.input"
	lines, err := readFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return 0, 0
	}

	var safe int = 0

	for _, line := range lines {
		splitLine := strings.Split(line, " ")
		isSafe := baselineSafeCheck(splitLine)

		if !isSafe {
			for i := 0; i < len(splitLine); i++ {
				splitLineCopy := make([]string, len(splitLine)-1)
				copy(splitLineCopy, splitLine[:i])
				copy(splitLineCopy[i:], splitLine[i+1:])
				if baselineSafeCheck(splitLineCopy) {
					isSafe = true
					break
				}
			}
		}
		if isSafe {
			safe++
		}
	}

	fmt.Println(safe)

	return 0, 0
}

func baselineSafeCheck(splitLine []string) bool {
	isIncreasing := true
	isDecreasing := true

	for i := 1; i < len(splitLine); i++ {
		num, _ := strconv.Atoi(splitLine[i])
		numBefore, _ := strconv.Atoi(splitLine[i-1])

		if num <= numBefore {
			isIncreasing = false
		}
		if num >= numBefore {
			isDecreasing = false
		}

		if math.Abs(float64(num-numBefore)) < 1 || math.Abs(float64(num-numBefore)) > 3 {
			return false
		}
	}
	return isIncreasing || isDecreasing
}

func main() {
	//day1p1()
	day1p2()
}
