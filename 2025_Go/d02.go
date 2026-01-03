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

// part 1

func invalid1(id int) bool {
	str := strconv.Itoa(id)

	if len(str)%2 == 1 {
		return false
	}

	half := len(str) / 2

	for i := 0; i < half; i++ {
		if str[i] != str[i+half] {
			return false
		}
	}

	return true
}

// part 2

/*
func chunk(s string, size int) []string {
	chunks := make([]string, 0, len(s)/size)

	for i := 0; i < len(s); i += size {
		chunks = append(chunks, s[i:(i+size-1)])
	}

	return chunks
    }
*/

func invalid2(id int) bool {
	str := strconv.Itoa(id)

outer:
	for size := 1; size < len(str); size++ {
		if len(str)%size != 0 {
			continue
		}

		for i := range size {
			char := str[i]

			for j := size + i; j < len(str); j += size {
				if str[j] != char {
					continue outer
				}
			}
		}

		return true
	}

	return false
}

// both parts

func countIDsInRange(str string) (int, int) {
	split := strings.Split(string(str), "-")

	start, err := strconv.Atoi(split[0])
	check(err)

	end, err := strconv.Atoi(split[1])
	check(err)

	p1, p2 := 0, 0

	for i := start; i <= end; i++ {
		if invalid1(i) {
			p1 += i
		}

		if invalid2(i) {
			p2 += i
		}
	}

	return p1, p2
}

func sols(dat []string) (int, int) {
	p1, p2 := 0, 0

	for _, v := range dat {
		p1a, p2a := countIDsInRange(v)
		p1 += p1a
		p2 += p2a
	}

	return p1, p2
}

func main() {
	path := os.Args[1]
	dat, err := os.ReadFile(path)
	check(err)

	if dat[len(dat)-1] == '\n' {
		dat = dat[:(len(dat) - 1)]
	}

	split := strings.Split(string(dat), ",")

	s1, s2 := sols(split)

	fmt.Println("Part 1:", s1)
	fmt.Println("Part 2:", s2)
}
