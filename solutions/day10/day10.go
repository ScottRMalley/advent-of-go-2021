package main

import (
	"advent-calendar/utils"
	"fmt"
	"sort"
	"strings"
)

type Data []string

type Day struct {
	data Data
}

func getCharacterBrackets() []string {
	return []string{
		"{}",
		"[]",
		"<>",
		"()",
	}
}

func getClosingCharacterPointsPart1() map[rune]int {
	return map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
}

func getClosingCharacterPointsPart2() map[rune]int {
	return map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}
}

func getClosingCharacterMap() map[rune]rune {
	return map[rune]rune{
		'{': '}',
		'(': ')',
		'[': ']',
		'<': '>',
	}
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
	sum := 0
	for _, line := range d.data {
		if _, end, ok := CheckCorrect(line); !ok {
			sum += end
		}
	}
	fmt.Printf("Part 1: %d\n", sum)
}

func CheckCorrect(line string) (string, int, bool) {
	toCheck := line
	l := len(toCheck)

	for {
		for _, bracket := range getCharacterBrackets() {
			toCheck = strings.ReplaceAll(toCheck, bracket, "")
		}
		if toCheck == "" {
			return "", 0, true
		} else if len(toCheck) == l {
			point, ok := GetFirstClosingPoint(toCheck)
			return toCheck, point, ok
		} else {
			l = len(toCheck)
		}
	}
}

func GetFirstClosingPoint(s string) (int, bool) {
	for _, c := range s {
		for closing, pointVal := range getClosingCharacterPointsPart1() {
			if c == closing {
				return pointVal, false
			}
		}
	}
	return 0, true
}

func RemoveAllCorrectBrackets(s string) string {
	out := s
	for _, bracket := range getCharacterBrackets() {
		out = strings.ReplaceAll(out, bracket, "")
	}
	return out
}

func Autocomplete(s string) int {
	out := ""
	toComplete := s

	for {
		if len(toComplete) == 0 {
			return getAutocorrectPoints(out)
		}
		lastCharacter := rune(toComplete[len(toComplete)-1])
		toComplete += string(getClosingCharacterMap()[lastCharacter])
		out += string(getClosingCharacterMap()[lastCharacter])
		toComplete = RemoveAllCorrectBrackets(toComplete)
	}
}

func getAutocorrectPoints(s string) int {
	out := 0
	for _, c := range s {
		out = 5*out + getClosingCharacterPointsPart2()[c]
	}
	return out
}

func (d *Day) RunPart2() {
	var scores []int
	for _, line := range d.data {
		if corrected, _, ok := CheckCorrect(line); ok {
			scores = append(scores, Autocomplete(corrected))
		}
	}
	sort.Ints(scores)
	out := scores[len(scores)/2]
	fmt.Printf("Part 2: %d\n", out)
}

func main() {
	day := NewDay("input/day10.txt")
	day.RunPart1()
	day.RunPart2()
}
