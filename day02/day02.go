package main

import (
	"io/ioutil"
	"math"
	"strings"
)

func main() {
	var pin string

	keypad := [3][3]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}

	var x, y = 1, 1 // coordinates on keypad

	var input, _ = ioutil.ReadFile("./day02.txt")
	var command_string = strings.TrimSpace(string(input))
	var lines = strings.Split(command_string, "\n")

	for _, line := range lines {
		for _, direction := range line {
			switch direction {
			case 'U':
				y = int(math.Max(0, float64(y)-1))
			case 'D':
				y = int(math.Min(2, float64(y)+1))
			case 'L':
				x = int(math.Max(0, float64(x)-1))
			case 'R':
				x = int(math.Min(2, float64(x)+1))
			}
		}

		pin = pin + keypad[y][x]
	}

	println(pin)
}
