package main

import (
	"container/heap"
	_ "embed"
	"fmt"
	adv "github.com/neilgarb/advent2023"
	"math"
)

//go:embed input.txt
var input []byte

const sample = `2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533`

func main() {
	part1(adv.Lines(sample))
	part1(adv.Lines(input))
	part2(adv.Lines(sample))
	part2(adv.Lines(input))
}

func part1(lines []string) {
	part(lines, 3, legal1)
}

func part2(lines []string) {
	part(lines, 10, legal2)
}

func part(lines []string, maxSteps int, legal func([]adv.Coord2D, adv.Coord2D) bool) {
	start := adv.Coord2D{0, 0}
	end := adv.Coord2D{len(lines[0]) - 1, len(lines) - 1}

	best := math.MaxInt

	var paths queue
	paths = append(paths, path{
		coords:   []adv.Coord2D{start},
		heatloss: 0,
		dist:     dist(start, end),
	})
	heap.Init(&paths)

	memo := map[memoKey]int{}

	for paths.Len() > 0 {
		p := heap.Pop(&paths).(path)

		if p.heatloss >= best {
			continue
		}

		last := p.coords[len(p.coords)-1]
		if last == end {
			if p.heatloss < best {
				best = p.heatloss
				continue
			}
		}

		seen := make(map[adv.Coord2D]bool)
		for _, c := range p.coords {
			seen[c] = true
		}

		for _, d := range dirs {
			n := last.Add(d)

			if last[1] == len(lines)-1 && d == (adv.Coord2D{-1, 0}) {
				continue
			}
			if last[0] == len(lines[0])-1 && d == (adv.Coord2D{0, -1}) {
				continue
			}

			if seen[n] || n[0] < 0 || n[1] < 0 || n[0] >= len(lines[0]) || n[1] >= len(lines) {
				continue
			}

			newHeatloss := p.heatloss + adv.Atoi(lines[n[1]][n[0]])
			if newHeatloss >= best {
				continue
			}

			if !legal(p.coords, d) {
				continue
			}

			newPath := append([]adv.Coord2D{}, p.coords...)
			newPath = append(newPath, n)

			memoLen := 10
			if len(newPath) >= memoLen+1 {
				var k memoKey
				for i := len(newPath) - 1; i >= max(0, len(newPath)-1-memoLen); i-- {
					k[len(newPath)-1-i] = newPath[i]
				}
				pb, ok := memo[k]
				if ok && newHeatloss >= pb {
					continue
				}
				memo[k] = newHeatloss
			}

			heap.Push(&paths, path{
				coords:   newPath,
				heatloss: newHeatloss,
				dist:     dist(n, end),
			})
		}
	}

	fmt.Println(best)
}

type memoKey [11]adv.Coord2D

type path struct {
	coords   []adv.Coord2D
	heatloss int
	dist     int
}

type queue []path

func (q *queue) Len() int {
	return len(*q)
}

func (q *queue) Less(i, j int) bool {
	if (*q)[i].dist == (*q)[j].dist {
		return (*q)[i].heatloss < (*q)[j].heatloss
	}
	return (*q)[i].dist < (*q)[j].dist
}

func (q *queue) Swap(i, j int) {
	(*q)[i], (*q)[j] = (*q)[j], (*q)[i]
}

func (q *queue) Push(x any) {
	*q = append(*q, x.(path))
}

func (q *queue) Pop() any {
	p := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]
	return p
}

func legal1(p []adv.Coord2D, d adv.Coord2D) bool {
	if len(p) < 4 {
		return true
	}
	for i := len(p) - 1; i >= len(p)-3; i-- {
		if d != p[i].Sub(p[i-1]) {
			return true
		}
	}
	return false
}

func legal2(p []adv.Coord2D, d adv.Coord2D) bool {
	if len(p) <= 1 {
		return true
	}

	cur := p[len(p)-1].Sub(p[len(p)-2])
	count := 1
	for i := len(p) - 2; i > 0; i-- {
		if p[i].Sub(p[i-1]) == cur {
			count++
		} else {
			break
		}
	}

	if count >= 10 && d == cur {
		return false
	}

	if count < 4 && d != cur {
		return false
	}

	return true
}

func dist(a, b adv.Coord2D) int {
	return max(a[1]-a[0], a[0]-a[1]) + max(b[1]-b[0], b[0]-b[1])
}

var dirs = []adv.Coord2D{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}
