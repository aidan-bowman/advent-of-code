package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// utils

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

type Pos [3]int

func (p1 Pos) GetDist(p2 Pos) float64 {
	dx := float64(abs(p1[0] - p2[0]))
	dy := float64(abs(p1[1] - p2[1]))
	dz := float64(abs(p1[2] - p2[2]))

	return math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2) + math.Pow(dz, 2))
}

func (p1 Pos) Equal(p2 Pos) bool {
	if p1[0] == p2[0] && p1[1] == p2[1] && p1[2] == p2[2] {
		return true
	}
	return false
}

type Connection struct {
	Dist float64
	P1   Pos
	P2   Pos
}

func makeConnection(p1, p2 Pos) Connection {
	return Connection{
		p1.GetDist(p2),
		p1,
		p2,
	}
}

func genConns(pl []Pos) (cl []Connection) {
	// skipping pregen because lazy
	// cl := make([]Connection, len(pos)*(len(pos)-1)/2)

	for i, a := range pl {
		for _, b := range pl[i+1:] {
			cl = append(cl, makeConnection(a, b))
		}
	}

	sort.Slice(cl, func(i, j int) bool {
		return cl[i].Dist < cl[j].Dist
	})

	return cl
}

func popConn(cl []Connection, p Pos) (outcl []Connection, outp []Pos) {
	for _, v := range cl {
		if v.P1 == p {
			outp = append(outp, v.P2)
		} else if v.P2 == p {
			outp = append(outp, v.P1)
		} else {
			outcl = append(outcl, v)
		}
	}

	return outcl, outp
}

func insSet(set []Pos, p Pos) []Pos {
	for _, v := range set {
		if v == p {
			return set
		}
	}

	return append(set, p)
}

// part 1

// if only this was in lisp ecks dee
func findCircuits(cl []Connection) [][]Pos {
	// too stupid to figure out how to impoooort modules so lets just implement sets and queues ourselves
	circuit := []Pos{cl[0].P1} // set
	active := []Pos{cl[0].P1}  // queue

	var out []Pos

	// while we have active nodes, get connected nodes and reduce connection list (for performance)
	for len(active) != 0 {
		cl, out = popConn(cl, active[0])
		active = active[1:]

		for _, p := range out {
			circuit = insSet(circuit, p)
			active = append(active, out...)
		}
	}

	if len(cl) == 0 {
		return [][]Pos{circuit}
	} else {
		rec := findCircuits(cl)
		return append(rec, circuit)
	}
}

func part1(pl []Pos, conns int) int {
	cl := genConns(pl)

	circuits := findCircuits(cl[:conns])

	sort.Slice(circuits, func(i, j int) bool {
		return len(circuits[i]) > len(circuits[j])
	})

	return len(circuits[0]) * len(circuits[1]) * len(circuits[2])
}

// part 2
// i should have just done this from the start... things have taken a tuna for the worse...

func checkSingle(circuits map[Pos]int) bool {
	var n int

	for _, v := range circuits {
		if n == 0 {
			n = v
		} else {
			if n != v {
				return false
			}
		}
	}

	return true
}

// in-place
func combCircuit(circuits map[Pos]int, p1, p2 Pos) {
	a := circuits[p1]
	b := circuits[p2]

	for k, v := range circuits {
		if v == a {
			circuits[k] = b
		}
	}
}

func part2(pl []Pos) int {
	cl := genConns(pl)

	circuits := make(map[Pos]int)

	for i, p := range pl {
		circuits[p] = i
	}

	for _, c := range cl {
		combCircuit(circuits, c.P1, c.P2)

		if checkSingle(circuits) {
			return c.P1[0] * c.P2[0]
		}
	}

	return -1
}

// both parts

func parseInput(dat string) (out []Pos) {
	lines := strings.Split(dat, "\n")

	for _, v := range lines {
		split := strings.Split(v, ",")

		x, err := strconv.Atoi(split[0])
		check(err)
		y, err := strconv.Atoi(split[1])
		check(err)
		z, err := strconv.Atoi(split[2])
		check(err)

		out = append(out, [3]int{x, y, z})
	}

	return out
}

func main() {
	dat, err := os.ReadFile(os.Args[1])
	check(err)

	if dat[len(dat)-1] == '\n' {
		dat = dat[:len(dat)-1]
	}

	conns, err := strconv.Atoi(os.Args[2])
	check(err)

	pl := parseInput(string(dat))

	p1 := part1(pl, conns)
	p2 := part2(pl)

	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}
