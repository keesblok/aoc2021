package main

import (
	"aoc2021/aocUtil"
	"fmt"
	"io"
	"log"
)

func main() {
	fmt.Println("Part 1 small: ", partOne(aocUtil.LoadInput(1, "_small")))
	fmt.Println("Part 1 final: ", partOne(aocUtil.LoadInput(1, "")))
	fmt.Println("Part 2 small: ", partTwo(aocUtil.LoadInput(1, "_small")))
	fmt.Println("Part 2 final: ", partTwo(aocUtil.LoadInput(1, "")))
}

func partOne(r io.Reader) int {
	numbers, err := aocUtil.ReadInts(r)
	if err != nil {
		log.Fatalf("Could not read the numbers form the file: %v", err)
	}

	increased := 0
	last := numbers[0]
	for _, num := range numbers {
		if num > last {
			increased++
		}
		last = num
	}

	return increased
}

func partTwo(r io.Reader) int {
	numbers, err := aocUtil.ReadInts(r)
	if err != nil {
		log.Fatalf("Could not read the numbers form the file: %v", err)
	}

	if len(numbers) <= 3 {
		return 0
	}

	increased, lastSum := 0, numbers[0] + numbers[1] + numbers[2]
	for i := 3; i < len(numbers); i++ {
		sum := numbers[i-2] + numbers[i-1] + numbers[i]
		if sum > lastSum {
			increased++
		}
		lastSum = sum
	}

	return increased
}