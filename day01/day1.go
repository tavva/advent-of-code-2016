package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	var input, _ = ioutil.ReadFile("./day01.txt")
	var command_string = strings.TrimSpace(string(input))

	var commands = strings.Split(command_string, ", ")

	var x, y float64 // defaults to 0
	var current_direction int

	for _, v := range commands {
		var direction = string(v[0])
		var distance, _ = strconv.Atoi(string(v[1]))

		if direction == "R" {
			current_direction++
		} else if direction == "L" {
			current_direction += -1 + 4
		}

		current_direction %= 4

		switch current_direction {
		case 0:
			y += float64(distance)
		case 1:
			x += float64(distance)
		case 2:
			y -= float64(distance)
		case 3:
			x -= float64(distance)
		}
	}

	fmt.Printf("%v", math.Abs(x)+math.Abs(y))
}
