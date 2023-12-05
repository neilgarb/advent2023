package main

import (
	_ "embed"
	"fmt"
	adv "github.com/neilgarb/advent2023"
	"math"
	"strings"
)

//go:embed input.txt
var input []byte

const sample = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

func main() {
	part1(adv.Lines(sample))
	part1(adv.Lines(input))
	part2(adv.Lines(sample))
	part2(adv.Lines(input))
}

func part1(lines []string) {
	d := parse(lines)
	lowest := math.MaxInt
	for _, seed := range d.seeds {
		seed = d.mapseed(seed)
		if seed < lowest {
			lowest = seed
		}
	}
	fmt.Println(lowest)
}

func part2(lines []string) {
	d := parse(lines)
	lowest := math.MaxInt
	for i := 0; i < len(d.seeds); i += 2 {
		for j := d.seeds[i]; j < d.seeds[i]+d.seeds[i+1]; j++ {
			seed := d.mapseed(j)
			if seed < lowest {
				lowest = seed
			}
		}
	}
	fmt.Println(lowest)
}

type data struct {
	seeds  []int
	ranges [][]range_
}

type range_ struct {
	from   int
	to     int
	length int
}

func parse(lines []string) (res data) {
	lines = append(lines, "")
	var cur []range_
	for _, l := range lines {
		if strings.HasPrefix(l, "seeds: ") {
			for _, f := range strings.Fields(strings.TrimPrefix(l, "seeds: ")) {
				res.seeds = append(res.seeds, adv.Atoi(f))
			}
			continue
		}

		if l == "" {
			if len(cur) > 0 {
				res.ranges = append(res.ranges, cur)
				cur = nil
			}
			continue
		}

		if strings.HasSuffix(l, "map:") {
			continue
		}

		fields := strings.Fields(l)
		cur = append(cur, range_{
			from:   adv.Atoi(fields[1]),
			to:     adv.Atoi(fields[0]),
			length: adv.Atoi(fields[2]),
		})
	}
	return res
}

func (d data) mapseed(seed int) int {
MapLoop:
	for _, mappings := range d.ranges {
		for _, r := range mappings {
			if seed >= r.from && seed < r.from+r.length {
				seed = r.to + (seed - r.from)
				continue MapLoop
			}
		}
	}
	return seed
}
