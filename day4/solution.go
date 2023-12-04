package main

import (
	_ "embed"
	"fmt"
	adv "github.com/neilgarb/advent2023"
	"strings"
)

//go:embed input.txt
var input []byte

const sample = `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

func main() {
	part1(adv.Lines(sample))
	part1(adv.Lines(input))
	part2(adv.Lines(sample))
	part2(adv.Lines(input))
}

func part1(lines []string) {
	var tot int
	for _, l := range lines {
		_, l, _ = strings.Cut(l, ": ")
		winning, actual, _ := strings.Cut(l, " | ")
		winningMap := make(map[int]bool)
		for _, n := range strings.Fields(winning) {
			winningMap[adv.Atoi(n)] = true
		}
		var score int
		for _, n := range strings.Fields(actual) {
			if winningMap[adv.Atoi(n)] {
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
			}
		}
		tot += score
	}
	fmt.Println(tot)
}

func part2(lines []string) {
	var copies []int
	var tot int
	for _, l := range lines {
		_, l, _ = strings.Cut(l, ": ")
		winning, actual, _ := strings.Cut(l, " | ")
		winningMap := make(map[int]bool)
		for _, n := range strings.Fields(winning) {
			winningMap[adv.Atoi(n)] = true
		}
		var score int
		for _, n := range strings.Fields(actual) {
			if winningMap[adv.Atoi(n)] {
				score++
			}
		}
		tot++
		add := 1
		if len(copies) > 0 {
			tot += copies[0]
			add += copies[0]
			copies = copies[1:]
		}
		for i := 0; i < score; i++ {
			if i < len(copies) {
				copies[i] += add
			} else {
				copies = append(copies, add)
			}
		}
	}
	fmt.Println(tot)
}
