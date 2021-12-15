package day2

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Data []string

type Day struct {
	data Data
}

func loadData(fname string) Data {
	out, err := ioutil.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	return lines
}

func NewDay(fname string) *Day {
	data := loadData(fname)
	return &Day{data}
}

func (d *Day) RunPart1() {
	depth := 0
	position := 0
	for _, line := range d.data {
		vals := strings.Split(line, " ")
		dis, _ := strconv.Atoi(vals[1])
		if vals[0] == "down"{
			depth += dis
		}
		if vals[0] == "up" {
			depth -= dis
		}
		if vals[0] == "forward" {
			position += dis
		}
	}
	fmt.Printf("Part 1: %d\n", depth*position)
}

func (d *Day) RunPart2() {
	depth := 0
	position := 0
	aim := 0
	for _, line := range d.data {
		vals := strings.Split(line, " ")
		dis, _ := strconv.Atoi(vals[1])
		if vals[0] == "down"{
			aim += dis
		}
		if vals[0] == "up" {
			aim -= dis
		}
		if vals[0] == "forward" {
			position += dis
			depth += aim*dis
		}
	}
	fmt.Printf("Part 2: %d\n", depth*position)
}

func main() {
	day := NewDay("input/day2.txt")
	day.RunPart1()
	day.RunPart2()
}
