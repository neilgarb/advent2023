package main

import (
	_ "embed"
	"fmt"
	adv "github.com/neilgarb/advent2023"
)

//go:embed input.txt
var input []byte

const sample = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

func main() {
	part1(adv.Lines(sample))
	part1(adv.Lines(input))
	part2(adv.Lines(sample))
	part2(adv.Lines(input))
}

func part1(lines []string) {
	part(lines, 1)
}

func part2(lines []string) {
	part(lines, 999999)
}

func part(lines []string, add int) {
	var galaxies []adv.Coord2D
	galRows := make(map[int]bool)
	galCols := make(map[int]bool)
	for y, l := range lines {
		for x, c := range l {
			if c == '#' {
				galaxies = append(galaxies, adv.Coord2D{x, y})
				galRows[y] = true
				galCols[x] = true
			}
		}
	}
	var tot int
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			tot += distance(galaxies[i], galaxies[j], galRows, galCols, add)
		}
	}
	fmt.Println(tot)
}

func distance(gal1, gal2 adv.Coord2D, galRows, galCols map[int]bool, add int) int {
	minX, maxX := min(gal1[0], gal2[0]), max(gal1[0], gal2[0])
	minY, maxY := min(gal1[1], gal2[1]), max(gal1[1], gal2[1])
	tot := (maxX - minX) + (maxY - minY)
	for y := minY + 1; y < maxY; y++ {
		if !galRows[y] {
			tot += add
		}
	}
	for x := minX + 1; x < maxX; x++ {
		if !galCols[x] {
			tot += add
		}
	}
	return tot
}
