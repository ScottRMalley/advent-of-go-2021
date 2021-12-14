package utils

import (
	"strings"
	"unicode"
)

func CheckInStringSlice(a []string, val string) bool {
	for _, s := range a {
		if val == s {
			return true
		}
	}
	return false
}

func CheckRuneInString(s string, r rune) bool {
	for _, rr := range s {
		if rr == r {
			return true
		}
	}
	return false
}

func CountUniqueRunes(s string) map[rune]int {
	count := make(map[rune]int)
	for _, ss := range s {
		if val, ok := count[ss]; !ok {
			count[ss] = 1
		} else {
			count[ss] = val + 1
		}
	}
	return count
}

func StringsMatchNoOrder(s1, s2 string) bool {
	for _, r1 := range s1 {
		if !CheckRuneInString(s2, r1) {
			return false
		}
	}
	for _, r2 := range s2 {
		if !CheckRuneInString(s1, r2) {
			return false
		}
	}
	return true
}

func GetOverlapRuneSlices(r1 []rune, r2 []rune) []rune {
	out := ""
	for _, r := range r1 {
		if strings.ContainsRune(string(r2), r) && !strings.ContainsRune(out, r) {
			out += string(r)
		}
	}
	return []rune(out)
}

func GetOverlapString(s1, s2 string) string {
	return string(GetOverlapRuneSlices([]rune(s1), []rune(s2)))
}

func GetSetComplementRuneSlices(r1, r2 []rune) []rune {
	out := ""
	for _, r := range r1 {
		if !strings.ContainsRune(string(r2), r) && !strings.ContainsRune(out, r) {
			out += string(r)
		}
	}

	for _, r := range r2 {
		if !strings.ContainsRune(string(r1), r) && !strings.ContainsRune(out, r) {
			out += string(r)
		}
	}
	return []rune(out)
}

func GetSetComplementString(s1, s2 string) string {
	return string(GetSetComplementRuneSlices([]rune(s1), []rune(s2)))
}

func GetUnionRuneSlices(r1, r2 []rune) []rune {
	out := ""
	for _, r := range r1 {
		if !strings.ContainsRune(out, r) {
			out += string(r)
		}
	}
	for _, r := range r2 {
		if !strings.ContainsRune(out, r) {
			out += string(r)
		}
	}
	return []rune(out)
}

func GetUnionString(s1, s2 string) string {
	return string(GetUnionRuneSlices([]rune(s1), []rune(s2)))
}

func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func ValInStringSlice(s []string, val string) bool {
	for _, ss := range s {
		if val == ss {
			return true
		}
	}
	return false
}

func CountOccurrenceStringSlice(s []string, val string) int {
	occurences := 0
	for _, ss := range s {
		if ss == val {
			occurences += 1
		}
	}
	return occurences
}
