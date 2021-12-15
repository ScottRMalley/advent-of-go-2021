package main

import (
	"advent-calendar/solutions/day1"
	"advent-calendar/solutions/day10"
	"advent-calendar/solutions/day11"
	"advent-calendar/solutions/day12"
	"advent-calendar/solutions/day13"
	"advent-calendar/solutions/day14"
	"advent-calendar/solutions/day15"
	"advent-calendar/solutions/day2"
	"advent-calendar/solutions/day3"
	"advent-calendar/solutions/day4"
	"advent-calendar/solutions/day5"
	"advent-calendar/solutions/day6"
	"advent-calendar/solutions/day7"
	"advent-calendar/solutions/day8"
	"advent-calendar/solutions/day9"
	"advent-calendar/utils"
	"fmt"
	"os"
)

type Day interface {
	RunPart1()
	RunPart2()
}

type DayMaker func(fname string) Day

var days = map[int]DayMaker{
	1:  func(s string) Day { return day1.NewDay(s) },
	2:  func(s string) Day { return day2.NewDay(s) },
	3:  func(s string) Day { return day3.NewDay(s) },
	4:  func(s string) Day { return day4.NewDay(s) },
	5:  func(s string) Day { return day5.NewDay(s) },
	6:  func(s string) Day { return day6.NewDay(s) },
	7:  func(s string) Day { return day7.NewDay(s) },
	8:  func(s string) Day { return day8.NewDay(s) },
	9:  func(s string) Day { return day9.NewDay(s) },
	10: func(s string) Day { return day10.NewDay(s) },
	11: func(s string) Day { return day11.NewDay(s) },
	12: func(s string) Day { return day12.NewDay(s) },
	13: func(s string) Day { return day13.NewDay(s) },
	14: func(s string) Day { return day14.NewDay(s) },
	15: func(s string) Day { return day15.NewDay(s) },
}

func runDay(i int, useTestData bool) {
	if useTestData {
		fmt.Printf("Day %d Solutions (Test Data)\n", i)
	} else {
		fmt.Printf("Day %d Solutions\n", i)
	}

	var fname string
	if useTestData {
		fname = fmt.Sprintf("input/day%d_test.txt", i)
	} else {
		fname = fmt.Sprintf("input/day%d.txt", i)
	}
	day := days[i](fname)
	fmt.Print("\t")
	day.RunPart1()
	fmt.Print("\t")
	day.RunPart2()
	fmt.Print("\n")
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		for i := 1; i <= len(days); i++ {
			runDay(i, false)
		}
	} else {
		i := utils.ParseInt(args[0])
		test := false
		if utils.CheckInStringSlice(args, "test") {
			test = true
		}
		runDay(i, test)
	}

	return
}
