package main

import (
	"advent-calendar/utils"
	"fmt"
)

type Data []int

type Day struct {
	data Data
}

func loadData(fname string) Data {
	lines := utils.ReadFile(fname)
	var out []int
	for _, line := range lines {
		out = utils.ParseToIntSlice(line, ",")
	}
	return out
}

func NewDay(fname string) *Day {
	data := loadData(fname)
	return &Day{
		data: data,
	}
}

func (d *Day) RunPart1() {
	fmt.Printf("Part 1: %d\n", calculateFuel(d.data, utils.MedianInt(d.data)))
}

func calculateFuel(positions []int, destination int) int {
	fuel := 0
	for _, pos := range positions {
		fuel += utils.Abs(pos - destination)
	}
	return fuel
}

func calculateFuelNonlinear(positions []int, destination int) int {
	fuel := 0
	for _, pos := range positions {
		fuel += utils.SumInts(utils.GetIntRange(0, utils.Abs(destination-pos)))
	}
	return fuel
}

func (d *Day) RunPart2() {
	var fuels []int
	for k := utils.Min(d.data); k < utils.Max(d.data); k++ {
		fuels = append(fuels, calculateFuelNonlinear(d.data, k))
	}
	fmt.Printf("Part 2: %d\n", utils.Min(fuels))
}

func main() {
	day := NewDay("input/day7.txt")
	day.RunPart1()
	day.RunPart2()
}
