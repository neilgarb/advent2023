package main

import (
	_ "embed"
	"fmt"
	adv "github.com/neilgarb/advent2023"
	"sort"
	"strings"
)

//go:embed input.txt
var input []byte

const sample = `1,0,1~1,2,1
0,0,2~2,0,2
0,2,3~2,2,3
0,0,4~0,2,4
2,0,5~2,2,5
0,1,6~2,1,6
1,1,8~1,1,9`

func main() {
	part1(adv.Lines(sample))
	part1(adv.Lines(input))
	part2(adv.Lines(sample))
	part2(adv.Lines(input))
}

type brick struct {
	id     int
	coords [2]adv.Coord3D
}

func part1(lines []string) {
	var bricks []brick
	for i, l := range lines {
		b := parseBrick(l)
		bricks = append(bricks, brick{id: i, coords: b})
	}
	move(bricks, false, nil)
	var tot int
	for _, b := range bricks {
		bricks = append([]brick{}, bricks...)
		if move(bricks, true, map[brick]bool{b: true}) == 0 {
			tot++
		}
	}
	fmt.Println(tot)
}

func part2(lines []string) {
	var bricks []brick
	for i, l := range lines {
		b := parseBrick(l)
		bricks = append(bricks, brick{id: i, coords: b})
	}
	move(bricks, false, nil)
	var tot int
	for _, b := range bricks {
		cpy := append([]brick{}, bricks...)
		tot += move(cpy, false, map[brick]bool{b: true})
	}
	fmt.Println(tot)
}

func parseBrick(l string) (b [2]adv.Coord3D) {
	from, to, _ := strings.Cut(l, "~")
	b[0] = parseCoord(from)
	b[1] = parseCoord(to)
	return b
}

func parseCoord(s string) adv.Coord3D {
	var c adv.Coord3D
	parts := strings.Split(s, ",")
	for i := 0; i < 3; i++ {
		c[i] = adv.Atoi(parts[i])
	}
	return c
}

func collide(a, b [2]adv.Coord3D) bool {
	for x := a[0][0]; x <= a[1][0]; x++ {
		for y := a[0][1]; y <= a[1][1]; y++ {
			for z := a[0][2]; z <= a[1][2]; z++ {
				if (x == b[0][0] && y == b[0][1] && z >= b[0][2] && z <= b[1][2]) ||
					(x == b[0][0] && y >= b[0][1] && y <= b[1][1] && z == b[0][2]) ||
					(x >= b[0][0] && x <= b[1][0] && y == b[0][1] && z == b[0][2]) {
					return true
				}
			}
		}
	}
	return false
}

func move(bricks []brick, exitEarly bool, excl map[brick]bool) int {
	moveMap := map[int]bool{}
	for {
		var m int
		for _, b := range bricks {
			m = max(m, b.coords[0][2])
			m = max(m, b.coords[1][2])
		}
		moved := false
		sort.Slice(bricks, func(i, j int) bool {
			return bricks[i].coords[0][2] < bricks[j].coords[0][2]
		})
		for i, b := range bricks {
			if b.coords[0][2] <= 1 || b.coords[1][2] <= 1 {
				continue // Already on the ground.
			}
			check := b
			check.coords[0][2]--
			check.coords[1][2]--
			var collided bool
			for _, col := range bricks {
				if excl[col] {
					continue
				}
				if b == col {
					continue
				}
				if collide(check.coords, col.coords) || collide(col.coords, check.coords) {
					collided = true
					break
				}
			}
			if collided {
				continue
			}
			if exitEarly {
				return 1
			}
			bricks[i] = check
			moved = true
			moveMap[check.id] = true
		}
		if !moved {
			break
		}
	}
	return len(moveMap)
}
