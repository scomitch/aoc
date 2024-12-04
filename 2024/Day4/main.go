package main

import (
	"bufio"
	"fmt"
	"os"
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
	inputFile := "input.input"
	lines, err := readFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return 0, 0
	}

	word := "XMAS"
	countFind := 0

	// L2R
	rows := len(lines)
	cols := len(lines[0])

	// Janky way of doing it, on reflection could probably have used a crawl method to check all position matrix directions when X is found.
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			// l2R
			if col+len(word) <= cols {
				match := true
				for i := 0; i < len(word); i++ {
					if lines[row][col+i] != word[i] {
						match = false
						break
					}
				}
				if match {
					countFind++
				}
			}

			// R2L
			if col-len(word)+1 >= 0 {
				match := true
				for i := 0; i < len(word); i++ {
					if lines[row][col-i] != word[i] {
						match = false
						break
					}
				}
				if match {
					countFind++
				}
			}

			// Row+1
			if row+len(word) <= rows {
				match := true
				for i := 0; i < len(word); i++ {
					if lines[row+i][col] != word[i] {
						match = false
						break
					}
				}
				if match {
					countFind++
				}
			}

			// row-1
			if row-len(word)+1 >= 0 {
				match := true
				for i := 0; i < len(word); i++ {
					if lines[row-i][col] != word[i] {
						match = false
						break
					}
				}
				if match {
					countFind++
				}
			}

			// col+1 row+1
			if row+len(word) <= rows && col+len(word) <= cols {
				match := true
				for i := 0; i < len(word); i++ {
					if lines[row+i][col+i] != word[i] {
						match = false
						break
					}
				}
				if match {
					countFind++
				}
			}

			// row-1 col-1
			if row-len(word)+1 >= 0 && col-len(word)+1 >= 0 {
				match := true
				for i := 0; i < len(word); i++ {
					if lines[row-i][col-i] != word[i] {
						match = false
						break
					}
				}
				if match {
					countFind++
				}
			}

			// row+1 col-1
			if row+len(word) <= rows && col-len(word)+1 >= 0 {
				match := true
				for i := 0; i < len(word); i++ {
					if lines[row+i][col-i] != word[i] {
						match = false
						break
					}
				}
				if match {
					countFind++
				}
			}

			// row-1 col+1
			if row-len(word)+1 >= 0 && col+len(word) <= cols {
				match := true
				for i := 0; i < len(word); i++ {
					if lines[row-i][col+i] != word[i] {
						match = false
						break
					}
				}
				if match {
					countFind++
				}
			}
		}
	}

	fmt.Println(countFind)

	return 0, 0
}

func day1p2() (int, int) {
	inputFile := "input.input"
	lines, err := readFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return 0, 0
	}

	//word := "MAS"
	countFind := 0

	rows := len(lines)
	cols := len(lines[0])

	// Could probably use a better crawl here, to revise
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			// Ignore A check on first and last row.
			if row == 0 || row == rows-1 {
				continue
			}

			// A is always in the middle. We expect S or M to either be top left or top right in either order
			if lines[row][col] == 'A' {
				if row-1 >= 0 && row+1 < rows && col-1 >= 0 && col+1 < cols {
					if (lines[row-1][col-1] == 'S' && lines[row+1][col+1] == 'M') && (lines[row-1][col+1] == 'M' && lines[row+1][col-1] == 'S') {
						fmt.Println("FOUND MAS A", "at", row, col)
						countFind++
					}
					if (lines[row-1][col-1] == 'M' && lines[row+1][col+1] == 'S') && (lines[row-1][col+1] == 'M' && lines[row+1][col-1] == 'S') {
						fmt.Println("FOUND MAS A", "at", row, col)
						countFind++
					}
					if (lines[row-1][col-1] == 'S' && lines[row+1][col+1] == 'M') && (lines[row-1][col+1] == 'S' && lines[row+1][col-1] == 'M') {
						fmt.Println("FOUND MAS A", "at", row, col)
						countFind++
					}
					if (lines[row-1][col-1] == 'M' && lines[row+1][col+1] == 'S') && (lines[row-1][col+1] == 'S' && lines[row+1][col-1] == 'M') {
						fmt.Println("FOUND MAS A", "at", row, col)
						countFind++
					}
				}
			}
		}
	}

	fmt.Println(countFind)

	return 0, 0
}

func main() {
	//day1p1()
	day1p2()
}
