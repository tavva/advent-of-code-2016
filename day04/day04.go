package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	var input, _ = ioutil.ReadFile("./day04.txt")
	var room_strings = strings.TrimSpace(string(input))
	var lines = strings.Split(room_strings, "\n")

	for _, line := range lines {
		println(line)
	}
}
