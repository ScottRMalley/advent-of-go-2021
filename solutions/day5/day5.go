package main

import (
	"advent-calendar/utils"
	"fmt"
	"strings"
)

type Direction struct {
	From []int
	To   []int
}

type Data []Direction

type Day struct {
	data Data
}

func loadData(fname string) Data {
	lines := utils.ReadFile(fname)
	var data []Direction
	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		from := utils.ParseToIntSlice(parts[0], ",")
		to := utils.ParseToIntSlice(parts[1], ",")
		data = append(data, Direction{To: to, From: from})
	}
	return data
}

func NewDay(fname string) *Day {
	data := loadData(fname)
	return &Day{
		data: data,
	}
}

func (d *Day) RunPart1() {
	X, Y := d.getMax()
	grid := utils.Zeros2D(X, Y)
	for _, line := range d.data {
		if line.From[0] == line.To[0] {
			for _, y := range utils.GetIntRange(line.From[1], line.To[1]) {
				grid[line.From[0]][y] += 1
			}
		} else if line.From[1] == line.To[1] {
			for _, x := range utils.GetIntRange(line.From[0], line.To[0]) {
				grid[x][line.From[1]] += 1
			}
		}
	}
	total := utils.CountWhere(utils.Flatten(grid), func(i int) bool { return i >= 2 })
	fmt.Printf("Part 1: %d\n", total)
}

func (d *Day) getMax() (int, int) {
	maxX := 0
	maxY := 0
	for _, point := range d.data {
		if point.From[0] >= maxX {
			maxX = point.From[0]
		}
		if point.From[1] >= maxY {
			maxY = point.From[1]
		}
		if point.To[0] >= maxX {
			maxX = point.To[0]
		}
		if point.To[1] >= maxY {
			maxY = point.To[1]
		}
	}
	return maxX + 1, maxY + 1
}

func (d *Day) RunPart2() {
	X, Y := d.getMax()
	grid := utils.Zeros2D(X, Y)
	for _, line := range d.data {
		xRange := utils.GetIntRange(line.From[0], line.To[0])
		yRange := utils.GetIntRange(line.From[1], line.To[1])
		if line.From[0] == line.To[0] {
			for _, y := range yRange {
				grid[line.From[0]][y] += 1
			}
		} else if line.From[1] == line.To[1] {
			for _, x := range xRange {
				grid[x][line.From[1]] += 1
			}
		} else {
			for i := 0; i < len(xRange); i++ {
				grid[xRange[i]][yRange[i]] += 1
			}
		}
	}
	total := utils.CountWhere(utils.Flatten(grid), func(i int) bool { return i >= 2 })
	fmt.Printf("Part 2: %d\n", total)
}

func main() {
	day := NewDay("input/day5.txt")
	day.RunPart1()
	day.RunPart2()
}
