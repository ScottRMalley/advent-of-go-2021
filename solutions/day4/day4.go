package main

import (
	"advent-calendar/utils"
	"fmt"
	"io/ioutil"
	"strings"
)

type Data []Board

type Board [][]int

type Day struct {
	seq  []int
	data Data
}

func loadData(fname string) ([]int, Data) {
	out, err := ioutil.ReadFile(fname)
	utils.Check(err)
	lines := strings.Split(strings.TrimSpace(string(out)), "\n\n")
	seq := utils.ParseToIntSlice(lines[0], ",")
	boards := make([]Board, len(lines[1:]))
	for i := 0; i < len(lines[1:]); i++ {
		boardString := strings.TrimSpace(lines[i+1])
		rows := strings.Split(boardString, "\n")
		board := make([][]int, len(rows))
		for k, row := range rows {
			rowString := strings.ReplaceAll(row, "  ", " ")
			intRow := utils.ParseToIntSlice(strings.TrimSpace(rowString), " ")
			board[k] = intRow
		}
		boards[i] = board
	}
	return seq, boards
}

func NewDay(fname string) *Day {
	seq, data := loadData(fname)
	return &Day{
		seq:  seq,
		data: data,
	}
}

func (d *Day) RunPart1() {
	choice, sum, lastChoice, lastSum := d.applySeq()
	fmt.Printf("Part 1: %d\n", choice*sum)
	fmt.Printf("Part 2: %d\n", lastChoice*lastSum)
}

func (d *Day) applySeq() (int, int, int, int) {
	firstBoardChoice, firstBoardSum := 0, 0
	lastBoardChoice, lastBoardSum := 0, 0
	numSolved := 0
	firstFound := false
	boards := d.data
	solved := make([]bool, len(boards))
	for _, choice := range d.seq {
		for k, board := range boards {
			boards[k] = d.applyChoice(board, choice)
			if d.checkBoard(boards[k]) {
				if !solved[k]{
					numSolved += 1
					solved[k] = true
				}
				if !firstFound {
					firstBoardChoice, firstBoardSum = choice, d.sumBoard(boards[k])
					firstFound = true
				}
				if numSolved == len(boards) {
					lastBoardChoice, lastBoardSum = choice, d.sumBoard(boards[k])
					return firstBoardChoice, firstBoardSum, lastBoardChoice, lastBoardSum
				}

			}
		}
	}
	return firstBoardChoice, firstBoardSum, lastBoardChoice, lastBoardSum
}

func (d *Day) applyChoice(board Board, choice int) Board {
	for i := range board {
		for j := range board[i] {
			if board[i][j] == choice {
				board[i][j] = -1
			}
		}
	}
	return board
}

func (d *Day) checkBoard(board Board) bool {
	for i := range board {
		if utils.SumInts(board[i]) == 0-len(board[i]) {
			return true
		}
	}
	for k := range board[0] {
		if utils.SumInts(utils.GetColumn(board, k)) == 0-len(utils.GetColumn(board, k)) {
			return true
		}
	}
	return false
}

func (d *Day) sumBoard(board Board) int {
	sum := 0
	for i := range board {
		for j := range board[i] {
			if board[i][j] > 0 {
				sum += board[i][j]
			}
		}
	}
	return sum
}


func (d *Day) RunPart2() {
}

func main() {
	day := NewDay("input/day4.txt")
	day.RunPart1()
	day.RunPart2()
}
