package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func rotate(str string, pos int) (npos int, zeros int) {
	dist, err := strconv.Atoi(str[1:])
	check(err)

	// fuarrrk everything i'm brooting this

	npos = pos
	zeros = 0

	if str[0] == 'R' {
		for dist != 0 {
			npos += 1
			dist -= 1

			if npos == 100 {
				npos = 0
				zeros += 1
			}
		}
	} else {
		for dist != 0 {
			npos -= 1
			dist -= 1

			if npos == 0 {
				zeros += 1
			}

			if npos == -1 {
				npos = 99
			}
		}
	}

	return npos, zeros
}

func part1(dat []string) int {
	var iter func(dat []string, pos, sol int) int

	iter = func(dat []string, pos, sol int) int {
		if len(dat) == 0 || len(dat[0]) == 0 {
			return sol
		} else {
			dist, err := strconv.Atoi(dat[0][1:])
			check(err)

			if dat[0][0] == 'R' {
				pos = (pos + dist) % 100
			} else {
				pos = (pos - dist) % 100
			}

			if pos == 0 {
				return iter(dat[1:], pos, sol+1)
			} else {
				return iter(dat[1:], pos, sol)
			}
		}
	}

	return iter(dat, 50, 0)
}

func part2(dat []string) int {
	var iter func(dat []string, pos, sol int) int

	iter = func(dat []string, pos, sol int) int {
		if len(dat) == 0 || len(dat[0]) == 0 {
			return sol
		} else {
			pos, zeros := rotate(dat[0], pos)

			return iter(dat[1:], pos, sol+zeros)
		}
	}

	return iter(dat, 50, 0)
}

func main() {
	path := os.Args[1]
	dat, err := os.ReadFile(path)
	check(err)

	split := strings.Split(string(dat), "\n")

	s1 := part1(split)
	s2 := part2(split)

	fmt.Println("Part 1:", s1)
	fmt.Println("Part 2:", s2)
}
