package main

import (
	"advent-calendar/utils"
	"fmt"
	"github.com/fatih/color"
	"strings"
)

type Data [][]int

type Day struct {
	data Data
}

func loadData(fname string) Data {
	lines := utils.ReadFile(fname)
	grid := make([][]int, len(lines))
	for i, line := range lines {
		nums := utils.ParseToIntSlice(strings.TrimSpace(line), "")
		grid[i] = nums
	}
	return grid
}

func NewDay(fname string) *Day {
	data := loadData(fname)
	return &Day{
		data: data,
	}
}

func (d *Day) findNeighbors(x, y int) []int {
	X := len(d.data)
	Y := len(d.data[0])
	var out []int
	for _, v := range []int{-1, 0, 1} {
		for _, w := range []int{-1, 0, 1} {
			if v == w {
				continue
			}
			xx := x + v
			yy := y + w
			if xx < 0 || yy < 0 {
				continue
			}
			if xx > X-1 || yy > Y-1 {
				continue
			}
			out = append(out, d.data[xx][yy])
		}
	}
	return out
}

func (d *Day) checkForLowPoint(x, y int) (int, bool) {
	n := d.findNeighbors(x, y)
	for _, nn := range n {
		if d.data[x][y] >= nn {
			return 0, false
		}
	}
	return d.data[x][y] + 1, true
}

func (d *Day) RunPart1() {
	sum := 0
	for x := range d.data {
		for y := range d.data[0] {
			if risk, ok := d.checkForLowPoint(x, y); ok {
				sum += risk
			}
		}
	}
	fmt.Printf("Part 1: %d\n", sum)
}

func (d *Day) findBasinSize(x, y int, print bool) (int, bool) {
	// first check for low point
	if _, ok := d.checkForLowPoint(x, y); !ok {
		return 0, false
	}

	size := 0
	var xs []int
	var ys []int
	for xx := range d.data {
		for yy := range d.data[0] {
			if d.isPointInBasin(xx, yy, x, y) {
				size += 1
				xs = append(xs, xx)
				ys = append(ys, yy)
			}
		}
	}

	if print {
		d.printBasin(xs, ys)
	}

	return size, true
}

func (d *Day) isPointInBasin(x, y, lowX, lowY int) bool {
	X := len(d.data)
	Y := len(d.data[0])

	// check if we are a 9
	if d.data[x][y] == 9 {
		return false
	}

	// check if we reached the basin
	if x == lowX && y == lowY {
		return true
	}

	// check if we reached another basin
	if _, ok := d.checkForLowPoint(x, y); ok {
		return false
	}

	// check all neighbors
	for _, v := range []int{-1, 0, 1} {
		for _, w := range []int{-1, 0, 1} {
			if v == w || v == -w {
				continue
			}
			xx := x + v
			yy := y + w
			if xx < 0 || yy < 0 {
				continue
			}

			if xx > X-1 || yy > Y-1 {
				continue
			}
			// if less than we can move there
			if d.data[xx][yy] < d.data[x][y] {
				if d.isPointInBasin(xx, yy, lowX, lowY) {
					return true
				}
			}
		}
	}
	return false
}

func (d *Day) RunPart2() {
	var lowX []int
	var lowY []int
	for x := range d.data {
		for y := range d.data[0] {
			if _, ok := d.checkForLowPoint(x, y); ok {
				lowX = append(lowX, x)
				lowY = append(lowY, y)
			}
		}
	}

	prod := 1
	var sizes []int
	for i := range lowX {
		if size, ok := d.findBasinSize(lowX[i], lowY[i], false); ok {
			sizes = append(sizes, size)
		}
	}

	for j := 0; j < 3; j++ {
		max := utils.Max(sizes)
		sizes = utils.RemoveFirstFromArray(sizes, max)
		prod *= max
	}

	fmt.Printf("Part 2: %d\n", prod)
}

func (d *Day) printBasin(xs, ys []int) {
	white := color.New(color.FgWhite)
	boldWhite := white.Add(color.Bold)

	for x := range d.data {
		for y := range d.data[0] {
			if checkIfPointIn(xs, ys, x, y) {
				_, _ = boldWhite.Printf(" %d", d.data[x][y])
			} else {
				fmt.Printf(" %d", d.data[x][y])
			}
		}
		fmt.Print("\n")
	}
}

func checkIfPointIn(xs, ys []int, x, y int) bool {
	for i := range xs {
		if xs[i] == x && ys[i] == y {
			return true
		}
	}
	return false
}

func main() {
	day := NewDay("input/day9.txt")
	day.RunPart1()
	day.RunPart2()
}
