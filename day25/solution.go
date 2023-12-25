package main

import (
	_ "embed"
	"fmt"
	adv "github.com/neilgarb/advent2023"
	"strings"
)

//go:embed input.txt
var input []byte

//go:embed sample.txt
var sample []byte

func main() {
	//part1(adv.Lines(sample))
	part1(adv.Lines(input))
}

func part1(lines []string) {
	conns := map[string][]string{}
	nodes := map[string]bool{}
	for _, l := range lines {
		from, to, _ := strings.Cut(l, ": ")
		conns[from] = append(conns[from], strings.Split(to, " ")...)
		nodes[from] = true
		for _, t := range strings.Split(to, " ") {
			nodes[t] = true
			conns[t] = append(conns[t], from)
		}
	}

	fmt.Println("graph G {")
	for from, tos := range conns {
		for _, to := range tos {
			fmt.Println(from, "--", to, ";")
		}
	}
	fmt.Println("}")

	// After running `go run . | dot -Tsvg -o input.svg`, inspect the graph visually.
	// Enter the links below.

	// fmt.Println(subGraphs(nodes, conns, map[string]bool{
	//	"jpn_vgf": true,
	//	"mnl_nmz": true,
	//	"fdb_txm": true,
	// }))
}

func subGraphs(nodes map[string]bool, conns map[string][]string, delMap map[string]bool) (res []int) {
	queue := map[string]bool{}
	for n := range nodes {
		queue[n] = true
	}

	for len(queue) > 0 {
		var start string
		for k := range queue {
			start = k
			break
		}
		graphQueue := []string{start}
		var s int
		for len(graphQueue) > 0 {
			c := graphQueue[0]
			graphQueue = graphQueue[1:]
			if !queue[c] {
				continue
			}
			delete(queue, c)
			s++
			for _, w := range conns[c] {
				if !delMap[delString(c, w)] {
					graphQueue = append(graphQueue, w)
				}
			}
		}
		res = append(res, s)
		if len(res) == 3 {
			return nil
		}
	}
	return res
}

func delString(a, b string) string {
	if b < a {
		return delString(b, a)
	}
	return a + "_" + b
}
