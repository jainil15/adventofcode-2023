package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
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
func getSeeds(input *Input) {
	removeIndex := []int{}
	removeIdx := 1
	n := len(input.seeds)
	for i := 0; i < n; i += 2 {
		removeIndex = append(removeIndex, removeIdx)
		input.seeds = append(input.seeds, input.seeds[i]+input.seeds[i+1])
		removeIdx++
		fmt.Println("removedIndex", removeIndex)
	}
	for _, i := range removeIndex {
		input.seeds = append(input.seeds[:i], input.seeds[i+1:]...)
	}
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
func addSeeds(input *Input, seedLower int, seedUpper int, Map [][]int) {
	for _, val := range Map {
		if checkBetween(seedLower, val[1], val[2]+val[1]) && !checkBetween(seedUpper, val[1], val[2]+val[1]) {
			input.seeds = append(input.seeds, val[1]+1)
		}
		if !checkBetween(seedLower, val[1], val[2]+val[1]) && checkBetween(seedUpper, val[1], val[2]+val[1]) {
			input.seeds = append(input.seeds, val[1]-1)
		}
		if checkBetween(seedLower, val[1], val[2]+val[1]) && checkBetween(seedUpper, val[1], val[2]+val[1]) {
			fmt.Println(seedLower, seedUpper, val[1], val[1]+val[2])
		}
	}
}
func generateSeeds(input *Input, Map [][]int, n int) {
	for i := 0; i < n; i += 2 {
		addSeeds(
			input,
			input.seeds[i],
			input.seeds[i+1]+input.seeds[i],
			Map,
		)
	}
}
func getRemoveIndex(input *Input) []int {
	removeIndex := []int{}
	removeIdx := 1
	n := len(input.seeds)
	for i := 0; i < n; i += 2 {
		removeIndex = append(removeIndex, removeIdx)
		removeIdx++
	}
	return removeIndex
}
func removeIndexes(input *Input, removeIndex []int) {
	for _, i := range removeIndex {
		input.seeds = append(input.seeds[:i], input.seeds[i+1:]...)
	}
}
func seedRanges(input *Input) []int {
	n := len(input.seeds)
	var output = make([]int, n)
	for i := 0; i < n; i += 2 {
		output[i] = input.seeds[i]
		output[i+1] = input.seeds[i] + input.seeds[i+1]
	}
	slices.Sort(output)
	return output
}
func toList(input [][]int) ([]int, []int) {
	source := []int{}
	destination := []int{}
	for _, val := range input {
		source = append(source, val[1], val[1]+val[2])
		destination = append(destination, val[0], val[0]+val[2])
	}
	// slices.Sort(source)
	// slices.Sort(destination)
	return source, destination
}
func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func mapRange(input []int, source []int, destination []int) []int {
	n := len(input)
	m := len(source)
	// for _, i := range source {
	// 	for j := 0; j < n; j += 2 {
	// 		if checkBetween(i, input[j], input[j+1]) {
	// 			input = append(input, i-1)
	// 			input = append(input, i)
	// 			break
	// 		}
	// 	}
	// 	// slices.Sort(input)
	// }
	for i := 0; i < m; i += 2 {
		for j := 0; j < n; j += 2 {
			if checkBetween(source[i], input[j], input[j+1]) {
				input = append(input, Max(0, source[i]-1))
				input = append(input, Max(0, source[i]))
			}
			if checkBetween(source[i+1], input[j], input[j+1]) {
				input = append(input, source[i+1]+1)
				input = append(input, Max(0, source[i+1]))
			}
		}
	}
	fmt.Printf("Input Before Sort: %v\n", input)
	slices.Sort(input)
	fmt.Printf("Input: %v\n", input)
	n = len(input)
	outputRange := make([]int, n)
	for i := 0; i < n; i++ {
		checked := false
		for j := 0; j < m; j += 2 {
			fmt.Printf("Source: %v, Destination: %v\n", source, destination)
			if checkBetween(input[i], source[j], source[j+1]) {
				outputRange[i] = destination[j] + input[i] - source[j]
				fmt.Printf("Range Mapping: %v, to %v \n", input[i], outputRange)
				checked = true
				break
			}
		}
		if checked {
			continue
		} else {
			outputRange[i] = input[i]
			fmt.Printf("Self Mapping: %v, to %v \n", input[i], outputRange)
		}
	}
	slices.Sort(outputRange)
	fmt.Println(seedCount(outputRange))
	fmt.Println()
	return outputRange

}
func seedCount(seedRange []int) int {
	n := len(seedRange)
	count := 0
	for i := 0; i < n; i += 2 {
		count += seedRange[i+1] - seedRange[i]
	}
	return count
}
func part2(content *string) {
	input := convertToInput(content)
	seedRanges := seedRanges(input)
	slices.Sort(seedRanges)
	output := []int{}

	seedToSoilMapSource, seedToSoilMapDestination := toList(input.seedToSoilMap)
	output = mapRange(seedRanges, seedToSoilMapSource, seedToSoilMapDestination)
	fmt.Println(output)

	soilToFertilizerMapSource, soilToFertilizerMapDestination := toList(input.soilToFertilizerMap)
	output = mapRange(output, soilToFertilizerMapSource, soilToFertilizerMapDestination)
	fmt.Println(output)

	fertilizerToWaterMapSource, fertilizerToWaterMapDestination := toList(input.fertilizerToWaterMap)
	output = mapRange(output, fertilizerToWaterMapSource, fertilizerToWaterMapDestination)
	fmt.Println(output)

	waterToLightMapSource, waterToLightMapDestination := toList(input.waterToLightMap)
	output = mapRange(output, waterToLightMapSource, waterToLightMapDestination)
	fmt.Println(output)

	lightToTemperatureMapSource, lightToTemperatureMapDestination := toList(input.lightToTemperatureMap)
	output = mapRange(output, lightToTemperatureMapSource, lightToTemperatureMapDestination)
	fmt.Printf("Output: %v\n", output)

	temperatureToHumidityMapSource, temperatureToHumidityMapDestination := toList(input.temperatureToHumidityMap)
	output = mapRange(output, temperatureToHumidityMapSource, temperatureToHumidityMapDestination)
	fmt.Println(output)

	humidityToLocationMapSource, humidityToLocationMapDestination := toList(input.humidityToLocationMap)
	output = mapRange(output, humidityToLocationMapSource, humidityToLocationMapDestination)
	fmt.Println(output)

	fmt.Println(output[0])
}
func main() {
	start := time.Now()
	input := readFile(&os.Args[1])
	part2(&input)
	fmt.Printf("Time taken: %s", time.Since(start))
}
