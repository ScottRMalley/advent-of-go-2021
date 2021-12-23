package day21

import (
	"advent-calendar/utils"
	"fmt"
	"strings"
)

type Data struct {
	p1 int
	p2 int
}

type Day struct {
	data      Data
	positions []int
	scores    []int
	dice      int
	nRolls    int
}

func loadData(fname string) Data {
	lines := utils.ReadFile(fname)
	p1 := utils.ParseInt(strings.Replace(lines[0], "Player 1 starting position: ", "", 1))
	p2 := utils.ParseInt(strings.Replace(lines[1], "Player 2 starting position: ", "", 1))
	return Data{p1, p2}
}

func NewDay(fname string) *Day {
	data := loadData(fname)
	return &Day{
		data:      data,
		positions: []int{data.p1, data.p2},
		scores:    make([]int, 2),
		dice:      0,
		nRolls:    0,
	}
}

func (d *Day) roll() int {
	res := d.dice + 1
	if res == 101 {
		res = 1
	}
	d.dice = res
	d.nRolls += 1
	return res
}

func (d *Day) move(player, dis int) {
	pos := d.positions[player]
	newPos := pos + dis
	for newPos > 10 {
		newPos -= 10
	}
	d.positions[player] = newPos
	d.scores[player] += newPos
}

func (d *Day) winner() (int, bool) {
	for player, score := range d.scores {
		if score >= 1000 {
			return player, true
		}
	}
	return -1, false
}

func (d *Day) play() {
	for {
		for player := range d.positions {
			toMove := 0
			for i := 0; i < 3; i++ {
				roll := d.roll()
				toMove += roll
			}
			d.move(player, toMove)
			if _, ok := d.winner(); ok {
				return
			}
		}
	}
}

var cache map[string][]int

func withCaching(p1, p2, score1, score2, player_turn, r1, r2 int) (int, int) {
	if cache == nil {
		cache = make(map[string][]int)
	}
	sig := fmt.Sprintf("%d,%d,%d,%d,%d", p1, p2, score1, score2, player_turn)
	cache[sig] = []int{r1, r2}
	return r1, r2
}

func fromCache(p1, p2, score1, score2, player_turn int) (int, int, bool) {
	if cache == nil {
		cache = make(map[string][]int)
	}
	sig := fmt.Sprintf("%d,%d,%d,%d,%d", p1, p2, score1, score2, player_turn)
	if val, ok := cache[sig]; ok {
		return val[0], val[1], true
	} else {
		return 0, 0, false
	}
}

func playQuantum(p1, p2, score1, score2, player_turn int) (int, int) {
	if score1 >= 21 {
		return 1, 0
	} else if score2 >= 21 {
		return 0, 1
	}
	if r1, r2, ok := fromCache(p1, p2, score1, score2, player_turn); ok {
		return r1, r2
	}

	p1wins := 0
	p2wins := 0
	for _, roll1 := range []int{1, 2, 3} {
		for _, roll2 := range []int{1, 2, 3} {
			for _, roll3 := range []int{1, 2, 3} {
				if player_turn == 0 {
					newPos := getNewPos(p1, roll1+roll2+roll3)
					w1, w2 := playQuantum(newPos, p2, score1+newPos, score2, 1)
					p1wins += w1
					p2wins += w2
				} else {
					newPos := getNewPos(p2, roll1+roll2+roll3)
					w1, w2 := playQuantum(p1, newPos, score1, score2+newPos, 0)
					p1wins += w1
					p2wins += w2
				}
			}
		}
	}
	return withCaching(p1, p2, score1, score2, player_turn, p1wins, p2wins)
}

func getNewPos(pos, dis int) int {
	newPos := pos + dis
	for newPos > 10 {
		newPos -= 10
	}
	return newPos
}

func (d *Day) RunPart1() {
	d.play()
	winner, _ := d.winner()
	losingScore := d.scores[(winner+1)%2]
	fmt.Printf("Part 1: %d\n", d.nRolls*losingScore)
}

func (d *Day) RunPart2() {
	w1, w2 := playQuantum(d.data.p1, d.data.p2, 0, 0, 0)
	max := w1
	if w2 > w1 {
		max = w2
	}
	fmt.Printf("Part 2: %d\n", max)
}
