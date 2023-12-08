package main

import (
	_ "embed"
	"fmt"
	adv "github.com/neilgarb/advent2023"
	"strings"
)

//go:embed input.txt
var input []byte

const sample = `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`

const sample2 = `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`

func main() {
	part1(adv.Lines(sample))
	part1(adv.Lines(input))
	part2(adv.Lines(sample2))
	part2(adv.Lines(input))
}

func part1(lines []string) {
	instructions, nodes := parse(lines)
	var counter int
	for start := "AAA"; start != "ZZZ"; counter++ {
		start = nodes[start][instructions[counter%len(instructions)]]
	}
	fmt.Println(counter)
}

func part2(lines []string) {
	instructions, nodes := parse(lines)
	steps := 1
	for k := range nodes {
		if strings.HasSuffix(k, "A") {
			var counter int
			for ; !strings.HasSuffix(k, "Z"); counter++ {
				k = nodes[k][instructions[counter%len(instructions)]]
			}
			steps = calcLCM(steps, counter)
		}
	}
	fmt.Println(steps)
}

func parse(lines []string) (string, map[string]map[byte]string) {
	res := make(map[string]map[byte]string)
	for i := 2; i < len(lines); i++ {
		l := lines[i]
		from, to, _ := strings.Cut(l, " = ")
		left, right, _ := strings.Cut(strings.Trim(to, "()"), ", ")
		res[from] = map[byte]string{'L': left, 'R': right}
	}
	return lines[0], res
}

func calcLCM(a, b int) int {
	return a * b / calcGCD(a, b)
}

func calcGCD(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
