package main

import (
	"aoc2021/aocUtil"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Part 1 small: ", partOne(aocUtil.LoadInput(2, "_small")))
	fmt.Println("Part 1 final: ", partOne(aocUtil.LoadInput(2, "")))
	fmt.Println("Part 2 small: ", partTwo(aocUtil.LoadInput(2, "_small")))
	fmt.Println("Part 2 final: ", partTwo(aocUtil.LoadInput(2, "")))
}

func partOne(r io.Reader) int {
	lines := aocUtil.ReadLines(r)

	pos, depth := 0, 0
	for _, line := range lines {
		words := strings.Fields(line)
		x, _ := strconv.Atoi(words[1])
		switch words[0] {
		case "forward":
			pos += x
		case "up":
			depth -= x
		case "down":
			depth += x
		}
	}
	return pos * depth
}

func partTwo(r io.Reader) int {
	lines := aocUtil.ReadLines(r)

	pos, depth, aim := 0, 0, 0
	for _, line := range lines {
		words := strings.Fields(line)
		x, _ := strconv.Atoi(words[1])
		switch words[0] {
		case "forward":
			pos += x
			depth += aim * x
		case "up":
			aim -= x
		case "down":
			aim += x
		}
	}
	return pos * depth
}