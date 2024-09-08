package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

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
func main() {
	start := time.Now()
	path := os.Args[1]
	finalCalibrationValue := 0
	content, err := fileReader(path)
	if err != nil {
		panic(err)
	}
	blocks := getBlock(content)
	fmt.Println(len(blocks))
	for _, block := range blocks {
		lines := getLine(block)
		for _, line := range lines {
			finalCalibrationValue += getCalibrationValue(line)
		}
	}
	fmt.Println(finalCalibrationValue)
	fmt.Println(time.Since(start))
}
