package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	// parse the smallest number from the left list and the largest number from the right list, put them in a list from smallest to largest
	var leftListSorted []int
	var rightListSorted []int

	for _, line := range lines {
		// Take the fist number, put it in the left list. Take the last number, put it in the right list.
		splitLine := strings.Split(line, "   ")
		num1, _ := strconv.Atoi(splitLine[0])
		num2, _ := strconv.Atoi(splitLine[1])
		leftListSorted = append(leftListSorted, num1)
		rightListSorted = append(rightListSorted, num2)
	}

	// Sort both lists from smallest to largest
	sort.Ints(leftListSorted)
	sort.Ints(rightListSorted)

	// Size of each list
	leftListSize := len(leftListSorted)
	rightListSize := len(rightListSorted)
	fmt.Println(leftListSize, rightListSize)

	// Find the numbered difference between each index of the two lists
	// Ensure the difference is not < 0
	var differenceList []int
	for i := 0; i < len(leftListSorted); i++ {
		if rightListSorted[i]-leftListSorted[i] < 0 {
			differenceList = append(differenceList, leftListSorted[i]-rightListSorted[i])
		} else {
			differenceList = append(differenceList, rightListSorted[i]-leftListSorted[i])
		}
	}

	fmt.Println(differenceList)

	// Add all the numbers in the difference list together.
	sum := 0
	for _, num := range differenceList {
		sum += num
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

	// parse the smallest number from the left list and the largest number from the right list, put them in a list from smallest to largest
	var leftListSorted []int
	var rightListSorted []int

	for _, line := range lines {
		// Take the fist number, put it in the left list. Take the last number, put it in the right list.
		splitLine := strings.Split(line, "   ")
		num1, _ := strconv.Atoi(splitLine[0])
		num2, _ := strconv.Atoi(splitLine[1])
		leftListSorted = append(leftListSorted, num1)
		rightListSorted = append(rightListSorted, num2)
	}

	// Sort both lists from smallest to largest
	sort.Ints(leftListSorted)
	sort.Ints(rightListSorted)

	// Size of each list
	leftListSize := len(leftListSorted)
	rightListSize := len(rightListSorted)
	fmt.Println(leftListSize, rightListSize)

	// Check first entry in left list, store number. Count how many times this number appears in the right list.
	// Multiply left list value with count of right list value.
	finalTotal := 0
	for i := 0; i < len(leftListSorted); i++ {
		count := 0
		for j := 0; j < len(rightListSorted); j++ {
			if rightListSorted[j] == leftListSorted[i] {
				count++
			}
		}
		fmt.Println("Count of ", leftListSorted[i], " is ", count)
		finalTotal += leftListSorted[i] * count
	}

	fmt.Println("Total is ", finalTotal)

	return 0, 0
}

func main() {
	day1p1()
	day1p2()
}
