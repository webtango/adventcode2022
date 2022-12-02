package day2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StartDay() {
	filePath := "./day2/input.txt"
	f, _ := os.Open(filePath)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	//Rock (A), Paper (Y)
	//Paper (B), Rock (X)
	//Scissors (C), Scissors (Z)
	//Rock worth 1
	//Paper worth 2
	//Scissors worth 3
	var totalScore int64 = 0
	for scanner.Scan() {
		line := scanner.Text()

		rps := strings.Split(line, " ")

		if rps[1] == "Y" {
			totalScore = totalScore + 2

			if rps[0] == "A" {
				totalScore = totalScore + 6
			}

			if rps[0] == "B" {
				totalScore = totalScore + 3
			}
		} else if rps[1] == "X" {
			totalScore = totalScore + 1

			if rps[0] == "C" {
				totalScore = totalScore + 6
			}

			if rps[0] == "A" {
				totalScore = totalScore + 3
			}
		} else if rps[1] == "Z" {
			totalScore = totalScore + 3

			if rps[0] == "B" {
				totalScore = totalScore + 6
			}

			if rps[0] == "C" {
				totalScore = totalScore + 3
			}
		}
	}

	//part 2
	//Rock (A), Paper (Y)
	//Paper (B), Rock (X)
	//Scissors (C), Scissors (Z)
	//Rock worth 1
	//Paper worth 2
	//Scissors worth 3
	totalScore = 0
	f, _ = os.Open(filePath)
	defer f.Close()
	scanner = bufio.NewScanner(f)
	//X means you need to lose add zero, Y means you need to end the round in a draw add 3, and Z means you need to win add 6
	for scanner.Scan() {
		line := scanner.Text()

		rps := strings.Split(line, " ")

		if rps[1] == "X" {

			if rps[0] == "A" {
				totalScore = totalScore + 3
			} else if rps[0] == "B" {
				totalScore = totalScore + 1
			} else if rps[0] == "C" {
				totalScore = totalScore + 2
			}
		} else if rps[1] == "Y" {
			totalScore = totalScore + 3

			if rps[0] == "A" {
				totalScore = totalScore + 1
			} else if rps[0] == "B" {
				totalScore = totalScore + 2
			} else if rps[0] == "C" {
				totalScore = totalScore + 3
			}
		} else if rps[1] == "Z" {
			totalScore = totalScore + 6

			if rps[0] == "A" {
				totalScore = totalScore + 2
			} else if rps[0] == "B" {
				totalScore = totalScore + 3
			} else if rps[0] == "C" {
				totalScore = totalScore + 1
			}
		}
	}

	fmt.Println(totalScore)
	return
}
