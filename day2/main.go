package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile(filename string) []string {
	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), ",")

	return lines
}

func isInvalidPart1(n string) bool {
	length := len(n)

	if length%2 != 0 {
		return false
	}

	mid := length / 2
	firstHalf := n[:mid]
	secondHalf := n[mid:]

	return firstHalf == secondHalf && firstHalf[0] != '0'
}

func isInvalidPart2(n string) bool {
	length := len(n)

	for patternLen := 1; patternLen <= length/2; patternLen++ {
		if length%patternLen != 0 {
			continue
		}
		pattern := n[:patternLen]

		if pattern[0] == '0' {
			continue
		}

		valid := true
		for i := patternLen; i < length; i += patternLen {
			if n[i:i+patternLen] != pattern {
				valid = false
				break
			}
		}
		if valid {
			return true
		}
	}
	return false
}

func findInvalidPart1(start int, end int) []int {
	invalid := []int{}
	for num := start; num <= end; num++ {
		numStr := strconv.Itoa(num)
		if isInvalidPart1(numStr) {
			invalid = append(invalid, num)
		}
	}
	return invalid
}

func findInvalidPart2(start int, end int) []int {
	invalid := []int{}
	for num := start; num <= end; num++ {
		numStr := strconv.Itoa(num)
		if isInvalidPart2(numStr) {
			invalid = append(invalid, num)
		}
	}
	return invalid
}

func main() {
	file := readFile("input.txt")
	totalPart1 := 0
	totalPart2 := 0

	for _, l := range file {
		part := strings.Split(l, "-")
		start, _ := strconv.Atoi(strings.TrimSpace(part[0]))
		end, _ := strconv.Atoi(strings.TrimSpace(part[1]))

		invalidIdsPart1 := findInvalidPart1(start, end)
		for _, id := range invalidIdsPart1 {
			totalPart1 += id
		}

		invalidIdsPart2 := findInvalidPart2(start, end)
		for _, id := range invalidIdsPart2 {
			totalPart2 += id
		}
	}

	fmt.Println("RESULTAT ⭐:", totalPart1)
	fmt.Println("RESULTAT ⭐⭐:", totalPart2)
}
