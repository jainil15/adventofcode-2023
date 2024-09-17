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

type wrapper struct {
	sides [][]int
}

func testPath2(network [][]rune, w *wrapper, sidesMap map[int][]int, direction direction, i, j, count int) int {
	if network[i][j] == 'S' {
		count++
		w.sides = append(w.sides, []int{i, j})
		sidesMap[i] = append(sidesMap[i], j)
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
	w.sides = append(w.sides, []int{i, j})
	c := testPath2(network, w, sidesMap, d, newI, newJ, count+1)
	if c == -1 {
		// fmt.Printf("Err: -1 pipe: %v direction: %v count: %v, i %v, j %v \n", string(network[i][j]), direction, count, i, j)
		return -1
	}
	sidesMap[i] = append(sidesMap[i], j)
	count = c
	return count
}
func checkNorthSouth(pipe rune) bool {
	// fmt.Printf("rune: %v\n", string(pipe))
	switch pipe {
	case '-':
		return false
	default:
		return true
	}
}
func checkPipeNS(pipe1 rune, pipe2 rune) bool {
	if pipe1 == 'S' {
		return false
	}
	if pipe2 == 'S' {
		return true
	}
	if pipe1 == '-' {
		if pipe2 == 'J' {
			return false
		}
		if pipe2 == '7' {
			return false
		}
		if pipe2 == '-' {
			return false
		}
		return true
	}
	if pipe2 == '-' {
		if pipe1 == 'L' {
			return false
		}
		if pipe1 == 'F' {
			return false
		}
		if pipe1 == '-' {
			return false
		}
		return true
	}
	if pipe1 == 'F' && pipe2 == '7' {
		return true
	}
	if pipe1 == 'F' && pipe2 == 'J' {
		return false
	}
	if pipe1 == 'L' && pipe2 == '7' {
		return false
	}
	if pipe1 == 'L' && pipe2 == 'J' {
		return true
	}
	return true
}
func checkIfInsideLoop(pipeMap [][]rune, i, k int, sides []int) bool {
	fmt.Printf("sides: %v\n\n", sides)
	count := 0
	n := len(sides[i:])
	fmt.Println(sides[i], sides[i+1:], n)
	for j := i; j < i+n-1; j++ {
		if sides[j+1]-sides[j] > 0 &&
			// checkNorthSouth(pipeMap[k][sides[j]]) &&
			// checkNorthSouth(pipeMap[k][sides[j+1]]) &&
			checkPipeNS(pipeMap[k][sides[j]], pipeMap[k][sides[j+1]]) {
			fmt.Printf("j is %d, j+1: %v, j: %v, pipe: %v, pipe: %v\n", j, sides[j+1], sides[j], string(pipeMap[k][sides[j]]), string(pipeMap[k][sides[j+1]]))
			count++
		}
		// } else {
		//
		// 	fmt.Printf("diff: %v first: %v second: %v\n", sides[j+1]-sides[j] > 0, checkNorthSouth(pipeMap[k][j]), checkNorthSouth(pipeMap[k][j+1]))
		// }
	}
	if count%2 == 0 {
		fmt.Println("IT IS FALSE!!!!!!!")
		return false
	}
	fmt.Println("IT IS TRUE!!!!!!!")
	return true
}
func getInsideLoopCount(pipeMap [][]rune, sideMap map[int][]int) int {
	count := 0
	for k := range sideMap {
		n := len(sideMap[k])
		slices.Sort(sideMap[k])
		if len(sideMap[k-1]) == 0 {
			fmt.Printf("INGORED FIRST k: %v\n", k)
			continue
		}
		if len(sideMap[k+1]) == 0 {
			fmt.Printf("IGNORED LAST k: %v\n", k)
			continue
		}
		fmt.Printf("\nk: %v\t\n", k)
		for i := 0; i < n-1; i++ {
			sub := sideMap[k][i+1] - sideMap[k][i] - 1
			if sub > 0 && checkIfInsideLoop(pipeMap, i, k, sideMap[k]) {
				count += sub
			}
		}
	}

	return count
}

func part2(input *string) {
	network, i, j := parseInput(input)
	fmt.Printf("Network: %v, sI: %v, sj: %v\n", network, i, j)
	newI, newJ := 0, 0
	newI, newJ = getIndex(south, i, j)
	sidesMap := make(map[int][]int)
	w := wrapper{}
	count := testPath2(network, &w, sidesMap, south, newI, newJ, 0)
	// fmt.Println(w)
	fmt.Println(getInsideLoopCount(network, sidesMap))
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
