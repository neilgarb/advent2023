package advent2023

import (
	"bufio"
	"strconv"
	"strings"
)

func Lines[T string | []byte](s T) []string {
	scanner := bufio.NewScanner(strings.NewReader(strings.TrimSpace(string(s))))
	var res []string
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	return res

}

func Atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

type Coord2D [2]int

func (c Coord2D) Add(c2 Coord2D) Coord2D {
	return Coord2D{c[0] + c2[0], c[1] + c2[1]}
}

type Coord3D [3]int

func (c Coord3D) Add(c2 Coord3D) Coord3D {
	return Coord3D{c[0] + c2[0], c[1] + c2[1], c[2] + c2[2]}
}
