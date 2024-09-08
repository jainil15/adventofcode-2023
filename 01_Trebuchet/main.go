package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var numberRegex = regexp.MustCompile(`one|two|three|four|five|six|seven|eight|nine|zero|\d`)
var reverseNumberRegex = regexp.MustCompile(`eno|owt|eerht|ruof|evif|xis|neves|thgie|enin|orez|\d`)

func fileReader(path string) (string, error) {
	inBytes, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(inBytes), nil
}

func getBlock(input string) []string {
	blocks := strings.Split(input, "\n\n")
	return blocks
}

func getLine(input string) []string {
	line := strings.Split(input, "\n")
	return line[:len(line)-1]
}
func isNumeric(w string) bool {
	_, err := strconv.Atoi(w)
	return err == nil
}
func getCalibrationValue(line string) int {
	var calibrationValueString string
	n := len(line)
	for _, word := range line {
		if isNumeric(string(word)) {
			calibrationValueString = string(word)
			break
		}
	}

	for i, _ := range line {
		if isNumeric(string(line[n-1-i])) {
			calibrationValueString += string(line[n-1-i])
			break
		}
	}
	calibrationValue, err := strconv.Atoi(calibrationValueString)
	if err != nil {
		panic(err)
	}
	return calibrationValue
}

func part1(content string) {

	finalCalibrationValue := 0
	blocks := getBlock(content)
	fmt.Println(len(blocks))
	for _, block := range blocks {
		lines := getLine(block)
		for _, line := range lines {
			finalCalibrationValue += getCalibrationValue(line)
		}
	}
	fmt.Println(finalCalibrationValue)
}
func convertToNumber(word string) string {
	switch word {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	case "zero":
		return "0"
	default:
		panic("Invalid word")
	}
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
func getCalibrationValueWithWords(line string) int {
	var calibrationValueString string
	wordsString := numberRegex.FindAllString(line, -1)
	fmt.Println(wordsString)
	if len(wordsString[0]) == 1 {
		calibrationValueString += wordsString[0]
	} else {
		calibrationValueString += convertToNumber(wordsString[0])
	}
	reversedLine := Reverse(line)
	fmt.Println(reversedLine)
	reverseWordsString := reverseNumberRegex.FindAllString(reversedLine, -1)
	fmt.Println(reverseWordsString)
	if len(reverseWordsString[0]) == 1 {
		calibrationValueString += reverseWordsString[0]
	} else {
		calibrationValueString += convertToNumber(Reverse(reverseWordsString[0]))
	}
	calibrationValue, err := strconv.Atoi(calibrationValueString)

	if err != nil {
		panic(err)
	}
	return calibrationValue
}
func part2(content string) {

	finalCalibrationValue := 0
	lines := getLine(content)
	for i, line := range lines {
		fmt.Println(i)
		finalCalibrationValue += getCalibrationValueWithWords(line)
	}
	fmt.Println(finalCalibrationValue)
}
func main() {
	start := time.Now()
	path := os.Args[1]
	content, err := fileReader(path)
	if err != nil {
		panic(err)
	}
	part2(content)
	fmt.Println(time.Since(start))
}
