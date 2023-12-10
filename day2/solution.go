package main

import (
	_ "embed"
	"fmt"
	adv "github.com/neilgarb/advent2023"
	"strings"
)

//go:embed input.txt
var input []byte

const sample = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

func main() {
	part1(adv.Lines(sample))
	part1(adv.Lines(input))
	part2(adv.Lines(sample))
	part2(adv.Lines(input))
}

func part1(lines []string) {
	games := parseGames(lines)
	var tot int
GameLoop:
	for i, game := range games {
		for _, set := range game {
			if set[0] > 12 || set[1] > 13 || set[2] > 14 {
				continue GameLoop
			}
		}
		tot += i + 1
	}
	fmt.Println(tot)
}

func part2(lines []string) {
	games := parseGames(lines)
	var tot int
	for _, game := range games {
		var c adv.Coord3D
		for _, set := range game {
			c[0] = max(c[0], set[0])
			c[1] = max(c[1], set[1])
			c[2] = max(c[2], set[2])
		}
		tot += c[0] * c[1] * c[2]
	}
	fmt.Println(tot)
}

func parseGames(lines []string) [][]adv.Coord3D {
	var res [][]adv.Coord3D
	for _, l := range lines {
		var row []adv.Coord3D
		_, sets, _ := strings.Cut(l, ": ")
		for _, set := range strings.Split(sets, "; ") {
			var c adv.Coord3D
			for _, pick := range strings.Split(set, ", ") {
				count, colour, _ := strings.Cut(pick, " ")
				if colour == "red" {
					c[0] = adv.Atoi(count)
				} else if colour == "green" {
					c[1] = adv.Atoi(count)
				} else {
					c[2] = adv.Atoi(count)
				}
			}
			row = append(row, c)
		}
		res = append(res, row)
	}
	return res
}
