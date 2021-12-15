package day8

import (
	"advent-calendar/utils"
	"strings"
)

const LETTERS = "abcdefg"

type Digit struct {
	Input           string
	PossibleMatches []string
}

func NewDigit(input string) *Digit {
	s := &Digit{
		Input:           input,
		PossibleMatches: getNumberMap(),
	}
	s.pruneMatches()
	return s
}

func (s *Digit) Reset() {
	s.PossibleMatches = getNumberMap()
	s.pruneMatches()
}

func (s *Digit) pruneMatches() {
	var newMatches []string
	for _, val := range s.PossibleMatches {
		if len(s.Input) == len(val) {
			newMatches = append(newMatches, val)
		}
	}
	s.PossibleMatches = newMatches
}

func (s *Digit) ApplyAssignment(original, assigned rune) bool {
	var newMatches []string
	for _, match := range s.PossibleMatches {
		if strings.ContainsRune(match, original) && strings.ContainsRune(s.Input, assigned) {
			newMatches = append(newMatches, match)
		} else if !strings.ContainsRune(match, original) && !strings.ContainsRune(s.Input, assigned) {
			newMatches = append(newMatches, match)
		}
	}
	s.PossibleMatches = newMatches
	return len(newMatches) != 0
}

func (s *Digit) ApplyAssignments(assignment string) (int, bool) {
	for i, r := range assignment {
		if ok := s.ApplyAssignment(rune(LETTERS[i]), r); !ok {
			return 0, false
		}
	}
	return s.GetRealDigit()
}

func (s *Digit) GetPossibleAssignments() map[rune]string {
	out := make(map[rune]string)
	if len(s.PossibleMatches) > 1 {
		return out
	}
	match := s.PossibleMatches[0]
	for _, original := range match {
		out[original] = s.Input
	}

	return out
}

func (s *Digit) GetRealDigit() (int, bool) {
	if len(s.PossibleMatches) > 1 {
		return 0, false
	}
	match := s.PossibleMatches[0]
	for k, numberString := range getNumberMap() {
		if utils.StringsMatchNoOrder(numberString, match) {
			return k, true
		}
	}
	return 0, false
}
