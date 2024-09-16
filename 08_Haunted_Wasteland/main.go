package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

var wordRegex = regexp.MustCompile("[A-Z][A-Z][A-Z]")

var zRegex = regexp.MustCompile("[A-Z][A-Z]Z")
var aRegex = regexp.MustCompile("[A-Z][A-Z]A")

func readFile(path string) string {
	o, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(o)
}
func parseInput(input *string) (string, map[string][]string) {
	output := make(map[string][]string)
	s := strings.Split(*input, "\n\n")
	dir := s[0]
	p := strings.Split(s[1], "\n")[:len(strings.Split(s[1], "\n"))-1]
	for _, o := range p {
		g := strings.Split(o, " = ")
		d := wordRegex.FindAllString(g[1], -1)
		output[g[0]] = d
	}
	return dir, output
}

func getStartingPos(network map[string][]string) []string {
	startingPos := []string{}
	for k := range network {
		if aRegex.MatchString(k) {
			startingPos = append(startingPos, k)
		}
	}
	return startingPos
}
func move(finalString string, directions string, network map[string][]string) string {
	for _, i := range directions {
		switch i {
		case 'L':
			finalString = network[finalString][0]
		case 'R':
			finalString = network[finalString][1]
		}
	}
	return finalString
}
func part1(input *string) {
	directions, network := parseInput(input)
	finalString := "AAA"
	total := 0
	for finalString != "ZZZ" {
		finalString = move(finalString, directions, network)
		total++
	}
	fmt.Printf("Total: %v\n", total*len(directions))
}
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)
	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}
	return result
}
func checkFinalPosReached(finalPos []string) bool {
	fmt.Println(finalPos)
	for _, f := range finalPos {
		if !zRegex.MatchString(f) {
			return false
		}
	}
	return true
}
func part2(input *string) {
	directions, network := parseInput(input)
	startingPos := getStartingPos(network)
	fmt.Printf("Directions: %v,\nStartingPos %v\n", directions, startingPos)
	counts := make([]int, len(startingPos))
	for i, s := range startingPos {
		total := 0
		ps := s
		for {
			ps = move(ps, directions, network)
			total++
			if zRegex.MatchString(ps) {
				counts[i] = total * len(directions)
				break
			}
		}
	}
	fmt.Printf("Total: %v", LCM(counts[0], counts[1], counts...))
}
func main() {
	start := time.Now()
	path := os.Args[1]
	input := readFile(path)
	part2(&input)
	fmt.Printf("Took: %v", time.Since(start))
}
