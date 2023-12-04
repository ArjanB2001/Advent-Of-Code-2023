package io

import (
	"bufio"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func GetInput(day int) []string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cookie := os.Getenv("COOKIE")
	client := http.Client{}
	url := fmt.Sprintf("https://adventofcode.com/2023/day/%d/input", day)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error opening input request to AoC!")
	}

	req.Header = http.Header{
		"Cookie": {fmt.Sprintf("session=%s", cookie)},
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error closing input request to AoC!")
	}

	b, err := io.ReadAll(res.Body)
	return strings.Split(string(b), "\n")
}

func ReadFile(day int) []string {
	file, err := os.Open(fmt.Sprintf("day%02d/example", day))
	if err != nil {
		log.Println(err)
		return nil
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
