package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const target = 2020

// DayTwo runs day two
func DayTwo() {
	file, err := os.Open("../internal/day-2-input-test.txt")
	check(err)
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	// count := 0
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		check(err)
		fmt.Println("ayyy", val)
	}

	err = file.Close()

	check(err)

	fmt.Println("No value found")
	os.Exit(1)
}
