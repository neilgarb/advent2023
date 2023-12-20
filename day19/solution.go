package main

import (
	_ "embed"
	"fmt"
	adv "github.com/neilgarb/advent2023"
	"strings"
)

//go:embed input.txt
var input []byte

const sample = `px{a<2006:qkq,m>2090:A,rfg}
pv{a>1716:R,A}
lnx{m>1548:A,A}
rfg{s<537:gd,x>2440:R,A}
qs{s>3448:A,lnx}
qkq{x<1416:A,crn}
crn{x>2662:A,R}
in{s<1351:px,qqz}
qqz{s>2770:qs,m<1801:hdj,R}
gd{a>3333:R,R}
hdj{m>838:A,pv}

{x=787,m=2655,a=1222,s=2876}
{x=1679,m=44,a=2067,s=496}
{x=2036,m=264,a=79,s=2244}
{x=2461,m=1339,a=466,s=291}
{x=2127,m=1623,a=2188,s=1013}`

func main() {
	part1(adv.Lines(sample))
	part1(adv.Lines(input))
	part2(adv.Lines(sample))
	part2(adv.Lines(input))
}

type part [4]int

func part1(lines []string) {
	rules, parts := parse(lines)

	var tot int
	for _, p := range parts {
		rule := "in"
	RuleLoop:
		for rule != "R" && rule != "A" {
			for _, r := range rules[rule] {
				if r == "R" || r == "A" {
					rule = r
					break RuleLoop
				}

				if !strings.Contains(r, ":") {
					rule = r
					continue RuleLoop
				}

				cmp, dest, _ := strings.Cut(r, ":")

				f := cmp[0]
				c := cmp[1]
				num := adv.Atoi(cmp[2:])
				idx := strings.Index(xmas, string(f))

				if c == '>' {
					if p[idx] > num {
						rule = dest
						continue RuleLoop
					}
				} else if c == '<' {
					if p[idx] < num {
						rule = dest
						continue RuleLoop
					}
				}
			}
		}
		if rule == "A" {
			tot += p[0] + p[1] + p[2] + p[3]
		}
	}

	fmt.Println(tot)
}

type partRange struct {
	ranges [4][2]int
	rule   string
}

func part2(lines []string) {
	rules, _ := parse(lines)
	parts := []partRange{
		{[4][2]int{{1, 4000}, {1, 4000}, {1, 4000}, {1, 4000}}, "in"},
	}

	var accepted []partRange

	for len(parts) > 0 {
		p := parts[0]
		parts = parts[1:]

		if p.rule == "A" {
			accepted = append(accepted, p)
			continue
		}

		for _, r := range rules[p.rule] {
			if strings.Contains(r, ":") {
				cmp, dest, _ := strings.Cut(r, ":")

				f := cmp[0]
				c := cmp[1]
				num := adv.Atoi(cmp[2:])
				idx := strings.Index(xmas, string(f))

				if c == '>' {
					if p.ranges[idx][0] > num {
						p.rule = dest
					} else {
						newPart := p
						newPart.ranges[idx][0] = num + 1
						newPart.rule = dest
						parts = append(parts, newPart)

						p.ranges[idx][1] = num
					}
				} else if c == '<' {
					if p.ranges[idx][1] < num {
						p.rule = dest
					} else {
						newPart := p
						newPart.ranges[idx][1] = num - 1
						newPart.rule = dest
						parts = append(parts, newPart)

						p.ranges[idx][0] = num
					}
				}
			} else {
				newPart := p
				newPart.rule = r
				parts = append(parts, newPart)
			}
		}
	}

	var tot int
	for _, r := range accepted {
		tot += (r.ranges[0][1] - r.ranges[0][0] + 1) * (r.ranges[1][1] - r.ranges[1][0] + 1) * (r.ranges[2][1] - r.ranges[2][0] + 1) * (r.ranges[3][1] - r.ranges[3][0] + 1)
	}
	fmt.Println(tot)
}

func parse(lines []string) (rules map[string][]string, parts []part) {
	rules = map[string][]string{}
	for _, l := range lines {
		if strings.HasPrefix(l, "{") {
			var p part
			fields := strings.Split(l, ",")
			for i, f := range fields {
				p[i] = adv.Atoi(strings.Trim(f, "{}xmas="))
			}
			parts = append(parts, p)
		} else if l != "" {
			fields := strings.Split(l, "{")
			rules[fields[0]] = strings.Split(strings.Trim(fields[1], "}"), ",")
		}
	}
	return rules, parts
}

const xmas = "xmas"
