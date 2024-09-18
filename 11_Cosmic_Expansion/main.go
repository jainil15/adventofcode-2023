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
			expandedGalaxy = append(expandedGalaxy, val, val)
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
			expandedGalaxy = append(expandedGalaxy, val, val)
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
func part1(content *string) {
	galaxyMap := parseInput(content)
	// fmt.Printf("Galaxy Map: %v\n", galaxyMap)
	expandedGalaxy := expandGalaxy(galaxyMap)
	total := getTotal(expandedGalaxy)

	fmt.Printf("Expanded Galaxy: %d\n", total)

}

func main() {
	start := time.Now()
	path := os.Args[1]
	content := readFile(path)
	part1(&content)
	fmt.Printf("Took: %v\n", time.Since(start))
}
