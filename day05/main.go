package day05

import (
	"aoc2023/io"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Map struct {
	destinationStart int
	sourceStart      int
	rangeLength      int
}

type FeatureMap struct {
	maps []Map
}

var Day int = 5

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

func recurse(input int, mappings []FeatureMap) int {
	for _, m := range mappings[0].maps {

		if input >= m.sourceStart && input < m.sourceStart+m.rangeLength {
			if len(mappings) == 1 {
				return m.destinationStart + (input - m.sourceStart)
			}

			return recurse(m.destinationStart+(input-m.sourceStart), mappings[1:])
		}

	}
	if len(mappings) == 1 {
		return input
	}
	return recurse(input, mappings[1:])

}

func Part1(input []string) (int, error) {
	seeds := strings.Split(strings.Split(input[0], ": ")[1], " ")

	fmt.Println(seeds)
	mappings := []FeatureMap{}

	for i := 1; i < len(input); i++ {
		if input[i] == "" {
			mappings = append(mappings, FeatureMap{})
			i += 2
		}
		values := strings.Split(input[i], " ")
		ds, _ := strconv.Atoi(values[0])
		ss, _ := strconv.Atoi(values[1])
		rl, _ := strconv.Atoi(values[2])
		mappings[len(mappings)-1].maps = append(mappings[len(mappings)-1].maps, Map{destinationStart: ds, sourceStart: ss, rangeLength: rl})

	}

	lowest := 0
	for _, seed := range seeds {

		seedInt, _ := strconv.Atoi(seed)
		result := recurse(seedInt, mappings)
		//fmt.Println(result)
		if lowest == 0 || result < lowest {
			lowest = result
		}

	}

	return lowest, nil
	return 0, errors.New("Not implemented")
}

func Part2(input []string) (int, error) {
	seeds := strings.Split(strings.Split(input[0], ": ")[1], " ")
	//fmt.Println(seeds)
	mappings := []FeatureMap{}

	for i := 1; i < len(input); i++ {
		if input[i] == "" {
			mappings = append(mappings, FeatureMap{})
			i += 2
		}
		values := strings.Split(input[i], " ")
		ds, _ := strconv.Atoi(values[0])
		ss, _ := strconv.Atoi(values[1])
		rl, _ := strconv.Atoi(values[2])
		mappings[len(mappings)-1].maps = append(mappings[len(mappings)-1].maps, Map{destinationStart: ds, sourceStart: ss, rangeLength: rl})

	}

	lowest := 0
	for i := 0; i < len(seeds); i += 2 {
		seed1, _ := strconv.Atoi(seeds[i])
		seed2, _ := strconv.Atoi(seeds[i+1])

		fmt.Println(seed1)
		for is := seed1; is < seed1+seed2; is++ {
			result := recurse(is, mappings)
			if lowest == 0 || result < lowest {
				lowest = result
			}
		} //fmt.Println(result)
	}
	return lowest, nil
	return 0, errors.New("Not implemented")
}
