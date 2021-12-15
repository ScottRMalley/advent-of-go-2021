package day13

import (
	"advent-calendar/utils"
	"fmt"
	"github.com/fatih/color"
	"strings"
)

type Data struct {
	Points [][]int
	Folds  [][]int
}

type Day struct {
	data Data
}

func loadData(fname string) Data {
	fileString := utils.RawFileString(fname)
	parts := strings.Split(fileString, "\n\n")

	pointLines := strings.Split(strings.TrimSpace(parts[0]), "\n")
	var points [][]int
	for _, line := range pointLines {
		xy := utils.ParseToIntSlice(line, ",")
		points = append(points, xy)
	}

	foldLines := strings.Split(strings.TrimSpace(parts[1]), "\n")
	var folds [][]int
	for _, line := range foldLines {
		foldParts := strings.Split(strings.Replace(line, "fold along ", "", 1), "=")
		var key int
		if foldParts[0] == "x" {
			key = 0
		} else {
			key = 1
		}
		val := utils.ParseInt(foldParts[1])
		folds = append(folds, []int{key, val})
	}
	return Data{
		Points: points,
		Folds:  folds,
	}
}

func NewDay(fname string) *Day {
	data := loadData(fname)
	return &Day{
		data: data,
	}
}

func (d *Day) getMaxPoint() (int, int) {
	maxX := 0
	maxY := 0
	for _, point := range d.data.Points {
		if point[0] > maxX {
			maxX = point[0]
		}
		if point[1] > maxY {
			maxY = point[1]
		}
	}
	return maxX, maxY
}

func (d *Day) getFirstFold() []int {
	var fold []int
	fold, d.data.Folds = d.data.Folds[0], d.data.Folds[1:]
	return fold
}

func (d *Day) fold() {
	maxX, maxY := d.getMaxPoint()
	nextFold := d.getFirstFold()

	foldLine := nextFold[1]
	var newPoints [][]int

	if nextFold[0] == 0 {
		for _, point := range d.data.Points {
			if point[0] > foldLine {
				newPoints = append(newPoints, []int{
					maxX - point[0],
					point[1],
				})
			} else {
				newPoints = append(newPoints, point)
			}
		}
	} else if nextFold[0] == 1 {
		// fold along y
		firstHalfSize := foldLine
		secondHalfSize := maxY - foldLine

		if firstHalfSize != secondHalfSize {
			maxY += firstHalfSize - secondHalfSize
		}

		for _, point := range d.data.Points {
			if point[1] > foldLine {
				newPoints = append(newPoints, []int{
					point[0],
					maxY - point[1],
				})
			} else {
				newPoints = append(newPoints, point)
			}
		}
	}
	d.data.Points = removeDuplicates(newPoints)
}

func removeDuplicates(points [][]int) [][]int {
	var out [][]int
	for _, point := range points {
		if !pointInSlice(out, point) {
			out = append(out, point)
		}
	}
	return out
}

func pointInSlice(s [][]int, point []int) bool {
	for _, p := range s {
		if p[0] == point[0] && p[1] == point[1] {
			return true
		}
	}
	return false
}

func countInSlice(s [][]int, point []int) int {
	count := 0
	for _, p := range s {
		if p[0] == point[0] && p[1] == point[1] {
			count += 1
		}
	}
	return count
}

func (d *Day) print() {
	maxX, maxY := d.getMaxPoint()
	cyan := color.New(color.FgHiGreen)
	boldCyan := cyan.Add(color.Bold)

	bg := color.New(color.FgBlack)

	for y := -2; y < maxY+3; y++ {
		for x := -2; x < maxX+3; x++ {
			if pointInSlice(d.data.Points, []int{x, y}) {
				_, _ = boldCyan.Print("# ")
			} else {
				_, _ = bg.Print(". ")
			}
		}
		fmt.Print("\n")
	}
}

func (d *Day) RunPart1() {
	d.fold()
	fmt.Printf("Part 1: %d\n", len(d.data.Points))
}

func (d *Day) RunPart2() {
	nFolds := len(d.data.Folds)
	for i := 0; i < nFolds; i++ {
		d.fold()
	}
	fmt.Println("Part 2:")
	d.print()
}

func main() {
	day := NewDay("input/day13.txt")
	day.RunPart1()
	day.RunPart2()
}
