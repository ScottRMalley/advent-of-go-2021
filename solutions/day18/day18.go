package day18

import (
	"advent-calendar/utils"
	"fmt"
	"strconv"
	"strings"
)

type Data []string

type Day struct {
	data Data
}

func loadData(fname string) Data {
	lines := utils.ReadFile(fname)
	return lines
}

func NewDay(fname string) *Day {
	data := loadData(fname)
	return &Day{
		data: data,
	}
}

func Add(s1 string, s2 string) string {
	if s1 == "" {
		return s2
	}
	if s2 == "" {
		return s1
	}
	return "[" + s1 + "," + s2 + "]"
}

func ExplodeAny(s1 string) (string, bool) {
	result := s1
	exploded := false

	leftBrackets := 0
	for i, c := range result {
		if c == '[' {
			leftBrackets += 1
		}
		if c == ']' {
			leftBrackets -= 1
		}
		if leftBrackets == 5 {
			result = Explode(result, i)
			exploded = true
			break
		}
	}
	return result, exploded
}

func SplitAny(s1 string) (string, bool) {
	result := s1
	split := false

	for i := range result {
		numberString, start, end := getNumberStringAt(result, i)
		if numberString != "" && len(numberString) > 1 {
			result = result[:start+1] + Split(numberString) + result[end:]
			split = true
			break
		}
	}
	return result, split
}

func Reduce(s1 string) string {
	result := s1
	exploded := false
	split := false
	for {
		result, exploded = ExplodeAny(result)
		if exploded {
			continue
		}
		result, split = SplitAny(result)
		if split {
			continue
		}
		if !exploded && !split {
			return result
		}
	}

}

func getNumberStringAt(s1 string, index int) (string, int, int) {
	if strings.ContainsRune("[],", rune(s1[index])) {
		return "", -1, -1
	}
	start := findNext(s1, "[],", index, -1)
	end := findNext(s1, "[],", index, +1)
	return s1[start+1 : end], start, end
}

func findNext(s1, chars string, pos, direction int) int {
	if direction > 0 {
		for k := pos; k < len(s1); k++ {
			if strings.ContainsRune(chars, rune(s1[k])) {
				return k
			}
		}
	} else {
		for k := pos; k >= 0; k-- {
			if strings.ContainsRune(chars, rune(s1[k])) {
				return k
			}
		}
	}
	return -1
}

func Magnitude(reducedString string) int {
	result := reducedString
	replaced := false
	for {
		for i := 0; i < len(result)-5; i++ {
			end := findNext(result, "]", i, +1)
			slice := result[i : end+1]
			if !strings.ContainsRune(slice[1:len(slice)-1], '[') && !strings.ContainsRune(slice[1:len(slice)-1], '[') {
				nums := utils.ParseToIntSlice(slice[1:len(slice)-1], ",")
				result = result[:i] + strconv.Itoa(3*nums[0]+2*nums[1]) + result[end+1:]
				replaced = true
				break
			}
		}
		if !replaced {
			return utils.ParseInt(result)
		} else {
			replaced = false
		}
	}

}

func Explode(s1 string, start int) string {
	stop := start
	for k := start; k < len(s1); k++ {
		if rune(s1[k]) == ']' {
			stop = k
			break
		}
	}

	left := s1[:start]
	right := s1[stop+1:]
	pairString := s1[start+1 : stop]

	pair := utils.ParseToIntSlice(pairString, ",")
	// move left
	for n := start - 1; n >= 0; n-- {
		if !strings.ContainsRune("[],", rune(left[n])) {
			numString := string(left[n])
			pos := n
			// check if next one is an int as well
			if !strings.ContainsRune("[],", rune(left[n-1])) {
				numString = string(left[n-1]) + numString
				pos -= 1
			}
			num := utils.ParseInt(numString) + pair[0]
			left = left[:pos] + strconv.Itoa(num) + left[n+1:]
			break
		}
	}
	// move right
	for n := 0; n < len(right); n++ {
		if !strings.ContainsRune("[],", rune(right[n])) {
			numString := string(right[n])
			pos := n
			// check if next one is an int as well
			if !strings.ContainsRune("[],", rune(right[n+1])) {
				numString += string(right[n+1])
				pos += 1
			}
			num := utils.ParseInt(numString) + pair[1]
			right = right[:n] + strconv.Itoa(num) + right[pos+1:]
			break
		}
	}
	return left + "0" + right
}

func Split(s1 string) string {
	i := utils.ParseInt(s1)
	left := i / 2
	right := i - left
	return "[" + strconv.Itoa(left) + "," + strconv.Itoa(right) + "]"
}

func (d *Day) RunPart1() {
	result := ""
	for _, snailNum := range d.data {
		result = Add(result, snailNum)
		result = Reduce(result)
	}
	fmt.Printf("Part 1: %d\n", Magnitude(result))
}

func (d *Day) RunPart2() {
	max := 0
	for i, s1 := range d.data {
		for j, s2 := range d.data {
			if i == j {
				continue
			}
			m1 := Magnitude(Reduce(Add(s1, s2)))
			if m1 > max {
				max = m1
			}
		}
	}
	fmt.Printf("Part 2: %d\n", max)
}
