package day5

import (
	"advent2022/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type craneMoves struct {
	start  int
	end    int
	crates int
}

func StartDay() {
	filePath := "./day5/input.txt"
	f, _ := os.Open(filePath)
	lines := utils.ReadFileAsLines(f)
	defer f.Close()

	Part1Solution(lines)

	Part2Solution(lines)

	return
}

func Part1Solution(lines []string) {

	stacks := map[int]string{}
	var movesList []craneMoves

	for i, line := range lines {
		if i < 8 {
			col := 0
			for j := 1; j < len(line); j = j + 4 {
				stacks[col] += string(line[j])
				col++
			}
		}

		if i > 9 {
			lineSet := strings.Split(line, " ")
			var moveLine craneMoves
			val, _ := strconv.Atoi(lineSet[3])
			moveLine.start = val - 1
			val2, _ := strconv.Atoi(lineSet[5])
			moveLine.end = val2 - 1
			val3, _ := strconv.Atoi(lineSet[1])
			moveLine.crates = val3

			movesList = append(movesList, moveLine)
		}
	}

	for stackNum, stack := range stacks {
		result := ""
		for _, v := range stack {
			result = string(v) + result
		}
		result = strings.Trim(result, " ")
		stacks[stackNum] = result
	}

	fmt.Println(stacks)

	for _, move := range movesList {
		doMove1AtTime(stacks, move)
	}

	tops := ""
	for x := 0; x < 9; x++ {
		stackheight := len(stacks[x])
		tops += stacks[x][stackheight-1 : stackheight]
	}

	fmt.Println("Part 1 ", tops)
}

func Part2Solution(lines []string) {
	stacks := map[int]string{}
	var movesList []craneMoves

	for i, line := range lines {
		if i < 8 {
			col := 0
			for j := 1; j < len(line); j = j + 4 {
				stacks[col] += string(line[j])
				col++
			}
		}

		if i > 9 {
			lineSet := strings.Split(line, " ")
			var moveLine craneMoves
			val, _ := strconv.Atoi(lineSet[3])
			moveLine.start = val - 1
			val2, _ := strconv.Atoi(lineSet[5])
			moveLine.end = val2 - 1
			val3, _ := strconv.Atoi(lineSet[1])
			moveLine.crates = val3

			movesList = append(movesList, moveLine)
		}
	}

	for stackNum, stack := range stacks {
		result := ""
		for _, v := range stack {
			result = string(v) + result
		}
		result = strings.Trim(result, " ")
		stacks[stackNum] = result
	}

	fmt.Println(stacks)

	for _, move := range movesList {
		doMoveAll(stacks, move)
	}

	tops := ""
	for x := 0; x < 9; x++ {
		stackheight := len(stacks[x])
		tops += stacks[x][stackheight-1 : stackheight]
	}

	fmt.Println("Part 2 ", tops)
}

func doMove1AtTime(stack map[int]string, move craneMoves) {

	//fmt.Println("start before ", move.start, stack[move.start])
	//fmt.Println("end before ", move.end, stack[move.end])

	stackheight := len(stack[move.start])
	items := stack[move.start][stackheight-move.crates : stackheight]

	result := ""
	for _, v := range items {
		result = string(v) + result
	}
	items = result

	stack[move.end] += items
	stack[move.start] = stack[move.start][:stackheight-move.crates]
	//fmt.Println("start after", stack[move.start])
	//fmt.Println("end after", stack[move.end])

}

func doMoveAll(stack map[int]string, move craneMoves) {

	//fmt.Println("start before ", move.start, stack[move.start])
	//fmt.Println("end before ", move.end, stack[move.end])

	stackheight := len(stack[move.start])
	items := stack[move.start][stackheight-move.crates : stackheight]

	stack[move.end] += items
	stack[move.start] = stack[move.start][:stackheight-move.crates]
	//fmt.Println("start after", stack[move.start])
	//fmt.Println("end after", stack[move.end])

}
