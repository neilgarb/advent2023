package main

import (
	_ "embed"
	"fmt"
	adv "github.com/neilgarb/advent2023"
	"regexp"
	"strings"
)

//go:embed input.txt
var input []byte

func main() {
	part1(adv.LinesFromString(`Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`))
	part1(adv.LinesFromBytes(input))
	part2(adv.LinesFromString(`Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`))
	part2(adv.LinesFromBytes(input))
}

func part1(lines []string) {
	constraint := adv.Coord3D{X: 12, Y: 13, Z: 14}
	games := parseGames(lines)
	var tot int
GameLoop:
	for i, game := range games {
		for _, set := range game {
			if set.X > constraint.X || set.Y > constraint.Y || set.Z > constraint.Z {
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
	numRe := regexp.MustCompile(`(\d+) (red|green|blue)`)
	var res [][]adv.Coord3D
	for _, l := range lines {
		var row []adv.Coord3D
		game := strings.Split(l, ": ")[1]
		sets := strings.Split(game, "; ")
		for _, set := range sets {
			matches := numRe.FindAllStringSubmatch(set, -1)
			var c adv.Coord3D
			for _, m := range matches {
				if m[2] == "red" {
					c.X = adv.Atoi(m[1])
				} else if m[2] == "green" {
					c.Y = adv.Atoi(m[1])
				} else {
					c.Z = adv.Atoi(m[1])
				}
			}
			row = append(row, c)
		}
		res = append(res, row)
	}
	return res
}
