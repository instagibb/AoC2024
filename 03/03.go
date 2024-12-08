package main

import (
	"fmt"
	"os"
	"regexp"
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

	corruptedMemory := string(text)

	// Regex to match mul(x,y)
	mulRegex := regexp.MustCompile(`mul\(\d*,\d*\)+`)
	var allMuls = mulRegex.FindAllString(corruptedMemory, -1)

	digitRegex := regexp.MustCompile(`\d+,\d+`)
	multiplied := make([]int, 0)
	for _, mul := range allMuls {
		pair := s.Split(digitRegex.FindString(mul), ",")
		x, _ := strconv.Atoi(pair[0])
		y, _ := strconv.Atoi(pair[1])
		multiplied = append(multiplied, x*y)
	}

	sum := 0
	for _, m := range multiplied {
		sum += m
	}
	fmt.Println(sum)

	// Part 2
	cleanedMemory := getClean(corruptedMemory, "don't()")
	fmt.Println(cleanedMemory)

	var allCleanedMuls = mulRegex.FindAllString(cleanedMemory, -1)
	//fmt.Println(allCleanedMuls)

	cleanMultiplied := make([]int, 0)
	for _, mul := range allCleanedMuls {
		pair := s.Split(digitRegex.FindString(mul), ",")
		x, _ := strconv.Atoi(pair[0])
		y, _ := strconv.Atoi(pair[1])
		cleanMultiplied = append(cleanMultiplied, x*y)
	}
	cleanSum := 0
	for _, m := range cleanMultiplied {
		cleanSum += m
	}
	fmt.Println(cleanSum)
}

func getClean(corruptedMemory string, target string) string {
	// Find target index
	cleanedMemory := ""
	targetIndex := s.Index(corruptedMemory, target)

	if targetIndex == -1 {
		return cleanedMemory + corruptedMemory
	}

	// Get the substring up to the index of the target
	if target == "don't()" {
		cleanedMemory += corruptedMemory[:targetIndex] + getClean(corruptedMemory[targetIndex:], "do()")
	} else {
		cleanedMemory += getClean(corruptedMemory[targetIndex+4:], "don't()")
	}
	return cleanedMemory
}
