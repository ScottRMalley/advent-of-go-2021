package day17

import (
	"advent-calendar/utils"
	"fmt"
	"sort"
	"strings"
)

type Data struct {
	Xlim []int
	Ylim []int
}

type Day struct {
	data Data
}

func loadData(fname string) Data {
	line := utils.RawFileString(fname)
	line = strings.Replace(strings.TrimSpace(line), "target area: ", "", 1)
	parts := strings.Split(strings.TrimSpace(line), ", ")
	xparts := utils.ParseToIntSlice(strings.Replace(parts[0], "x=", "", 1), "..")
	yparts := utils.ParseToIntSlice(strings.Replace(parts[1], "y=", "", 1), "..")
	sort.Ints(xparts)
	sort.Ints(yparts)
	return Data{
		Xlim: xparts,
		Ylim: yparts,
	}
}

func NewDay(fname string) *Day {
	data := loadData(fname)
	return &Day{
		data: data,
	}
}

func step(pos, velocity []int) ([]int, []int) {
	pos[0] += velocity[0]
	pos[1] += velocity[1]
	velocity[0] = utils.Sgn(velocity[0]) * (utils.Abs(velocity[0]) - 1)
	velocity[1] -= 1
	return pos, velocity
}

func (d *Day) HitsTarget(v0 []int) bool {
	pos := []int{0, 0}
	data := d.data
	velocity := v0
	for {
		if pos[0] > data.Xlim[1] || pos[1] < data.Ylim[0] {
			return false
		}
		if d.InTarget(pos) {
			return true
		}
		pos, velocity = step(pos, velocity)
	}
}

func (d *Day) InTarget(pos []int) bool {
	return (pos[0] >= d.data.Xlim[0]) && (pos[0] <= d.data.Xlim[1]) && (pos[1] >= d.data.Ylim[0]) && (pos[1] <= d.data.Ylim[1])
}

func (d *Day) findAllPaths(vxrange, vyrange []int) [][]int {
	var out [][]int
	for vx := vxrange[0]; vx <= vxrange[1]; vx++ {
		for vy := vyrange[0]; vy <= vyrange[1]; vy++ {
			if d.HitsTarget([]int{vx, vy}) {
				out = append(out, []int{vx, vy})
			}
		}
	}
	return out
}

func (d *Day) HighestPoint(v0 []int) int {
	pos := []int{0,0}
	velocity := v0
	max := 0
	for {
		if d.InTarget(pos) {
			return max
		}
		pos, velocity = step(pos, velocity)
		if pos[1] > max {
			max = pos[1]
		}
	}
}

func (d *Day) RunPart1() {
	velocities := d.findAllPaths([]int{1, d.data.Xlim[1]}, []int{1, 1000})
	var heights []int
	for _, v0 := range velocities {
		heights = append(heights, d.HighestPoint(v0))
	}
	fmt.Printf("Part 1: %d\n", utils.Max(heights))
}

func (d *Day) RunPart2() {
	vs := d.findAllPaths([]int{1, d.data.Xlim[1]}, []int{d.data.Ylim[0], 1000})
	fmt.Printf("Part 2: %d\n", len(vs))
}
