package main

import (
	"io/ioutil"
	"regexp"
	"strings"
)

func repeatCheck(s string) bool {
	if len(s) < 4 {
		return false
	}

	for i := 0; i <= len(s)-4; i++ {
		if s[i] != s[i+1] && s[i] == s[i+3] && s[i+1] == s[i+2] {
			return true
		}
	}

	return false
}

func main() {
	input, _ := ioutil.ReadFile("./day07.txt")
	raw_lines := strings.TrimSpace(string(input))
	lines := strings.Split(raw_lines, "\n")

	good_re := regexp.MustCompile("(?:^|])([a-z]+)(?:$|\\[)")
	bad_re := regexp.MustCompile("\\[([a-z]+)\\]")

Line:
	for _, line := range lines {
		println(line)
	}
}
