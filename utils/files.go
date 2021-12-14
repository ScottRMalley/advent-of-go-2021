package utils

import (
	"io/ioutil"
	"strings"
)

func ReadFile(fname string) []string {
	out, err := ioutil.ReadFile(fname)
	Check(err)
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	return lines
}

func RawFileString(fname string) string {
	out, err := ioutil.ReadFile(fname)
	Check(err)
	return strings.TrimSpace(string(out))
}
