package main

import (
	_ "embed"
	"fmt"
	adv "github.com/neilgarb/advent2023"
	"unicode"
)

//go:embed input.txt
var input []byte

const sample = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

func main() {
	part1(adv.Lines(sample))
	part1(adv.Lines(input))
	part2(adv.Lines(sample))
	part2(adv.Lines(input))
}

func part1(lines []string) {
	var num string
	var adjacent bool
	var tot int
	add := func() {
		if adjacent {
			tot += adv.Atoi(num) // num might be "".
		}
		num = ""
		adjacent = false
	}
	for y, l := range lines {
		for x, c := range l {
			if unicode.IsDigit(c) {
				num += string(c)
			AdjLoop:
				for i := x - 1; i <= x+1; i++ {
					for j := y - 1; j <= y+1; j++ {
						if (i != x || j != y) && j >= 0 && i >= 0 && j <= len(lines)-1 && i <= len(l)-1 {
							if !unicode.IsDigit(rune(lines[j][i])) && lines[j][i] != '.' {
								adjacent = true
								continue AdjLoop
							}
						}
					}
				}
			} else {
				add()
			}
		}
		add()
	}
	fmt.Println(tot)
}

func part2(lines []string) {
	gears := make(map[adv.Coord2D]int)
	var num string
	var gear *adv.Coord2D
	var tot int
	add := func() {
		if gear != nil {
			if gears[*gear] > 0 {
				tot += adv.Atoi(num) * gears[*gear]
			} else {
				gears[*gear] = adv.Atoi(num)
			}
		}
		num = ""
		gear = nil
	}
	for y, l := range lines {
		for x, c := range l {
			if unicode.IsDigit(c) {
				num += string(c)
			AdjLoop:
				for i := x - 1; i <= x+1; i++ {
					for j := y - 1; j <= y+1; j++ {
						if (i != x || j != y) && j >= 0 && i >= 0 && j <= len(lines)-1 && i <= len(l)-1 {
							if lines[j][i] == '*' {
								gear = &adv.Coord2D{i, j}
								continue AdjLoop
							}
						}
					}
				}
			} else {
				add()
			}
		}
		add()
	}
	fmt.Println(tot)
}
