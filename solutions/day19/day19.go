package day19

import (
	"advent-calendar/utils"
	"fmt"
	"strings"
)

type Data []Scanner

type Scanner struct {
	Beacons []Loc
}

type Loc struct {
	X, Y, Z int
}

type Rotation struct {
	Yaw, Pitch, Roll float64
}

func (r Rotation) Reverse() Rotation {
	return Rotation{-r.Yaw, -r.Pitch, -r.Roll}
}

func (l Loc) Equals(loc Loc) bool {
	return l.X == loc.X && l.Y == loc.Y && l.Z == loc.Z
}

func (l Loc) Add(loc Loc) Loc {
	return Loc{l.X + loc.X, l.Y + loc.Y, l.Z + loc.Z}
}

func (l Loc) Sub(loc Loc) Loc {
	return Loc{l.X - loc.X, l.Y - loc.Y, l.Z - loc.Z}
}

func (l Loc) Abs() Loc {
	return Loc{utils.Abs(l.X), utils.Abs(l.Y), utils.Abs(l.Z)}
}

func (l Loc) Sum() int {
	return l.X + l.Y + l.Z
}

func (l Loc) Rot(r Rotation) Loc {
	nx, ny, nz := utils.Rotate(float64(l.X), float64(l.Y), float64(l.Z), r.Yaw, r.Pitch, r.Roll)
	return Loc{utils.Round(nx), utils.Round(ny), utils.Round(nz)}
}

type Day struct {
	data        Data
	scannerLocs []Loc
}

func loadData(fname string) Data {
	rawString := utils.RawFileString(fname)
	scannerParts := strings.Split(strings.TrimSpace(rawString), "\n\n")
	var scanners []Scanner
	for _, scannerData := range scannerParts {
		lines := strings.Split(strings.TrimSpace(scannerData), "\n")
		var beacons []Loc
		for i, line := range lines {
			if i == 0 {
				continue
			}
			xyz := utils.ParseToIntSlice(line, ",")
			beacons = append(beacons, Loc{xyz[0], xyz[1], xyz[2]})
		}
		scanners = append(scanners, Scanner{Beacons: beacons})
	}
	return scanners
}

func NewDay(fname string) *Day {
	data := loadData(fname)
	return &Day{
		data: data,
	}
}

func (d *Day) RunPart1() {
	scanner, scannerLocs := MatchAllScanners(d.data)
	d.scannerLocs = scannerLocs
	fmt.Printf("Part 1: %d\n", len(scanner.Beacons))
}

func (d *Day) RunPart2() {
	max := 0
	for _, l1 := range d.scannerLocs {
		for _, l2 := range d.scannerLocs {
			dis := l1.Sub(l2).Abs().Sum()
			if dis > max {
				max = dis
			}
		}
	}
	fmt.Printf("Part 2: %d\n", max)
}

func MatchAllScanners(scanners []Scanner) (Scanner, []Loc) {
	rootScanner := scanners[0]
	rootBeacons := rootScanner.Beacons
	scannerLocs := []Loc{{0, 0, 0}}

	done := make([]bool, len(scanners))
	done[0] = true
	toCheck := 0
	for {
		targetScanner := scanners[toCheck]
		if AllDone(done) {
			return Scanner{Beacons: rootBeacons}, scannerLocs
		}
		if done[toCheck] {
			toCheck = (toCheck + 1) % len(scanners)
			continue
		}
		for _, rootBeacon := range rootBeacons {
			for _, targetBeacon := range targetScanner.Beacons {
				locs, orientations := GetPossibleScannerLocs(rootBeacon, targetBeacon)
				for k := range locs {
					if shifted, ok := VerifyScannerLoc(rootBeacons, targetScanner, locs[k], orientations[k]); ok {
						rootBeacons = AddBeaconsNoRepeats(rootBeacons, shifted)
						scannerLocs = append(scannerLocs, locs[k])
						done[toCheck] = true
						goto finally
					}
				}
			}
		}
	finally:
		toCheck = (toCheck + 1) % len(scanners)
	}
}

func AllDone(done []bool) bool {
	for _, d := range done {
		if !d {
			return false
		}
	}
	return true
}

func GetPossibleScannerLocs(
	rootBeacon,
	targetBeacon Loc,
) ([]Loc, []Rotation) {
	rotations, orientations := GetAllRotations(targetBeacon)
	var scannerLocs []Loc
	for _, rotation := range rotations {
		scannerLocs = append(scannerLocs, rootBeacon.Sub(rotation))
	}
	return scannerLocs, orientations
}

func VerifyScannerLoc(rootBeacons []Loc, scanner Scanner, scannerLoc Loc, scannerOrientation Rotation) ([]Loc, bool) {
	inRange := GetBeaconsInRange(scannerLoc, rootBeacons)
	rotatedScannerBeacons := RotateBeacons(scanner.Beacons, scannerOrientation)
	shiftedScannerBeacons := ShiftBeacons(rotatedScannerBeacons, scannerLoc)
	matches := 0
	for _, beacon := range inRange {
		if LocInLocs(beacon, shiftedScannerBeacons) {
			matches += 1
		}
	}
	return shiftedScannerBeacons, matches >= 12
}

func RotateBeacons(beacons []Loc, rotation Rotation) []Loc {
	var out []Loc
	for _, beacon := range beacons {
		out = append(out, beacon.Rot(rotation))
	}
	return out
}

func ShiftBeacons(beacons []Loc, scannerLoc Loc) []Loc {
	var out []Loc
	for _, beacon := range beacons {
		out = append(out, scannerLoc.Add(beacon))
	}
	return out
}

func GetBeaconsInRange(scannerLoc Loc, beacons []Loc) []Loc {
	var inRange []Loc
	for _, beacon := range beacons {
		dLoc := beacon.Sub(scannerLoc)
		if utils.Abs(dLoc.X) <= 1000 && utils.Abs(dLoc.Y) <= 1000 && utils.Abs(dLoc.Z) <= 1000 {
			inRange = append(inRange, beacon)
		}
	}
	return inRange
}

func GetAllRotations(loc Loc) ([]Loc, []Rotation) {
	var distances []Loc
	var rotations []Rotation
	for _, phi := range []float64{0, 90, 180, 270} {
		for _, psi := range []float64{0, 90, 180, 270} {
			for _, theta := range []float64{0, 90, 180, 270} {
				yaw, pitch, roll := utils.DegreeToRad(phi), utils.DegreeToRad(psi), utils.DegreeToRad(theta)
				rotation := Rotation{yaw, pitch, roll}
				distances = append(distances, loc.Rot(rotation))
				rotations = append(rotations, rotation)
			}
		}
	}
	return distances, rotations
}

func LocInLocs(loc Loc, locs []Loc) bool {
	for _, locIn := range locs {
		if locIn.Equals(loc) {
			return true
		}
	}
	return false
}

func AddBeaconsNoRepeats(b1, b2 []Loc) []Loc {
	var out []Loc
	for _, b := range b1 {
		if !LocInLocs(b, out) {
			out = append(out, b)
		}
	}
	for _, bn := range b2 {
		if !LocInLocs(bn, out) {
			out = append(out, bn)
		}
	}
	return out
}
