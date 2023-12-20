package main

import (
	_ "embed"
	"fmt"
	adv "github.com/neilgarb/advent2023"
	"strings"
)

//go:embed input.txt
var input []byte

const sample = `broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a`

func main() {
	part1(adv.Lines(sample))
	part1(adv.Lines(input))
	part2(adv.Lines(input))
}

type node struct {
	typ   byte
	dests []string
	on    bool
	last  map[string]bool
}

func part1(lines []string) {
	nodes := parse(lines)
	var lows, highs int
	for i := 1; i <= 1000; i++ {
		l, h, _ := push(nodes, i)
		lows += l
		highs += h
	}
	fmt.Println(lows * highs)
}

func part2(lines []string) {
	nodes := parse(lines)
	m := map[string]int{}
	var nums []int
	for i := 1; i < 10000; i++ {
		_, _, highToLX := push(nodes, i)
		if highToLX != "" && m[highToLX] == 0 {
			m[highToLX] = i
			nums = append(nums, i)
		}
	}
	fmt.Println(adv.LCM(nums...))
}

func parse(lines []string) map[string]node {
	nodes := map[string]node{}
	for _, l := range lines {
		name, dests, _ := strings.Cut(l, " -> ")
		var n node
		n.dests = strings.Split(dests, ", ")
		if name[0] == '&' || name[0] == '%' {
			n.typ = name[0]
			name = name[1:]
		}
		nodes[name] = n
	}

	for k, n := range nodes {
		if n.typ == '&' {
			last := map[string]bool{}
			for kk, nn := range nodes {
				for _, d := range nn.dests {
					if d == k {
						last[kk] = false
					}
				}
			}
			n.last = last
			nodes[k] = n
		}
	}
	return nodes
}

type pulse struct {
	from   string
	high   bool
	target string
}

func push(nodes map[string]node, idx int) (lows int, highs int, highToLX string) {
	pulses := []pulse{{"button", false, "broadcaster"}}
	for len(pulses) > 0 {
		p := pulses[0]
		if p.high && p.target == "lx" {
			highToLX = p.from
		}

		pulses = pulses[1:]
		if p.high {
			highs++
		} else {
			lows++
		}

		t := nodes[p.target]
		if t.typ == 0 {
			for _, d := range t.dests {
				pulses = append(pulses, pulse{p.target, p.high, d})
			}
		} else if t.typ == '%' {
			if p.high {
				continue
			}
			high := !t.on
			t.on = high
			nodes[p.target] = t
			for _, d := range t.dests {
				pulses = append(pulses, pulse{p.target, high, d})
			}
		} else if t.typ == '&' {
			t.last[p.from] = p.high
			allHigh := true
			for _, l := range t.last {
				if !l {
					allHigh = false
					break
				}
			}
			for _, d := range t.dests {
				pulses = append(pulses, pulse{p.target, !allHigh, d})
			}
		}
	}
	return lows, highs, highToLX
}
