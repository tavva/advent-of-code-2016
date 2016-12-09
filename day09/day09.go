package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("./day09.txt")
	data := strings.TrimSpace(string(input))
	println(input)

	for _, c := range data {
		println(c)
	}
}
