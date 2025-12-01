package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func readFile(filename string) ([]string) {
	file, err := os.ReadFile(filename)
	if(err != nil) {
		panic(err)
	}

	lines := strings.Split(string(file), "\n")

	return lines
}

func partOne(lines []string) (int32) {
	start := 50
	count := 0

	for _, line := range lines {
		direction := string(line[0])
		distance, _ := strconv.Atoi(line[1:])
		
		if direction == "L" {
			start -= distance
		} else if direction == "R" {
			start += distance
		}

		start = ((start % 100) + 100) % 100

		if start == 0 {
			count++
		}
	}

	return int32(count)
}

func partTwo(lines []string) (int32) {
	start := 50
	count := 0

	for _, line := range lines {
		direction := string(line[0])
		distance, _ := strconv.Atoi(line[1:])
		
		if direction == "L" {
			if start == 0 {
				count += (distance - 1) / 100
			} else {
				bonus := 0
                if distance%100 >= start {
                    bonus = 1
                }
                count += distance/100 + bonus
			}
			start -= distance
		} else if direction == "R" {
			count += (start + distance)/100 - start/100
			start += distance
		}

		start = ((start % 100) + 100) % 100
	}

	return int32(count)
}

func main() {
	lines := readFile("input.txt")

	resultatPartOne := partOne(lines)
	resultatPartTwo := partTwo(lines)

	fmt.Println("RESULTAT ⭐:", resultatPartOne)
	fmt.Println("RESULTAT ⭐⭐:", resultatPartTwo)
}