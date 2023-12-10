package main

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

const input4 = "../inputs/d4.txt"

func d4p1() int {
	sum := 0

	file, err := os.Open(input4)
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	r := regexp.MustCompile("\\d+")

	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Split(strings.Split(line, ":")[1], "|")
		winners := sliceAtoI(r.FindAllString(data[0], -1))
		given := sliceAtoI(r.FindAllString(data[1], -1))
		count := len(intersect_ints(winners, given))
		if count >= 1 {
			sum += 1 << (count - 1)
		}
	}
	return sum
}

func d4p2() int {
	sum := 0
	var count [256]int
	i := 0

	file, err := os.Open(input4)
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	r := regexp.MustCompile("\\d+")

	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Split(strings.Split(line, ":")[1], "|")
		winners := sliceAtoI(r.FindAllString(data[0], -1))
		given := sliceAtoI(r.FindAllString(data[1], -1))
		wins := len(intersect_ints(winners, given))
		if wins >= 1 {
			for j := 1; j < wins+1; j++ {
				count[i+j] += count[i] + 1
			}
		}
		sum += count[i] + 1
		i++
	}
	return sum
}
