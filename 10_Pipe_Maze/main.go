package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"
	"time"
)

type direction int

const (
	north direction = iota
	south
	west
	east
)

func parseInput(input *string) ([][]rune, int, int) {
	lines := strings.Split(*input, "\n")
	lines = lines[:len(lines)-1]
	var out = make([][]rune, len(lines))
	sRegex := regexp.MustCompile(`S`)
	sI, sJ := 0, 0
	for i, line := range lines {
		out[i] = []rune(line)
		if sRegex.MatchString(line) {
			sJ = strings.Index(line, "S")
			sI = i
		}
	}
	return out, sI, sJ
}

func readFile(path string) string {
	out, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(out)
}
func validateNextPipe(pipe rune, direction direction) bool {
	switch direction {
	case north:
		switch pipe {
		case '|', '7', 'F':
			return true
		default:
			return false
		}
	case south:
		switch pipe {
		case '|', 'L', 'J':
			return true
		default:
			return false
		}
	case west:
		switch pipe {
		case '-', 'L', 'F':
			return true
		default:
			return false
		}
	case east:
		switch pipe {
		case '-', '7', 'J':
			return true
		default:
			return false
		}
	}
	return false
}
func pipeDirections(pipe rune, direction direction) direction {
	switch direction {
	case north:
		switch pipe {
		case '|':
			return direction
		case '7':
			return west
		case 'F':
			return east
		default:
			panic(direction)
		}
	case south:
		switch pipe {
		case '|':
			return direction
		case 'L':
			return east
		case 'J':
			return west
		default:
			panic(direction)
		}
	case west:
		switch pipe {
		case '-':
			return direction
		case 'L':
			return north
		case 'F':
			return south
		default:
			panic(direction)
		}
	case east:
		switch pipe {
		case '-':
			return direction
		case '7':
			return south
		case 'J':
			return north
		default:
			panic(direction)
		}
	}
	return direction
}
func getIndex(direction direction, i, j int) (int, int) {
	switch direction {
	case north:
		return i - 1, j
	case south:
		return i + 1, j
	case east:
		return i, j + 1
	case west:
		return i, j - 1
	}
	panic(direction)
}

func testPath(network [][]rune, direction direction, i, j, count int) int {
	if network[i][j] == 'S' {
		count++
		return count
	}
	if network[i][j] == '.' {
		return -1
	}

	if !validateNextPipe(network[i][j], direction) {
		// fmt.Printf("Err: pipe: %v direction: %v count: %v, i %v, j %v \n", string(network[i][j]), direction, count, i, j)
		return -1
	}
	// fmt.Printf("PERFECT PIPE: pipe: %v direction: %v count: %v, i %v, j %v \n", string(network[i][j]), direction, count, i, j)
	d := pipeDirections(network[i][j], direction)
	newI, newJ := getIndex(d, i, j)
	c := testPath(network, d, newI, newJ, count+1)
	if c == -1 {
		// fmt.Printf("Err: -1 pipe: %v direction: %v count: %v, i %v, j %v \n", string(network[i][j]), direction, count, i, j)
		return -1
	}
	count = c
	return count
}

func testPath2(network [][]rune, sides, topDown map[int][]int, direction direction, i, j, count int) int {
	if network[i][j] == 'S' {
		count++
		sides[i] = append(sides[i], j)
		topDown[i] = append(topDown[i], j)
		return count
	}
	if network[i][j] == '.' {
		return -1
	}

	if !validateNextPipe(network[i][j], direction) {
		// fmt.Printf("Err: pipe: %v direction: %v count: %v, i %v, j %v \n", string(network[i][j]), direction, count, i, j)
		return -1
	}
	// fmt.Printf("PERFECT PIPE: pipe: %v direction: %v count: %v, i %v, j %v \n", string(network[i][j]), direction, count, i, j)
	d := pipeDirections(network[i][j], direction)
	newI, newJ := getIndex(d, i, j)
	c := testPath2(network, sides, topDown, d, newI, newJ, count+1)
	if c == -1 {
		// fmt.Printf("Err: -1 pipe: %v direction: %v count: %v, i %v, j %v \n", string(network[i][j]), direction, count, i, j)
		return -1
	}

	// sides[newI] = append(sides[newI], newJ)
	sides[i] = append(sides[i], j)
	topDown[j] = append(topDown[j], i)
	count = c
	return count
}
func subArr(arr []int) int {
	slices.Sort(arr)
	sub := 0
	n := len(arr)
	fmt.Println(arr, n)
	for i := 0; i < n-1; i += 2 {
		if (arr[i+1] - arr[i] - 1) < 0 {
			continue
		}

		if (arr[i+1] - arr[i] - 1) > 0 {
			fmt.Printf("arr: %v, i : %v, j : %v , sub: %v\n", arr, arr[i+1], arr[i], arr[i+1]-arr[i]-1)
		}
		sub += (arr[i+1] - arr[i] - 1)
	}
	return sub
}
func getCountInsideLoop(sides map[int][]int) (sum int) {
	for k := range sides {
		fmt.Println(k)
		sum += subArr(sides[k])
	}
	return sum
}
func part2(input *string) {
	network, i, j := parseInput(input)
	fmt.Printf("Network: %v, sI: %v, sj: %v\n", network, i, j)
	newI, newJ := 0, 0
	newI, newJ = getIndex(south, i, j)
	sides := map[int][]int{}
	topDown := map[int][]int{}
	count := testPath2(network, sides, topDown, south, newI, newJ, 0)
	insideLoopCount := getCountInsideLoop(sides)
	fmt.Println(insideLoopCount)
	fmt.Printf("Count : %d\n", count/2)
}
func part1(input *string) {
	network, i, j := parseInput(input)
	fmt.Printf("Network: %v, sI: %v, sj: %v\n", network, i, j)
	count := 0
	newI, newJ := 0, 0

	newI, newJ = getIndex(north, i, j)
	count = max(testPath(network, north, newI, newJ, 0), count)

	newI, newJ = getIndex(east, i, j)
	count = max(testPath(network, east, newI, newJ, 0), count)

	newI, newJ = getIndex(south, i, j)
	count = max(testPath(network, south, newI, newJ, 0), count)

	newI, newJ = getIndex(west, i, j)
	count = max(testPath(network, west, newI, newJ, 0), count)
	fmt.Printf("Count: %d\n", count/2)
}

func main() {
	start := time.Now()
	path := os.Args[1]
	content := readFile(path)
	part2(&content)

	fmt.Printf("Took: %v\n", time.Since(start))
}
