package day08

import (
	"aoc2023/io"
	"errors"
	"fmt"
)

var Day int = 8

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
	instructions := input[0]
	nodes := input[2:]

	nodeMap := make(map[string][]string)

	for _, node := range nodes {
		baseNode := node[:3]
		leftNode := node[7:10]
		rightNode := node[12:15]
		nodeMap[baseNode] = []string{leftNode, rightNode}

	}

	stepCount := 0
	currentNode := "AAA"
	index := 0
	for currentNode != "ZZZ" {
		if instructions[index] == 'L' {
			currentNode = nodeMap[currentNode][0]
		} else {
			currentNode = nodeMap[currentNode][1]
		}
		index++
		if index == len(instructions) {
			index = 0
		}
		stepCount++
	}
	return stepCount, nil
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func Part2(input []string) (int, error) {
	instructions := input[0]
	nodes := input[2:]

	nodeMap := make(map[string][]string)

	travelNodes := []string{}

	for _, node := range nodes {
		baseNode := node[:3]
		leftNode := node[7:10]
		rightNode := node[12:15]
		nodeMap[baseNode] = []string{leftNode, rightNode}

		if baseNode[2] == 'A' {
			travelNodes = append(travelNodes, baseNode)
		}
	}

	stepCounts := []int{}
	largest := 0
	//fmt.Println(travelNodes)
	for _, currentNode := range travelNodes {
		//fmt.Println(currentNode)
		stepCount := 0
		index := 0
		for string(currentNode[2]) != "Z" {
			if instructions[index] == 'L' {
				currentNode = nodeMap[currentNode][0]
			} else {
				currentNode = nodeMap[currentNode][1]
			}
			index++
			if index == len(instructions) {
				index = 0
			}
			stepCount++
		}
		stepCounts = append(stepCounts, stepCount)
		if stepCount < largest || largest == 0 {
			largest = stepCount
		}
	}

	//fmt.Println(stepCounts)

	return LCM(stepCounts[0], stepCounts[1], stepCounts[2:]...), nil

	return 0, errors.New("Not implemented")
}
