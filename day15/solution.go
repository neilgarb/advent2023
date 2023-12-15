package main

import (
	_ "embed"
	"fmt"
	adv "github.com/neilgarb/advent2023"
	"strings"
)

//go:embed input.txt
var input []byte

const sample = `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`

func main() {
	part1(adv.Lines(sample))
	part1(adv.Lines(input))
	part2(adv.Lines(sample))
	part2(adv.Lines(input))
}

func part1(lines []string) {
	var tot int
	for _, p := range strings.Split(lines[0], ",") {
		tot += hash(p)
	}
	fmt.Println(tot)
}

func part2(lines []string) {
	type Lens struct {
		label string
		focal int
	}
	boxes := make([][]Lens, 256)
	for _, p := range strings.Split(lines[0], ",") {
		label, focal, _ := strings.Cut(p, "=")
		if strings.HasSuffix(label, "-") {
			label = strings.TrimSuffix(label, "-")
			index := hash(label)
			box := boxes[index]
			for i := 0; i < len(box); i++ {
				if box[i].label == label {
					box = append(box[:i], box[i+1:]...)
					break
				}
			}
			boxes[index] = box
		} else {
			box := hash(label)
			var found bool
			for i, l := range boxes[box] {
				if l.label == label {
					boxes[box][i].focal = adv.Atoi(focal)
					found = true
					break
				}
			}
			if !found {
				boxes[box] = append(boxes[box], Lens{label, adv.Atoi(focal)})
			}
		}
	}
	var tot int
	for i, box := range boxes {
		for j, lens := range box {
			tot += (i + 1) * (j + 1) * lens.focal
		}
	}
	fmt.Println(tot)
}

func hash(s string) int {
	var tot int
	for _, c := range s {
		tot += int(c)
		tot *= 17
		tot %= 256
	}
	return tot
}
