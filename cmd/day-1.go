package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func checkDayOne(e error) {
	if e != nil {
		panic(e)
	}
}

const dayOneTarget = 2020

// DayOne runs the 1st part of day 1
func DayOne() {
	file, err := os.Open("../internal/day-1-input.txt")
	check(err)
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var diffMap = map[int]bool{}

	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		check(err)
		diffMap[val] = true
	}

	// Get the keys from the map to a string slice
	allValues := make([]int, 0, len(diffMap))
	for key := range diffMap {
		allValues = append(allValues, key)
	}

	for _, currVal := range allValues {
		diffValue := dayOneTarget - currVal

		if diffMap[diffValue] {
			fmt.Println("Found value", currVal * diffValue)

			os.Exit(0)
		}
	}

	err = file.Close()

	check(err)

	fmt.Println("No value found")
	os.Exit(1)
}
