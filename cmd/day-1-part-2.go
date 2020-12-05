package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const dayTwoTarget = 2020

func checkDayTwo(e error) {
	if e != nil {
		panic(e)
	}
}

func getTwoValues(currTarget int, values []int) (int, int) {
	var diffMap = map[int]bool{}
	for _, currVal := range values {
		diffMap[currVal] = true
	}

	for _, currVal := range values {
		diffVal := currTarget - currVal

		if diffMap[diffVal] {
			return currVal, diffVal
		}
	}

	return 0, 0
}

// DayOnePartTwo runs the 2nd part of day 1
func DayOnePartTwo() {
	file, err := os.Open("../internal/day-1-input.txt")
	checkDayTwo(err)
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var diffMap = map[int]bool{}

	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		checkDayTwo(err)
		diffMap[val] = true
	}

	// Get the keys from the map to a string slice
	allValues := make([]int, 0, len(diffMap))
	for key := range diffMap {
		allValues = append(allValues, key)
	}

	for _, currVal := range allValues {
		currTarget := dayTwoTarget - currVal
		first, second := getTwoValues(currTarget, allValues)

		if first == 0 && second == 0 {
			continue
		}

		diffVal := dayTwoTarget - first - second

		if diffMap[diffVal] {
			fmt.Println("Found values", currVal  * first  * second)

			os.Exit(0)
		}
	}

	err = file.Close()

	checkDayTwo(err)

	fmt.Println("No value found")
	os.Exit(1)
}