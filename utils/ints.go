package utils

import (
	"sort"
	"strconv"
	"strings"
)

func ParseToIntSlice(line string, separator string) []int {
	a := strings.Split(strings.TrimSpace(line), separator)
	var out []int
	for _, val := range a {
		out = append(out, ParseInt(val))
	}
	return out
}

func ToString(a []int) string {
	out := ""
	for _, val := range a {
		out += strconv.Itoa(val)
	}
	return out
}

func ParseInt(a string) int {
	val, err := strconv.Atoi(a)
	Check(err)
	return val
}

func RemoveFromArray(a []int, valToRemove int) []int {
	var out []int
	for _, val := range a {
		if val != valToRemove {
			out = append(out, val)
		}
	}
	return out
}

func RemoveFirstFromArray(a []int, valToRemove int) []int {
	var out []int
	found := false
	for _, val := range a {
		if !found && val == valToRemove {
			found = true
		} else {
			out = append(out, val)
		}
	}
	return out
}

func CheckInIntArray(a []int, val int) bool {
	for _, s := range a {
		if val == s {
			return true
		}
	}
	return false
}

func SumInts(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func GetColumn(amat [][]int, col int) []int {
	if col > len(amat) {
		return nil
	}
	var column []int
	for _, row := range amat {
		column = append(column, row[col])
	}
	return column
}

func IntRange(start, stop int) []int {
	var out []int
	for i := start; i < stop; i++ {
		out = append(out, i)
	}
	return out
}

func GetIntLocs(a []int, locs []int) []int {
	var out []int
	for _, loc := range locs {
		out = append(out, a[loc])
	}
	return out
}

// GetIntRange returns a sequence of integers between
// start and stop where stop is inclusive
func GetIntRange(start, stop int) []int {
	var out []int
	diff := Sgn(stop - start)
	k := start
	for {
		out = append(out, k)
		if k == stop {
			return out
		} else {
			k += diff
		}
	}
}

func Zeros2D(X, Y int) [][]int {
	grid := make([][]int, X)
	for i := 0; i < X; i++ {
		grid[i] = make([]int, Y)
	}
	return grid
}

func Sgn(a int) int {
	switch {
	case a < 0:
		return -1
	case a > 0:
		return +1
	}
	return 0
}

func Flatten(grid [][]int) []int {
	X := len(grid)
	Y := len(grid[0])
	out := make([]int, X*Y)
	for x := 0; x < X; x++ {
		for y := 0; y < Y; y++ {
			out = append(out, grid[x][y])
		}
	}
	return out
}

func CountWhere(a []int, condition func(int) bool) int {
	count := 0
	for _, val := range a {
		if condition(val) {
			count += 1
		}
	}
	return count
}

func Min(a []int) int {
	min := a[0]
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min
}

func Max(a []int) int {
	max := a[0]
	for _, v := range a {
		if v > max {
			max = v
		}
	}
	return max
}

func MedianInt(a []int) int {
	sort.Ints(a)
	n := len(a) / 2
	return a[n]
}

func AddConstant(a []int, c int) []int {
	out := make([]int, len(a))
	for i, val := range a {
		out[i] = val + c
	}
	return out
}

func AllValuesLessThan(a [][]int, v int) bool {
	for i := range a {
		for j := range a[0] {
			if a[i][j] >= v {
				return false
			}
		}
	}
	return true
}
