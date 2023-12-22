package main

import (
	"bufio"
	"os"
	"regexp"
	"time"
)

const input21 = "../inputs/d21.txt"

type garden [][]int

func make_garden_grid(s bufio.Scanner) (g garden, start point) {
	r := regexp.MustCompile("S")
	for i := 0; s.Scan(); i++ {
		if r.MatchString(s.Text()) {
			start.x = i
			start.y = r.FindStringIndex(s.Text())[0]
		}
		var row []int
		for _, c := range []rune(s.Text()) {
			if c == '#' {
				row = append(row, -3)
			} else if c == '.' || c == 'S' {
				row = append(row, -1)
			} else {
				panic("invalid rune during parsing")
			}
		}
		g = append(g, row)
	}
	return
}

func (p point) neighbours() []point {
	n := point{x: p.x - 1, y: p.y}
	e := point{x: p.x, y: p.y + 1}
	s := point{x: p.x + 1, y: p.y}
	w := point{x: p.x, y: p.y - 1}
	return []point{n, e, s, w}
}

func (g garden) walk(max_steps int, start point) {
	maxx := len(g)
	maxy := len(g[0])

	// fmt.Println(maxx, maxy)
	var queue []point
	queue = append(queue, start)
	g[start.x][start.y] = 0

	var current point

	for {
		if len(queue) == 0 {
			break
		} else {
			current = queue[0]
			queue = queue[1:]
		}
		if g[current.x][current.y] == max_steps {
			continue
		}
		for _, p := range current.neighbours() {
			if p.x < 0 || p.x >= maxx || p.y < 0 || p.y >= maxy {
				continue
			} else if g[p.x][p.y] != -1 {
				continue
			} else {
				g[p.x][p.y] = g[current.x][current.y] + 1
				queue = append(queue, p)
			}
		}
	}
}

func (g garden) copy() garden {
	c := make([][]int, len(g))
	for i := range c {
		c[i] = make([]int, len(g[i]))
		copy(c[i], g[i])
	}
	return c
}

func (g garden) count_squares(steps int, start point) []int {
	copyg := g.copy()
	copyg.walk(steps, start)
	evens := 0
	odds := 0
	for _, row := range copyg {
		for _, cell := range row {
			if cell%2 == 0 {
				evens++
			} else if cell%2 == 1 {
				odds++
			}
		}
	}
	return []int{evens, odds}
}

func d21p1() int {
	defer timeTrack(time.Now(), "d21p1")
	f, err := os.Open(input21)
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	g, start := make_garden_grid(*scanner)
	parities := g.count_squares(64, start)
	return parities[0]
}

func d21p2() int {
	defer timeTrack(time.Now(), "d21p2")
	f, err := os.Open(input21)
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	g, start := make_garden_grid(*scanner)

	parities := g.count_squares(1000, start)

	// constants used
	height := len(g)
	total_steps := 26501365
	squares_seen := parities[1]

	full_grids := total_steps/height - 1     // 202299 full grids - the last grid isn't fully filled
	boundary_steps := total_steps%height - 1 // 64 steps left from corner points

	// completely filled grids (inside of diamond)
	squares_seen += (full_grids - 1) * (full_grids + 1) * parities[1] // counting same parity grids - using geometry this is (n-1)(n+1) where n is the "radius"
	squares_seen += (full_grids + 1) * (full_grids + 1) * parities[0] // counting different parity grids - similar geoemtry yields (n+1)(n+1)

	// boundary grids on each diagonal side
	corners := []point{{0, 0}, {0, height - 1}, {height - 1, 0}, {height - 1, height - 1}}
	for _, corner := range corners {

		// grids that are 7/8 filled, there are n of them on each diagonal side
		parities = g.count_squares(boundary_steps+height, corner)
		squares_seen += parities[1] * full_grids

		// grids that are 1/8 filled - there are n+1 of them on each diagonal side
		parities = g.count_squares(boundary_steps, corner)
		squares_seen += parities[0] * (full_grids + 1)
	}

	// boundary grids on axis
	boundary_steps = total_steps - height/2 - full_grids*height - 1

	centres := []point{{0, height / 2}, {height - 1, height / 2}, {height / 2, 0}, {height / 2, height - 1}}
	for _, centre := range centres {
		parities = g.count_squares(boundary_steps, centre)
		squares_seen += parities[0]
	}

	return squares_seen
}
