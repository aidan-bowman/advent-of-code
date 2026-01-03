package main

import (
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

func maxInt(ns []int) (value, index int) {
	for i, v := range ns {
		if v > value {
			index = i
			value = v
		}
	}

	return value, index
}

func numify(s string) []int {
	nums := make([]int, len(s))

	for i, v := range s {
		temp, err := strconv.Atoi(string(v))
		nums[i] = temp
		check(err)
	}

	return nums
}

// are you fuggin wit me, do i really have to do this
func raise(x, y int) int {
	switch {
	case y == 0:
		return 1
	case y == 1:
		return x
	case y%2 == 0:
		return raise(x*x, y/2)
	default:
		return x * raise(x*x, (y-1)/2)
	}
}

// part 1

func jolts2(s string) int {
	nums := numify(s)

	tens, ind := maxInt(nums[:(len(s) - 1)])

	ones, _ := maxInt(nums[ind+1:])

	return (tens * 10) + ones
}

// part 2

func jolts12(s string) int {
	// right goes from 11 to 0, it's how many digits we need from the right
	var rec func([]int, int) int

	rec = func(nums []int, right int) int {
		if right == 0 {
			dig, _ := maxInt(nums)
			return dig
		}

		dig, ind := maxInt(nums[:(len(nums) - right)])

		return (dig * (raise(10, right))) + rec(nums[ind+1:], right-1)
	}

	return rec(numify(s), 11)
}

// both parts

func sumBanks(strs []string) (p1, p2 int) {
	for _, v := range strs {
		p1 += jolts2(v)
		p2 += jolts12(v)
	}

	return p1, p2
}

func main() {
	dat, err := os.ReadFile(os.Args[1])
	check(err)

	if dat[len(dat)-1] == '\n' {
		dat = dat[:(len(dat) - 1)]
	}

	split := strings.Split(string(dat), "\n")
	p1, p2 := sumBanks(split)

	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}
