package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	blueCubes  int = 14
	greenCubes int = 13
	redCubes   int = 12
	totalCubes int = 39
)

var gameStringRegex = regexp.MustCompile("Game [0-9]+")
var gameNumberRegex = regexp.MustCompile("[0-9]+")
var blueCubesRegex = regexp.MustCompile("[0-9]+ blue")
var redCubesRegex = regexp.MustCompile("[0-9]+ red")
var greenCubesRegex = regexp.MustCompile("[0-9]+ green")

func readFile(path string) string {
	input, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(input)
}

func getLines(input string) []string {
	return strings.Split(input, "\n")[0 : len(strings.Split(input, "\n"))-1]
}

// Bad Performance :: Improvement needed
func getGameRegex(input string) int {
	gameNumberString := gameNumberRegex.FindString(gameStringRegex.FindString(input))
	gameNumber, err := strconv.Atoi(gameNumberString)
	if err != nil {
		panic(err)
	}
	return gameNumber
}

func getTotalBlueCubes(input string) bool {
	blueCubesString := blueCubesRegex.FindAllString(input, -1)
	// fmt.Println(strings.Join(blueCubesString, ", "))
	for _, blueCubesStr := range blueCubesString {
		blueCubesInt, err := strconv.Atoi(strings.Split(blueCubesStr, " ")[0])
		if err != nil {
			panic(err)
		}
		if blueCubesInt > blueCubes {
			return false
		}
	}
	return true
}

func getTotalGreenCubes(input string) bool {
	GreenCubesString := greenCubesRegex.FindAllString(input, -1)
	// fmt.Println(strings.Join(GreenCubesString, ", "))
	for _, GreenCubesStr := range GreenCubesString {
		GreenCubesInt, err := strconv.Atoi(strings.Split(GreenCubesStr, " ")[0])
		if err != nil {
			panic(err)
		}
		if GreenCubesInt > greenCubes {
			return false
		}
	}
	return true
}

func getTotalRedCubes(input string) bool {
	RedCubesString := redCubesRegex.FindAllString(input, -1)
	// fmt.Println(strings.Join(RedCubesString, ", "))
	for _, RedCubesStr := range RedCubesString {
		RedCubesInt, err := strconv.Atoi(strings.Split(RedCubesStr, " ")[0])
		if err != nil {
			panic(err)
		}
		if RedCubesInt > redCubes {
			return false
		}
	}
	return true
}

func getGameSplit(input string) int {
	gameString := strings.Split(input, ": ")[0]
	gameNumberString := strings.Split(gameString, " ")[1]
	gameNumber, err := strconv.Atoi(gameNumberString)
	if err != nil {
		panic(err)
	}
	return gameNumber
}

func main() {
	start := time.Now()
	path := os.Args[1]
	input := readFile(path)
	lines := getLines(input)
	sumGoodGames := 0
	for index, line := range lines {
		// fmt.Printf("Game: %d : Blue: %d, Green: %d, Red: %d\n", getGameSplit(line), getTotalBlueCubes(line), getTotalGreenCubes(line), getTotalRedCubes(line))
		if !getTotalBlueCubes(line) {
			fmt.Printf("Blue Cubes: %d\n", getTotalBlueCubes(line))
			continue
		}
		if !getTotalGreenCubes(line) {
			fmt.Printf("Green Cubes: %d\n", getTotalGreenCubes(line))
			continue
		}
		if !getTotalRedCubes(line) {
			fmt.Printf("Red Cubes: %d\n", getTotalRedCubes(line))
			continue
		}
		fmt.Printf("\n\n************* Game %d ***************\n\n", index+1)
		sumGoodGames += index + 1
	}

	fmt.Println(sumGoodGames)
	fmt.Println(time.Since(start))
}
