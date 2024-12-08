package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
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

	// Regex to match mul(x,y)
	mulRegex := regexp.MustCompile(`mul\(\d*,\d*\)+`)
	var allMuls = mulRegex.FindAllString(string(text), -1)

	digitRegex := regexp.MustCompile(`\d+,\d+`)
	multiplied := make([]int, 0)
	for _, mul := range allMuls {
		pair := strings.Split(digitRegex.FindString(mul), ",")
		x, _ := strconv.Atoi(pair[0])
		y, _ := strconv.Atoi(pair[1])
		multiplied = append(multiplied, x*y)
	}

	sum := 0
	for _, m := range multiplied {
		sum += m
	}
	fmt.Println(sum)

}
