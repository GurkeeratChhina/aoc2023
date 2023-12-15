package main

import (
	"bufio"
	"os"
	"strings"
)

const input13 = "../inputs/d13.txt"

func find_horizontal_symmetry(rows []string, differences int) int {
	for i := 1; i < len(rows); i++ {
		if has_symmetry(rows, i, differences) == differences {
			return i
		}
	}
	return 0
}

func has_symmetry(rows []string, reflection_axis, max int) int {
	count := 0
	for i := 0; i < reflection_axis && i+reflection_axis < len(rows); i++ {
		count += num_diffs(rows[reflection_axis-i-1], rows[reflection_axis+i], max-count)
		if count > max {
			return count
		}
	}
	return count
}

func num_diffs(row1, row2 string, max_diffs int) int {
	count := 0
	row1chars := strings.Split(row1, "")
	row2chars := strings.Split(row2, "")
	for i := range row1chars {
		if row1chars[i] != row2chars[i] {
			count++
			if count > max_diffs {
				return count
			}
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

func score(rows []string, diffs int) int {
	horizontal := find_horizontal_symmetry(rows, diffs)
	if horizontal == 0 {
		return find_horizontal_symmetry(mirror_grid(rows), diffs)
	}
	return 100 * horizontal
}

func d13(num_differences int) int {
	f, err := os.Open(input13)
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	sum := 0

	var grid []string
	for scanner.Scan() {
		if scanner.Text() == "" {
			sum += score(grid, num_differences)
			grid = nil
		} else {
			grid = append(grid, scanner.Text())
		}
	}
	sum += score(grid, num_differences)

	return sum
}

func d13p1() int {
	return d13(0)
}

func d13p2() int {
	return d13(1)
}
