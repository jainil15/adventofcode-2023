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
	handType int
	str      int
	max      int
	pos      float64
	bet      int
	hand     string
	maxCard  string
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
	if (cardMap[keys[0].String()] == 3 && cardMap[keys[1].String()] == 2) || (cardMap[keys[0].String()] == 2 && cardMap[keys[1].String()] == 3) {
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
	maxCard := ""
	total := 0
	pos := ""
	for _, c := range hand {
		cardMap[string(c)]++
		total += cardValue[c]
		pos += strconv.FormatInt(int64(cardValue[c]), 16)
		if max < cardMap[string(c)] {
			max = cardMap[string(c)]
			maxCard = string(c)
		}
	}
	handType := 0
	if isFive(cardMap) {
		handType = 6
	} else if isFour(cardMap) {
		handType = 5
	} else if isFullHouse(cardMap) {
		handType = 4
	} else if isThree(cardMap) {
		handType = 3
	} else if isDoublePair(cardMap) {
		handType = 2
	} else if isSinglePair(cardMap) {
		handType = 1
	} else {
		handType = 0
	}

	fmt.Println(pos)
	p, err := strconv.ParseUint(pos, 16, 64)
	if err != nil {
		panic(err)
	}
	sp := strconv.FormatInt(int64(p), 10)
	sp = "0." + sp
	pf, err := strconv.ParseFloat(sp, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(pf)
	handStrength := strength{
		max:      max,
		handType: handType,
		str:      total,
		bet:      bet,
		pos:      pf,
		hand:     hand,
		maxCard:  maxCard,
	}
	fmt.Printf("Hand: %v, Str: %v Pos: %v HandStr: %v\n", hand, total, handStrength)
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
func isGreater(hand1 string, hand2 string) bool {
	n := len(hand1)
	for i := 0; i < n; i++ {
		if cardValue[rune(hand1[i])] > cardValue[rune(hand2[i])] {
			return false
		} else if cardValue[rune(hand1[i])] < cardValue[rune(hand2[i])] {
			return true
		}
	}
	return true
}
func matchHandSort(handSort1, handSort2 string) bool {
	n := len(handSort1)
	for i := 0; i < n; i++ {
		if handSort1[i] != handSort2[i] {
			fmt.Printf("\nHandsSort1: %v\nHandsSort2: %v\n\n\n", handSort1[i-40:i+80], handSort2[i-40:i+80])
			return false
		}
	}
	return true
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
	o, err := os.ReadFile("output.txt")
	if err != nil {
		panic(err)
	}
	os.WriteFile("output.txt",
		[]byte(fmt.Sprintf("%v", handsSort)),
		0644)
	if matchHandSort(string(o), fmt.Sprintf("%v", handsSort)) {
		fmt.Println("Output matches")
	}
	if string(o) != fmt.Sprintf("%v", handsSort) {
		panic("Output does not match")
	}
	sort.SliceStable(handsSort, func(i int, j int) bool {
		if handsSort[i].handType == handsSort[j].handType {
			// return handsSort[i].pos < handsSort[j].pos
			return isGreater(handsSort[i].hand, handsSort[j].hand)
		}
		return handsSort[i].handType < handsSort[j].handType
	})
	var total int64 = 0
	for i, hand := range handsSort {
		fmt.Println(hand.bet * (i + 1))
		total = total + int64(hand.bet*(i+1))
	}
	fmt.Printf("Hands Sorted: %+v\n", handsSort)
	fmt.Println(total)
	fmt.Printf("Time Req: %v\n", time.Since(start))
}
