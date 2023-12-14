package main

import (
	"bufio"
	"os"
	"strings"
)

const input13 = "../inputs/d13.txt"

func find_horizontal_symmetry(rows []string, differences int) int {
	for i := 1; i < len(rows); i++ {
		if has_symmetry(rows, i) == differences {
			return i
		}
	}
	return 0
}

func has_symmetry(rows []string, reflection_axis int) int {
	count := 0
	for i := 0; i < reflection_axis && i+reflection_axis < len(rows); i++ {
		if rows[reflection_axis-i-1] != rows[reflection_axis+i] {
			count++
		}
	}
	return count
}

func mirror_grid(rows []string) []string {
	num_cols := len(rows[0])
	var newgrid []string
	var split_rows [][]string
	for _, row := range rows {
		split_rows = append(split_rows, strings.Split(row, ""))
	}
	for i := 0; i < num_cols; i++ {
		newgrid = append(newgrid, "")
		for j := 0; j < len(rows); j++ {
			newgrid[i] += split_rows[j][i]
		}
	}
	return newgrid
}

func score(rows []string) int {
	horizontal := find_horizontal_symmetry(rows)
	if horizontal == 0 {
		return find_horizontal_symmetry(mirror_grid(rows))
	}
	return 100 * horizontal
}

func d13p1() int {
	f, err := os.Open(input13)
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	sum := 0

	var grid []string
	for scanner.Scan() {
		if scanner.Text() == "" {
			sum += score(grid)
			grid = nil
		} else {
			grid = append(grid, scanner.Text())
		}
	}
	sum += score(grid)

	return sum
}

func d13p2() int {
	f, err := os.Open(input13)
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	sum := 0

	var grid []string
	for scanner.Scan() {
		if scanner.Text() == "" {
			sum += score(grid)
			grid = nil
		} else {
			grid = append(grid, scanner.Text())
		}
	}
	sum += score(grid)

	return sum
}
