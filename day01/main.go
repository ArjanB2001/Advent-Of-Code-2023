package day01

import (
	"aoc2023/io"
	"errors"
	"fmt"
	"log"
	"slices"
	"strconv"
)

var Day = 1

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
	//out:
	for _, line := range input {

		runes := []rune(line)
		var i int = 0
		result := ""
		swap := false
		for len(result) < 2 {
			_, err := strconv.Atoi(string(runes[i]))
			if err != nil {
				if swap {
					i--
				} else {
					i++
				}

				continue
			} else {
				result += string(runes[i])
				i = len(runes) - 1
				swap = true
			}

		}
		//fmt.Println(result)
		v, err := strconv.Atoi(result)
		if err != nil {
			log.Println("shit")
		}
		total += v
		//fmt.Println(total)
	}
	//fmt.Println("huh")

	return total, nil
	return 0, errors.New("Not implemented")
}

func Part2(input []string) (int, error) {
	total := 0
	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for _, line := range input {
		foundChars := ""
		chars := []rune(line)
		//	Find first
		index := 0
		length := 1
		for {
			if slices.Contains(numbers, string(chars[index:index+length])) {
				// Found first
				_, err := strconv.Atoi(string(chars[index : index+length]))
				if err != nil {
					i := slices.Index(numbers, string(chars[index:index+length]))
					foundChars = foundChars + string(strconv.Itoa(i+1))
				} else {
					foundChars = foundChars + string(chars[index:index+length])
				}
				//fmt.Println(string(chars[index : index+length]))
				break
			}

			if length >= 5 {
				index++
				length = 1
				continue
			} else {
				length += 1
				continue
			}
		}

		//	Find second
		index = len(chars)
		length = 1
		for {
			if slices.Contains(numbers, string(chars[index-length:index])) {
				_, err := strconv.Atoi(string(chars[index-length : index]))
				if err != nil {
					i := slices.Index(numbers, string(chars[index-length:index]))
					foundChars = foundChars + string(strconv.Itoa(i+1))
				} else {
					foundChars = foundChars + string(chars[index-length:index])
				}
				break
			}

			if length >= 5 || index-length == 0 {
				index--
				length = 1
				continue
			} else {
				length += 1
				continue
			}
		}
		v, _ := strconv.Atoi(foundChars)
		total += v
	}
	return total, nil
	return 0, errors.New("Not implemented")
}
