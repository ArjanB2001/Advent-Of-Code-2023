package day06

import (
	"aoc2023/io"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var Day int = 6

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
	times := strings.Fields(strings.Split(input[0], ":")[1])
	distances := strings.Fields(strings.Split(input[1], ":")[1])
	//return 0, errors.New("Not implemented")
	total := 1
	for raceIndex, time := range times {
		record, _ := strconv.Atoi(distances[raceIndex])
		ways := 0
		timeInt, _ := strconv.Atoi(time)
		for i := 1; i < timeInt; i++ {
			timeRemaining := timeInt - i
			distTraveled := timeRemaining * i
			if distTraveled > record {
				ways++
			}

		}

		total *= ways
	}
	return total, nil
}

func Part2(input []string) (int, error) {
	time, _ := strconv.Atoi(strings.Join(strings.Fields(strings.Split(input[0], ":")[1]), ""))
	record, _ := strconv.Atoi(strings.Join(strings.Fields(strings.Split(input[1], ":")[1]), ""))
	ways := 0
	for i := 1; i < time; i++ {
		timeRemaining := time - i
		distTraveled := timeRemaining * i
		if distTraveled > record {
			ways++
		}
	}
	return ways, nil
	return 0, errors.New("Not implemented")
}
