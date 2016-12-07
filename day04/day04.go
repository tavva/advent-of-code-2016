package main

import (
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("./day04.txt")
	room_strings := strings.TrimSpace(string(input))
	lines := strings.Split(room_strings, "\n")

	re := regexp.MustCompile("^([a-z-]+)([0-9]+)\\[([a-z]+)\\]$")

	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		name := strings.Replace(matches[1], "-", "", -1)
		sector_id := matches[2]
		checksum := matches[3]
		println(name)
		println(sector_id)
		println(checksum)
	}
}
