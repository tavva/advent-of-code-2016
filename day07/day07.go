package main

import (
	"fmt"
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

func abaCheck(s string) (bool, []string) {
	var matches []string

	if len(s) < 3 {
		return false, nil
	}

	for i := 0; i <= len(s)-3; i++ {
		if s[i] == s[i+2] && s[i] != s[i+1] {
			matches = append(matches, s[i:i+3])
		}
	}

	if len(matches) > 0 {
		return true, matches
	}

	return false, nil
}

func babCheck(s string, check string) bool {
	if len(s) < 3 {
		return false
	}

	for i := 0; i <= len(s)-3; i++ {
		if s[i] == check[1] && s[i+1] == check[0] && s[i+2] == check[1] {
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

		for _, match := range good_bits {
			s := match[1]

			success, aba := abaCheck(s)
			if !success {
				continue
			}

			for _, match := range bad_bits {
				for _, good := range aba {
					if babCheck(match[1], good) {
						fmt.Println(line)
						continue Line
					}
				}
			}
		}
	}
}
