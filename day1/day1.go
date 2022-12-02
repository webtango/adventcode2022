package day1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func StartDay() {
	filePath := "./day1/input.txt"
	f, _ := os.Open(filePath)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var topCalories []int64
	var elfCalories int64 = 0
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			topCalories = append(topCalories, elfCalories)
			elfCalories = 0
			continue
		}
		val, _ := strconv.ParseInt(line, 10, 0)
		elfCalories = elfCalories + val
	}

	sort.Slice(topCalories, func(i, j int) bool {
		return topCalories[i] > topCalories[j]
	})

	fmt.Println("Top elf calories: ", topCalories[0])
	top3 := topCalories[0] + topCalories[1] + topCalories[2]
	fmt.Println("Combined top 3 calories: ", top3)
	return
}
