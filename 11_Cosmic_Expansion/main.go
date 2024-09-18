package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"
)

func readFile(path string) string {
	output, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(output)
}

func parseInput(content *string) [][]rune {
	output := strings.Split(*content, "\n")
	output = output[:len(output)-1]
	outputRunes := make([][]rune, len(output))
	for i, _ := range outputRunes {
		outputRunes[i] = []rune(output[i])
	}
	return outputRunes
}

func transpose(arr [][]rune) [][]rune {
	n, m := len(arr), len(arr[0])
	transposed := make([][]rune, m)
	for i, _ := range transposed {
		transposed[i] = make([]rune, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			transposed[j][i] = arr[i][j]
		}
	}
	// fmt.Printf("Normal: %v\nTransposed: %v\n", arr, transposed)
	return transposed
}

func expandGalaxy(galaxyMap [][]rune) [][]rune {
	n, m := len(galaxyMap), len(galaxyMap[0])
	expandedGalaxy := [][]rune{}
	_ = n
	emptyGalaxyRow := make([]rune, m)
	for i := 0; i < m; i++ {
		emptyGalaxyRow[i] = '.'
	}
	// fmt.Println(n, m)
	// ROW Expansion
	for _, val := range galaxyMap {
		if reflect.DeepEqual(val, emptyGalaxyRow) {
			// fmt.Println(val)
			for l := 0; l < 2; l++ {
				expandedGalaxy = append(expandedGalaxy, val)
			}
			continue
		}
		expandedGalaxy = append(expandedGalaxy, val)
	}
	transposedGalaxy := transpose(expandedGalaxy)
	emptyGalaxyCol := make([]rune, len(transposedGalaxy[0]))
	expandedGalaxy = [][]rune{}
	for i := 0; i < len(transposedGalaxy[0]); i++ {
		emptyGalaxyCol[i] = '.'
	}
	for _, val := range transposedGalaxy {
		// fmt.Println(val)
		// fmt.Println(emptyGalaxyCol)
		// fmt.Println()
		if reflect.DeepEqual(val, emptyGalaxyCol) {
			// fmt.Println(val)
			for l := 0; l < 2; l++ {
				expandedGalaxy = append(expandedGalaxy, val)
			}
			continue
		}
		expandedGalaxy = append(expandedGalaxy, val)
	}
	// COL Expansion
	// fmt.Println(transpose(expandedGalaxy))
	return transpose(expandedGalaxy)
}

type Position struct {
	x int
	y int
}

func getLocations(galaxyMap [][]rune) map[int]Position {
	locations := make(map[int]Position)
	index := 0
	for i, val := range galaxyMap {
		for j, space := range val {
			if space == '#' {
				locations[index] = Position{
					i, j,
				}
				index++
			}
		}
	}
	return locations
}
func mathMod(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func getTotal(galaxyMap [][]rune) int {
	locations := getLocations(galaxyMap)
	n := len(locations)
	total := 0
	for k, _ := range locations {
		for i := k + 1; i < n; i++ {
			total += mathMod(locations[k].x-locations[i].x) + mathMod(locations[k].y-locations[i].y)
		}
	}
	return total
}

func getTotalV2(galaxyMap [][]rune, locations map[int]Position, count int) int {
	n, m := len(galaxyMap), len(galaxyMap[0])
	currentPos := Position{
		0, 0,
	}
	// TODO : DELETE THIS LINE
	_ = n
	emptyGalaxyRow := make([]rune, m)
	for i := 0; i < m; i++ {
		emptyGalaxyRow[i] = '.'
	}
	//::ROW::

	for _, val := range galaxyMap {
		if reflect.DeepEqual(val, emptyGalaxyRow) {
			for k, l := range locations {
				if l.x > currentPos.x {
					locations[k] = Position{
						l.x + count - 1,
						l.y,
					}
				}
			}
			currentPos.x += count
			continue
		}
		currentPos.x++
	}
	galaxyMap = transpose(galaxyMap)
	for _, val := range galaxyMap {
		if reflect.DeepEqual(val, emptyGalaxyRow) {
			for k, l := range locations {
				if l.y > currentPos.y {
					locations[k] = Position{
						l.x,
						l.y + count - 1,
					}
				}
			}
			currentPos.y += count
			continue
		}
		currentPos.y++
	}
	n = len(locations)
	total := 0
	for k, _ := range locations {
		for i := k + 1; i < n; i++ {
			total += mathMod(locations[k].x-locations[i].x) + mathMod(locations[k].y-locations[i].y)
		}
	}
	return total

}
func part2(content *string) {
	galaxyMap := parseInput(content)
	locations := getLocations(galaxyMap)
	total := getTotalV2(galaxyMap, locations, 1000000)
	fmt.Printf("Expanded Galaxy: %d\n", total)

}
func part1(content *string) {
	galaxyMap := parseInput(content)
	locations := getLocations(galaxyMap)
	// fmt.Printf("Galaxy Map: %v\n", galaxyMap)
	total := getTotalV2(galaxyMap, locations, 2)
	fmt.Printf("Expanded Galaxy: %d\n", total)

}

func main() {
	start := time.Now()
	path := os.Args[1]
	content := readFile(path)
	part2(&content)
	fmt.Printf("Took: %v\n", time.Since(start))
}
