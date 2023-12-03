package main

import (
	_ "embed"
	"fmt"
	adv "github.com/neilgarb/advent2023"
	"strings"
	"unicode"
)

//go:embed input.txt
var input []byte

func main() {
	part1(adv.Lines(`1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`))
	part1(adv.Lines(input))
	part2(adv.Lines(`two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`))
	part2(adv.Lines(input))
}

func part1(lines []string) {
	var tot int
	for _, l := range lines {
		var nums []int
		for _, c := range l {
			if unicode.IsDigit(c) {
				nums = append(nums, adv.Atoi(string(c)))
			}
		}
		tot += nums[0]*10 + nums[len(nums)-1]
	}
	fmt.Println(tot)
}

func part2(lines []string) {
	var tot int
	for _, l := range lines {
		var nums []int
	LineLoop:
		for i, c := range l {
			for k, v := range digits {
				if strings.HasPrefix(l[i:], k) {
					nums = append(nums, v)
					continue LineLoop
				}
			}
			if unicode.IsDigit(c) {
				nums = append(nums, adv.Atoi(string(c)))
			}
		}
		tot += nums[0]*10 + nums[len(nums)-1]
	}
	fmt.Println(tot)
}

var digits = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}
