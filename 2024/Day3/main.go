package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func day1p1() (int, int) {
	inputFile := "input.input"
	line, err := readFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return 0, 0
	}

	// Regex match to find all mul x,y
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(line, -1)
	total := 0
	for _, match := range matches {
		// get the numbers
		fmt.Println(match)
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		fmt.Println(num1, num2)
		total += num1 * num2
	}

	fmt.Println(total)

	return 0, 0
}

func day1p2() (int, int) {
	inputFile := "input.input"
	line, err := readFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return 0, 0
	}
	fmt.Println(line[1])

	total := 0
	split := strings.Split(line, "mul(")
	calculate := true

	for i := 1; i < len(split); i++ {
		part := split[i]
		end := strings.Index(part, ")")

		if end == -1 {
			continue
		}

		nums := part[:end]
		numPair := strings.Split(nums, ",")
		if len(numPair) != 2 {
			continue
		}

		beforeMul := split[i-1]

		if strings.Contains(beforeMul, "don't()") {
			calculate = false
		} else if strings.Contains(beforeMul, "do()") {
			calculate = true
		}

		if calculate {
			first, _ := strconv.Atoi(numPair[0])
			last, _ := strconv.Atoi(numPair[1])
			fmt.Println(first, last, "< CALCD NUMBERS")
			total += first * last
		}
	}

	fmt.Println(total)

	return 0, 0
}

func main() {
	//day1p1()
	day1p2()
}
