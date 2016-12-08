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
		good_bits := good_re.FindAllStringSubmatch(line, -1)
		bad_bits := bad_re.FindAllStringSubmatch(line, -1)

		for _, match := range bad_bits {
			if repeatCheck(match[1]) {
				continue Line
			}
		}

		for _, match := range good_bits {
			if repeatCheck(match[1]) {
				println(line)
				continue Line
			}
		}
	}
}
