package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

var wordRegex = regexp.MustCompile("[A-Z][A-Z][A-Z]")

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
func main() {
	start := time.Now()
	path := os.Args[1]
	input := readFile(path)
	part1(&input)
	fmt.Printf("Took: %v", time.Since(start))
}
