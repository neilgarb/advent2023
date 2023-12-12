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

const sample = `???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`

func main() {
	part1(adv.Lines(sample))
	part1(adv.Lines(input))
	part2(adv.Lines(sample))
	part2(adv.Lines(input))
}

func part1(lines []string) {
	var tot int
	for _, l := range lines {
		x := permutations(l, 1)
		tot += x
	}
	fmt.Println(tot)
}

func part2(lines []string) {
	var tot int
	for _, l := range lines {
		tot += permutations(l, 5)
	}
	fmt.Println(tot)
}

func permutations(line string, repeat int) int {
	springs, records, _ := strings.Cut(line, " ")
	var springsR, recordsR []string
	for i := 0; i < repeat; i++ {
		springsR = append(springsR, springs)
		recordsR = append(recordsR, records)
	}
	springs = strings.Join(springsR, "?")
	records = strings.Join(recordsR, ",")
	var groups []int
	for _, r := range strings.Split(records, ",") {
		groups = append(groups, adv.Atoi(r))
	}
	return permute(springs, 0, groups, map[string]int{})
}

func permute(springs string, index int, groups []int, known map[string]int) int {
	key := springs + "__" + strconv.Itoa(index)
	for _, g := range groups {
		key += "__" + strconv.Itoa(g)
	}
	if i, ok := known[key]; ok {
		return i
	}
	if len(groups) == 0 {
		for i := index; i < len(springs); i++ {
			if springs[i] == '#' {
				return 0
			}
		}
		return 1
	}
	if index >= len(springs) {
		return 0
	}
	var tot int
	if placeable(groups[0], springs, index) {
		tot += permute(springs, index+groups[0]+1, groups[1:], known)
	}
	if springs[index] != '#' {
		tot += permute(springs, index+1, groups, known)
	}
	known[key] = tot
	return tot
}

func placeable(group int, springs string, index int) bool {
	if index > 0 && springs[index-1] == '#' {
		return false
	}
	if index+group < len(springs) && springs[index+group] == '#' {
		return false
	}
	for i := index; i < index+group; i++ {
		if i >= len(springs) {
			return false
		}
		if springs[i] == '.' {
			return false
		}
	}
	return true
}
