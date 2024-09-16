package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func readFile(path string) string {
	o, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(o)
}

func parseIntArr(inputs []string) []int {
	output := make([]int, len(inputs))
	for i, input := range inputs {
		o, err := strconv.Atoi(input)
		if err != nil {
			panic(err)
		}
		output[i] = o
	}
	return output
}

func parseInput(input *string) [][]int {
	lines := strings.Split(*input, "\n")
	lines = lines[:len(lines)-1]
	n := len(strings.Split(lines[0], " "))
	output := make([][]int, len(lines))
	for i := 0; i < len(lines); i++ {
		output[i] = make([]int, n)
	}
	for i, line := range lines {
		l := strings.Split(line, " ")
		output[i] = parseIntArr(l)
	}
	return output
}
func getLine(line []int) ([]int, bool) {
	n := len(line)
	output := make([]int, n-1)
	isZero := true
	for i := 0; i < n-1; i++ {
		output[i] = line[i+1] - line[i]
		if output[i] != 0 && isZero {
			isZero = false
		}
	}
	return output, isZero
}
func getHistory(line []int) []int {
	output := [][]int{line}
	i := 0
	n := len(line) - 1
	endVals := []int{line[n]}
	for {
		o, isZero := getLine(output[i])
		output = append(output, o)
		i++
		endVals = append(endVals, o[len(o)-1])
		if isZero {
			break
		}
	}
	fmt.Printf("Endvals: %v\n", endVals)
	return endVals
}

func sumArr(list []int) int {
	sum := 0
	for _, val := range list {
		sum += val
	}
	return sum
}

func part1(input *string) {
	lines := parseInput(input)
	sum := 0
	for _, line := range lines {
		output := getHistory(line)
		sum += sumArr(output)
		fmt.Println(output)
	}
	fmt.Printf("Sum: %v\n", sum)
}
func main() {
	start := time.Now()
	path := os.Args[1]
	input := readFile(path)
	part1(&input)
	fmt.Printf("Took: %v\n", time.Since(start))
}
