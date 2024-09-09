package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

var numberRegex = regexp.MustCompile(`\d+`)

func readFile(path *string) string {
	i, err := os.ReadFile(*path)
	if err != nil {
		panic(err)
	}
	return string(i)
}
func getLines(content *string) []string {
	return strings.Split(*content, "\n")[:len(strings.Split(*content, "\n"))-1]
}
func getWinningNumbers(content *string) []string {
	actualString := strings.Split(*content, ": ")[1]

	winningString := numberRegex.FindAllString(strings.Split(actualString, "|")[0], -1)
	return winningString
}
func getYourNumbers(content *string) []string {
	actualString := strings.Split(*content, ": ")[1]
	cardString := numberRegex.FindAllString(strings.Split(actualString, "|")[1], -1)
	return cardString
}
func in(a string, list []string) bool {
	for _, i := range list {
		if i == a {
			return true
		}
	}
	return false
}
func part1(input *string) {
	lines := getLines(input)
	sumScore := 0
	for _, line := range lines {
		winningNumbers := getWinningNumbers(&line)
		yourNumbers := getYourNumbers(&line)
		score := 0
		for _, yourNumber := range yourNumbers {
			if in(yourNumber, winningNumbers) {
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
			}
		}
		sumScore += score
	}
	fmt.Println(sumScore)
}
func main() {
	start := time.Now()
	path := os.Args[1]
	input := readFile(&path)
	part1(&input)
	fmt.Println(time.Since(start))
}
