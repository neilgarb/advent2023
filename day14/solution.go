package main

import (
	_ "embed"
	"fmt"
	adv "github.com/neilgarb/advent2023"
)

//go:embed input.txt
var input []byte

const sample = `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`

func main() {
	part1(adv.Lines(sample))
	part1(adv.Lines(input))
	part2(adv.Lines(sample))
	part2(adv.Lines(input))
}

func part1(lines []string) {
	roll(lines)
	fmt.Println(load(lines))
}

func part2(lines []string) {
	for i := 0; i < 1_000; i++ {
		roll(lines) // N
		lines = rotate(lines)
		roll(lines) // W
		lines = rotate(lines)
		roll(lines) // S
		lines = rotate(lines)
		roll(lines) // E
		lines = rotate(lines)

		// fmt.Println(load(lines))
		// TODO(neil): Find the repeating pattern programmatically.
	}
}

func roll(lines []string) {
	for y := 0; y < len(lines); y++ {
		for x, c := range lines[y] {
			if c == 'O' {
				newLine := 0
				for p := y - 1; p >= 0; p-- {
					if lines[p][x] != '.' {
						newLine = p + 1
						break
					}
				}
				lines[y] = lines[y][:x] + "." + lines[y][x+1:]
				lines[newLine] = lines[newLine][:x] + "O" + lines[newLine][x+1:]
			}
		}
	}
}

func load(lines []string) int {
	var tot int
	for y := 0; y < len(lines); y++ {
		for _, c := range lines[y] {
			if c == 'O' {
				tot += len(lines) - y
			}
		}
	}
	return tot
}

func dump(lines []string) {
	for _, l := range lines {
		fmt.Println(l)
	}
	fmt.Println()
}

func rotate(lines []string) []string {
	cols := make([]string, len(lines[0]))
	for y := len(lines) - 1; y >= 0; y-- {
		for x := 0; x < len(lines[y]); x++ {
			cols[x] += string(lines[y][x])
		}
	}
	return cols
}
