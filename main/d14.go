package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const input14 = "../inputs/d14.txt"

func load_north(grid [][]int) (sum int) {
	for i, row := range grid {
		sum += (i + 1) * rock_count(row)
	}
	return
}

func rock_count(row []int) (count int) {
	for _, x := range row {
		if x == 1 {
			count++
		}
	}
	return
}

func rotate90(grid [][]int) (newgrid [][]int) {
	num_cols := len(grid[0])
	num_rows := len(grid)
	for i := 0; i < num_cols; i++ {
		newgrid = append(newgrid, []int{})
		for j := num_rows - 1; j >= 0; j-- {
			newgrid[i] = append(newgrid[i], grid[j][i])
		}
	}
	return
}

func move_rock_next(grid [][]int) [][]int {
	new_grid := rotate90(grid)
	move_east(new_grid)
	return new_grid
}

func rock_spin_cycle(grid [][]int) [][]int {
	for i := 0; i < 4; i++ {
		grid = move_rock_next(grid)
	}
	return grid
}

func move_east(grid [][]int) {
	for _, row := range grid {
		align_right(row)
	}
}

func align_right(row []int) {
	for j := len(row) - 2; j >= 0; j-- {
		if row[j] == 1 {
			// found a rock
			for k := j + 1; k < len(row)+1; k++ {
				//look for right-most empty space
				if k == len(row) {
					row[j] = 0
					row[k-1] = 1
				} else if row[k] > 0 {
					// found obstruction
					row[j] = 0
					row[k-1] = 1
					break
				}
			}
		}
	}
}

func grid_is_equal(grid1, grid2 [][]int) bool {
	for i := range grid1 {
		for j := range grid1[0] {
			if grid1[i][j] != grid2[i][j] {
				return false
			}
		}
	}
	return true
}

func print_grid(grid [][]int) {
	fmt.Println()
	for _, row := range grid {
		fmt.Println(row)
	}
}

func parse_rocks(line string) (result []int) {
	for _, char := range strings.Split(line, "") {
		if char == "." {
			result = append(result, 0)
		} else if char == "O" {
			result = append(result, 1)
		} else if char == "#" {
			result = append(result, 2)
		} else {
			panic("invalid character found")
		}
	}
	return
}

func d14p1() int {
	f, err := os.Open(input14)
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var grid [][]int
	for scanner.Scan() {
		grid = append(grid, parse_rocks(scanner.Text()))
	}

	grid = move_rock_next(grid)
	grid = rotate90(grid)
	return load_north(grid)
}

func d14p2() int {
	f, err := os.Open(inputtest)
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var grid [][]int
	for scanner.Scan() {
		grid = append(grid, parse_rocks(scanner.Text()))
	}

	state_cache := make(map[int][][]int)

	for i := 0; ; i++ {
		fmt.Println(load_north(rotate90(rotate90(grid))))
		grid = rock_spin_cycle(grid)
		// print_grid(grid)
		for key, val := range state_cache {
			if grid_is_equal(val, grid) {
				// found cycle
				cycle_length := i - key
				cycle_start := key
				steps_into_cycle := (1000000000 - cycle_start) % cycle_length
				fmt.Println("found cycle - start & length", cycle_start, cycle_length, steps_into_cycle)
				answer_grid := state_cache[cycle_start+steps_into_cycle-1]
				answer_grid = rotate90(rotate90(answer_grid))
				return load_north(answer_grid)
			}
		}
		state_cache[i] = grid
	}
}
