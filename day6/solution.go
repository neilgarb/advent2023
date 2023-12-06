package main

import (
	_ "embed"
	"fmt"
	adv "github.com/neilgarb/advent2023"
	"strings"
)

//go:embed input.txt
var input []byte

const sample = `Time:      7  15   30
Distance:  9  40  200`

func main() {
	part1(adv.Lines(sample))
	part1(adv.Lines(input))
	part2(adv.Lines(sample))
	part2(adv.Lines(input))
}

func part1(lines []string) {
	times := strings.Fields(strings.TrimPrefix(lines[0], "Time:"))
	distances := strings.Fields(strings.TrimPrefix(lines[1], "Distance:"))
	tot := 1
	for i := 0; i < len(times); i++ {
		tot *= victories(adv.Atoi(times[i]), adv.Atoi(distances[i]))
	}
	fmt.Println(tot)
}

func part2(lines []string) {
	time := adv.Atoi(strings.Join(strings.Fields(strings.TrimPrefix(lines[0], "Time:")), ""))
	distance := adv.Atoi(strings.Join(strings.Fields(strings.TrimPrefix(lines[1], "Distance:")), ""))
	fmt.Println(victories(time, distance))
}

func victories(time, distance int) int {
	var count int
	for j := 0; j <= time; j++ {
		d := (time - j) * j
		if d > distance {
			count++
		}
	}
	return count
}
