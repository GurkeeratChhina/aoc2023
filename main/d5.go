package main

import (
	"bufio"
	"os"
	"slices"
)

const input5 = "../inputs/d5.txt"

func destination(source int, transform_data []int) int {
	return source + transform_data[0] - transform_data[1]
}

func d5p1() int {
	file, err := os.Open(input5)
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	current := extract_nums(scanner.Text())
	var next []int

	for scanner.Scan() {
		data := extract_nums(scanner.Text())
		if len(data) == 0 {
			current = append(current, next...)
			next = nil
		} else {
			for i := 0; i < len(current); {
				if current[i] >= data[1] && current[i] < data[1]+data[2] {
					next = append(next, destination(current[i], data))
					current[i] = current[len(current)-1]
					current = current[:len(current)-1]
				} else {
					i++
				}
			}
		}
	}
	current = append(current, next...)

	return slices.Min(current)
}

func d5p2() int {
	file, err := os.Open(input5)
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	seeds := extract_nums(scanner.Text())
	var current [][]int
	var next [][]int

	for i := 0; i < len(seeds); i = i + 2 {
		current = append(current, []int{seeds[i], seeds[i] + seeds[i+1]})
	}

	for scanner.Scan() {
		data := extract_nums(scanner.Text())
		if len(data) == 0 {
			current = append(current, next...)
			next = nil
		} else {
			interval2 := []int{data[1], data[1] + data[2]}
			for i := 0; i < len(current); {
				intersection := intersect_intervals(current[i], interval2)
				if intersection == nil {
					i++
				} else {
					next = append(next, []int{destination(intersection[0], data), destination(intersection[1], data)})
					leftovers := interval_subtract(current[i], interval2)
					current[i] = current[len(current)-1]
					current = current[:len(current)-1]
					current = append(current, leftovers...)
				}
			}
		}
	}
	current = append(current, next...)

	min_loc := current[0][0]
	for _, interval := range current {
		min_loc = min(min_loc, interval[0])
	}
	return min_loc
}
