package main

import (
	_ "embed"
	"fmt"
	adv "github.com/neilgarb/advent2023"
	"strconv"
	"strings"
)

//go:embed input.txt
var input []byte

//go:embed sample.txt
var sample []byte

func main() {
	part1(adv.Lines(sample), 7, 27)
	part1(adv.Lines(input), 200000000000000, 400000000000000)
	part2(adv.Lines(input))
}

func part1(lines []string, areaMin, areaMax float64) {
	hails := parse(lines)
	var tot int
	for i := 0; i < len(hails); i++ {
		for j := i + 1; j < len(hails); j++ {
			c, _, ok := futureCollide(hails[i], hails[j], true)
			if !ok {
				continue
			}
			if c[0] >= areaMin && c[0] <= areaMax && c[1] >= areaMin && c[1] <= areaMax {
				tot++
			}
		}
	}
	fmt.Println(tot)
}

func part2(lines []string) {
	hails := parse(lines)

	for _, h := range hails {
		// vx * t + sx = h[3] * t + h[0]
		//  vx * t - h[3] * t = h[0] - sx
		//  t = (h[0] - sx) / (vx - h[3])
		// vy * t + sy = h[4] * t + h[1]
		//  vy * t - h[4] * t = h[1] - sy
		//  t = (h[1] - sy) / (vy - h[4])
		// vz * t + sz = h[5] * t + h[2]
		//  vz * t - h[5] * t = h[2] - sz
		//  t = (h[2] - sz) / (vz - h[5])

		//  (h[0] - sx) * (vy - h[4]) - (h[1] - sy) * (vx - h[3]) = 0
		//  (h[1] - sy) * (vz - h[5]) - (h[2] - sz) * (vy - h[4]) = 0

		fmt.Printf("((%d) - x) * (b - (%d)) - ((%d) - y) * (a - (%d)), ", int(h[0]), int(h[4]), int(h[1]), int(h[3]))
		fmt.Printf("((%d) - y) * (c - (%d)) - ((%d) - z) * (b - (%d)), ", int(h[1]), int(h[5]), int(h[2]), int(h[4]))

		// Plugin into sympy.
	}
}

func parse(lines []string) (res []hail) {
	for _, l := range lines {
		l = strings.Replace(l, "@", ",", 1)
		l = strings.ReplaceAll(l, " ", "")
		var h hail
		for i, n := range strings.Split(l, ",") {
			h[i], _ = strconv.ParseFloat(n, 64)
		}
		res = append(res, h)
	}
	return res
}

type hail [6]float64

func futureCollide(h1, h2 [6]float64, ignoreZ bool) ([3]float64, float64, bool) {

	// x: h1[3] * t + h1[0] = h2[3] * s + h2[0]
	// y: h1[4] * t + h1[1] = h2[4] * s + h2[1]
	// z: h1[5] * t + h1[2] = h2[5] * s + h2[2]

	t := (-h1[1]*h2[3] + h2[4]*h1[0] - h2[4]*h2[0] + h2[1]*h2[3]) / (h1[4]*h2[3] - h2[4]*h1[3])

	if h1[4]*h2[3] == h2[4]*h1[3] {
		return [3]float64{}, 0, false
	}

	x := h1[0] + t*h1[3]
	y := h1[1] + t*h1[4]
	z := h1[2] + t*h1[5]

	if x < h1[0] && h1[3] > 0 {
		return [3]float64{}, 0, false
	}
	if x > h1[0] && h1[3] < 0 {
		return [3]float64{}, 0, false
	}
	if x < h2[0] && h2[3] > 0 {
		return [3]float64{}, 0, false
	}
	if x > h2[0] && h2[3] < 0 {
		return [3]float64{}, 0, false
	}
	if y < h1[1] && h1[4] > 0 {
		return [3]float64{}, 0, false
	}
	if y > h1[1] && h1[4] < 0 {
		return [3]float64{}, 0, false
	}
	if y < h2[1] && h2[4] > 0 {
		return [3]float64{}, 0, false
	}
	if y > h2[1] && h2[4] < 0 {
		return [3]float64{}, 0, false
	}
	if !ignoreZ {
		if z < h1[2] && h1[5] > 0 {
			return [3]float64{}, 0, false
		}
		if z > h1[2] && h1[5] < 0 {
			return [3]float64{}, 0, false
		}
		if z < h2[2] && h2[5] > 0 {
			return [3]float64{}, 0, false
		}
		if z > h2[2] && h2[5] < 0 {
			return [3]float64{}, 0, false
		}
	}

	return [3]float64{x, y, z}, t, true
}
