package main

import (
	"advent-calendar/utils"
	"fmt"
	"strings"
)

type Data map[string][]string

type Day struct {
	data Data
}

func loadData(fname string) Data {
	lines := utils.ReadFile(fname)
	out := make(map[string][]string)
	for _, line := range lines {
		parts := strings.Split(line, "-")
		from := strings.TrimSpace(parts[0])
		to := strings.TrimSpace(parts[1])
		out[from] = append(out[from], to)
		out[to] = append(out[to], from)
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
	paths := getAllPathsPart1(d.data)
	fmt.Printf("Part 1: %d\n", len(paths))
}

func getAllPathsPart1(connections map[string][]string) []string {
	var output []string
	var findAll func(current []string, toAdd string)

	findAll = func(current []string, toAdd string) {
		if toAdd != "" {
			current = append(current, toAdd)
			if current[len(current)-1] == "end" {
				output = append(output, strings.Join(current, ","))
				return
			}
		}
		options := connections[current[len(current)-1]]
		for _, option := range options {
			if utils.IsLower(option) && utils.ValInStringSlice(current, option) {
				continue
			}
			findAll(current, option)
		}
	}

	findAll([]string{"start"}, "")
	return output
}

func getAllPathsPart2(connections map[string][]string) []string {
	var output []string
	var findAll func(current []string, toAdd string)

	findAll = func(current []string, toAdd string) {
		if toAdd != "" {
			current = append(current, toAdd)
			if current[len(current)-1] == "end" {
				output = append(output, strings.Join(current, ","))
				return
			}
		}
		options := connections[current[len(current)-1]]
		for _, option := range options {
			if option == "start" {
				continue
			}
			if utils.IsLower(option) && checkIfDoubleLowerUsed(current) && utils.ValInStringSlice(current, option) {
				continue
			}
			findAll(current, option)
		}
	}

	findAll([]string{"start"}, "")
	return output
}

func checkIfDoubleLowerUsed(current []string) bool {
	for _, s := range current {
		if utils.IsLower(s) {
			occurrences := utils.CountOccurrenceStringSlice(current, s)
			if occurrences > 1 {
				return true
			}
		}
	}
	return false
}

func (d *Day) RunPart2() {
	paths := getAllPathsPart2(d.data)
	fmt.Printf("Part 2: %d\n", len(paths))
}

func main() {
	day := NewDay("input/day12.txt")
	day.RunPart1()
	day.RunPart2()
}
