package day3

import (
	"advent2022/utils"
	"fmt"
	"os"
	"unicode"
)

type backpack struct {
	left  string
	right string
	total int
}

func StartDay() {
	filePath := "./day3/input.txt"
	f, _ := os.Open(filePath)
	defer f.Close()

	Part1Solution(f)

	f, _ = os.Open(filePath)
	defer f.Close()
	Part2Solution(f)

	return
}

func Part1Solution(f *os.File) {
	lines := utils.ReadFileAsLines(f)

	priorityCount := 0
	for _, line := range lines {

		var bp backpack

		halfLength := len(line) / 2
		bp.left = line[:halfLength]
		bp.right = line[halfLength:]

		intersections := intersection(bp.left, bp.right)

		for _, c := range intersections {
			priorityCount += toValue(c)
		}
	}

	fmt.Println(priorityCount)
}

func Part2Solution(f *os.File) {
	lines := utils.ReadFileAsLines(f)

	elf1, elf2, elf3 := "", "", ""
	priorityCount := 0
	groupCount := 0
	for _, line := range lines {
		if groupCount == 0 {
			elf1 = line
		}
		if groupCount == 1 {
			elf2 = line
		}
		if groupCount == 2 {
			elf3 = line
		}

		if elf1 != "" && elf2 != "" && elf3 != "" {
			intersections := intersectionGroup(elf1, elf2, elf3)

			for _, c := range intersections {
				priorityCount += toValue(c)
			}

			elf1 = ""
			elf2 = ""
			elf3 = ""
			groupCount = 0
		} else {
			groupCount++
		}

	}
	fmt.Println(priorityCount)
}
func toValue(char rune) int {
	if unicode.IsLower(char) {
		return int(char) - 96
	} else {
		return int(char) - (65 - 27)
	}
}

func intersection(a, b string) []rune {
	set := make(map[rune]bool)
	for _, c := range a {
		for _, c2 := range b {
			if c2 == c {
				set[c] = true
			}
		}
	}

	var keys []rune
	for r, _ := range set {
		keys = append(keys, r)
	}

	return keys
}

func intersectionGroup(a, b, d string) []rune {
	set := make(map[rune]bool)
	for _, c := range a {
		for _, c2 := range b {
			for _, c3 := range d {
				if c == c2 && c2 == c3 {
					set[c] = true
				}
			}
		}
	}

	var keys []rune
	for r, _ := range set {
		keys = append(keys, r)
	}

	return keys
}
