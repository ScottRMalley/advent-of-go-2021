package day15

import (
	"advent-calendar/utils"
	"fmt"
	"sort"
)

type Data [][]int

type Day struct {
	data Data
}

func loadData(fname string) Data {
	lines := utils.ReadFile(fname)
	grid := make([][]int, len(lines))
	for i, line := range lines {
		grid[i] = utils.ParseToIntSlice(line, "")
	}
	return grid
}

func NewDay(fname string) *Day {
	data := loadData(fname)
	return &Day{
		data: data,
	}
}

type cell [][]int

func (c cell) Len() int {
	return len(c)
}

func (c cell) Less(i, j int) bool {
	return c[i][2] < c[j][2]
}

func (c cell) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func shortest(grid [][]int) int {
	maxX, maxY := len(grid), len(grid[0])
	dis := utils.Zeros2D(maxX, maxY)
	for i := range dis {
		for j := range dis[0] {
			dis[i][j] = 100000000000000000
		}
	}

	st := [][]int{{0, 0, 0}}
	dis[0][0] = grid[0][0]

	inGrid := func(i, j int) bool {
		return i >= 0 && i < maxX && j >= 0 && j < maxY
	}

	for {
		if len(st) == 0 {
			break
		}
		k := st[0]
		st = st[1:]

		options := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
		for i := range options {
			x := k[0] + options[i][0]
			y := k[1] + options[i][1]
			if !inGrid(x, y) {
				continue
			}
			if dis[x][y] > dis[k[0]][k[1]]+grid[x][y] {
				dis[x][y] = dis[k[0]][k[1]] + grid[x][y]
				st = append(st, []int{x, y, dis[x][y]})
			}
		}
		sort.Sort(cell(st))
	}
	return dis[maxX-1][maxY-1]
}

func (d *Day) RunPart1() {
	min := shortest(d.data) - d.data[0][0]
	fmt.Printf("Part 1: %d\n", min)
}

func wrap(val, i, j int) int {
	if val+i+j < 10 {
		return val + i + j
	}
	return 1 + ((val + i + j) % 10)
}

func (d *Day) RunPart2() {
	maxX, maxY := len(d.data), len(d.data[0])
	grid := utils.Zeros2D(5*maxX, 5*maxY)
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			for k := range d.data {
				for l := range d.data {
					grid[maxX*i+k][maxY*j+l] = wrap(d.data[k][l], i, j)
				}
			}
		}
	}
	min := shortest(grid) - grid[0][0]
	fmt.Printf("Part 2: %d\n", min)
}
