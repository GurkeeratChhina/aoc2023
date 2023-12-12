package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const input10 = "../inputs/d10.txt"

// direction is 3=N, 0=S, 4=E, 2=W
var direction_coords = map[int][]int{
	3: {-1, 0},
	0: {1, 0},
	4: {0, 1},
	2: {0, -1},
}

// each pipe is given a value as the sum of the two directions it connects, mod 6
// this enables us to easily locate pipes combinates that, when combined, form a N-S wall
// since EW combinations sum to 0 and NS sum to 3 by construction
var pipe_vals = map[string]int{
	"|": 3,
	"-": 0,
	"J": 5,
	"L": 1,
	"7": 2,
	"F": 4,
	".": -1,
}

func find_S(grid [][]string) (int, int) {
	for i, row := range grid {
		for j, cell := range row {
			if cell == "S" {
				return i, j
			}
		}
	}
	panic("No s found")
}

func replace_S(grid [][]string) int {
	row, col := find_S(grid)
	north := pipe_vals[grid[row-1][col]]
	east := pipe_vals[grid[row][col+1]]
	west := pipe_vals[grid[row][col-1]]

	S_val := 0
	if north == 2 || north == 3 || north == 4 {
		S_val += 3
	}
	if east == 0 || east == 2 || east == 5 {
		S_val += 4
	}
	if west == 0 || west == 1 || west == 4 {
		S_val += 2
	}
	S_val = S_val % 6

	for _, val := range pipe_vals {
		if S_val == val {
			grid[row][col] = strconv.Itoa(val)
			return val
		}
	}
	panic("Start doesn't match any pipe")
}

func build_grid(file string) (grid [][]string) {
	f, err := os.Open(file)
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := "." + scanner.Text() + "."
		grid = append(grid, strings.Split(line, ""))
	}
	width := len(grid[0])
	grid = append(grid, strings.Split(strings.Repeat(".", width), ""))
	grid = append([][]string{strings.Split(strings.Repeat(".", width), "")}, grid...)
	return
}

func traverse_loop(grid [][]string) int {
	rowS, colS := find_S(grid)
	S := replace_S(grid)

	dir := 0
	if S == 1 || S == 3 || S == 5 {
		dir = 3
	} else if S == 0 || S == 1 || S == 4 {
		dir = 4
	} else if S == 0 || S == 2 || S == 5 {
		dir = 2
	}

	row := rowS
	col := colS
	for i := 0; ; i++ {
		deltas := direction_coords[dir]
		row = row + deltas[0]
		col = col + deltas[1]
		if row == rowS && col == colS {
			return i + 1
		}
		if dir%3 == 0 {
			dir = (dir + 3 + pipe_vals[grid[row][col]]) % 6
		} else {
			dir = (dir + pipe_vals[grid[row][col]]) % 6
		}
		grid[row][col] = strconv.Itoa(pipe_vals[grid[row][col]])
	}
}

func d10p1() int {
	grid := build_grid(input10)

	loop_len := traverse_loop(grid)

	return loop_len / 2
}

func d10p2() int {
	grid := build_grid(input10)

	traverse_loop(grid)

	count := 0
	parity := 0
	for _, row := range grid {
		for _, cell := range row {
			num, err := strconv.Atoi(cell)
			if err == nil && num < 7 {
				parity = (parity + num) % 6
			} else {
				if parity == 3 {
					count++
				} else if parity != 0 {
					panic("wrong parity")
				}
			}
		}
	}
	return count
}
