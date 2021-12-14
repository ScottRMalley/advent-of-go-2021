package main

import (
	"advent-calendar/utils"
	"fmt"
)

type Data [][]int

type Day struct {
	data   Data
	backup Data
}

func loadData(fname string) Data {
	lines := utils.ReadFile(fname)
	out := make([][]int, 10)
	for k, line := range lines {
		out[k] = utils.ParseToIntSlice(line, "")
	}
	return out
}

func NewDay(fname string) *Day {
	data := loadData(fname)
	backup := loadData(fname)
	return &Day{
		data:   data,
		backup: backup,
	}
}

func (d *Day) reset() {
	d.data = d.backup
}

func (d *Day) RunPart1() {
	sum := 0
	for i := 0; i < 100; i++ {
		//d.printGrid()
		sum += d.step()
	}
	fmt.Printf("Part 1: %d\n", sum)
}

func replace(in [][]int, val int, newVal int) [][]int {
	out := utils.Zeros2D(10, 10)
	for i := range in {
		for j := range in[0] {
			if in[i][j] == val {
				out[i][j] = newVal
			} else {
				out[i][j] = in[i][j]
			}
		}
	}
	return out
}

func (d *Day) step() int {
	grid := make([][]int, len(d.data))
	for i := 0; i < len(d.data); i++ {
		grid[i] = utils.AddConstant(d.data[i], 1)
	}
	flashes := 0
	for {
		if utils.AllValuesLessThan(grid, 10) {
			d.data = replace(grid, -1, 0)
			return flashes
		}
		newGrid, newFlashes := d.IncreaseNeighbors(grid)
		grid = newGrid
		flashes += newFlashes
	}
}

func (d *Day) getNeighbors(x, y int) ([]int, []int) {
	var xs []int
	var ys []int
	for _, dx := range []int{-1, 0, 1} {
		for _, dy := range []int{-1, 0, 1} {
			if dx == 0 && dy == 0 {
				continue
			} else if x+dx > len(d.data)-1 || x+dx < 0 {
				continue
			} else if y+dy > len(d.data[0])-1 || y+dy < 0 {
				continue
			}
			xs = append(xs, x+dx)
			ys = append(ys, y+dy)
		}
	}
	return xs, ys
}

func (d *Day) IncreaseNeighbors(grid [][]int) ([][]int, int) {
	flashes := 0
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			if grid[x][y] > 9 {
				flashes += 1
				xs, ys := d.getNeighbors(x, y)
				for i := range xs {
					if grid[xs[i]][ys[i]] != -1 {
						grid[xs[i]][ys[i]] += 1
					}
				}
				grid[x][y] = -1
			}
		}
	}
	return grid, flashes
}

func (d *Day) printGrid() {
	for _, line := range d.data {
		fmt.Println(line)
	}
	fmt.Println()
}

func (d *Day) synchronize() int {
	i := 0
	for {
		i += 1
		d.step()
		if utils.AllValuesLessThan(d.data, 1) {
			return i
		}
	}
}

func (d *Day) RunPart2() {
	d.reset()
	result := d.synchronize()
	fmt.Printf("Part 2: %d\n", result)
}

func main() {
	day := NewDay("input/day11.txt")
	day.RunPart1()
	day.RunPart2()
}
