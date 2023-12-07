package main

import (
	_ "embed"
	"fmt"
	adv "github.com/neilgarb/advent2023"
	"sort"
	"strings"
)

//go:embed input.txt
var input []byte

const sample = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

func main() {
	part1(adv.Lines(sample))
	part1(adv.Lines(input))
	part2(adv.Lines(sample))
	part2(adv.Lines(input))
}

func part1(lines []string) {
	part(lines, cards)
}

func part2(lines []string) {
	part(lines, cards2)
}

func part(lines []string, ranks string) {
	var scores [][3]int
	var hands []string
	for i, l := range lines {
		hand, bid, _ := strings.Cut(l, " ")
		hands = append(hands, hand)
		scores = append(scores, [3]int{score(hand, ranks[0] == 'J'), adv.Atoi(bid), i})
	}
	sort.Slice(scores, func(i, j int) bool {
		if scores[i][0] == scores[j][0] {
			hand1 := hands[scores[i][2]]
			hand2 := hands[scores[j][2]]
			for k := 0; k < 5; k++ {
				if hand1[k] == hand2[k] {
					continue
				}
				return strings.IndexByte(ranks, hand1[k]) < strings.IndexByte(ranks, hand2[k])
			}
			return false
		}
		return scores[i][0] < scores[j][0]
	})
	var tot int
	for i, hand := range scores {
		tot += (i + 1) * hand[1]
	}
	fmt.Println(tot)
}

const (
	FiveOfAKind  = 7
	FourOfAKind  = 6
	FullHouse    = 5
	ThreeOfAKind = 4
	TwoPair      = 3
	Pair         = 2
	HighCard     = 1

	cards  = "23456789TJQKA"
	cards2 = "J23456789TQKA"
)

func score(hand string, useJokers bool) int {
	counts := make(map[rune]int)
	for _, c := range hand {
		counts[c]++
	}
	var jokers int
	if useJokers {
		jokers = counts['J']
		delete(counts, 'J')
	}
	if jokers == 5 {
		return FiveOfAKind
	}
	var three bool
	var pairs int
	for _, cnt := range counts {
		if cnt == 5 {
			return FiveOfAKind
		} else if cnt == 4 {
			if jokers == 1 {
				return FiveOfAKind
			}
			return FourOfAKind
		} else if cnt == 3 {
			three = true
		} else if cnt == 2 {
			pairs++
		}
	}
	if three {
		if jokers == 2 {
			return FiveOfAKind
		} else if jokers == 1 {
			return FourOfAKind
		}
		if pairs == 1 {
			return FullHouse
		}
		return ThreeOfAKind
	} else if pairs == 2 {
		if jokers == 1 {
			return FullHouse
		}
		return TwoPair
	} else if pairs == 1 {
		if jokers == 3 {
			return FiveOfAKind
		} else if jokers == 2 {
			return FourOfAKind
		} else if jokers == 1 {
			return ThreeOfAKind
		}
		return Pair
	}
	if jokers == 4 {
		return FiveOfAKind
	} else if jokers == 3 {
		return FourOfAKind
	} else if jokers == 2 {
		return ThreeOfAKind
	} else if jokers == 1 {
		return Pair
	}
	return HighCard
}
