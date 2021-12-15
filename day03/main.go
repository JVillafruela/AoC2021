package main

import (
	"aoc2021/utils"
	"fmt"
	"strconv"
)

func main() {
	part1("input.txt")
}

func part1(filename string) {
	data := utils.ReadLines(filename)
	length := len(data[0])
	ones := make([]uint, length)
	zeroes := make([]uint, length)
	for _, str := range data {
		for i := 0; i < length; i++ {
			switch str[i] {
			case '0':
				zeroes[i]++
			case '1':
				ones[i]++
			}
		}
	}

	gamma := ""
	epsilon := ""
	for i := 0; i < length; i++ {
		if ones[i] > zeroes[i] {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}
	fmt.Println("part1", gamma, epsilon)

	g, _ := strconv.ParseInt(gamma, 2, 64)
	e, _ := strconv.ParseInt(epsilon, 2, 64)
	fmt.Println("part1", g, e, "answer", g*e)
}
