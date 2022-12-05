package day4

import (
	"advent2022/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type backpack struct {
	left  string
	right string
	total int
}

func StartDay() {
	filePath := "./day4/input.txt"
	f, _ := os.Open(filePath)
	defer f.Close()

	Part1Solution(f)

	f, _ = os.Open(filePath)
	defer f.Close()

	Part2Solution(f)

	return
}

func Part1Solution(f *os.File) {

	contains := 0
	lines := utils.ReadFileAsLines(f)

	for _, line := range lines {
		elfList := strings.Split(line, ",")

		elf1FirstLast := strings.Split(elfList[0], "-")
		elf2FirstLast := strings.Split(elfList[1], "-")

		elf1Low, _ := strconv.Atoi(elf1FirstLast[0])
		elf1High, _ := strconv.Atoi(elf1FirstLast[1])
		elf2Low, _ := strconv.Atoi(elf2FirstLast[0])
		elf2High, _ := strconv.Atoi(elf2FirstLast[1])

		if elf1Low <= elf2Low && elf1High >= elf2High {
			contains++
		} else if elf2Low <= elf1Low && elf2High >= elf1High {
			contains++
		}
	}

	fmt.Println("Part 1 solution: ", contains)
}

func Part2Solution(f *os.File) {

	contains := 0
	lines := utils.ReadFileAsLines(f)

	for _, line := range lines {
		elfList := strings.Split(line, ",")

		elf1FirstLast := strings.Split(elfList[0], "-")
		elf2FirstLast := strings.Split(elfList[1], "-")

		elf1Low, _ := strconv.Atoi(elf1FirstLast[0])
		elf1High, _ := strconv.Atoi(elf1FirstLast[1])
		elf2Low, _ := strconv.Atoi(elf2FirstLast[0])
		elf2High, _ := strconv.Atoi(elf2FirstLast[1])

		if (elf2High >= elf1Low && elf2Low <= elf1High) ||
			(elf1High >= elf2Low && elf1Low <= elf2High) {
			contains++
		}
	}

	fmt.Println("Part 2 solution: ", contains)
}
