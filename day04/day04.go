package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Pair struct {
	Key   string
	Value int
}

func (p Pair) String() string {
	return fmt.Sprintf("%s: %d", p.Key, p.Value)
}

type PairList []Pair

func (p PairList) Len() int      { return len(p) }
func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool {
	if p[i].Value == p[j].Value {
		return p[i].Key > p[j].Key
	}
	return p[i].Value < p[j].Value
}

func main() {
	input, _ := ioutil.ReadFile("./day04.txt")
	room_strings := strings.TrimSpace(string(input))
	lines := strings.Split(room_strings, "\n")

	re := regexp.MustCompile("^([a-z-]+)([0-9]+)\\[([a-z]+)\\]$")

	sector_sum := 0

	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		name := matches[1]
		println(name)
		sector_id, _ := strconv.Atoi(matches[2])

		decrypted_name := []rune(name)

		for i := 0; i < sector_id; i++ {
			for j, x := range decrypted_name {
				var decrypted_letter rune

				if x == 'z' {
					decrypted_letter = 'a'
				} else if x == ' ' {
					decrypted_letter = ' '
				} else if x == '-' {
					decrypted_letter = ' '
				} else {
					decrypted_letter = rune(int(x) + 1)
				}

				decrypted_name[j] = decrypted_letter
			}
		}

		name = strings.Replace(name, "-", "", -1)

		checksum := matches[3]

		letter_counts := make(map[rune][]int)

		for _, x := range name {
			letter_counts[x] = append(letter_counts[x], 1)
		}

		list := PairList{}

		for k, v := range letter_counts {
			list = append(list, Pair{string(k), len(v)})
		}
		sort.Sort(sort.Reverse(list))

		sum := ""

		for _, p := range list[:5] {
			sum += p.Key
		}

		if sum == checksum {
			sector_sum += sector_id
		}

		println(string(decrypted_name))
	}

	println(sector_sum)
}
