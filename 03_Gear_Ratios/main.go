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

func readFile(path string) string {
	input, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(input)
}
func part1v2(input string) int {
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
func part1(input string) int {
	inputSplit := strings.Split(input, "\n")[0 : len(strings.Split(input, "\n"))-1]
	n := len(inputSplit[0]) + 1
	idx := numberRegex.FindAllStringSubmatchIndex(input, -1)
	numbers := numberRegex.FindAllString(input, -1)
	sum := 0
	fmt.Println(idx)
	fmt.Println(numbers)

	for k, i := range idx {
		// Check Up and down
		checked := false
		for j := i[0]; j < i[1]; j++ {
			col := j % n
			row := j / n
			if row > 0 && inputSplit[row-1][col] != '.' && !checked {
				fmt.Printf("Up: %s\n", numbers[k])
				fmt.Printf("Row: %d, Col: %d\n", row, col)
				s, err := strconv.Atoi(numbers[k])
				if err != nil {
					panic(err)
				}
				checked = true
				sum += s
				break
			}
			if row < len(inputSplit)-1 && inputSplit[row+1][col] != '.' && !checked {
				fmt.Printf("Down: %s\n", numbers[k])
				fmt.Printf("Row: %d, Col: %d\n", row, col)
				s, err := strconv.Atoi(numbers[k])
				if err != nil {
					panic(err)
				}
				sum += s
				checked = true
				break
			}
		}

		leftRow := i[0] / n
		leftCol := i[0] % n
		rightRow := (i[1] - 1) / n
		rightCol := (i[1] - 1) % n
		// Check Left and Right

		if leftCol > 0 && inputSplit[leftRow][leftCol-1] != '.' && !checked {
			fmt.Printf("Left: %s\n", numbers[k])
			fmt.Printf("Row: %d, Col: %d\n", leftRow, leftCol)
			s, err := strconv.Atoi(numbers[k])
			if err != nil {
				panic(err)
			}
			sum += s
			checked = true
			continue
		}

		if rightCol < n-1 && inputSplit[rightRow][rightCol+1] != '.' && !checked {
			fmt.Printf("Right: %s\n", numbers[k])
			s, err := strconv.Atoi(numbers[k])
			if err != nil {
				panic(err)
			}
			sum += s
			checked = true
			continue
		}

		// Check Adjacent
		// Top Right
		if leftRow > 0 && leftCol < n-1 && inputSplit[leftRow-1][leftCol+1] != '.' && !checked {
			fmt.Printf("Top Right: %s\n", numbers[k])
			s, err := strconv.Atoi(numbers[k])
			if err != nil {
				panic(err)
			}
			sum += s
			checked = true
			continue
		}
		// Top Left
		if leftRow > 0 && leftCol > 0 && inputSplit[leftRow-1][leftCol-1] != '.' && !checked {
			fmt.Printf("Top Left: %s\n", numbers[k])
			s, err := strconv.Atoi(numbers[k])
			if err != nil {
				panic(err)
			}
			sum += s
			checked = true
			continue

		}
		// Bottom Right
		if rightRow < len(inputSplit)-1 && rightCol < n-1 && inputSplit[rightRow+1][rightCol+1] != '.' && !checked {
			fmt.Printf("Bottom Right: %s\n", numbers[k])
			s, err := strconv.Atoi(numbers[k])
			if err != nil {
				panic(err)
			}

			sum += s
			checked = true
			continue
		}
		// Bottom Left
		if rightRow < len(inputSplit)-1 && rightCol > 0 && inputSplit[rightRow+1][rightCol-1] != '.' && !checked {
			fmt.Printf("Bottom Left: %s\n", numbers[k])

			s, err := strconv.Atoi(numbers[k])
			if err != nil {
				panic(err)
			}
			sum += s
			checked = true
			continue
		}

		// if i[0] > 0 && inputSplit != '.' && !checked {

		fmt.Printf("No Adjacent: %s\n", numbers[k])
	}
	fmt.Println(sum)
	return 0
}

func main() {
	start := time.Now()
	path := os.Args[1]
	input := readFile(path)
	part1v2(input)

	fmt.Println(time.Since(start))
}
