package main

import (
	"aoc2021/utils"
	"fmt"
)

func main() {
	part1("input.txt")
	part2("input.txt")
}

func part1(filename string) {
	ints := utils.ReadInts(filename)

	prev := 0
	increase := 0
	for _, num := range ints {
		if prev > 0 && num > prev {
			increase++
		}
		prev = num
	}
	fmt.Println("part1", increase)
}

func part2(filename string) {
	ints := utils.ReadInts(filename)

	increase := 0
	for i := 3; i < len(ints); i++ {
		if window(ints, i) > window(ints, i-1) {
			increase++
		}
	}
	fmt.Println("part2", increase)
}

func window(a []int, i int) int {
	return a[i-2] + a[i-1] + a[i]
}
