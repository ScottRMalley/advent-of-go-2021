package day22

import (
	"advent-calendar/utils"
	"fmt"
	"strings"
)

type Data []Instruction

type Day struct {
	data Data
}
type Instruction struct {
	On     bool
	Xrange []int
	Yrange []int
	Zrange []int
}

type Cuboid struct {
	Xrange []int
	Yrange []int
	Zrange []int
}

func (c Cuboid) Equals(cuboid Cuboid) bool {
	return c.Xrange[0] == cuboid.Xrange[0] && c.Xrange[1] == cuboid.Xrange[1] &&
		c.Yrange[0] == cuboid.Yrange[0] && c.Yrange[1] == cuboid.Yrange[1] &&
		c.Zrange[0] == cuboid.Zrange[0] && c.Zrange[1] == cuboid.Zrange[1]
}

func (c Cuboid) Volume() int {
	return (c.Xrange[1] - c.Xrange[0]) * (c.Yrange[1] - c.Yrange[0]) * (c.Zrange[1] - c.Zrange[0])
}

func loadData(fname string) Data {
	lines := utils.ReadFile(fname)
	var data []Instruction
	for _, line := range lines {
		in := Instruction{}
		var parts []string
		if line[0:3] == "on " {
			in.On = true
			parts = strings.Split(strings.Replace(strings.TrimSpace(line), "on ", "", 1), ",")
		} else {
			in.On = false
			parts = strings.Split(strings.Replace(strings.TrimSpace(line), "off ", "", 1), ",")
		}
		in.Xrange = utils.ParseToIntSlice(strings.Replace(parts[0], "x=", "", 1), "..")
		in.Yrange = utils.ParseToIntSlice(strings.Replace(parts[1], "y=", "", 1), "..")
		in.Zrange = utils.ParseToIntSlice(strings.Replace(parts[2], "z=", "", 1), "..")
		data = append(data, in)
	}
	return data
}

func NewDay(fname string) *Day {
	data := loadData(fname)
	return &Day{
		data: data,
	}
}

func CheckOverlap(c1, c2 Cuboid) (Cuboid, bool) {
	xRanges, xOk := CheckAxisOverlap(c1.Xrange, c2.Xrange)
	yRanges, yOk := CheckAxisOverlap(c1.Yrange, c2.Yrange)
	zRanges, zOk := CheckAxisOverlap(c1.Zrange, c2.Zrange)
	if !(xOk && yOk && zOk) {
		return Cuboid{}, false
	}
	return Cuboid{xRanges, yRanges, zRanges}, true
}

func CheckAxisOverlap(r1, r2 []int) ([]int, bool) {
	startMax := utils.Max([]int{r1[0], r2[0]})
	stopMin := utils.Min([]int{r1[1], r2[1]})

	if startMax > stopMin {
		return nil, false
	}
	return []int{startMax, stopMin}, true
}

func AddCuboid(c Cuboid, cs []Cuboid) []Cuboid {
	var cuboidsWithSplits []Cuboid
	for _, cuboid := range cs {
		overlap, ok := CheckOverlap(cuboid, c)
		if ok {
			splits := SplitCuboid(cuboid, overlap)
			cuboidsWithSplits = append(cuboidsWithSplits, splits...)
		} else {
			cuboidsWithSplits = append(cuboidsWithSplits, cuboid)
		}
	}
	return append(cuboidsWithSplits, c)
}

func RemoveCuboid(c Cuboid, cs []Cuboid) []Cuboid {
	var cuboidsWithSplits []Cuboid
	for _, cuboid := range cs {
		overlap, ok := CheckOverlap(cuboid, c)
		if ok {
			splits := SplitCuboid(cuboid, overlap)
			cuboidsWithSplits = append(cuboidsWithSplits, splits...)
		} else {
			cuboidsWithSplits = append(cuboidsWithSplits, cuboid)
		}
	}
	return cuboidsWithSplits
}

func SplitCuboid(c Cuboid, overlap Cuboid) []Cuboid {
	var split []Cuboid
	xRanges := [][]int{
		{c.Xrange[0], overlap.Xrange[0]},
		{overlap.Xrange[0], overlap.Xrange[1]},
		{overlap.Xrange[1], c.Xrange[1]},
	}

	yRanges := [][]int{
		{c.Yrange[0], overlap.Yrange[0]},
		{overlap.Yrange[0], overlap.Yrange[1]},
		{overlap.Yrange[1], c.Yrange[1]},
	}

	zRanges := [][]int{
		{c.Zrange[0], overlap.Zrange[0]},
		{overlap.Zrange[0], overlap.Zrange[1]},
		{overlap.Zrange[1], c.Zrange[1]},
	}
	for _, xRange := range xRanges {
		for _, yRange := range yRanges {
			for _, zRange := range zRanges {
				c1 := Cuboid{xRange, yRange, zRange}
				if c1.Volume() > 0 && !c1.Equals(overlap) {
					split = append(split, c1)
				}
			}
		}
	}
	return split
}
func (d *Day) RunPart1() {
	startX, stopX, startY, stopY, startZ, stopZ := -50, 50, -50, 50, -50, 50
	rangeCuboid := Cuboid{[]int{startX, stopX}, []int{startY, stopY}, []int{startZ, stopZ}}
	var relevantIns []Instruction
	for _, in := range d.data {
		overlap, ok := CheckOverlap(rangeCuboid, Cuboid{in.Xrange, in.Yrange, in.Zrange})
		if !ok {
			continue
		} else {
			relevantIns = append(relevantIns, Instruction{On: in.On, Xrange: overlap.Xrange, Yrange: overlap.Yrange, Zrange: overlap.Zrange})
		}
	}
	grid := NewSparseMat(stopX-startX, stopY-startY, stopZ-startZ)
	for _, in := range relevantIns {
		for x := in.Xrange[0]; x <= in.Xrange[1]; x++ {
			for y := in.Yrange[0]; y <= in.Yrange[1]; y++ {
				for z := in.Zrange[0]; z <= in.Zrange[1]; z++ {
					if in.On {
						grid.On(x-startX, y-startY, z-startZ)
					} else {
						grid.Off(x-startX, y-startY, z-startZ)
					}
				}
			}
		}
	}
	fmt.Printf("Part 1: %d\n", grid.Len())
}

func InstructionToCuboid(in Instruction) Cuboid {
	return Cuboid{[]int{in.Xrange[0], in.Xrange[1] + 1}, []int{in.Yrange[0], in.Yrange[1] + 1}, []int{in.Zrange[0], in.Zrange[1] + 1}}
}

func (d *Day) RunPart2() {
	/*
		startX, stopX, startY, stopY, startZ, stopZ := -50, 50+1, -50, 50+1, -50, 50+1
		rangeCuboid := Cuboid{[]int{startX, stopX}, []int{startY, stopY}, []int{startZ, stopZ}}
		var relevantIns []Instruction
		for _, in := range d.data {
			overlap, ok := CheckOverlap(rangeCuboid, InstructionToCuboid(in))
			if !ok {
				continue
			} else {
				relevantIns = append(relevantIns, Instruction{On: in.On, Xrange: overlap.Xrange, Yrange: overlap.Yrange, Zrange: overlap.Zrange})
			}
		}*/
	var allCuboids []Cuboid
	for _, in := range d.data {
		if in.On {
			allCuboids = AddCuboid(InstructionToCuboid(in), allCuboids)
		} else {
			allCuboids = RemoveCuboid(InstructionToCuboid(in), allCuboids)
		}
	}
	area := 0
	for _, c := range allCuboids {
		area += c.Volume()
	}
	fmt.Printf("Part 2: %d\n", area)
}
