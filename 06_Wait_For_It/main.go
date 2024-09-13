package main

import (
	"fmt"
	"math"
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
func parseInput2(input string) Input {
	var output Input
	lines := strings.Split(input, "\n")
	time := numberRegex.FindAllString(lines[0], -1)
	distance := numberRegex.FindAllString(lines[1], -1)
	t, err := strconv.Atoi(strings.Join(time, ""))
	if err != nil {
		panic(err)
	}
	d, err := strconv.Atoi(strings.Join(distance, ""))
	if err != nil {
		panic(err)
	}
	output.time = []int{t}
	output.distance = []int{d}
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
			output = append(output, time-i)
			break
		}
	}
	return output
}

func calculateRange2(distance int, time int) []int {
	// equation: output = t + (t^2 - 4d)^0.5 / 2
	distanceF := float64(distance)
	timeF := float64(time)
	output := int(math.Ceil((timeF - math.Sqrt(timeF*timeF-4*distanceF)) / 2))
	return []int{output, time - output}
}

func part2(input Input) int {

	fmt.Println(input)
	prod := 1
	for i := 0; i < len(input.time); i++ {
		time := input.time[i]
		distance := input.distance[i]
		rangeTime := calculateRange2(distance, time)
		diff := calculateDiff(rangeTime[0], rangeTime[1])
		prod *= diff
	}
	return prod
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
	input := parseInput2(readFile(path))
	fmt.Println("Part 1: ", part2(input))
	fmt.Println("Time: ", time.Since(start))
}
