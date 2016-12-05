package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	var input, _ = ioutil.ReadFile("./day01.txt")
	var command_string = strings.TrimSpace(string(input))

	var commands = strings.Split(command_string, ", ")

	for _, v := range commands {
		var direction = string(v[0])
		var distance, _ = strconv.Atoi(string(v[1]))
	}
}
