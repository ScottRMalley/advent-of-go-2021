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
	fmt.Printf("Part 1: %d\n", getResult(d.step(d.data.template, 10)))
}

func (d *Day) RunPart2() {
	fmt.Printf("Part 2: %d\n", getResult(d.step(d.data.template, 40)))
}

var cache map[string]map[rune]int

func withCaching(pair string, nsteps int, result map[rune]int) map[rune]int {
	if cache == nil {
		cache = make(map[string]map[rune]int)
	}
	sig := fmt.Sprintf("%s%d", pair, nsteps)
	cache[sig] = result
	return result
}

func getFromCache(pair string, nsteps int) (map[rune]int, bool) {
	if cache == nil {
		cache = make(map[string]map[rune]int)
	}
	sig := fmt.Sprintf("%s%d", pair, nsteps)
	if val, ok := cache[sig]; !ok {
		return nil, false
	} else {
		return val, true
	}
}

func (d *Day) stepPair(pair string, nsteps int) map[rune]int {
	if val, ok := getFromCache(pair, nsteps); ok {
		return val
	}
	if nsteps == 0 {
		return withCaching(pair, nsteps, utils.CountUniqueRunes(pair))
	} else {
		pair1 := string(pair[0]) + d.data.replacements[pair]
		pair2 := d.data.replacements[pair] + string(pair[1])
		return withCaching(
			pair,
			nsteps,
			combineCounts(d.stepPair(pair1, nsteps-1), d.stepPair(pair2, nsteps-1),
				rune(d.data.replacements[pair][0])),
		)
	}
}

func (d *Day) step(template string, nsteps int) map[rune]int {
	var sumCounts map[rune]int
	pairs := getPairs(template)
	for i, pair := range pairs {
		counts := d.stepPair(pair, nsteps)
		if i == 0 {
			sumCounts = counts
		} else {
			sumCounts = combineCounts(sumCounts, counts, rune(pair[0]))
		}
	}
	return sumCounts
}

func getPairs(s string) []string {
	var out []string
	for i := 0; i < len(s)-1; i++ {
		out = append(out, s[i:i+2])
	}
	return out
}

func getResult(counts map[rune]int) int {
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

func combineCounts(c1, c2 map[rune]int, common rune) map[rune]int {
	out := make(map[rune]int)
	for k, v := range c1 {
		if val, ok := out[k]; !ok {
			out[k] = v
		} else {
			out[k] = val + v
		}
	}
	for k, v := range c2 {
		if val, ok := out[k]; !ok {
			out[k] = v
		} else {
			out[k] = val + v
		}
	}
	out[common] -= 1
	return out
}

func main() {
	day := NewDay("input/day14.txt")
	day.RunPart1()
	day.RunPart2()
}
