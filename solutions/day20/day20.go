package day20

import (
	"advent-calendar/utils"
	"fmt"
	"github.com/fatih/color"
	"strconv"
	"strings"
)

type Data struct {
	enhancementMap []int
	inputImage     [][]int
}

type Day struct {
	data Data
}

func loadData(fname string) Data {
	lines := utils.RawFileString(fname)
	parts := strings.Split(strings.TrimSpace(lines), "\n\n")
	enhancementString := parts[0]
	enhancementMap := make([]int, len(enhancementString))
	for i, c := range enhancementString {
		if c == '#' {
			enhancementMap[i] = 1
		} else {
			enhancementMap[i] = 0
		}
	}
	imgLines := strings.Split(strings.TrimSpace(parts[1]), "\n")
	inputImage := make([][]int, len(imgLines))
	for k, line := range imgLines {
		imy := make([]int, len(line))
		for j, c := range line {
			if c == '#' {
				imy[j] = 1
			} else {
				imy[j] = 0
			}
		}
		inputImage[k] = imy
	}
	return Data{
		enhancementMap: enhancementMap,
		inputImage:     inputImage,
	}
}

func NewDay(fname string) *Day {
	data := loadData(fname)
	return &Day{
		data: data,
	}
}

func (d *Day) RunPart1() {
	img := d.reconstruct(1)
	nonzero := utils.SumInts(utils.Flatten(img))
	/*fmt.Println("")
	printImage(img)*/
	fmt.Printf("Part 1: %d\n", nonzero)
}

func (d *Day) RunPart2() {
	img := d.reconstruct(49)
	nonzero := utils.SumInts(utils.Flatten(img))
	fmt.Printf("Part 2: %d\n", nonzero)
}

func printImage(img [][]int) {
	cyan := color.New(color.FgHiGreen)
	boldCyan := cyan.Add(color.Bold)
	bg := color.New(color.FgBlack)
	for i := range img {
		for j := range img[0] {
			if img[i][j] > 0 {
				_, _ = boldCyan.Print("# ")
			} else {
				_, _ = bg.Print(". ")
			}
		}
		fmt.Println("")
	}
}

func (d *Day) reconstruct(iterations int) [][]int {
	s1, s2 := len(d.data.inputImage), len(d.data.inputImage[0])
	s1pad := 2 * iterations
	s2pad := 2 * iterations
	out := utils.Zeros2D(s1+2*s1pad, s2+2*s1pad)
	for i := 0; i < s1+2*s1pad; i++ {
		for j := 0; j < s2+2*s2pad; j++ {
			out[i][j] = d.getPixelAt(i-s1pad, j-s1pad, iterations)
		}
	}
	return out
}

var cache map[string]int

func withCaching(i, j, iteration, result int) int {
	if cache == nil {
		cache = make(map[string]int)
	}
	sig := fmt.Sprintf("%d,%d,%d", i, j, iteration)
	cache[sig] = result
	return result
}

func getFromCache(i, j, iteration int) (int, bool) {
	if cache == nil {
		cache = make(map[string]int)
	}
	sig := fmt.Sprintf("%d,%d,%d", i, j, iteration)
	if val, ok := cache[sig]; ok {
		return val, true
	} else {
		return 0, false
	}
}

func (d *Day) getPixelAt(i, j, iteration int) int {
	if val, ok := getFromCache(i, j, iteration); ok {
		return val
	}
	Y := len(d.data.inputImage)
	X := len(d.data.inputImage[0])
	binary := ""
	for _, k := range []int{-1, 0, 1} {
		for _, l := range []int{-1, 0, 1} {
			x := i + k
			y := j + l
			if iteration > 0 {
				binary += strconv.Itoa(d.getPixelAt(x, y, iteration-1))
			} else {
				if (x >= X || x < 0) || (y >= Y || y < 0) {
					binary += "0"
				} else {
					binary += strconv.Itoa(d.data.inputImage[x][y])
				}
			}
		}
	}
	loc := utils.ParseBinaryInt(binary)
	return withCaching(i, j, iteration, d.data.enhancementMap[loc])
}
