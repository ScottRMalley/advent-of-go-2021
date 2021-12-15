package day1

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Data []int

type Day struct {
	data Data
}

func loadData(fname string) Data {
	out, err := ioutil.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	var output []int
	for _, line := range lines {
		i, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		output = append(output, i)
	}
	return output
}

func NewDay(fname string) *Day {
	data := loadData(fname)
	return &Day{data}
}

func (d *Day) RunPart1() {
	increased := 0
	for i := 1; i<len(d.data); i++ {
		if d.data[i-1] < d.data[i] {
			increased += 1
		}
	}
	fmt.Printf("Part 1: %d\n", increased)
}

func (d *Day) RunPart2() {
	var a []int
	for i := 0; i<len(d.data)-2; i++ {
		a = append(a, d.data[i] + d.data[i+1] + d.data[i+2])
	}
	increased := 0
	for i := 1; i<len(a); i++ {
		if a[i-1] < a[i] {
			increased += 1
		}
	}
	fmt.Printf("Part 2: %d\n", increased)

}

func main() {
	day := NewDay("input/day1.txt")
	day.RunPart1()
	day.RunPart2()
}
