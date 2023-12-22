package main

import (
	_ "embed"
	"fmt"
	adv "github.com/neilgarb/advent2023"
	"math"
)

//go:embed input.txt
var input []byte

const sample = `...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........`

func main() {
	part1(adv.Lines(sample), 6)
	part1(adv.Lines(input), 64)
	part2(adv.Lines(input), 26501365)
}

func part1(lines []string, maxSteps int) {
	dist := dijkstra(lines, findStart(lines))
	var tot int
	for _, d := range dist {
		if d%2 == 0 && d <= maxSteps {
			tot++
		}
	}
	fmt.Println(tot)
}

func part2(lines []string, steps int) {
	start := findStart(lines)
	dist := dijkstra(lines, start)
	ignoreList := map[adv.Coord2D]bool{}
	for y, l := range lines {
		for x, c := range l {
			if c == '#' {
				continue
			}
			d, ok := dist[adv.Coord2D{x, y}]
			if !ok || d == math.MaxInt {
				ignoreList[adv.Coord2D{x, y}] = true
			}
		}
	}

	oddCounts := map[int]int{}
	evenCounts := map[int]int{}

	// Full.
	for _, d := range dist {
		if d%2 == 1 {
			oddCounts[0]++
		} else {
			evenCounts[0]++
		}
	}

	// NE corner
	for y := 0; y < len(lines)/2; y++ {
		for x := len(lines[0])/2 + y + 1; x < len(lines[0]); x++ {
			c := adv.Coord2D{x, y}
			if ignoreList[c] {
				continue
			}
			if lines[c[1]][c[0]] == '#' {
				continue
			}
			if c[0]%2 != c[1]%2 {
				oddCounts[1]++
			} else {
				evenCounts[1]++
			}
		}
	}

	// NW corner
	for y := 0; y < len(lines)/2; y++ {
		for x := 0; x < len(lines)/2-y; x++ {
			c := adv.Coord2D{x, y}
			if ignoreList[c] {
				continue
			}
			if lines[c[1]][c[0]] == '#' {
				continue
			}
			if c[0]%2 != c[1]%2 {
				oddCounts[2]++
			} else {
				evenCounts[2]++
			}
		}
	}

	// SW
	for y := len(lines)/2 + 1; y < len(lines); y++ {
		for x := 0; x < y-len(lines)/2; x++ {
			c := adv.Coord2D{x, y}
			if ignoreList[c] {
				continue
			}
			if lines[c[1]][c[0]] == '#' {
				continue
			}
			if c[0]%2 != c[1]%2 {
				oddCounts[3]++
			} else {
				evenCounts[3]++
			}
		}
	}

	// SE corner
	for y := len(lines)/2 + 1; y < len(lines); y++ {
		for x := len(lines[0]) - 1 - (y - len(lines)/2 - 1); x < len(lines[0]); x++ {
			c := adv.Coord2D{x, y}
			if ignoreList[c] {
				continue
			}
			if lines[c[1]][c[0]] == '#' {
				continue
			}
			if c[0]%2 != c[1]%2 {
				oddCounts[4]++
			} else {
				evenCounts[4]++
			}
		}
	}

	// Full blocks ✔️
	var tot int
	tot += 202299 * 202300 / 2 * oddCounts[0]
	tot += 202298 * 202299 / 2 * oddCounts[0]
	tot += 202300 * 202301 / 2 * evenCounts[0]
	tot += 202299 * 202300 / 2 * evenCounts[0]

	// Apexes ✔️
	tot += oddCounts[0] - oddCounts[4] - oddCounts[1]
	tot += oddCounts[0] - oddCounts[1] - oddCounts[2]
	tot += oddCounts[0] - oddCounts[2] - oddCounts[3]
	tot += oddCounts[0] - oddCounts[3] - oddCounts[4]

	// Corners ✔️
	tot += 202300 * evenCounts[1]
	tot += 202300 * evenCounts[2]
	tot += 202300 * evenCounts[3]
	tot += 202300 * evenCounts[4]

	// Wodges ✔️
	tot += 202299 * (oddCounts[0] - oddCounts[1])
	tot += 202299 * (oddCounts[0] - oddCounts[2])
	tot += 202299 * (oddCounts[0] - oddCounts[3])
	tot += 202299 * (oddCounts[0] - oddCounts[4])

	fmt.Println(tot)
}

func findStart(lines []string) adv.Coord2D {
	for y, l := range lines {
		for x, c := range l {
			if c == 'S' {
				return adv.Coord2D{x, y}
			}
		}
	}
	return adv.Coord2D{}
}

func dijkstra(lines []string, start adv.Coord2D) map[adv.Coord2D]int {
	dist := map[adv.Coord2D]int{}
	prev := map[adv.Coord2D]adv.Coord2D{}
	var queue []adv.Coord2D
	queueMap := map[adv.Coord2D]bool{}
	for y, l := range lines {
		for x := range l {
			queue = append(queue, adv.Coord2D{x, y})
			queueMap[adv.Coord2D{x, y}] = true
		}
	}
	dist[start] = 0
	for len(queue) > 0 {
		var m int
		minDist := math.MaxInt
		for i := 0; i < len(queue); i++ {
			distQ, ok := dist[queue[i]]
			if ok && distQ < minDist {
				minDist = distQ
				m = i
			}
		}

		if minDist == math.MaxInt {
			break
		}

		delete(queueMap, queue[m])
		u := queue[m]
		queue = append(queue[:m], queue[m+1:]...)

		for _, d := range dirs {
			v := u.Add(d)
			if !queueMap[v] || lines[v[1]][v[0]] == '#' {
				continue
			}
			alt := dist[u] + 1
			distV, ok := dist[v]
			if !ok || alt < distV {
				dist[v] = alt
				prev[v] = u
			}
		}
	}

	return dist
}

var dirs = []adv.Coord2D{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}
