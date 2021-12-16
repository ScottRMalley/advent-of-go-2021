package day17

import (
	"advent-calendar/utils"
	"fmt"
)

type Data []string

type Day struct {
	data    Data
}

func loadData(fname string) Data {
	lines := utils.ReadFile(fname)
	return lines
}

func NewDay(fname string) *Day {
	data := loadData(fname)
	return &Day{
		data: data,
	}
}

func (d *Day) RunPart1() {
	fmt.Printf("Part 1: %s\n", "not-implemented")
}

func (d *Day) RunPart2() {
	fmt.Printf("Part 2: %s\n", "not-implemented")
}
