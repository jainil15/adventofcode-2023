package main

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

var cardValue = map[rune]int{
	'A': 13,
	'K': 12,
	'Q': 11,
	'J': 10,
	'T': 9,
	'9': 8,
	'8': 7,
	'7': 6,
	'6': 5,
	'5': 4,
	'4': 3,
	'3': 2,
	'2': 1,
}

const (
	singlePair = 20
	doublePair = 40
	triple     = 80
	fullHouse  = 160
	four       = 320
	five       = 640
)

type strength struct {
	str int
	max int
	pos float64
	bet int
}

func readFile(path string) string {
	out, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(out)
}
func isFive(cardMap map[string]int) bool {
	for k := range cardMap {
		if cardMap[k] == 5 {
			return true
		}
		break
	}
	return false
}
func isFour(cardMap map[string]int) bool {
	for k := range cardMap {
		if cardMap[k] == 4 {
			return true
		}
	}
	return false
}
func isFullHouse(cardMap map[string]int) bool {
	keys := reflect.ValueOf(cardMap).MapKeys()
	if len(keys) != 2 {
		return false
	}
	if (cardMap[keys[0].String()] == 3 && cardMap[keys[1].String()] == 2) || (cardMap[keys[1].String()] == 2 && cardMap[keys[0].String()] == 3) {
		return true
	}
	return false
}
func isThree(cardMap map[string]int) bool {
	for k := range cardMap {
		if cardMap[k] == 3 {
			return true
		}
	}
	return false
}
func isDoublePair(cardMap map[string]int) bool {
	doubleCount := 0
	for k := range cardMap {
		if cardMap[k] == 2 {
			doubleCount++
		}
		if doubleCount == 2 {
			return true
		}
	}
	return false
}
func isSinglePair(cardMap map[string]int) bool {
	for k := range cardMap {
		if cardMap[k] == 2 {
			return true
		}
	}
	return false
}
func calcStr(hand string, bet int) strength {
	cardMap := make(map[string]int)
	max := 0
	total := 0
	pos := "0."
	for _, c := range hand {
		cardMap[string(c)]++
		total += cardValue[c]
		pos += strconv.Itoa(cardValue[c])
		if max < cardMap[string(c)] {
			max = cardMap[string(c)]
		}
	}
	if isFive(cardMap) {
		total = total * five
	} else if isFour(cardMap) {
		total = total * four
	} else if isFullHouse(cardMap) {
		total = total * fullHouse
	} else if isThree(cardMap) {
		total = total * triple
	} else if isDoublePair(cardMap) {
		total = total * doublePair
	} else if isSinglePair(cardMap) {
		total = total * singlePair
	} else {
		total = total * max
	}

	fmt.Println(pos)
	p, err := strconv.ParseFloat(pos, 64)
	if err != nil {
		panic(err)
	}
	handStrength := strength{
		max: max,
		pos: p,
		str: total,
		bet: bet,
	}
	fmt.Printf("Hand: %v, Str: %v Pos: %v HandStr: %v\n", hand, total, p, handStrength)
	return handStrength
}

func parseInput(input *string) ([]string, []int) {
	lines := strings.Split(*input, "\n")
	lines = lines[:len(lines)-1]
	hand := make([]string, len(lines))
	bet := make([]int, len(lines))
	for i, val := range lines {
		o := strings.Split(val, " ")
		hand[i] = o[0]
		b, err := strconv.Atoi(o[1])
		if err != nil {
			panic(err)
		}
		bet[i] = b
	}
	return hand, bet
}

func main() {
	start := time.Now()
	path := os.Args[1]
	input := readFile(path)
	hands, bet := parseInput(&input)
	fmt.Println(hands, bet)
	handsSort := []strength{}
	for i, hand := range hands {
		strength := calcStr(hand, bet[i])
		handsSort = append(handsSort, strength)
	}
	sort.Slice(handsSort, func(i int, j int) bool {
		return handsSort[i].str > handsSort[j].str
	})
	total := 0
	for i, hand := range handsSort {
		total = total + (hand.bet * (i + 1))
	}
	fmt.Printf("Hands Sorted: %v\n", handsSort)
	fmt.Println(total)
	fmt.Printf("Time Req: %v\n", time.Since(start))
}
