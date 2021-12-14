package main

import (
	"advent-calendar/utils"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Signal struct {
	Patterns []string
	Outputs  []string
}

func getNumberMap() []string {
	return []string{
		"abcefg",
		"cf",
		"acdeg",
		"acdfg",
		"bcdf",
		"abdfg",
		"abdefg",
		"acf",
		"abcdefg",
		"abcdfg",
	}
}

func getNumberLengths() []int {
	out := make([]int, len(getNumberMap()))
	for k, val := range getNumberMap() {
		out[k] = len(val)
	}
	return out
}

type Data []Signal

type Day struct {
	data Data
}

func loadData(fname string) Data {
	lines := utils.ReadFile(fname)
	var out []Signal
	for _, line := range lines {
		parts := strings.Split(strings.TrimSpace(line), "|")
		patterns := strings.Split(strings.TrimSpace(parts[0]), " ")
		output := strings.Split(strings.TrimSpace(parts[1]), " ")
		out = append(out, Signal{
			Patterns: patterns,
			Outputs:  output,
		})
	}
	return out
}

func NewDay(fname string) *Day {
	data := loadData(fname)
	return &Day{
		data: data,
	}
}

func fullAssignmentRange() map[rune]string {
	out := make(map[rune]string)
	for _, r := range LETTERS {
		out[r] = LETTERS
	}
	return out
}

func filterBySize(digits []*Digit) map[rune]string {
	out := fullAssignmentRange()
	for _, digit := range digits {
		if len(digit.PossibleMatches) > 0 {
			continue
		}
		assignments := digit.GetPossibleAssignments()
		for _, r := range "abcdefg" {
			if _, ok := assignments[r]; ok {
				out[r] = utils.GetOverlapString(out[r], assignments[r])
			}
		}
	}
	return out
}

func decode(signal Signal) int {
	digits := getDigits(signal.Patterns)
	assignments := getAllAssignments(filterBySize(digits))
	correctAssignment := ""
	for _, assignment := range assignments {
		if checkAssignment(digits, assignment) {
			correctAssignment = assignment
			break
		}
	}

	outputDigits := getDigits(signal.Outputs)
	out := ""
	for _, od := range outputDigits {
		d, ok := od.ApplyAssignments(correctAssignment)
		if !ok {
			panic("Unsolvable")
		}
		out += strconv.Itoa(d)
	}
	o, _ := strconv.Atoi(out)
	return o
}

func getAllAssignments(possibleAssignments map[rune]string) []string {
	letters := []rune("abcdefg")
	var output []string
	var findAll func(current string, toAdd rune)

	findAll = func(current string, toAdd rune) {
		if current != "" || toAdd != 0 {
			current = current + string(toAdd)
			if len(current) == len(letters) {
				output = append(output, current)
				return
			}
		}
		options := utils.GetSetComplementString(possibleAssignments[letters[len(current)]], current)
		for _, option := range options {
			findAll(current, option)
		}
	}
	findAll("", 0)
	return output
}

func checkAssignment(digits []*Digit, assignment string) bool {
	original := "abcdefg"
	for i, assigned := range assignment {
		for _, digit := range digits {
			if ok := digit.ApplyAssignment(rune(original[i]), assigned); !ok {
				for _, digit := range digits {
					digit.Reset()
				}
				return false
			}
		}
	}
	return true
}

func getDigits(patterns []string) []*Digit {
	var digits []*Digit
	for _, val := range patterns {
		digits = append(digits, NewDigit(val))
	}
	return digits
}

func (d *Day) RunPart1() {
	counts := 0
	numberLengths := getNumberLengths()
	sizes := []int{
		numberLengths[1],
		numberLengths[4],
		numberLengths[7],
		numberLengths[8],
	}

	for _, signal := range d.data {
		for _, output := range signal.Outputs {
			if utils.CheckInIntArray(sizes, len(output)) {
				counts += 1
			}
		}
	}

	fmt.Printf("Part 1: %d\n", counts)
}

func (d *Day) RunPart2() {
	sum := 0
	for _, signal := range d.data {
		sum += decode(signal)
	}
	fmt.Printf("Part 2: %d\n", sum)
}

func main() {
	rand.Seed(time.Now().Unix())
	day := NewDay("input/day8.txt")
	day.RunPart1()
	day.RunPart2()
}
