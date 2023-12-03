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
			if set.X > 12 || set.Y > 13 || set.Z > 14 {
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
			c.X = max(c.X, set.X)
			c.Y = max(c.Y, set.Y)
			c.Z = max(c.Z, set.Z)
		}
		tot += c.X * c.Y * c.Z
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
					c.X = adv.Atoi(count)
				} else if colour == "green" {
					c.Y = adv.Atoi(count)
				} else {
					c.Z = adv.Atoi(count)
				}
			}
			row = append(row, c)
		}
		res = append(res, row)
	}
	return res
}
