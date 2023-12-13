package main

import (
	_ "embed"
	"fmt"
	adv "github.com/neilgarb/advent2023"
)

//go:embed input.txt
var input []byte

const sample = `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`

func main() {
	part1(adv.Lines(sample))
	part1(adv.Lines(input))
	part2(adv.Lines(sample))
	part2(adv.Lines(input))
}

func part1(lines []string) {
	var tot int
	var cur []string
	for _, l := range append(lines, "") {
		if l == "" {
			tot += calculate(cur, 0)
			cur = nil
			continue
		}
		cur = append(cur, l)
	}
	fmt.Println(tot)
}

func part2(lines []string) {
	var tot int
	var cur []string
	for _, l := range append(lines, "") {
		if l == "" {
			before := calculate(cur, 0)
		CurLoop:
			for y, row := range cur {
				for x, c := range row {
					if c == '.' {
						c = '#'
					} else {
						c = '.'
					}
					next := calculate(flip(cur, y, x, c), before)
					if next > 0 && next != before {
						tot += next
						break CurLoop
					}
				}
			}
			cur = nil
			continue
		}
		cur = append(cur, l)
	}
	fmt.Println(tot)
}

func flip(lines []string, y, x int, newChar rune) []string {
	var res []string
	res = append(res, lines[:y]...)
	res = append(res, lines[y][:x]+string(newChar)+lines[y][x+1:])
	res = append(res, lines[y+1:]...)
	return res
}

func calculate(lines []string, before int) int {
	if c := calculateRows(lines, before, 100); c > 0 {
		return c
	}
	return calculateRows(rotate(lines), before, 1)
}

func calculateRows(lines []string, before, mult int) int {
	for i := 1; i < len(lines); i++ {
		if lines[i] == lines[i-1] {
			if mirrorRows(lines[:i], lines[i:]) && i*mult != before {
				return i * mult
			}
		}
	}
	return 0
}

func mirrorRows(above, below []string) bool {
	for i := len(above) - 1; i >= 0; i-- {
		if len(above)-1-i >= len(below) {
			return true
		}
		if above[i] != below[len(above)-1-i] {
			return false
		}
	}
	return true
}

func rotate(lines []string) []string {
	cols := make([]string, len(lines[0]))
	for _, l := range lines {
		for x, c := range l {
			cols[x] += string(c)
		}
	}
	return cols
}
