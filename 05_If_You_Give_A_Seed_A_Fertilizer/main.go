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
	seeds                    []int
	seedToSoilMap            [][]int
	soilToFertilizerMap      [][]int
	fertilizerToWaterMap     [][]int
	waterToLightMap          [][]int
	lightToTemperatureMap    [][]int
	temperatureToHumidityMap [][]int
	humidityToLocationMap    [][]int
}

func convertListTo2DInt(input []string) [][]int {
	n := len(input) / 3
	var output = make([][]int, n)
	for i := 0; i < n; i++ {
		output[i] = make([]int, 3)
	}
	for i, val := range input {
		row := i / 3
		col := i % 3
		o, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		output[row][col] = o
	}
	return output
}

func convertListToInt(input []string) []int {
	var output []int
	for _, i := range input {
		val, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		output = append(output, val)
	}
	return output
}
func checkBetween(input int, lower int, upper int) bool {
	if input >= lower && input < upper {
		// fmt.Printf("Checked Between: %d, %d, %d\n", input, lower, upper)
		return true
	}
	return false
}
func convertTo(input int, Map [][]int) int {
	for _, val := range Map {
		if checkBetween(input, val[1], val[1]+val[2]) {
			// fmt.Printf("val: %v\n", val)
			return input + val[0] - val[1]
		}
	}
	return input
}
func getLocation(seed int, input *Input, loc chan int) {
	output := seed
	output = convertTo(output, input.seedToSoilMap)
	output = convertTo(output, input.soilToFertilizerMap)
	output = convertTo(output, input.fertilizerToWaterMap)
	output = convertTo(output, input.waterToLightMap)
	output = convertTo(output, input.lightToTemperatureMap)
	output = convertTo(output, input.temperatureToHumidityMap)
	output = convertTo(output, input.humidityToLocationMap)
	fmt.Printf("%d, Output: %d,  \n", seed, output)
	loc <- output
}

func convertListTo2DIntRoutine(input []string, c chan [][]int) {
	n := len(input) / 3
	var output = make([][]int, n)
	for i := 0; i < n; i++ {
		output[i] = make([]int, 3)
	}
	for i, val := range input {
		row := i / 3
		col := i % 3
		o, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		output[row][col] = o
	}
	c <- output
}

func convertToInput(input *string) *Input {
	inputLines := strings.Split(*input, "\n\n")
	seeds := numberRegex.FindAllString(inputLines[0], -1)
	seedToSoilChan := make(chan [][]int)
	soilToFertilizerChan := make(chan [][]int)
	fertilizerToWaterChan := make(chan [][]int)
	waterToLightChan := make(chan [][]int)
	lightToTemperatureChan := make(chan [][]int)
	temperatureToHumidityChan := make(chan [][]int)
	humidityToLocationChan := make(chan [][]int)
	go convertListTo2DIntRoutine(numberRegex.FindAllString(inputLines[2], -1), soilToFertilizerChan)
	soilToFertilizer := <-soilToFertilizerChan
	go convertListTo2DIntRoutine(numberRegex.FindAllString(inputLines[3], -1), fertilizerToWaterChan)
	fertilizerToWater := <-fertilizerToWaterChan
	go convertListTo2DIntRoutine(numberRegex.FindAllString(inputLines[4], -1), waterToLightChan)
	waterToLight := <-waterToLightChan
	go convertListTo2DIntRoutine(numberRegex.FindAllString(inputLines[5], -1), lightToTemperatureChan)
	lightToTemperature := <-lightToTemperatureChan
	go convertListTo2DIntRoutine(numberRegex.FindAllString(inputLines[6], -1), temperatureToHumidityChan)
	temperatureToHumidity := <-temperatureToHumidityChan
	go convertListTo2DIntRoutine(numberRegex.FindAllString(inputLines[7], -1), humidityToLocationChan)
	humidityToLocation := <-humidityToLocationChan
	go convertListTo2DIntRoutine(numberRegex.FindAllString(inputLines[1], -1), seedToSoilChan)
	seedToSoil := <-seedToSoilChan
	return &Input{
		seeds:                    convertListToInt(seeds),
		seedToSoilMap:            seedToSoil,
		soilToFertilizerMap:      soilToFertilizer,
		fertilizerToWaterMap:     fertilizerToWater,
		waterToLightMap:          waterToLight,
		lightToTemperatureMap:    lightToTemperature,
		temperatureToHumidityMap: temperatureToHumidity,
		humidityToLocationMap:    humidityToLocation,
	}
}
func readFile(path *string) string {
	input, err := os.ReadFile(*path)
	if err != nil {
		panic(err)
	}
	return string(input)
}
func part1(input *string) {
	convertedInput := convertToInput(input)
	fmt.Printf("%v\n", *convertedInput)
	min := 999999999999999999
	for _, i := range convertedInput.seeds {
		c := make(chan int)
		go getLocation(i, convertedInput, c)
		loc := <-c
		if min > loc {
			min = loc
		}
	}
	fmt.Println(min)
}

func main() {
	start := time.Now()
	input := readFile(&os.Args[1])
	part1(&input)
	fmt.Printf("Time taken: %s", time.Since(start))
}
