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
	fish := d.data
	for i := 0; i < 80; i++ {
		fish = dayPasses(fish)
	}
	res := len(fish)
	fmt.Printf("Part 1: %d\n", res)
}

func dayPasses(fishies []int) []int {
	var output []int
	for _, fish := range fishies {
		if fish == 0 {
			// fish
			output = append(output, 6)
			// new fish
			output = append(output, 8)
		} else {
			output = append(output, fish-1)
		}
	}
	return output
}

func organizeFishies(fishies []int) []int {
	organized := make([]int, 9)
	for _, fish := range fishies {
		organized[fish] += 1
	}
	return organized
}

func dayPassesForOrganizedFishies(organizedFishies []int) []int {
	output := make([]int, 9)
	for daysLeft, numFishies := range organizedFishies {
		if daysLeft == 0 {
			output[6] += numFishies
			output[8] += numFishies
		} else {
			output[daysLeft-1] += numFishies
		}
	}
	return output
}

func (d *Day) RunPart2() {
	organizedFishies := organizeFishies(d.data)
	for i := 0; i < 256; i++ {
		organizedFishies = dayPassesForOrganizedFishies(organizedFishies)
	}
	res := utils.SumInts(organizedFishies)
	fmt.Printf("Part 2: %d\n", res)
}

func main() {
	day := NewDay("input/day6.txt")
	day.RunPart1()
	day.RunPart2()
}
