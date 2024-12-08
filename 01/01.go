package main

import (
	"fmt"
	"os"
	"slices"
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
	var lines = s.Split(string(text), "\n")
	var first = make([]string, len(lines))
	var second = make([]string, len(lines))
	for i := 0; i < len(lines); i++ {
		var split = s.Split(lines[i], "   ")
		if len(split) != 2 {
			continue
		}
		first[i] = split[0]
		second[i] = split[1]
	}
	//fmt.Println(first)
	//fmt.Println(second)

	// Order first, second arrays
	slices.Sort(first)
	slices.Sort(second)

	//fmt.Println(first)
	//fmt.Println(second)

	var differences = make([]int, len(lines))
	// Match first and second arrays
	for i := 0; i < len(first); i++ {
		//var diff = first[i], second[i])
		var f, _ = strconv.Atoi(first[i])
		var s, _ = strconv.Atoi(second[i])

		if f == s {
			differences[i] = 0
		} else if first[i] > second[i] {
			differences[i] = f - s
		} else {
			differences[i] = s - f
		}
	}

	// Add up differences
	sum := 0
	for _, diff := range differences {
		sum += diff
	}
	fmt.Println(sum)

	// Part 2
	var similarity = make([]int, len(lines))
	for i := 0; i < len(first); i++ {
		var f = first[i]

		if slices.Contains(second, f) {
			// get the count of times f appears in second
			count := 0
			for _, test := range second {
				if test == f {
					count += 1
				}
			}
			var fInt, _ = strconv.Atoi(f)
			similarity[i] = fInt * count
		}

	}
	simsum := 0
	for _, sim := range similarity {
		simsum += sim
	}
	fmt.Println(simsum)
}
