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

// part 1

func handleSplit(beams map[int]bool, line string) (cnt int, out map[int]bool) {
	out = map[int]bool{}

	for k := range beams {
		if line[k] == '^' {
			cnt++
			out[k-1] = true
			out[k+1] = true
		} else {
			out[k] = true
		}
	}

	return cnt, out
}

func part1(lines []string) (sum int) {
	cur := map[int]bool{}
	tmp := 0

	for i, v := range lines[0] {
		if v == 'S' {
			cur[i] = true
			break
		}
	}

	// splitters are only on every OTHER line
	for i := 2; i < len(lines); i += 2 {
		tmp, cur = handleSplit(cur, lines[i])
		sum += tmp
	}

	return sum
}

// part 2

func handleTimeSplit(beams map[int]int, line string) map[int]int {
	out := map[int]int{}

	for k, v := range beams {
		if line[k] == '^' {
			out[k-1] = out[k-1] + v
			out[k+1] = out[k+1] + v // could this even contain any values yet? whatever
		} else {
			out[k] = out[k] + v
		}
	}

	return out
}

func part2(lines []string) (sum int) {
	cur := map[int]int{}

	for i, v := range lines[0] {
		if v == 'S' {
			cur[i] = 1
			break
		}
	}

	for i := 2; i < len(lines); i += 2 {
		cur = handleTimeSplit(cur, lines[i])
	}

	for _, v := range cur {
		sum += v
	}

	return sum
}

// both parts

func main() {
	dat, err := os.ReadFile(os.Args[1])
	check(err)

	if dat[len(dat)-1] == '\n' {
		dat = dat[:len(dat)-1]
	}

	split := strings.Split(string(dat), "\n")

	p1 := part1(split)
	p2 := part2(split)

	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}
