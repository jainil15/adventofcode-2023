package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var numberRegex = regexp.MustCompile(`\d+`)

type Input struct {
	time     []int
	distance []int
}

func convertToInt(input []string) []int {
	var output []int
	for _, i := range input {
		o, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		output = append(output, o)
	}
	return output
}
func parseInput(input string) Input {
	var output Input
	lines := strings.Split(input, "\n")
	time := numberRegex.FindAllString(lines[0], -1)
	distance := numberRegex.FindAllString(lines[1], -1)
	output.time = convertToInt(time)
	output.distance = convertToInt(distance)
	return output
}
func readFile(path string) string {
	output, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(output)
}
func calculateDiff(a int, b int) int {
	return b - a + 1
}
func calculateRange(distance int, time int) []int {
	var output []int
	for i := 0; i < time; i++ {
		if distance < (i * (time - i)) {
			output = append(output, i)
			break
		}
	}
	for i := time; i > 0; i-- {
		if distance < (i * (time - i)) {
			output = append(output, i)
			break
		}
	}
	return output
}
func part1(input Input) int {
	fmt.Println(input)
	prod := 1
	for i := 0; i < len(input.time); i++ {
		time := input.time[i]
		distance := input.distance[i]
		rangeTime := calculateRange(distance, time)
		fmt.Println(rangeTime)
		diff := calculateDiff(rangeTime[0], rangeTime[1])
		prod *= diff
	}
	return prod

}
func main() {
	start := time.Now()
	path := os.Args[1]
	input := parseInput(readFile(path))
	fmt.Println("Part 1: ", part1(input))
	fmt.Println("Time: ", time.Since(start))
}
