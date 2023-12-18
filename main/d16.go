package main

import (
	"bufio"
	"os"
	"time"
)

const input16 = "../inputs/d16.txt"

// 0 is north, 1 is east, 2 is south, 3 is west
var direction_table = map[int][]int{0: {-1, 0}, 1: {0, 1}, 2: {1, 0}, 3: {0, -1}}

func step_forward(x, y, dir, mirror int) (newx, newy, newdir int) {
	if mirror == 0 {
		// mirror = . or splitter that's ignored
		newdir = dir
	} else if mirror == 1 {
		// mirror = /
		newdir = (1 - dir + 4) % 4
	} else if mirror == 2 {
		// mirror = \
		newdir = (3 - dir + 4) % 4
	} else {
		panic("invalid direction for stepping")
	}
	// fmt.Println("new direction", newdir)
	newx = x + direction_table[newdir][0]
	newy = y + direction_table[newdir][1]
	return
}

type light_grid struct {
	mirrors [][]int
	height  int
	width   int
}

func (l light_grid) illuminate_non_parallel(startx, starty, startdir int, e [][]bool) int {
	count := 0
	for i, j, dir := startx, starty, startdir; i >= 0 && j >= 0 && i < l.height && j < l.width; {
		curr_mirror := l.mirrors[i][j]
		//3 and 4 are splitters, - and | respectively
		if curr_mirror > 2 {
			// fmt.Println("encountered splitter")
			if e[i][j] {
				return count
			} else {
				e[i][j] = true
				count++
				if dir%2 == 1 {
					//left or right
					if curr_mirror == 3 {
						// continue going
						i, j, dir = step_forward(i, j, dir, 0)
					} else {
						//split
						count += l.illuminate_non_parallel(i-1, j, 0, e)
						count += l.illuminate_non_parallel(i+1, j, 2, e)
						return count
					}
				} else {
					//up or down
					if curr_mirror == 4 {
						//continue going
						i, j, dir = step_forward(i, j, dir, 0)
					} else {
						//split
						count += l.illuminate_non_parallel(i, j+1, 1, e)
						count += l.illuminate_non_parallel(i, j-1, 3, e)
						return count
					}
				}
			}
		} else {
			// fmt.Println("not a splitter", i, j, dir)
			if !e[i][j] {
				e[i][j] = true
				count++
			}
			i, j, dir = step_forward(i, j, dir, curr_mirror)
		}
	}
	return count
}

func build_light_grid(s bufio.Scanner) (grid light_grid) {
	for s.Scan() {
		mirror_row := []int{}
		for _, char := range []rune(s.Text()) {
			if char == '.' {
				mirror_row = append(mirror_row, 0)
			} else if char == '/' {
				mirror_row = append(mirror_row, 1)
			} else if char == '\\' {
				mirror_row = append(mirror_row, 2)
			} else if char == '-' {
				mirror_row = append(mirror_row, 3)
			} else if char == '|' {
				mirror_row = append(mirror_row, 4)
			} else {
				panic("invalid input during parsing")
			}
		}
		grid.mirrors = append(grid.mirrors, mirror_row)
	}

	grid.height = len(grid.mirrors)
	grid.width = len(grid.mirrors[0])
	return
}

func (grid light_grid) make_new_energized() (energized [][]bool) {
	for i := 0; i < grid.height; i++ {
		boolrow := []bool{}
		for j := 0; j < grid.width; j++ {
			boolrow = append(boolrow, false)
		}
		energized = append(energized, boolrow)
	}
	return
}

func illuminate_non_parallel_wrapper(x, y, dir int, g light_grid, c chan int) {
	c <- g.illuminate_non_parallel(x, y, dir, g.make_new_energized())
}

func d16p1() int {
	defer timeTrack(time.Now(), "d16p1")
	f, err := os.Open(input16)
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	grid := build_light_grid(*scanner)

	return grid.illuminate_non_parallel(0, 0, 1, grid.make_new_energized())
}

func d16p2() int {
	defer timeTrack(time.Now(), "d16p2")
	f, err := os.Open(input16)
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	grid := build_light_grid(*scanner)

	ch := make(chan int)

	for i := 0; i < grid.height; i++ {
		go illuminate_non_parallel_wrapper(i, 0, 1, grid, ch)
		go illuminate_non_parallel_wrapper(i, grid.width-1, 3, grid, ch)
	}
	for i := 0; i < grid.width; i++ {
		go illuminate_non_parallel_wrapper(0, i, 2, grid, ch)
		go illuminate_non_parallel_wrapper(grid.height-1, i, 0, grid, ch)
	}

	max := 0
	for i := 0; i < 2*(grid.height+grid.width); i++ {
		max = my_max(max, <-ch)
	}
	return max
}
