package main

import (
	"fmt"
	"os"
	"strings"
)

// utils

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func makeDir(pos, size int) []int {
	dir := []int{0}

	if pos != 0 {
		dir = append(dir, -1)
	}

	if pos < (size - 1) {
		dir = append(dir, 1)
	}

	return dir
}

func checkRoll(lines [][]rune, x, y int) bool {
	xdir := makeDir(x, len(lines[0]))
	ydir := makeDir(y, len(lines))
	cnt := 0

	for _, dx := range xdir {
		for _, dy := range ydir {
			if lines[y+dy][x+dx] == '@' || lines[y+dy][x+dx] == 'x' {
				cnt++
			}
		}
	}

	// the roll at x, y is counted, so we raise the limit by one
	return (cnt < 5)
}

// part 1

func sumRolls(lines [][]rune) (sum int) {
	for y, s := range lines {
		for x, v := range s {
			if (v == '@') && checkRoll(lines, x, y) {
				sum++
			}
		}
	}

	return sum
}

// part 2
func remRolls(lines [][]rune) (sum int) {
	type Pos [2]int
	var poss []Pos

	for y, s := range lines {
		for x, v := range s {
			if (v == '@') && checkRoll(lines, x, y) {
				lines[y][x] = 'x'
				sum++
				poss = append(poss, [2]int{x, y})
			}
		}
	}

	for _, v := range poss {
		lines[v[1]][v[0]] = '.'
	}

	if sum == 0 {
		return 0
	} else {
		return sum + remRolls(lines)
	}
}

// both parts

func main() {
	dat, err := os.ReadFile(os.Args[1])
	check(err)

	if dat[len(dat)-1] == '\n' {
		dat = dat[:(len(dat) - 1)]
	}

	lines := strings.Split(string(dat), "\n")
	layout := make([][]rune, len(lines))

	// string to slice of runes
	for i, v := range lines {
		layout[i] = make([]rune, len(v))

		for j, r := range v {
			layout[i][j] = r
		}
	}

	fmt.Println("Part 1:", sumRolls(layout))
	fmt.Println("Part 2:", remRolls(layout))
}
