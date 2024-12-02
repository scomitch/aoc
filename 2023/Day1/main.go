package main

import (
	"bufio"
	"fmt"
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
	inputFile := "d1p1.input"
	lines, err := readFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return 0, 0
	}

	var numbers string = "0123456789"
	var sum int = 0
	for _, line := range lines {
		var tens, ones int = 0, 0

		for i := 0; i < len(line); i++ {
			if strings.Contains(numbers, string(line[i])) {
				tens, _ = strconv.Atoi(string(line[i]))
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if strings.Contains(numbers, string(line[i])) {
				ones, _ = strconv.Atoi(string(line[i]))
				break
			}
		}

		sum += tens*10 + ones
		fmt.Println(sum, tens, ones)
	}

	fmt.Println(sum)

	return 0, 0
}

func day1p2() (int, int) {
	inputFile := "d1p2.input"
	lines, err := readFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return 0, 0
	}

	var numbers string = "0123456789"
	acceptedNumbers := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"zero":  0,
	}
	var sum int = 0
	for _, line := range lines {
		var tens, ones int = 0, 0

		for i := 0; i < len(line); i++ {
			for word, num := range acceptedNumbers {
				if strings.HasPrefix(line[i:], word) {
					tens = num
					break
				}
			}

			if tens != 0 {
				break
			}

			if strings.Contains(numbers, string(line[i])) {
				tens, _ = strconv.Atoi(string(line[i]))
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			for word, num := range acceptedNumbers {
				if strings.HasPrefix(line[i+1:], word) {
					ones = num
					break
				}
			}

			if ones != 0 {
				break
			}

			if strings.Contains(numbers, string(line[i])) {
				ones, _ = strconv.Atoi(string(line[i]))
				break
			}
		}

		sum += tens*10 + ones
		fmt.Println(sum, tens, ones)
	}

	fmt.Println(sum)

	return 0, 0
}

func main() {
	//day1p1()
	day1p2()
}
