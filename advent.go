package advent2023

import (
	"bufio"
	"bytes"
	"strconv"
	"strings"
)

func LinesFromString(s string) []string {
	scanner := bufio.NewScanner(strings.NewReader(s))
	var res []string
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	return res

}

func LinesFromBytes(b []byte) []string {
	scanner := bufio.NewScanner(bytes.NewReader(b))
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

type Coord2D struct {
	X, Y int
}

func (c Coord2D) Add(c2 Coord2D) Coord2D {
	return Coord2D{c.X + c2.X, c.Y + c2.Y}
}

type Coord3D struct {
	X, Y, Z int
}

func (c Coord3D) Add(c2 Coord3D) Coord3D {
	return Coord3D{c.X + c2.X, c.Y + c2.Y, c.Z + c2.Z}
}
