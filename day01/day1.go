package main

import (
	"io/ioutil"
)

func main() {
	var input, _ = ioutil.ReadFile("./day01.txt")

	println(string(input))
}
