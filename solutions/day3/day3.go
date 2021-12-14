package main

import (
	"advent-calendar/utils"
	"fmt"
	"strconv"
)

type Data [][]int

type Day struct {
	data    Data
}

func loadData(fname string) Data {
	lines := utils.ReadFile(fname)
	output := make([][]int, len(lines))
	for k, line := range lines {
		a := utils.ParseToIntSlice(line, "")
		output[k] = a
	}
	return output
}

func NewDay(fname string) *Day {
	data := loadData(fname)
	return &Day{
		data: data,
	}
}

func (d *Day) RunPart1() {
	l := len(d.data[0])
	epsilon := ""
	gamma := ""
	for pos := 0; pos < l; pos++ {
		sum := utils.SumInts(utils.GetColumn(d.data, pos))
		if sum > len(d.data)/2 {
			epsilon += "0"
			gamma += "1"
		} else {
			epsilon += "1"
			gamma += "0"
		}
	}
	g, _ := strconv.ParseInt(gamma, 2, 0)
	e, _ := strconv.ParseInt(epsilon, 2, 0)
	fmt.Printf("Part 1: %d\n", g*e)
}

func (d *Day) RunPart2() {
	oxygen := d.getOxygen()
	co2 := d.getCo2()
	// convert to strings
	oxygenString := utils.ToString(oxygen)
	co2String := utils.ToString(co2)
	// parse stings as base 2
	ob2, _ := strconv.ParseInt(oxygenString, 2, 0)
	cb2, _ := strconv.ParseInt(co2String, 2, 0)
	fmt.Printf("Part 2: %d", ob2*cb2)
}

func (d *Day) getCo2() []int {
	valid := utils.IntRange(0, len(d.data))
	pos := 0
	for {
		if len(valid) == 1 {
			return d.data[valid[0]]
		}
		leastCommon := d.leastCommonBit(valid, pos)
		newValid := valid
		for _, i := range valid {
			if d.data[i][pos] != leastCommon {
				newValid = utils.RemoveFromArray(newValid, i)
			}
		}
		valid = newValid
		pos += 1
	}
}

func (d *Day) getOxygen() []int {
	var valid []int
	for i := range d.data {
		valid = append(valid, i)
	}
	pos := 0
	for {
		if len(valid) == 1 {
			return d.data[valid[0]]
		}
		mostCommon := d.mostCommonBit(valid, pos)
		newValid := valid
		for _, i := range valid {
			if d.data[i][pos] != mostCommon {
				newValid = utils.RemoveFromArray(newValid, i)
			}
		}
		valid = newValid
		pos += 1
	}
}

func (d *Day) mostCommonBit(a []int, pos int) int {
	ones := utils.SumInts(utils.GetIntLocs(utils.GetColumn(d.data, pos), a))
	zeros := len(a) - ones
	if ones > zeros {
		return 1
	} else if zeros > ones {
		return 0
	} else {
		return 1
	}
}

func (d *Day) leastCommonBit(a []int, pos int) int {
	ones := utils.SumInts(utils.GetIntLocs(utils.GetColumn(d.data, pos), a))
	zeros := len(a) - ones
	if ones > zeros {
		return 0
	} else if zeros > ones {
		return 1
	} else {
		return 0
	}
}

func main() {
	day := NewDay("input/day3.txt")
	day.RunPart1()
	day.RunPart2()
}
