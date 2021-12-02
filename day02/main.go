package main

import (
	"aoc2021/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	part1("input.txt")
	part2("input.txt")
}

func part1(filename string) {
	data := utils.ReadLines(filename)
	hpos := 0
	vpos := 0
	for _, instruction := range data {
		s := strings.Fields(instruction) // forward 5
		number, _ := strconv.Atoi(s[1])
		switch s[0] {
		case "forward":
			hpos += number
		case "up":
			vpos -= number
		case "down":
			vpos += number
		}
	}
	fmt.Println("hpos", hpos, "vpos", vpos)
	fmt.Println("part1", hpos*vpos)
}

func part2(filename string) {
	data := utils.ReadLines(filename)
	hpos := 0
	vpos := 0
	aim := 0
	for _, instruction := range data {
		s := strings.Fields(instruction) // forward 5
		number, _ := strconv.Atoi(s[1])
		switch s[0] {
		case "forward":
			hpos += number
			vpos += aim * number
		case "up":
			aim -= number
		case "down":
			aim += number
		}
	}
	fmt.Println("hpos", hpos, "vpos", vpos)
	fmt.Println("part2", hpos*vpos)
}
