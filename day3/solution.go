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
	part1(adv.LinesFromString(sample))
	part1(adv.LinesFromBytes(input))
	part2(adv.LinesFromString(sample))
	part2(adv.LinesFromBytes(input))
}

func part1(lines []string) {
	var curNum string
	var isAdjacent bool
	var tot int
	add := func() {
		if isAdjacent {
			tot += adv.Atoi(curNum) // curNum might be "".
		}
		curNum = ""
		isAdjacent = false
	}
	for y, l := range lines {
		for x, c := range l {
			if unicode.IsDigit(c) {
				curNum += string(c)
			AdjLoop:
				for i := x - 1; i <= x+1; i++ {
					for j := y - 1; j <= y+1; j++ {
						if (i != x || j != y) && j >= 0 && i >= 0 && j <= len(lines)-1 && i <= len(l)-1 {
							if !unicode.IsDigit(rune(lines[j][i])) && lines[j][i] != '.' {
								isAdjacent = true
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
	var curNum string
	var adjacentToGear *adv.Coord2D
	var tot int
	add := func() {
		if adjacentToGear != nil {
			if gears[*adjacentToGear] > 0 {
				tot += adv.Atoi(curNum) * gears[*adjacentToGear]
			} else {
				gears[*adjacentToGear] = adv.Atoi(curNum)
			}
		}
		curNum = ""
		adjacentToGear = nil
	}
	for y, l := range lines {
		for x, c := range l {
			if unicode.IsDigit(c) {
				curNum += string(c)
			AdjLoop:
				for i := x - 1; i <= x+1; i++ {
					for j := y - 1; j <= y+1; j++ {
						if (i != x || j != y) && j >= 0 && i >= 0 && j <= len(lines)-1 && i <= len(l)-1 {
							if lines[j][i] == '*' {
								adjacentToGear = &adv.Coord2D{X: i, Y: j}
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
