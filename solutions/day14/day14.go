package main

import (
	"advent-calendar/utils"
	"fmt"
	"strings"
)

type Data struct {
	template     string
	replacements map[string]string
}

type Day struct {
	data Data
}

func loadData(fname string) Data {
	fstring := utils.RawFileString(fname)
	parts := strings.Split(strings.TrimSpace(fstring), "\n\n")
	template := parts[0]
	lines := strings.Split(strings.TrimSpace(parts[1]), "\n")
	replacements := make(map[string]string)
	for _, line := range lines {
		mols := strings.Split(strings.TrimSpace(line), " -> ")
		from := mols[0]
		to := mols[1]
		replacements[from] = to
	}
	return Data{
		template:     template,
		replacements: replacements,
	}
}

func NewDay(fname string) *Day {
	data := loadData(fname)
	return &Day{
		data: data,
	}
}

func (d *Day) RunPart1() {
	for i := 0; i < 10; i++ {
		d.step()
	}
	result := getResult(d.data.template)
	fmt.Printf("Part 1: %d\n", result)
}

func (d *Day) step() {
	pairs := getPairs(d.data.template)
	d.data.template = d.insertReplacements(pairs)
}

func getPairs(s string) []string {
	var out []string
	for i := 0; i < len(s)-1; i++ {
		out = append(out, s[i:i+2])
	}
	return out
}

func getResult(s string) int {
	counts := utils.CountUniqueRunes(s)
	max := 0
	min := -1
	for _, v := range counts {
		if min == -1 {
			min = v
		}
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return max - min
}

func (d *Day) insertReplacements(pairs []string) string {
	out := ""
	for i, pair := range pairs {
		if i < len(pairs)-1 {
			out += string(pair[0]) + d.data.replacements[pair]
		} else {
			out += string(pair[0]) + d.data.replacements[pair] + string(pair[1])
		}
	}
	return out
}

func (d *Day) RunPart2() {
	for i := 0; i < 30; i++ {
		d.step()
	}
	result := getResult(d.data.template)
	fmt.Printf("Part 2: %d\n", result)
}

func main() {
	day := NewDay("input/day14.txt")
	day.RunPart1()
	day.RunPart2()
}
