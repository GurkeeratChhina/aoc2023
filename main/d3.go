package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

const input3 = "../inputs/d3.txt"

func d3p1() int {
	sum := 0
	n := regexp.MustCompile("\\d+")
	s := regexp.MustCompile("[^\\d\\.]")

	file, err := os.Open(input3)
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	previous := ""
	prevdigits := n.FindAllStringIndex(previous, -1)
	var prevused []int

	scanner.Scan()
	current := scanner.Text()
	currdigits := n.FindAllStringIndex(current, -1)
	var currused []int
	currsymbols := s.FindAllStringIndex(current, -1)

	scanner.Scan()
	next := scanner.Text()
	nextdigits := n.FindAllStringIndex(next, -1)
	var nextused []int
	nextsymbols := s.FindAllStringIndex(next, -1)

	for scanner.Scan() {
		previous = current
		prevdigits = currdigits
		prevused = currused

		current = next
		currdigits = nextdigits
		currused = nextused
		currsymbols = nextsymbols

		next = scanner.Text()
		nextdigits = n.FindAllStringIndex(next, -1)
		nextused = []int{}
		nextsymbols = s.FindAllStringIndex(next, -1)

		for i := 0; i < len(currsymbols); i++ {
			column := currsymbols[i][0]
			for j := 0; j < len(prevdigits); j++ {
				if !contains(prevused, prevdigits[j][0]) && prevdigits[j][0]-1 <= column && column <= prevdigits[j][1] {
					prevused = append(prevused, prevdigits[j][0])
					x, _ := strconv.Atoi(n.FindString(previous[prevdigits[j][0]:]))
					sum += x
				}
			}

			for j := 0; j < len(currdigits); j++ {
				if !contains(currused, currdigits[j][0]) && currdigits[j][0]-1 <= column && column <= currdigits[j][1] {
					currused = append(currused, currdigits[j][0])
					x, _ := strconv.Atoi(n.FindString(current[currdigits[j][0]:]))
					sum += x
				}
			}

			for j := 0; j < len(nextdigits); j++ {
				if !contains(nextused, nextdigits[j][0]) && nextdigits[j][0]-1 <= column && column <= nextdigits[j][1] {
					nextused = append(nextused, nextdigits[j][0])
					x, _ := strconv.Atoi(n.FindString(next[nextdigits[j][0]:]))
					sum += x
				}
			}
		}
	}
	return sum
}

func d3p2() int {
	sum := 0
	n := regexp.MustCompile("\\d+")
	s := regexp.MustCompile("\\*")

	file, err := os.Open(input3)
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	previous := ""
	prevdigits := n.FindAllStringIndex(previous, -1)

	scanner.Scan()
	current := scanner.Text()
	currdigits := n.FindAllStringIndex(current, -1)
	currsymbols := s.FindAllStringIndex(current, -1)

	scanner.Scan()
	next := scanner.Text()
	nextdigits := n.FindAllStringIndex(next, -1)
	nextsymbols := s.FindAllStringIndex(next, -1)

	for scanner.Scan() {
		previous = current
		prevdigits = currdigits

		current = next
		currdigits = nextdigits
		currsymbols = nextsymbols

		next = scanner.Text()
		nextdigits = n.FindAllStringIndex(next, -1)
		nextsymbols = s.FindAllStringIndex(next, -1)

		for i := 0; i < len(currsymbols); i++ {
			column := currsymbols[i][0]
			nums := []int{}

			for j := 0; j < len(prevdigits); j++ {
				if prevdigits[j][0]-1 <= column && column <= prevdigits[j][1] {
					x, _ := strconv.Atoi(n.FindString(previous[prevdigits[j][0]:]))
					nums = append(nums, x)
				}
			}

			for j := 0; j < len(currdigits); j++ {
				if currdigits[j][0]-1 <= column && column <= currdigits[j][1] {
					x, _ := strconv.Atoi(n.FindString(current[currdigits[j][0]:]))
					nums = append(nums, x)
				}
			}

			for j := 0; j < len(nextdigits); j++ {
				if nextdigits[j][0]-1 <= column && column <= nextdigits[j][1] {
					x, _ := strconv.Atoi(n.FindString(next[nextdigits[j][0]:]))
					nums = append(nums, x)
				}
			}

			if len(nums) == 2 {
				sum += nums[0] * nums[1]
			}
		}
	}
	return sum
}
