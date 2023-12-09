package main

import (
	_ "embed"
	"fmt"
	adv "github.com/neilgarb/advent2023"
	"strings"
)

//go:embed input.txt
var input []byte

const sample = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

func main() {
	part1(adv.Lines(sample))
	part1(adv.Lines(input))
	part2(adv.Lines(sample))
	part2(adv.Lines(input))
}

func part1(lines []string) {
	var tot int
	for _, l := range lines {
		tot += extrapolate(l, false)
	}
	fmt.Println(tot)
}

func part2(lines []string) {
	var tot int
	for _, l := range lines {
		tot += extrapolate(l, true)
	}
	fmt.Println(tot)
}

func extrapolate(l string, first bool) int {
	var nums []int
	for _, n := range strings.Fields(l) {
		nums = append(nums, adv.Atoi(n))
	}
	var stack []int
	for {
		if first {
			stack = append(stack, nums[0])
		} else {
			stack = append(stack, nums[len(nums)-1])
		}
		var diffs []int
		zeroes := true
		for i := 1; i < len(nums); i++ {
			diffs = append(diffs, nums[i]-nums[i-1])
			if nums[i] != nums[i-1] {
				zeroes = false
			}
		}
		if zeroes {
			break
		}
		nums = diffs
	}
	stack = append(stack, 0)
	var tot int
	for i := len(stack) - 1; i >= 0; i-- {
		if first {
			tot = stack[i] - tot
		} else {
			tot += stack[i]
		}
	}
	return tot
}
