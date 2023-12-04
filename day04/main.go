package main

import (
	"aoc2023/io"
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

var day int = 4

func main() {
	fmt.Printf("Running day %d\n", day)
	//input := io.GetInput(day)
	input := io.ReadFile(day)
	//fmt.Printf("%v", input)

	answer1, err := part1(input)
	if err != nil {
		fmt.Printf("Error running part 1\n")
	} else {
		fmt.Printf("Part 1 answer: %d \n", answer1)
	}

	answer2, err := part2(input)
	if err != nil {
		fmt.Printf("Error running part 2\n")
	} else {
		fmt.Printf("Part 2 answer: %d \n", answer2)
	}
}

func PowInts(x, n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	y := PowInts(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return x * y * y
}

func part1(input []string) (int, error) {
	total := 0
	for _, row := range input {
		count := 0
		row = strings.Split(row, ": ")[1]
		lists := strings.Split(row, " | ")
		winning := []int{}
		for _, stringNumber := range strings.Split(lists[0], " ") {
			v, _ := strconv.Atoi(stringNumber)
			winning = append(winning, v)
		}

		for _, stringNumber := range strings.Split(lists[1], " ") {
			v, _ := strconv.Atoi(stringNumber)
			if slices.Contains(winning, v) && stringNumber != "" {
				//fmt.Println(v)
				count++
			}
		}
		if count > 0 {

			total += PowInts(2, count-1)
		}
		//fmt.Println(row)
		//fmt.Println(count)

	}
	return total, nil
	return 0, errors.New("Not implemented")
}

func getCountScratchcards(winningmap map[int]int, card int) int {
	if winningmap[card] == 0 {
		return 0
	}

	sum := 0
	for i := 1; i <= winningmap[card]; i++ {
		sum += getCountScratchcards(winningmap, card+i) + 1
	}
	return sum
}

func part2(input []string) (int, error) {
	winningMap := make(map[int]int)

	for card, row := range input {
		count := 0
		row = strings.Split(row, ": ")[1]
		lists := strings.Split(row, " | ")
		winning := []int{}
		for _, stringNumber := range strings.Split(lists[0], " ") {
			v, _ := strconv.Atoi(stringNumber)
			winning = append(winning, v)
		}

		for _, stringNumber := range strings.Split(lists[1], " ") {
			v, _ := strconv.Atoi(stringNumber)
			if slices.Contains(winning, v) && stringNumber != "" {
				//fmt.Println(v)
				count++
			}
		}
		if count > 0 {
			winningMap[card] = count
		}
	}

	total := 0
	for k, _ := range winningMap {
		total += getCountScratchcards(winningMap, k)
		//fmt.Print(k+1, ": ")
		//fmt.Println(getCountScratchcards(winningMap, k))
	}

	total += len(input)

	return total, nil
	return 0, errors.New("Not implemented")
}
