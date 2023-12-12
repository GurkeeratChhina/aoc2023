package main

import (
	"bufio"
	"os"
	"regexp"
)

const input8 = "../inputs/d8.txt"

func binary_ditree(s bufio.Scanner) map[string][]string {
	ditree := make(map[string][]string)
	l := regexp.MustCompile("[A-Z]+")
	for s.Scan() {
		nodes := l.FindAllString(s.Text(), -1)
		ditree[nodes[0]] = []string{nodes[1], nodes[2]}
	}
	return ditree
}

func lr_to_binary(dir byte) int {
	if dir == "L"[0] {
		return 0
	} else {
		return 1
	}
}

func d8p1() int {
	return d8(func(key string) bool { return key == "AAA" }, func(key string) bool { return key == "ZZZ" })
}

func d8(input_condition, end_condition func(string) bool) int {
	file, err := os.Open(input8)
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	instructions := scanner.Text()
	scanner.Scan()

	ditree := binary_ditree(*scanner)
	var cycles []int
	for key, _ := range ditree {
		if input_condition(key) {
			current_key := key
			for i := 0; ; i++ {
				current_key = ditree[current_key][lr_to_binary(instructions[i%len(instructions)])]
				if end_condition(current_key) {
					cycles = append(cycles, i+1)
					break
				}
			}
		}
	}
	return LCM_slice(cycles)
}

func d8p2() int {
	return d8(func(key string) bool { return key[2] == "A"[0] }, func(key string) bool { return key[2] == "Z"[0] })
}
