package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func process_marker(s string) {
	end_index := strings.Index(s, ")")
	marker := s[:end_index]

	bits := strings.Split(marker, "x")
	seq_len, _ := strconv.Atoi(bits[0])
	seq_rep, _ := strconv.Atoi(bits[1])

	for i := 0; i < seq_rep; i++ {
		fmt.Printf(s[end_index+1 : end_index+1+seq_len])
	}

	process_data(s[end_index+1+seq_len:])
}

func process_data(s string) {
For:
	for i := 0; i < len(s); i++ {
		r := rune(s[i])

		switch r {
		case ' ':
			continue
		case '(':
			process_marker(s[i+1:])
			break For
		default:
			fmt.Printf("%v", string(r))
		}
	}
}

func main() {
	input, _ := ioutil.ReadFile("./day09-short.txt")
	data := strings.TrimSpace(string(input))

	process_data(data)
}
