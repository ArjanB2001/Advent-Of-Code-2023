package day07

import (
	"aoc2023/io"
	"fmt"
	"reflect"
	"slices"
	"sort"
	"strconv"
	"strings"
)

var Day int = 7

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

func betterSort(firstCounts []int, secondCounts []int, firstHand string, secondHand string, types []rune) bool {
	if reflect.DeepEqual(firstCounts, secondCounts) {
		i := 0
		for firstHand[i] == secondHand[i] {
			i++
		}
		return slices.Index(types, rune(firstHand[i])) < slices.Index(types, rune(secondHand[i]))
	}

	if firstCounts[0] == secondCounts[0] {
		return firstCounts[1] > secondCounts[1]
	} else {
		return firstCounts[0] > secondCounts[0]
	}
}

func Part1(input []string) (int, error) {
	types := []rune{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}

	sort.Slice(input, func(first, second int) bool {
		firstHand := strings.Split(input[first], " ")[0]
		secondHand := strings.Split(input[second], " ")[0]

		firstCounts := make([]int, len(types))
		for _, char := range firstHand {
			firstCounts[slices.Index(types, char)]++
		}

		sort.Sort(sort.Reverse(sort.IntSlice(firstCounts)))

		secondCounts := make([]int, len(types))
		for _, char := range secondHand {
			secondCounts[slices.Index(types, char)]++
		}

		sort.Sort(sort.Reverse(sort.IntSlice(secondCounts)))

		//fmt.Println(firstHand)
		//fmt.Println(secondHand)

		better := betterSort(firstCounts, secondCounts, firstHand, secondHand, types)
		//fmt.Println(better)
		//fmt.Println("=================")
		return better
	})

	total := 0
	for index, hand := range input {
		bid, _ := strconv.Atoi(strings.Split(hand, " ")[1])
		total += bid * (len(input) - index)
	}
	//fmt.Println(input)
	return total, nil
}

func Part2(input []string) (int, error) {
	types := []rune{'A', 'K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2', 'J'}

	sort.Slice(input, func(first, second int) bool {
		firstHand := strings.Split(input[first], " ")[0]
		secondHand := strings.Split(input[second], " ")[0]

		firstCounts := make([]int, len(types))
		Jcount := 0
		for _, char := range firstHand {
			if char != 'J' {
				firstCounts[slices.Index(types, char)]++
			} else {
				Jcount++
			}

		}

		sort.Sort(sort.Reverse(sort.IntSlice(firstCounts)))
		firstCounts[0] += Jcount
		Jcount = 0

		secondCounts := make([]int, len(types))
		for _, char := range secondHand {
			if char != 'J' {
				secondCounts[slices.Index(types, char)]++
			} else {
				Jcount++
			}
		}

		sort.Sort(sort.Reverse(sort.IntSlice(secondCounts)))
		secondCounts[0] += Jcount

		//fmt.Println(firstHand)
		//fmt.Println(secondHand)

		better := betterSort(firstCounts, secondCounts, firstHand, secondHand, types)
		//fmt.Println(better)
		//fmt.Println("=================")
		return better
	})

	total := 0
	for index, hand := range input {
		bid, _ := strconv.Atoi(strings.Split(hand, " ")[1])
		total += bid * (len(input) - index)
	}
	//fmt.Println(input)
	return total, nil
}
