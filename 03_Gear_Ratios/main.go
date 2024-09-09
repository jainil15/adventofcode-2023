package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var numberRegex = regexp.MustCompile("[0-9]+")
var starRegex = regexp.MustCompile("[*]")

func readFile(path string) string {
	input, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(input)
}
func part1(input string) int {
	inputSplit := strings.Split(input, "\n")[0 : len(strings.Split(input, "\n"))-1]
	n := len(inputSplit[0]) + 1
	indexes := numberRegex.FindAllStringSubmatchIndex(input, -1)
	numbers := numberRegex.FindAllString(input, -1)
	sum := 0
	for i, val := range indexes {
		checked := false
		for j := val[0]; j < val[1]; j++ {
			row := j / n
			col := j % n
			fmt.Printf("Index: %d, Row: %d, Col: %d\n", j, row, col)
			if row > 0 {

				// fmt.Println(string(inputSplit[row-1][col]))
			}
			if row > 0 && inputSplit[row-1][col] != '.' && !checked {
				s, err := strconv.Atoi(numbers[i])
				if err != nil {
					panic(err)
				}
				sum += s
				checked = true
				break
			}
			if row < n-2 && inputSplit[row+1][col] != '.' && !checked {
				s, err := strconv.Atoi(numbers[i])
				if err != nil {
					panic(err)
				}
				sum += s
				checked = true
				break
			}
		}
		leftRow := val[0] / n
		leftCol := val[0] % n
		rightRow := (val[1] - 1) / n
		rightCol := (val[1] - 1) % n
		if leftCol > 0 && inputSplit[leftRow][leftCol-1] != '.' && !checked {
			s, err := strconv.Atoi(numbers[i])
			if err != nil {

				panic(err)
			}
			sum += s
			checked = true
			continue
		}
		if rightCol < n-2 && inputSplit[rightRow][rightCol+1] != '.' && !checked {
			s, err := strconv.Atoi(numbers[i])
			if err != nil {
				panic(err)
			}
			sum += s
			checked = true
			continue
		}
		//Adjacent
		// Bottom left
		if leftRow < n-2 && leftCol > 0 && inputSplit[leftRow+1][leftCol-1] != '.' && !checked {
			s, err := strconv.Atoi(numbers[i])
			if err != nil {
				panic(err)
			}
			sum += s
			checked = true
			continue
		}
		// Top Left
		if leftRow > 0 && leftCol > 0 && inputSplit[leftRow-1][leftCol-1] != '.' && !checked {
			s, err := strconv.Atoi(numbers[i])
			if err != nil {
				panic(err)
			}
			sum += s
			checked = true
			continue
		}
		// Bottom Right
		if rightRow < n-2 && rightCol < n-2 && inputSplit[rightRow+1][rightCol+1] != '.' && !checked {
			s, err := strconv.Atoi(numbers[i])
			if err != nil {
				panic(err)
			}
			sum += s
			checked = true
			continue
		}
		// Bottom Left
		if rightRow > 0 && rightCol < n-2 && inputSplit[rightRow-1][rightCol+1] != '.' && !checked {
			s, err := strconv.Atoi(numbers[i])
			if err != nil {
				panic(err)
			}
			sum += s
			checked = true
			continue
		}

		if !checked {
			fmt.Println("Nothinggl", numbers[i])
		}
	}
	fmt.Println(sum)
	return sum
}
func checkBetween(index1 []int, index2 int) bool {
	if index2 >= index1[0] && index2 <= index1[1] {
		fmt.Printf("Index1: %v, Index2: %d\n", index1, index2)
		return true
	}
	return false
}

func part2(input string) int {
	inputSplit := strings.Split(input, "\n")[0 : len(strings.Split(input, "\n"))-1]
	n := len(inputSplit[0]) + 1
	fmt.Println(n)
	indexes := numberRegex.FindAllStringSubmatchIndex(input, -1)
	starIndexes := starRegex.FindAllStringSubmatchIndex(input, -1)
	numbers := numberRegex.FindAllString(input, -1)
	sum := 0
	for _, val := range starIndexes {
		prod := 1
		count := 0
		fmt.Println(val)
		topUpperBound := (val[0]/n-1)*n + (val[0] % n) - 1
		topLowerBound := (val[0]/n-1)*n + (val[0] % n) + 1

		bottomUpperBound := (val[0]/n+1)*n + (val[0] % n) - 1
		bottomLowerBound := (val[0]/n+1)*n + (val[0] % n) + 1
		middleUpperBound := (val[0]/n)*n + (val[0] % n) - 1
		middleLowerBound := (val[0]/n)*n + (val[0] % n) + 1
		for index, val := range indexes {
			for j := val[0]; j < val[1]; j++ {

				if checkBetween([]int{topUpperBound, topLowerBound}, j) {
					p, err := strconv.Atoi(numbers[index])
					if err != nil {
						panic(err)
					}
					prod *= p
					count++
					break
				}
				if checkBetween([]int{bottomUpperBound, bottomLowerBound}, j) {
					p, err := strconv.Atoi(numbers[index])
					if err != nil {
						panic(err)
					}
					prod *= p
					count++
					break

				}
				if checkBetween([]int{middleUpperBound, middleLowerBound}, j) {
					p, err := strconv.Atoi(numbers[index])
					if err != nil {
						panic(err)
					}
					prod *= p
					count++
					break
				}
			}
		}
		if count == 2 {
			sum += prod
		}
		if count > 2 {
			panic("More than 2")
		}
		fmt.Println(topUpperBound, topLowerBound, bottomUpperBound, bottomLowerBound, middleUpperBound, middleLowerBound)
	}
	fmt.Println(indexes, numbers)
	fmt.Println(sum)
	return sum
}

func main() {
	start := time.Now()
	path := os.Args[1]
	input := readFile(path)
	// part1(input)
	part2(input)
	fmt.Println(time.Since(start))
}
