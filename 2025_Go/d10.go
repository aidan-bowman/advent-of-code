package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// utils

var rd *regexp.Regexp
var rb *regexp.Regexp

type Machine struct {
	ind int
	btn []int
	jlt []int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseLine(l string) Machine {
	ind := 0
	for i, r := range l[1:] {
		if r == ']' {
			break
		}

		if r == '#' {
			ind += 1 << i
		}
	}

	btns := rb.FindAllString(l, -1)
	btn := make([]int, len(btns))

	for i, s := range btns {
		nums := rd.FindAllString(s, -1)
		val := 0
		for _, v := range nums {
			tmp, err := strconv.Atoi(v)
			check(err)
			val += 1 << tmp
		}
		btn[i] = val
	}

	pos := strings.IndexRune(l, '{')
	jlts := rd.FindAllString(l[pos:], -1)
	jlt := make([]int, len(jlts))

	for i, n := range jlts {
		tmp, err := strconv.Atoi(n)
		check(err)
		jlt[i] = tmp
	}

	return Machine{ind, btn, jlt}
}

func parseData(dat string) []Machine {
	spl := strings.Split(dat, "\n")

	ms := make([]Machine, len(spl))

	for i, v := range spl {
		ms[i] = parseLine(v)
	}

	return ms
}

// part 1

func recInd(goal, cur, prs int, btns []int) int {
	if len(btns) == 0 {
		if goal == cur {
			return prs
		} else {
			return 99999
		}
	}

	np := recInd(goal, cur, prs, btns[1:])
	yp := recInd(goal, cur^btns[0], prs+1, btns[1:])

	if yp < np {
		return yp
	} else {
		return np
	}
}

func part1(ms []Machine) (sum int) {
	for _, v := range ms {
		sum += recInd(0, v.ind, 0, v.btn)
	}

	return sum
}

// part 2

// both parts

func main() {

	rd = regexp.MustCompile("\\d+")
	rb = regexp.MustCompile("\\([\\d,]+\\)")

	dat, err := os.ReadFile(os.Args[1])
	check(err)

	if dat[len(dat)-1] == '\n' {
		dat = dat[:len(dat)-1]
	}

	ms := parseData(string(dat))

	p1 := part1(ms)
	// p2 := part2(ms)

	fmt.Println("Part 1:", p1)
	// fmt.Println("Part 2:", p2)
}
