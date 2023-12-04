package main

import (
	"aoc2023/io"
	"errors"
	"fmt"
	"strconv"
)

var day int = 3

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

func isGear(input []string, lineId int, beginning int, end int) (bool, int, int) {
	// Top row
	for i := beginning - 1; i <= end+1; i++ {
		if i < 0 || i >= len(input[lineId]) {
			continue
		}
		if lineId-1 < 0 {
			break
		}
		if input[lineId-1][i] == '*' {
			return true, lineId - 1, i
		}
	}

	// Bottom row
	for i := beginning - 1; i <= end+1; i++ {
		if i < 0 || i >= len(input[lineId]) {
			continue
		}

		if lineId+1 >= len(input) {
			break
		}
		if input[lineId+1][i] == '*' {
			return true, lineId + 1, i
		}
	}

	if beginning-1 >= 0 && input[lineId][beginning-1] == '*' {
		return true, lineId, beginning - 1
	}

	if end+1 < len(input[lineId]) && input[lineId][end+1] == '*' {
		return true, lineId, end + 1
	}

	return false, 0, 0
}

func isPartNumber(input []string, lineId int, beginning int, end int) bool {

	// Top row
	for i := beginning - 1; i <= end+1; i++ {
		if i < 0 || i >= len(input[lineId]) {
			continue
		}
		if lineId-1 < 0 {
			break
		}
		if input[lineId-1][i] != '.' {
			return true
		}
	}

	// Bottom row
	for i := beginning - 1; i <= end+1; i++ {
		if i < 0 || i >= len(input[lineId]) {
			continue
		}

		if lineId+1 >= len(input) {
			break
		}
		if input[lineId+1][i] != '.' {
			return true
		}
	}

	if beginning-1 >= 0 && input[lineId][beginning-1] != '.' {
		return true
	}

	if end+1 < len(input[lineId]) && input[lineId][end+1] != '.' {
		return true
	}

	return false
}

func part1(input []string) (int, error) {

	total := 0
	for lineId, line := range input {
		currentNumber := ""
		beginning := 0
		end := 0
		for i, char := range []rune(line) {
			if _, err := strconv.Atoi(string(char)); err == nil {
				currentNumber += string(char)
				if i == len(line)-1 {
					end = i
					if isPartNumber(input, lineId, beginning, end) {
						v, _ := strconv.Atoi(currentNumber)
						total += v
					}
				}
				continue
			} else if len(currentNumber) > 0 {
				//fmt.Println(currentNumber)

				end = i - 1
				if isPartNumber(input, lineId, beginning, end) {
					v, _ := strconv.Atoi(currentNumber)
					total += v
				}
				currentNumber = ""
				beginning = i + 1
			} else {
				beginning = i + 1
			}
		}
	}
	return total, nil
	return 0, errors.New("Not implemented")
}

func part2(input []string) (int, error) {
	gearMap := make(map[int][]int)
	total := 0
	for lineId, line := range input {
		currentNumber := ""
		beginning := 0
		end := 0
		for i, char := range []rune(line) {
			if _, err := strconv.Atoi(string(char)); err == nil {
				currentNumber += string(char)
				if i == len(line)-1 {
					end = i
					v, _ := strconv.Atoi(currentNumber)
					if yes, row, column := isGear(input, lineId, beginning, end); yes {
						if val, ok := gearMap[len(line)*row+column]; ok {
							gearMap[len(line)*row+column] = append(val, v)
						} else {
							gearMap[len(line)*row+column] = []int{v}
						}
					}
				}
				continue
			} else if len(currentNumber) > 0 {
				//fmt.Println(currentNumber)

				end = i - 1
				v, _ := strconv.Atoi(currentNumber)
				if yes, row, column := isGear(input, lineId, beginning, end); yes {
					if val, ok := gearMap[len(line)*row+column]; ok {
						gearMap[len(line)*row+column] = append(val, v)
					} else {
						gearMap[len(line)*row+column] = []int{v}
					}
				}
				currentNumber = ""
				beginning = i + 1
			} else {
				beginning = i + 1
			}
		}
	}

	for _, v := range gearMap {
		if len(v) == 2 {
			total += v[0] * v[1]
		}

	}
	return total, nil
	return 0, errors.New("Not implemented")
}
