package main

import (
	_ "embed"
	"fmt"
	adv "github.com/neilgarb/advent2023"
)

//go:embed input.txt
var input []byte

const sample = `-L|F7
7S-7|
L|7||
-L-J|
L|-JF`

const sample2 = `FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L`

func main() {
	part1(adv.Lines(sample))
	part1(adv.Lines(input))
	part2(adv.Lines(sample2))
	part2(adv.Lines(input))
}

func part1(lines []string) {
	start := findStart(lines)
	loop := findLoop(lines, start)
	fmt.Println(len(loop)/2 + 1)
}

func part2(lines []string) {
	start := findStart(lines)
	loop := findLoop(lines, start)
	var tot int
	for y, l := range lines {
		var inside bool
		var wall byte
		for x, c := range l {
			if c == 'S' {
				c = 'L'
			}
			if loop[adv.Coord2D{X: x, Y: y}] {
				if c == '|' {
					inside = !inside
					wall = '|'
				} else if c == 'F' {
					wall = 'F'
				} else if c == 'L' {
					wall = 'L'
				} else if c == 'J' {
					if wall == 'F' {
						inside = !inside
					}
					wall = 'J'
				} else if c == '7' {
					if wall == 'L' {
						inside = !inside
					}
					wall = '7'
				}
			} else if inside {
				tot++
			}
		}
	}
	fmt.Println(tot)
}

func move(lines []string, start adv.Coord2D, seen map[adv.Coord2D]bool) (next adv.Coord2D, ok bool) {
	c := charat(lines, start)
	up := adv.Coord2D{X: start.X, Y: start.Y - 1}
	down := adv.Coord2D{X: start.X, Y: start.Y + 1}
	left := adv.Coord2D{X: start.X - 1, Y: start.Y}
	right := adv.Coord2D{X: start.X + 1, Y: start.Y}
	if start.Y > 0 && (charat(lines, up) == '|' || charat(lines, up) == '7' || charat(lines, up) == 'F') && (c == 'S' || c == 'J' || c == '|' || c == 'L') && !seen[up] {
		return up, true
	}
	if start.X > 0 && (charat(lines, left) == '-' || charat(lines, left) == 'L' || charat(lines, left) == 'F') && (c == 'S' || c == 'J' || c == '-' || c == '7') && !seen[left] {
		return left, true
	}
	if start.Y < len(lines)-1 && (charat(lines, down) == '|' || charat(lines, down) == 'L' || charat(lines, down) == 'J') && (c == 'S' || c == '7' || c == '|' || c == 'F') && !seen[down] {
		return down, true
	}
	if start.X < len(lines[0])-1 && (charat(lines, right) == '-' || charat(lines, right) == '7' || charat(lines, right) == 'J') && (c == 'S' || c == 'F' || c == '-' || c == 'L') && !seen[right] {
		return right, true
	}
	return adv.Coord2D{}, false
}

func charat(lines []string, coord adv.Coord2D) byte {
	return lines[coord.Y][coord.X]
}

func findStart(lines []string) adv.Coord2D {
	for y, l := range lines {
		for x, c := range l {
			if c == 'S' {
				return adv.Coord2D{X: x, Y: y}
			}
		}
	}
	return adv.Coord2D{}
}

func findLoop(lines []string, start adv.Coord2D) map[adv.Coord2D]bool {
	seen := map[adv.Coord2D]bool{start: true}
	for {
		var ok bool
		start, ok = move(lines, start, seen)
		if !ok {
			break
		}
		seen[start] = true
	}
	return seen
}
