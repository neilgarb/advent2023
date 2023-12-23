package main

import (
	"container/heap"
	_ "embed"
	"fmt"
	adv "github.com/neilgarb/advent2023"
)

//go:embed input.txt
var input []byte

//go:embed sample.txt
var sample []byte

func main() {
	part1(adv.Lines(sample))
	part1(adv.Lines(input))
	part2(adv.Lines(sample))
	part2(adv.Lines(input))
}

func part1(lines []string) {
	start := adv.Coord2D{1, 0}
	end := adv.Coord2D{len(lines[0]) - 2, len(lines) - 1}

	q := queue{{start}}
	heap.Init(&q)

	var best int

	for q.Len() > 0 {
		p := q.Pop().(path)

		last := p[len(p)-1]
		if last == end {
			if len(p)-1 > best {
				best = len(p) - 1
			}
			continue
		}

		seen := map[adv.Coord2D]bool{}
		for _, c := range p {
			seen[c] = true
		}

		for _, d := range []adv.Coord2D{{0, 1}, {1, 0}, {-1, 0}, {0, -1}} {
			n := last.Add(d)
			if seen[n] || n[0] < 0 || n[1] < 0 || n[0] >= len(lines[0]) || n[1] >= len(lines) {
				continue
			}

			c := lines[n[1]][n[0]]
			if c == '#' {
				continue
			}

			newPath := append(path{}, p...)
			newPath = append(newPath, n)

			if c == '.' {
				q.Push(newPath)
				continue
			}

			if d[0] == 1 && c != '>' {
				continue
			} else if d[1] == 1 && c != 'v' {
				continue
			} else if d[0] == -1 && c != '^' {
				continue
			} else if d[1] == -1 && c != '<' {
				continue
			}

			q.Push(newPath)
		}
	}

	fmt.Println(best)
}

func part2(lines []string) {
	start := adv.Coord2D{1, 0}
	end := adv.Coord2D{len(lines[0]) - 2, len(lines) - 1}

	nodes := []adv.Coord2D{start, end}
	nodeMap := map[adv.Coord2D]bool{start: true, end: true}
	for y, l := range lines {
		for x := range l {
			c := adv.Coord2D{x, y}
			if lines[c[1]][c[0]] == '#' {
				continue
			}
			var cnt int
			for _, d := range dirs {
				n := c.Add(d)
				if n[0] >= 0 && n[1] >= 0 && n[0] < len(lines[0]) && n[1] < len(lines) && lines[n[1]][n[0]] != '#' {
					cnt++
				}
			}
			if cnt > 2 {
				nodes = append(nodes, c)
				nodeMap[c] = true
			}
		}
	}

	edges := map[adv.Coord2D]map[adv.Coord2D]int{}

	for _, c := range nodes {
		for _, d := range dirs {
			n := c.Add(d)
			if n[0] < 0 || n[1] < 0 || n[0] >= len(lines[0]) || n[1] >= len(lines) || lines[n[1]][n[0]] == '#' {
				continue
			}

			seen := map[adv.Coord2D]bool{c: true, n: true}
			dist := 1
			var deadEnd bool
			for !nodeMap[n] {
				var added bool
				for _, dd := range dirs {
					nn := n.Add(dd)
					if seen[nn] || nn[0] < 0 || nn[1] < 0 || nn[0] >= len(lines[0]) || nn[1] >= len(lines) || lines[nn[1]][nn[0]] == '#' {
						continue
					}
					n = nn
					seen[n] = true
					dist++
					added = true
					break
				}
				if !added {
					deadEnd = true
					break
				}
			}

			if !deadEnd {
				if edges[c] == nil {
					edges[c] = map[adv.Coord2D]int{}
				}
				edges[c][n] = dist
			}
		}
	}

	q := queue{{start}}
	heap.Init(&q)

	var best int

	for q.Len() > 0 {
		p := q.Pop().(path)

		var dist int
		for i := 1; i < len(p); i++ {
			dist += edges[p[i-1]][p[i]]
		}

		last := p[len(p)-1]
		if last == end {
			if dist > best {
				best = dist
			}
			continue
		}

		seen := make(map[adv.Coord2D]bool)
		for _, c := range p {
			seen[c] = true
		}

		for k := range edges[last] {
			if !seen[k] {
				newPath := append(path{}, p...)
				newPath = append(newPath, k)
				q.Push(newPath)
			}
		}
	}

	fmt.Println(best)
}

var dirs = []adv.Coord2D{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

type path []adv.Coord2D

type queue []path

func (q *queue) Len() int {
	return len(*q)
}

func (q *queue) Less(i, j int) bool {
	return len((*q)[i]) > len((*q)[j])
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
