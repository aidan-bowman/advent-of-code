package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// utils

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseCables(dat string) map[string][]string {
	out := make(map[string][]string)
	r := regexp.MustCompile("(\\w){3}")

	split := strings.Split(dat, "\n")

	for _, v := range split {
		out[v[:3]] = r.FindAllString(v[3:], -1)
	}

	return out
}

// part 1

func part1(cbs map[string][]string) int {
	var rec func(string) int

	rec = func(d string) (sum int) {
		if d == "out" {
			return 1
		}

		for _, v := range cbs[d] {
			sum += rec(v)
		}

		return sum
	}

	return rec("you")
}

// part 2

type Signals struct {
	blank  int
	dac    int
	fft    int
	dacfft int
}

func getNext(sgn map[string]Signals) string {
	for k, v := range sgn {
		if k == "out" {
			continue
		}

		if v.blank != 0 || v.dac != 0 || v.fft != 0 || v.dacfft != 0 {
			return k
		}
	}

	return "done"
}

func handleDAC(sgn map[string]Signals) {
	old := sgn["dac"]

	sgn["dac"] = Signals{0, old.dac + old.blank, 0, old.dacfft + old.fft}
}

func handleFFT(sgn map[string]Signals) {
	old := sgn["fft"]

	sgn["fft"] = Signals{0, 0, old.fft + old.blank, old.dacfft + old.dac}
}

// xd recursive solution is naive even when checking for loops and only searching between each point
func part2(cbs map[string][]string) int {
	sgn := make(map[string]Signals)

	for k := range cbs {
		sgn[k] = Signals{}
	}

	sgn["svr"] = Signals{1, 0, 0, 0}

	for {
		nxt := getNext(sgn)

		if nxt == "done" {
			return sgn["out"].dacfft
		}

		if nxt == "dac" {
			handleDAC(sgn)
		}

		if nxt == "fft" {
			handleFFT(sgn)
		}

		new := sgn[nxt]

		for _, v := range cbs[nxt] {
			old := sgn[v]

			sgn[v] = Signals{old.blank + new.blank, old.dac + new.dac, old.fft + new.fft, old.dacfft + new.dacfft}
		}

		sgn[nxt] = Signals{}
	}
}

// both parts

func main() {
	dat, err := os.ReadFile(os.Args[1])
	check(err)

	if dat[len(dat)-1] == '\n' {
		dat = dat[:len(dat)-1]
	}

	cbs := parseCables(string(dat))

	p1 := part1(cbs)
	p2 := part2(cbs)

	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}
