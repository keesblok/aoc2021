package main

import (
	"aoc2021/aocUtil"
	"fmt"
	"io"
	"math"
	"strconv"
)

//Note: My solution for this day was a complete mess and I didn't clean it (yet)
func main() {
	fmt.Println("Part 1 small: ", partOne(aocUtil.LoadInput(3, "_small")))
	fmt.Println("Part 1 final: ", partOne(aocUtil.LoadInput(3, "")))
	fmt.Println("Part 2 small: ", partTwo(aocUtil.LoadInput(3, "_small")))
	fmt.Println("Part 2 final: ", partTwo(aocUtil.LoadInput(3, "")))
}

func partOne(r io.Reader) int {
	lines := aocUtil.ReadLines(r)
	var sums [100]int
	for _, line := range lines {
		for i, num := range line {
			res, _ := strconv.Atoi(string(num))
			sums[i] += res
		}
	}
	gamma := 0
	epsilon := 0
	for i, sum := range sums {
		if sum > len(lines)/2 {
			gamma += int(math.Pow(2, float64(len(lines[0])-i-1)))
		} else {
			epsilon += int(math.Pow(2, float64(len(lines[0])-i-1)))
		}
	}
	return gamma * epsilon
}

func partTwo(r io.Reader) int {
	lines := aocUtil.ReadLines(r)

	sums := make([]int, len(lines[0]))
	var toSave = make([]bool, len(lines))
	for i := 0; i < len(lines); i++ {
		toSave[i] = true
	}
	sums = updateSums(lines, toSave, sums)

	lastTrues := countTrues(toSave)
	for i:=0; countTrues(toSave) > 1; i++ {
		for j, line := range lines {
			if countTrues(toSave) == 1 {
				break
			}
			num, _ := strconv.Atoi(string(line[i]))
			if sums[i]*2 == lastTrues {
				if num == 0 {
					toSave[j] = false
				}
			} else {
				if num == 0 && sums[i]*2 > lastTrues {
					toSave[j] = false
				} else if num == 1 && sums[i]*2 < lastTrues {
					toSave[j] = false
				}
			}
		}
		sums = updateSums(lines, toSave, sums)
		lastTrues = countTrues(toSave)
	}
	ox, _ := strconv.ParseInt(lines[getTruth(toSave)], 2, 64)

	for i := 0; i < len(lines); i++ {
		toSave[i] = true
	}
	sums = updateSums(lines, toSave, sums)

	lastTrues = countTrues(toSave)
	for i:=0; countTrues(toSave) > 1; i++ {
		for j, line := range lines {
			if countTrues(toSave) == 1 {
				break
			}
			num, _ := strconv.Atoi(string(line[i]))
			if sums[i]*2 == lastTrues {
				if num == 1 {
					toSave[j] = false
				}
			} else {
				if num == 1 && sums[i]*2 > lastTrues {
					toSave[j] = false
				} else if num == 0 && sums[i]*2 < lastTrues {
					toSave[j] = false
				}
			}
		}
		sums = updateSums(lines, toSave, sums)
		lastTrues = countTrues(toSave)
	}
	co2, _ := strconv.ParseInt(lines[getTruth(toSave)], 2, 64)

	fmt.Printf("co2 %d  ox %d\n", co2, ox)
	return int(co2 * ox)
}

func countTrues(bools []bool) int {
	num := 0
	for _, b := range bools {
		if b {
			num++
		}
	}
	return num
}

func getTruth(bools []bool) int {
	for i, b := range bools {
		if b {
			return i
		}
	}
	return -1
}

func updateSums(lines []string, toKeep []bool, sums []int) []int {
	l := 0
	for i, _ := range sums {
		sums[i] = 0
	}
	for k, line := range lines {
		if toKeep[k] {
			for i, num := range line {
				res, _ := strconv.Atoi(string(num))
				sums[i] += res
			}
		}
		l++
	}
	return sums
}