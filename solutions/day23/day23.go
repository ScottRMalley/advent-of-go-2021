package day23

import (
	"advent-calendar/utils"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/fatih/color"
	"sort"
	"strconv"
	"strings"
)

type Data struct {
	grid   Grid
	solved Grid
}

type Day struct {
	data Data
}

func GetTypeMap() map[rune]int {
	return map[rune]int{
		'A': -1,
		'B': -2,
		'C': -3,
		'D': -4,
	}
}

func GetEnergyMap() map[int]int {
	return map[int]int{
		-1: 1,
		-2: 10,
		-3: 100,
		-4: 1000,
	}
}

func GetHomePositionMap() map[int]int {
	return map[int]int{
		-1: 2,
		-2: 4,
		-3: 6,
		-4: 8,
	}
}

type Grid struct {
	grid   [][]int
	energy int
}

func (g Grid) Equals(grid Grid) bool {
	for i := range g.grid {
		for j := range g.grid[0] {
			if g.grid[i][j] != grid.grid[i][j] {
				return false
			}
		}
	}
	return true
}

type Grids []Grid

func (g Grids) Len() int {
	return len(g)
}

func (g Grids) Less(i, j int) bool {
	return g[i].energy < g[j].energy
}

func (g Grids) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}

type Hash [32]byte

func ToHash(grid [][]int) Hash {
	s := ""
	for i := range grid {
		for j := range grid[0] {
			s += strconv.Itoa(grid[i][j])
		}
	}
	return sha256.Sum256([]byte(s))
}

func loadData(fname string) Data {
	lines := utils.ReadFile(fname)
	grid := utils.Zeros2D(11, 5)
	solved := utils.Zeros2D(11, 5)
	for i := range grid {
		grid[i][0] = 1
		solved[i][0] = 1
	}
	dl := strings.Split(strings.ReplaceAll(strings.TrimSpace(lines[2]), "#", ""), "")
	for i, c := range dl {
		grid[2+2*i][0] = 2
		solved[2+2*i][0] = 2
		solved[2+2*i][1] = -(i + 1)
		grid[2+2*i][1] = GetTypeMap()[[]rune(c)[0]]
	}
	dt := strings.Split(strings.ReplaceAll(strings.TrimSpace(lines[3]), "#", ""), "")
	for i, c := range dt {
		grid[2+2*i][2] = GetTypeMap()[[]rune(c)[0]]
		solved[2+2*i][2] = -(i + 1)
	}
	return Data{
		grid:   Grid{grid, 0},
		solved: Grid{solved, 0},
	}
}

func NewDay(fname string) *Day {
	data := loadData(fname)
	return &Day{
		data: data,
	}
}

func FindAllMoves(startPos, currentPos []int, moves int, grid Grid, positions [][]int) [][]int {
	X, Y := len(grid.grid), len(grid.grid[0])
	directions := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	startType := grid.grid[startPos[0]][startPos[1]]
	// if already home, dont go anywhere
	if currentPos[0] == GetHomePositionMap()[startType] {
		if currentPos[1] == 2 {
			// deep home, don't leave
			return nil
		} else if currentPos[1] == 1 && grid.grid[currentPos[0]][2] == startType {
			// home with roommate
			return nil
		}
	}
	var newPositions [][]int
	for _, v := range directions {
		newPos := []int{currentPos[0] + v[0], currentPos[1] + v[1]}

		if PositionInPositions(newPos, positions) {
			// wont go somewhere it's already been
			continue
		} else if newPos[0] >= X || newPos[0] < 0 {
			// wont go out of bounnds
			continue
		} else if newPos[1] >= Y || newPos[1] < 0 {
			// wont go out of bounds
			continue
		} else if grid.grid[newPos[0]][newPos[1]] <= 0 {
			// wont go to a wall or occupued space
			continue
		} else if newPos[1] > 0 && !IsHomePosition(newPos[0], startType) && newPos[0] != startPos[0] {
			// wont go into a room that isn't its own
			continue
		} else if newPos[1] > 0 && IsHomePosition(newPos[0], startType) && (grid.grid[newPos[0]][2] != 1 && grid.grid[newPos[0]][2] != startType) {
			// wont go into its home if it has someone already there who doesn't live with them
			continue
		}
		// now we know we can move there (but need to check if we can stop
		if grid.grid[newPos[0]][newPos[1]] == 2 {
			// we can move but can't stop
			newPositions = append(newPositions, FindAllMoves(startPos, newPos, moves+1, grid, append(positions, newPos))...)
		} else if startPos[1] == 0 && newPos[1] < 1 {
			// if we are in the hall, we can't stay in the hall
			newPositions = append(newPositions, FindAllMoves(startPos, newPos, moves+1, grid, append(positions, newPos))...)
		} else {
			// we can stop there or keep going
			newPositions = append(newPositions, []int{newPos[0], newPos[1], moves + 1})
			newPositions = append(newPositions, FindAllMoves(startPos, newPos, moves+1, grid, append(positions, newPos))...)
		}
	}
	// if there is a position that goes home
	var startPositions [][]int
	for _, np := range newPositions {
		if IsHomePosition(np[0], startType) {
			startPositions = append(startPositions, np)
		}
	}
	if len(startPositions) > 0 {
		return startPositions
	} else {
		return newPositions
	}
}

func PositionInPositions(pos []int, positions [][]int) bool {
	for _, p := range positions {
		if p[0] == pos[0] && p[1] == pos[1] {
			return true
		}
	}
	return false
}

func IsHomePosition(xPos, startType int) bool {
	return xPos == utils.Abs(startType)*2
}

func FindAllAmphipoda(grid [][]int) [][]int {
	var positions [][]int
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] < 0 {
				positions = append(positions, []int{i, j, grid[i][j]})
			}
		}
	}
	return positions
}

func SwapPositions(grid [][]int, oldPos []int, newPos []int) [][]int {
	newGrid := utils.Zeros2D(len(grid), len(grid[1]))
	val := grid[oldPos[0]][oldPos[1]]
	for i := range grid {
		for j := range grid[0] {
			if i == oldPos[0] && j == oldPos[1] {
				newGrid[i][j] = 1
			} else if i == newPos[0] && j == newPos[1] {
				newGrid[i][j] = val
			} else {
				replace := grid[i][j]
				newGrid[i][j] = replace
			}
		}
	}
	return newGrid
}

func FindAllNextGrids(grid Grid) Grids {
	var grids []Grid
	positions := FindAllAmphipoda(grid.grid)
	for _, pos := range positions {
		startPos, currentPos := []int{pos[0], pos[1]}, []int{pos[0], pos[1]}
		newPositions := FindAllMoves(startPos, currentPos, 0, grid, [][]int{currentPos})
		for _, newPos := range newPositions {
			grids = append(grids, Grid{SwapPositions(grid.grid, pos, newPos), grid.energy + GetEnergyMap()[pos[2]]*newPos[2]})
		}
	}
	return grids
}

func shortest(grid Grid, solved Grid) int {
	weights := make(map[Hash]int)
	toVisit := []Grid{grid}
	weights[ToHash(grid.grid)] = 0

	for {
		if len(toVisit) == 0 {
			break
		}
		currentGrid := toVisit[0]
		toVisit = toVisit[1:]

		options := FindAllNextGrids(currentGrid)
		for _, newGrid := range options {
			sig := ToHash(newGrid.grid)
			weight, ok := weights[sig]
			if !ok || weight > newGrid.energy {
				weights[sig] = newGrid.energy
				toVisit = append(toVisit, newGrid)
			}
		}
		sort.Sort(Grids(toVisit))
	}
	solvedSig := ToHash(solved.grid)
	if w, ok := weights[solvedSig]; !ok {
		panic(errors.New("could not find solution"))
	} else {
		return w
	}

}

func FindMinimumSolution(grid Grid, solved Grid) int {
	solvedHash := ToHash(solved.grid)
	cache := make(map[Hash]int)
	options := []Grid{grid}

	for {
		if len(options) < 1 {
			break
		}
		currentGrid := options[0]
		options = options[1:]
		for _, next := range FindAllNextGrids(currentGrid) {
			sig := ToHash(next.grid)
			if val, ok := cache[sig]; ok {
				if val < next.energy {
					continue
				} else {
					cache[sig] = next.energy
					options = append(options, next)
				}
			} else {
				cache[sig] = next.energy
				options = append(options, next)
			}
		}
		sort.Sort(Grids(options))
	}
	if val, ok := cache[solvedHash]; !ok {
		panic(errors.New("no solution found"))
	} else {
		return val
	}
}

func PrintGrid(g Grid) {
	cyan := color.New(color.FgHiGreen)
	boldCyan := cyan.Add(color.Bold)
	bg := color.New(color.FgBlack)

	fmt.Println("")
	for i := range g.grid[0] {
		for k := range g.grid {
			c := g.grid[k][i]
			if c == 1 || c == 2 {
				_, _ = cyan.Print(".")
			} else if c == 0 {
				_, _ = bg.Print("#")
			} else {
				for r, val := range GetTypeMap() {
					if val == c {
						_, _ = boldCyan.Print(string(r))
					}
				}

			}
		}
		fmt.Print("\n")
	}
	fmt.Println("")
}

func (d *Day) RunPart1() {
	//min := shortest(d.data.grid, d.data.solved)
	PrintGrid(d.data.grid)
	fmt.Printf("Part 1: %d\n", 0)
}

func (d *Day) RunPart2() {
	fmt.Printf("Part 2: %s\n", "not-implemented")
}
