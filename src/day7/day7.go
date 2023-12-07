package day7

import (
	"sort"
	"strings"

	"github.com/snopan/aoc-2023/src/helpers"
	"golang.org/x/exp/maps"
)

func SolutionPart1(input string) int {
	return Solution(input, false)
}

func SolutionPart2(input string) int {
	return Solution(input, true)
}

func Solution(input string, handleJoker bool) int {
	lines := strings.Split(input, "\n")

	handBids := map[string]int{}
	typeHands := make(map[HandType][]string)

	for _, line := range lines {
		hand := line[:5]
		bid := helpers.GetNumbersFromText(line[5:])[0]
		handBids[hand] = bid

		handType := GetHandType(hand, handleJoker)
		println(hand, handType.String())
		typeHands[handType] = append(typeHands[handType], hand)
	}

	currRank := 1
	winnings := 0
	for _, handType := range HandTypeRanking {
		hands, ok := typeHands[handType]
		if !ok {
			continue
		}

		if handleJoker {
			sort.Sort(SortJokerHands(hands))
		} else {
			sort.Sort(SortHands(hands))
		}

		for _, hand := range hands {
			bid := handBids[hand]
			winnings += currRank * bid
			currRank++
		}
	}

	return winnings
}

func GetHandType(hand string, handlerJoker bool) HandType {
	joker := 0
	labelsCount := make(map[rune]int)
	for _, label := range hand {
		if handlerJoker && label == 'J' {
			joker++
		} else {
			labelsCount[label]++
		}
	}

	if joker == 5 {
		return FiveOfAKind
	}

	counts := maps.Values(labelsCount)
	sort.Sort(sort.Reverse(sort.IntSlice(counts)))
	prevC := 0
	for _, c := range counts {
		if joker > 0 {
			c += joker
			joker = 0
		}
		switch c {
		case 5:
			return FiveOfAKind
		case 4:
			return FourOfAKind
		case 2:
			if prevC == 3 {
				return FullHouse
			} else if prevC == 2 {
				return TwoPair
			} else {
				prevC = c
			}
		case 1:
			if prevC == 3 {
				return ThreeOfAKind
			} else if prevC == 2 {
				return OnePair
			} else {
				return HighCard
			}
		default:
			prevC = c
		}
	}
	panic("failed to parse hand type")
}
