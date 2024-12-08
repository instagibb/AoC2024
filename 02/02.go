package main

import (
	"fmt"
	"os"
	"strconv"
	s "strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Break up data
	text, err := os.ReadFile("./input.txt")
	check(err)
	var reports = s.Split(string(text), "\n")

	// Safe count
	safe := 0

	// Check each report
	for _, report := range reports {
		if len(report) == 0 {
			continue
		}
		if checkReport(s.Split(report, " ")) {
			safe += 1
		}
	}

	// Print safe count
	fmt.Println(safe)
}

func checkReport(report []string) bool {
	//fmt.Println("Checking report", report)
	var first, _ = strconv.Atoi(report[0])
	var second, _ = strconv.Atoi(report[1])
	// Bail early if first and second are the same
	if first == second {
		return false
	}

	// Check if the report is incrementing or decrementing
	isIncrementing := first < second

	var valid = false
	for i := 0; i < len(report)-1; i++ {
		if i+1 == len(report) {
			break
		}

		var f, _ = strconv.Atoi(report[i])
		var s, _ = strconv.Atoi(report[i+1])

		if isIncrementing {
			valid = checkIncrementing(f, s)
		} else {
			valid = checkDecrementing(f, s)
		}

		if !valid {
			return false
		}
	}

	//fmt.Println("Report", report, valid)
	return valid
}

func checkIncrementing(first int, second int) bool {
	//fmt.Println("Checking incrementing", first, second)
	if first > second {
		return false
	}

	if first == second {
		return false
	}

	if second-first > 3 {
		return false
	}

	return true
}

func checkDecrementing(first int, second int) bool {
	//fmt.Println("Checking decrementing", first, second)
	if first < second {
		return false
	}

	if first == second {
		return false
	}

	if first-second > 3 {
		return false
	}

	return true
}
