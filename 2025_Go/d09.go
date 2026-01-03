package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// utils

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Pos [2]int

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func parsePos(s string) Pos {
	split := strings.Split(s, ",")

	x, err := strconv.Atoi(split[0])
	check(err)
	y, err := strconv.Atoi(split[1])
	check(err)

	return [2]int{x, y}
}

// part 1

func redTiles(lines []string) []Pos {
	out := make([]Pos, len(lines))

	for i, v := range lines {
		out[i] = parsePos(v)
	}

	return out
}

func part1(lines []string) (best int) {
	pl := redTiles(lines)

	// small enough we can just brute-force check every permutation
	for i, a := range pl {
		for _, b := range pl[i+1:] {
			dx := abs(a[0] - b[0])
			dy := abs(a[1] - b[1])
			area := (dx + 1) * (dy + 1)

			if area > best {
				best = area
			}
		}
	}

	return best
}

// part 2
// input is too fuggin BIG NOT XD NOT XD
// i'm spent and i have my flight fuarrrk this

func tilesBetween(grnm map[Pos]bool, a, b Pos) {
	if a[0] == b[0] {
		if a[1] < b[1] {
			for i := a[1]; i <= b[1]; i++ {
				grnm[[2]int{a[0], i}] = true
			}
		} else {
			for i := b[1]; i <= a[1]; i++ {
				grnm[[2]int{a[0], i}] = true
			}
		}
	} else {
		if a[0] < b[0] {
			for i := a[0]; i <= b[0]; i++ {
				grnm[[2]int{i, a[1]}] = true
			}
		} else {
			for i := b[0]; i <= a[0]; i++ {
				grnm[[2]int{i, a[1]}] = true
			}
		}
	}
}

// to find the center, lets shoot a beam diagonally from 0, 0. this coould fail but w/e
func findInner(grnmap map[Pos]bool) Pos {
	var cur Pos
	for grnmap[cur] == false {
		cur[0]++
		cur[1]++
	}
	return [2]int{cur[0] + 1, cur[1] + 1}
}

// will fail for very thin corridors but w/e
func fillInner(grnmap map[Pos]bool) {
	inner := findInner(grnmap)

	active := []Pos{inner}

	for len(active) != 0 {
		if grnmap[active[0]] {
			active = active[1:]
			continue
		}

		if active[0][0] == 0 || active[0][1] == 0 {
			// screw your idiomatic error handling
			panic(errors.New("fug"))
		}

		grnmap[active[0]] = true

		active = append(active[1:],
			[2]int{active[0][0] + 1, active[0][1]},
			[2]int{active[0][0] - 1, active[0][1]},
			[2]int{active[0][0], active[0][1] + 1},
			[2]int{active[0][0], active[0][1] - 1},
		)
	}
}

func debugPrintMap(grn map[Pos]bool) {
	for y := 0; y < 14; y++ {
		for x := 0; x < 14; x++ {
			if grn[[2]int{x, y}] {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

func allTiles(lines []string) ([]Pos, map[Pos]bool) {
	red := make([]Pos, len(lines))
	grn := make(map[Pos]bool)

	for i, v := range lines {
		red[i] = parsePos(v)

		if i == 0 {
			continue
		}

		tilesBetween(grn, red[i], red[i-1])

		if i == len(lines)-1 {
			tilesBetween(grn, red[i], red[0])
		}
	}

	fillInner(grn)

	return red, grn
}

func checkBounds(grnm map[Pos]bool, a, b Pos) bool {
	if a[0] < b[0] {
		for i := a[0]; i < b[0]; i++ {
			if !grnm[[2]int{i, a[1]}] || !grnm[[2]int{i, b[1]}] {
				return false
			}
		}
	} else {
		for i := b[0]; i < a[0]; i++ {
			if !grnm[[2]int{i, a[1]}] || !grnm[[2]int{i, b[1]}] {
				return false
			}
		}
	}

	if a[1] < b[1] {
		for i := a[1]; i < b[1]; i++ {
			if !grnm[[2]int{a[0], i}] || !grnm[[2]int{b[0], i}] {
				return false
			}
		}
	} else {
		for i := b[1]; i < a[1]; i++ {
			if !grnm[[2]int{a[0], i}] || !grnm[[2]int{b[0], i}] {
				return false
			}
		}
	}

	return true
}

func part2(lines []string) (best int) {
	red, grn := allTiles(lines)

	for i, a := range red {
		for _, b := range red[i+1:] {
			if !checkBounds(grn, a, b) {
				continue
			}

			dx := abs(a[0] - b[0])
			dy := abs(a[1] - b[1])
			area := (dx + 1) * (dy + 1)

			if area > best {
				best = area
			}
		}
	}

	return best
}

// both parts

func main() {
	dat, err := os.ReadFile(os.Args[1])
	check(err)

	if dat[len(dat)-1] == '\n' {
		dat = dat[:len(dat)-1]
	}

	lines := strings.Split(string(dat), "\n")

	p1 := part1(lines)
	p2 := part2(lines)

	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}
