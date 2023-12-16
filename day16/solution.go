package main

import (
	_ "embed"
	"fmt"
	adv "github.com/neilgarb/advent2023"
)

//go:embed input.txt
var input []byte

const sample = `.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`

func main() {
	part1(adv.Lines(sample))
	part1(adv.Lines(input))
	part2(adv.Lines(sample))
	part2(adv.Lines(input))
}

type beam [2]adv.Coord2D

func part1(lines []string) {
	e := energised(lines, beam{{-1, 0}, {1, 0}})
	fmt.Println(e)
}

func part2(lines []string) {
	var e int
	width := len(lines[0])
	for y := 0; y < len(lines); y++ {
		e = max(e, energised(lines, beam{{-1, y}, {1, 0}}))
		e = max(e, energised(lines, beam{{width, y}, {-1, 0}}))
	}
	for x := 0; x < width; x++ {
		e = max(e, energised(lines, beam{{x, -1}, {0, 1}}))
		e = max(e, energised(lines, beam{{x, len(lines)}, {0, -1}}))
	}
	fmt.Println(e)
}

func energised(lines []string, start beam) int {
	beams := []beam{start}
	seen := make(map[adv.Coord2D]map[adv.Coord2D]bool)
	for {
		var moved bool
		var newBeams []beam
		for i, b := range beams {
			pos := b[0].Add(b[1])
			if seen[pos][b[1]] {
				continue
			}
			if pos[0] < 0 || pos[0] >= len(lines[0]) || pos[1] < 0 || pos[1] >= len(lines) {
				continue
			}
			if seen[pos] == nil {
				seen[pos] = make(map[adv.Coord2D]bool)
			}
			seen[pos][b[1]] = true
			moved = true
			beams[i][0] = pos
			c := lines[pos[1]][pos[0]]
			if c == '/' || c == '\\' {
				beams[i][1] = reflect(beams[i][1], c)
			} else if c == '-' && b[1][1] != 0 {
				beams[i][1] = adv.Coord2D{-1, 0}
				newBeams = append(newBeams, beam{pos, adv.Coord2D{1, 0}})
			} else if c == '|' && b[1][0] != 0 {
				beams[i][1] = adv.Coord2D{0, -1}
				newBeams = append(newBeams, beam{pos, adv.Coord2D{0, 1}})
			}
		}
		if !moved {
			break
		}
		beams = append(beams, newBeams...)
	}
	return len(seen)
}

func reflect(dir adv.Coord2D, c byte) adv.Coord2D {
	if c == '/' {
		if dir[0] == 0 {
			dir[0], dir[1] = -dir[1], 0
		} else {
			dir[0], dir[1] = 0, -dir[0]
		}
	} else if c == '\\' {
		dir[0], dir[1] = dir[1], dir[0]
	}
	return dir
}
