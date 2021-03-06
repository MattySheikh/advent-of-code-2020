package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

// DayTwo runs day two
func DayTwo() {
	file, err := os.Open("../internal/day-2-input.txt")
	check(err)
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	count := 0
	for scanner.Scan() {
		split := strings.Fields(scanner.Text())
		numVals, ruleWithColon, str :=  split[0], split[1], split[2]
		numValsSplit := strings.Split(numVals, "-")
		min, _ := strconv.Atoi(numValsSplit[0])
		max, _ := strconv.Atoi(numValsSplit[1])

		rule := strings.Replace(ruleWithColon, ":", "", -1)
		numOccurences := strings.Count(str, rule)

		if numOccurences >= min && numOccurences <= max {
			count++
		}
	}

	err = file.Close()

	check(err)

	fmt.Println("Num valid", count)
	os.Exit(1)
}
