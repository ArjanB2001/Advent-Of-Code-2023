package day02

import (
	"aoc2023/io"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var Day int = 2

func main() {
	RunBothParts()
}

func RunBothParts() {
	fmt.Printf("Running Day %d\n", Day)
	//input := io.GetInput(Day)
	input := io.ReadFile(Day)
	//fmt.Printf("%v", input)

	answer1, err := Part1(input)
	if err != nil {
		fmt.Printf("Error running part 1\n")
	} else {
		fmt.Printf("Part 1 answer: %d \n", answer1)
	}

	answer2, err := Part2(input)
	if err != nil {
		fmt.Printf("Error running part 2\n")
	} else {
		fmt.Printf("Part 2 answer: %d \n", answer2)
	}
}

func Part1(input []string) (int, error) {
	total := 0
singleGame:
	for gameId, line := range input {
		//fmt.Println(line)
		sets := strings.Split(strings.Split(line, ":")[1][1:], "; ")

		for _, set := range sets {
			colors := strings.Split(set, ", ")
			for _, color := range colors {
				combo := strings.Split(color, " ")
				num, err := strconv.Atoi(combo[0])
				if err != nil {
					fmt.Printf("Error converting %s to int\n", combo[0])
				}
				if combo[1] == "red" {
					if num > 12 {
						continue singleGame
					}
				} else if combo[1] == "green" {
					if num > 13 {
						continue singleGame
					}
				} else if combo[1] == "blue" {
					if num > 14 {
						continue singleGame
					}
				} else {
					fmt.Printf("Error converting %s to color\n", combo[1])
				}

			}
		}
		//fmt.Printf("Game %d is valid\n", gameId+1)
		total += gameId + 1
	}
	return total, nil
	return 0, errors.New("Not implemented")
}

func Part2(input []string) (int, error) {
	total := 0

	for _, line := range input {
		maxRed := 0
		maxGreen := 0
		maxBlue := 0
		//fmt.Println(line)
		sets := strings.Split(strings.Split(line, ":")[1][1:], "; ")

		for _, set := range sets {
			colors := strings.Split(set, ", ")
			for _, color := range colors {
				combo := strings.Split(color, " ")
				num, err := strconv.Atoi(combo[0])
				if err != nil {
					fmt.Printf("Error converting %s to int\n", combo[0])
				}
				if combo[1] == "red" {
					if num > maxRed {
						maxRed = num
					}
				} else if combo[1] == "green" {
					if num > maxGreen {
						maxGreen = num
					}
				} else if combo[1] == "blue" {
					if num > maxBlue {
						maxBlue = num
					}
				} else {
					fmt.Printf("Error converting %s to color\n", combo[1])
				}

			}
		}
		power := maxRed * maxGreen * maxBlue
		total += power
		//fmt.Printf("Game %d is valid\n", gameId+1)
		//total += gameId + 1
	}
	return total, nil
	return 0, errors.New("Not implemented")
}
