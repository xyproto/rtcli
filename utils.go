package main

import (
	"io/ioutil"
	"strings"
)

// Checks if the given filename can be read, panics if not
func Exists(filename string) bool {
	_, err := ioutil.ReadFile(filename)
	return err == nil
}

// Reads the first two lines of a file and returns two strings, panics if not
func ReadTwoLines(filename string) (string, string) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("Could not read " + filename)
	}
	s := string(b)
	lines := strings.Split(s, "\n")
	if len(lines) < 2 {
		panic("Could not read two lines from " + filename)
	}
	return lines[0], lines[1]
}
