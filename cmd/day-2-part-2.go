package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const target = 2020

// DayTwoPartTwo runs day two part two
func DayTwoPartTwo() {
	file, err := os.Open("../internal/day-2-input.txt")
	check(err)
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	count := 0
	for scanner.Scan() {
		split := strings.Fields(scanner.Text())
		numVals, letterWithColon, str := split[0], split[1], split[2]
		numValsSplit := strings.Split(numVals, "-")
		min, _ := strconv.Atoi(numValsSplit[0])
		max, _ := strconv.Atoi(numValsSplit[1])

		letter := strings.Replace(letterWithColon, ":", "", -1)

		if (string(str[min - 1]) == letter) {
			if (string(str[max - 1]) != letter) {
				count++
			}
		} else if (string(str[max - 1]) == letter) {
			count++
		}
	}

	err = file.Close()

	check(err)

	fmt.Println("Num valid", count)
	os.Exit(1)
}
