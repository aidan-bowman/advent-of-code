package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// utils

type Bound [2]int

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseLines(lines []string) (bnds []Bound, vals []int) {
	mode := false
	for _, v := range lines {
		if mode {
			tmp, err := strconv.Atoi(v)
			check(err)
			vals = append(vals, tmp)
		} else {
			if len(v) == 0 {
				mode = true
				continue
			}

			spl := strings.Split(v, "-")

			low, err := strconv.Atoi(spl[0])
			check(err)
			hgh, err := strconv.Atoi(spl[1])
			check(err)
			bnds = append(bnds, [2]int{low, hgh})
		}
	}

	return bnds, vals
}

// part 1

func IDInRange(bnds []Bound, id int) bool {
	for _, v := range bnds {
		if v[0] <= id && id <= v[1] {
			return true
		}
	}

	return false
}

func part1(bnds []Bound, vals []int) (sum int) {
	for _, v := range vals {
		if IDInRange(bnds, v) {
			sum++
		}
	}

	return sum
}

// part 2

func mergeBound(srtd []Bound, in Bound) []Bound {
	// first insert
	if len(srtd) == 0 {
		return []Bound{in}
	}

	last := len(srtd) - 1
	// sequential
	if srtd[last][1] == (in[0] - 1) {
		srtd[last][1] = in[1]
		return srtd
	}

	// append
	if srtd[last][1] < in[0] {
		return append(srtd, in)
	}

	// 0 is way below
	// 1 is just below
	// 2 is just above
	// 3 is way above
	mode := 0
	// may be 1 more, same size, or fewer, so we'll use appends and set a liberal capacity
	out := make([]Bound, 0, len(srtd)+1)

	// holy fuarrk
	for i, v := range srtd {
		switch mode {
		case 0:
			// in is way above e.g [3 6] <- [7 9]
			if v[1] < in[0]-1 {
				out = append(out, v)
				continue
			}

			// issues
			fallthrough
		case 1:
			if v[0] > in[1]+1 {
				// in fits snugly e.g. [12 15] <- [7 9]
				out = append(out, in)
				out = append(out, v)
				mode = 3
			} else if v[0] <= in[0] && in[1] <= v[1] {
				// in is fully contained e.g. [7 12] <- [7 9]
				// at this point we know that we can just return srtd with no modifications
				return srtd
			} else if in[0] <= v[0] && v[1] <= in[1] {
				// in fully contains our value e.g. [8 8] <- [7 9]
				// we can't just return sorted because the next value might be relevant as well
				// e.g. [7 9] <- [10 11]
				out = append(out, in)
				mode = 2
			} else if v[0] > in[0] {
				// in is just below/overlaps below e.g. [8 11] <- [7 9]
				out = append(out, [2]int{in[0], v[1]})
				mode = 3
			} else {
				// in is just above/overlaps above e.g. [3 8] <- [7 9]
				out = append(out, [2]int{v[0], in[1]})
				mode = 2
			}
		case 2:
			// only happens if the last value has a new upper bound we could conflict with
			if v[0] > out[i-1][1]+1 {
				// no conflict e.g. [7 9] <- [11 13]
				out = append(out, v)
				mode = 3
			} else if v[1] > out[i-1][1] {
				// conflict and not contained, update bound and finish filling in
				out[i-1][1] = v[1]
				mode = 3
				continue
			}
			// last cond failed, v is contained e.g. [7 9] <- [9 9]
		case 3:
			// worry free!
			out = append(out, v)
		}
	}

	return out
}

func part2(bnds []Bound) (sum int) {
	// current plan: make sorted list of combined, disjoint bounds
	// do this by inserting bounds one-by-one
	var srtd []Bound
	for _, v := range bnds {
		srtd = mergeBound(srtd, v)
	}

	for _, v := range srtd {
		sum += v[1] - v[0] + 1
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

	bnds, vals := parseLines(strings.Split(string(dat), "\n"))

	p1 := part1(bnds, vals)
	p2 := part2(bnds)

	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}
