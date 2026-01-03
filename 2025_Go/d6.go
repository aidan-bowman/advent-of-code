package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

// utils

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// part 1

func parse1(lines []string) [][]int {
	out := make([][]int, len(lines))

	for i := range len(lines) {
		out[i] = []int{}

		conv := strings.Split(lines[i], " ")

		for _, v := range conv {
			if len(v) == 0 {
				continue
			}

			tmp, err := strconv.Atoi(v)
			check(err)

			out[i] = append(out[i], tmp)
		}
	}

	return out
}

func convOps(line string) (ops []rune) {
	conv := strings.Split(line, " ")

	for _, v := range conv {
		if len(v) == 0 {
			continue
		}

		r, _ := utf8.DecodeRuneInString(v)

		ops = append(ops, r)
	}

	return ops
}

func doOps(dat [][]int, ops []rune) (sum int) {
	for i, c := range ops {
		var op func(int, int) int
		var tmp int

		if c == '+' {
			tmp = 0
			op = func(a, b int) int {
				return a + b
			}
		} else {
			tmp = 1
			op = func(a, b int) int {
				return a * b
			}
		}

		for _, v := range dat {
			tmp = op(tmp, v[i])
		}

		sum += tmp
	}

	return sum
}

func part1(lines []string) int {
	dat := parse1(lines[:len(lines)-1])

	ops := convOps(lines[len(lines)-1])

	return doOps(dat, ops)
}

// part 2

func part2(lines []string) (sum int) {
	var op func(int, int) int
	var cur int

	for col := 0; ; col++ {
		if col == len(lines[0]) {
			return sum + cur
		}

		if (col != len(lines[0])-1) && (lines[len(lines)-1][col+1] != ' ') {
			sum += cur
			continue
		}

		if lines[len(lines)-1][col] != ' ' {
			if lines[len(lines)-1][col] == '+' {
				cur = 0
				op = func(a, b int) int {
					return a + b
				}
			} else {
				cur = 1
				op = func(a, b int) int {
					return a * b
				}
			}
		}

		reading := false
		val := 0

		for i := range len(lines) - 1 {
			if !reading && lines[i][col] != ' ' {
				reading = true
			}

			if reading {
				if lines[i][col] == ' ' {
					break
				}

				tmp, err := strconv.Atoi(string(lines[i][col]))
				check(err)

				val = (val * 10) + tmp
			}
		}

		cur = op(cur, val)
	}
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
