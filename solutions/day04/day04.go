package main

import (
	"aoc2021/aocUtil"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Part 1 small: ", partOne(aocUtil.LoadInput(4, "_small")))
	fmt.Println("Part 1 final: ", partOne(aocUtil.LoadInput(4, "")))
	fmt.Println("Part 2 small: ", partTwo(aocUtil.LoadInput(4, "_small")))
	fmt.Println("Part 2 final: ", partTwo(aocUtil.LoadInput(4, "")))
}

func partOne(r io.Reader) int {
	lines := aocUtil.ReadLines(r)
	draws := getDraws(lines[0])
	boards := makeBoards(lines)

	lastNumber, winner := 0, 0
	for _, num := range draws {
		boards = draw(boards, num)
		win, winners := getWinners(boards)
		if win {
			lastNumber = num
			winner = winners[0]
			break
		}
	}

	sums := sum(boards, false)
	return sums[winner]*lastNumber
}

func partTwo(r io.Reader) int {
	lines := aocUtil.ReadLines(r)
	draws := getDraws(lines[0])
	boards := makeBoards(lines)

	lastNumber, lastWinner := -1, -1
	wins := make([]bool, len(boards))

	for _, num := range draws {
		boards = draw(boards, num)
		win, winners := getWinners(boards)
		if win {
			for _, winnerBoard := range winners {
				if !wins[winnerBoard] {
					wins[winnerBoard] = true
					lastNumber = num
					lastWinner = winnerBoard
				}
			}

			if len(winners) == len(boards) {
				break
			}
		}
	}
	sums := sum(boards, false)
	return sums[lastWinner]*lastNumber
}

func getDraws(line string) []int {
	drawsString := strings.Split(line, ",")
	draws := make([]int, len(drawsString))

	for i := 0; i < len(drawsString); i++ {
		var err error
		draws[i], err = strconv.Atoi(drawsString[i])
		if err != nil {
			log.Fatalf("could not convert list to int: %s\n", drawsString[i])
		}
	}

	return draws
}

func makeBoards(lines []string) [][][]place {
	boards := make([][][]place, (len(lines)-1)/6)

	for boardNumber, _ := range boards {
		boards[boardNumber] = make([][]place, 5)

		for row := 0; row < 5; row++ {
			boards[boardNumber][row] = make([]place, 5)

			line := strings.Replace(lines[2 + boardNumber * 6 + row], "  ", " ", -1)
			nums := strings.Split(strings.Trim(line, " "), " ")

			for column := 0; column < 5; column++ {
				var err error
				boards[boardNumber][row][column].value, err = strconv.Atoi(strings.Trim(nums[column], " "))
				if err != nil {
					log.Printf("could not convert %s from board: %d, row: %d, column: %d to int\n", strings.Trim(nums[column], " "), boardNumber, row, column)
				}
			}
		}
	}

	return boards
}

func draw(boards [][][]place, num int) [][][]place {
	for boardNumber, board := range boards {
		for rowNumber, row := range board {
			for columnNumber, spot := range row {
				if spot.value == num {
					boards[boardNumber][rowNumber][columnNumber].drawn = true
				}
			}
		}
	}

	return boards
}

func getWinners(boards [][][]place) (bool, []int) {
	win := false
	var winBoards []int

	for boardNumber, board := range boards {
		drawnCollum := make([]int, 5)

		for _, row := range board {
			drawnRow := 0

			for j, spot := range row {
				if spot.drawn {
					drawnCollum[j]++
					drawnRow ++
				}
			}

			if drawnRow == 5 {
				win = true
				winBoards = append(winBoards, boardNumber)
			}
		}

		for _, column := range drawnCollum {
			if column == 5 {
				win = true
				winBoards = append(winBoards, boardNumber)
			}
		}
	}

	return win, winBoards
}

func sum(boards [][][]place, marked bool) []int {
	sums := make([]int, len(boards))

	for boardNumber, board := range boards {
		for _, row := range board {
			for _, spot := range row {
				if spot.drawn == marked {
					sums[boardNumber] += spot.value
				}
			}
		}
	}

	return sums
}

type place struct {
	value int
	drawn bool
}