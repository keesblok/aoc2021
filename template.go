package main

import (
	"aoc2021/aocUtil"
	"fmt"
	"io"
)

func main() {
	fmt.Println("Part 1 small: ", partOne(aocUtil.LoadInput(1, "_small")))
	fmt.Println("Part 1 final: ", partOne(aocUtil.LoadInput(1, "")))
	fmt.Println("Part 2 small: ", partTwo(aocUtil.LoadInput(1, "_small")))
	fmt.Println("Part 2 final: ", partTwo(aocUtil.LoadInput(1, "")))
}

func partOne(r io.Reader) int {
	numbers, _ := aocUtil.ReadInts(r)
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func partTwo(r io.Reader) int {
	lines := aocUtil.ReadLines(r)
	for _, line := range lines {
		println(line)
	}
	return len(lines)
}