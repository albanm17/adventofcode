package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile(filename string) string {
	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return string(file)
}

func partOne(s string) int {
	max := 0

	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] >= '0' && s[i] <= '9' && s[j] >= '0' && s[j] <= '9' {
				d1 := int(s[i] - '0')
				d2 := int(s[j] - '0')
				v := d1*10 + d2
				if v > max {
					max = v
				}
			}
		}
	}
	return max
}

func partTwo(s string) int64 {
	result := ""
	start := 0
	max := 12

	for max > 0 {
		maxDigit := byte('0' - 1)
		maxId := start

		for i := start; i <= len(s)-max; i++ {
			if s[i] > maxDigit {
				maxDigit = s[i]
				maxId = i
			}
		}

		result += string(maxDigit)
		start = maxId + 1
		max--
	}

	num, _ := strconv.ParseInt(result, 10, 64)
	return num
}

func main() {
	data := readFile("input.txt")

	lines := strings.Split(data, "\n")

	total := 0
	var total2 int64 = 0

	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		total += partOne(l)
		total2 += partTwo(l)
	}

	fmt.Println("RESULTAT ⭐:", total)
	fmt.Println("RESULTAT ⭐⭐:", total2)
}
