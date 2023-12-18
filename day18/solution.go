package main

import (
	_ "embed"
	"fmt"
	adv "github.com/neilgarb/advent2023"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input []byte

const sample = `R 6 (#70c710)
D 5 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)`

func main() {
	part1(adv.Lines(sample))
	part1(adv.Lines(input))
	part2(adv.Lines(sample))
	part2(adv.Lines(input))
}

func part1(lines []string) {
	at := adv.Coord2D{0, 0}
	edge := map[adv.Coord2D]bool{at: true}
	var left, top, bottom, right int
	for _, l := range lines {
		parts := strings.Fields(l)
		dir := parts[0]
		count := adv.Atoi(parts[1])

		var offset adv.Coord2D
		switch dir {
		case "U":
			offset = adv.Coord2D{0, -1}
		case "D":
			offset = adv.Coord2D{0, 1}
		case "L":
			offset = adv.Coord2D{-1, 0}
		case "R":
			offset = adv.Coord2D{1, 0}
		}

		for i := 0; i < count; i++ {
			left = min(at[0], left)
			top = min(at[1], top)
			right = max(at[0], right)
			bottom = max(at[1], bottom)
			at = at.Add(offset)
			edge[at] = true
		}
	}

	dug := map[adv.Coord2D]bool{}
	for y := top; y <= bottom; y++ {
		var inside bool
		var wall byte
		for x := left; x <= right; x++ {
			cur := adv.Coord2D{x, y}
			edgeUp := edge[cur.Add(dirUp)]
			edgeDown := edge[cur.Add(dirDown)]
			edgeLeft := edge[cur.Add(dirLeft)]
			edgeRight := edge[cur.Add(dirRight)]

			if edge[cur] {
				if !edgeLeft && !edgeUp && edgeRight && edgeDown {
					wall = 'F'
				} else if edgeLeft && !edgeUp && !edgeDown && edgeRight {
					// "-"
				} else if !edgeLeft && edgeUp && edgeRight && !edgeDown {
					wall = 'L'
				} else if !edgeLeft && edgeUp && edgeDown && !edgeRight {
					inside = !inside
				} else if edgeLeft && edgeUp && !edgeDown && !edgeRight {
					if wall == 'F' {
						inside = !inside
						wall = 0
					}
				} else if edgeLeft && !edgeUp && edgeDown && !edgeRight {
					if wall == 'L' {
						inside = !inside
						wall = 0
					}
				}
			} else if inside {
				dug[cur] = true
			}
		}
	}
	fmt.Println(len(edge) + len(dug))
}

func part2(lines []string) {
	at := adv.Coord2D{0, 0}
	yMap := map[int]bool{0: true}

	var edges [][4]int
	var wallLen int

	for _, l := range lines {
		parts := strings.Fields(l)
		colour := parts[2]
		colour = strings.Trim(colour, "#()")
		c, _ := strconv.ParseInt(colour[:5], 16, 64)
		count := int(c)
		wallLen += count

		var offset adv.Coord2D
		switch colour[5] {
		case '3':
			offset = adv.Coord2D{0, -count}
			edges = append(edges, [4]int{at[0], at[1] - count, at[1], -1})
		case '1':
			offset = adv.Coord2D{0, count}
			edges = append(edges, [4]int{at[0], at[1], at[1] + count, 1})
		case '2':
			offset = adv.Coord2D{-count, 0}
		case '0':
			offset = adv.Coord2D{count, 0}
		}

		at = at.Add(offset)
		yMap[at[1]] = true
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i][0] < edges[j][0]
	})

	var ys []int
	for k := range yMap {
		ys = append(ys, k)
	}
	sort.Ints(ys)

	tot := wallLen
	for y := ys[0]; y <= ys[len(ys)-1]; y++ {
		var inside bool

		var hits [][4]int
		for _, e := range edges {
			if y >= e[1] && y <= e[2] {
				hits = append(hits, e)
			}
		}

		var prev int
		for x, e := range hits {
			if y > e[1] && y < e[2] {
				if inside {
					tot += e[0] - hits[x-1][0] - 1
				}
				inside = !inside
			} else if y == e[1] {
				if prev == 2 {
					if inside {
						tot += hits[x-1][0] - hits[x-2][0] - 1
					}
					inside = !inside
					prev = 0
				} else if prev == 1 && inside {
					tot += hits[x-1][0] - hits[x-2][0] - 1
					prev = 0
				} else {
					prev = 1
				}
			} else if y == e[2] {
				if prev == 1 {
					if inside {
						tot += hits[x-1][0] - hits[x-2][0] - 1
					}
					inside = !inside
					prev = 0
				} else if prev == 2 && inside {
					tot += hits[x-1][0] - hits[x-2][0] - 1
					prev = 0
				} else {
					prev = 2
				}
			}
		}
	}
	fmt.Println(tot)
}

var (
	dirUp    = adv.Coord2D{0, -1}
	dirDown  = adv.Coord2D{0, 1}
	dirLeft  = adv.Coord2D{-1, 0}
	dirRight = adv.Coord2D{1, 0}
)
