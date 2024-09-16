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
	// 'J': 10,
	'J': 0,
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
		if k == "J" {
			continue
		}
		if cardMap[k] == 5 {
			return true
		}
		break
	}
	return false
}
func isFour(cardMap map[string]int) bool {
	for k := range cardMap {
		if k == "J" {
			continue
		}
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
	if keys[0].String() == "J" || keys[1].String() == "J" {
		return false
	}
	if (cardMap[keys[0].String()] == 3 && cardMap[keys[1].String()] == 2) || (cardMap[keys[0].String()] == 2 && cardMap[keys[1].String()] == 3) {
		return true
	}
	return false
}
func isThree(cardMap map[string]int) bool {
	for k := range cardMap {
		if k == "J" {
			continue
		}
		if cardMap[k] == 3 {
			return true
		}
	}
	return false
}
func isDoublePair(cardMap map[string]int) bool {
	doubleCount := 0
	for k := range cardMap {
		if k == "J" {
			continue
		}
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
		if k == "J" {
			continue
		}
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
		handType = 5 + cardMap["J"]
	} else if isFullHouse(cardMap) {
		handType = 4 + cardMap["J"]
	} else if isThree(cardMap) {
		handType = 3 + cardMap["J"]
		if cardMap["J"] > 0 {
			handType += 1
		}
	} else if isDoublePair(cardMap) {
		handType = 2 + cardMap["J"]
		if cardMap["J"] > 0 {
			handType += 1
		}
	} else if isSinglePair(cardMap) {
		// handType = 1 + cardMap["J"]
		switch cardMap["J"] {
		case 0:
			handType = 1
		case 1:
			handType = 3
		case 2:
			handType = 5
		case 3:
			handType = 6
		}
		cardMap["J"] = 0
	} else {
		switch cardMap["J"] {
		case 0:
			handType = 0
		case 1:
			handType = 1
		case 2:
			handType = 3
		case 3:
			handType = 5
		case 4:
			handType = 6
		case 5:
			handType = 6
		}

	}

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
	handStrength := strength{
		max:      max,
		handType: handType,
		str:      total,
		bet:      bet,
		pos:      pf,
		hand:     hand,
		maxCard:  maxCard,
	}
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
	handsSort := []strength{}
	for i, hand := range hands {
		strength := calcStr(hand, bet[i])
		handsSort = append(handsSort, strength)
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
		total = total + int64(hand.bet*(i+1))
	}
	fmt.Println(total)
	fmt.Printf("Time Req: %v\n", time.Since(start))
}
