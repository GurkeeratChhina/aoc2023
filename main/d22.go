package main

import (
	"bufio"
	"os"
	"sort"
	"time"
)

const input22 = "../inputs/d22.txt"

type tetris struct {
	supporting   map[*tetris]bool
	supported_by map[*tetris]bool
	points       []*threepoint
	lowest_z     int
}

type threepoint struct {
	x, y, z int
}

type tetrisgrid map[int]map[int]map[int]*tetris

func (tg tetrisgrid) addpoint(p *threepoint, data *tetris) {
	_, ok := tg[p.x]
	if !ok {
		tg[p.x] = make(map[int]map[int]*tetris)
		tg[p.x][p.y] = make(map[int]*tetris)
		tg[p.x][p.y][p.z] = data
		return
	}
	_, ok = tg[p.x][p.y]
	if !ok {
		tg[p.x][p.y] = make(map[int]*tetris)
		tg[p.x][p.y][p.z] = data
		return
	}
	_, ok = tg[p.x][p.y][p.z]
	if ok {
		panic("value already exists at key!")
	}
	tg[p.x][p.y][p.z] = data
	return
}

func (t *tetris) movedown(g tetrisgrid) {
	found := false
	for {
		for _, p := range t.points {
			if p.z-1 == 0 {
				// reached bottom
				for _, q := range t.points {
					g.addpoint(q, t)
				}
				return
			}
			below, ok := g[p.x][p.y][p.z-1]
			if ok {
				// found a tetris piece below
				t.supported_by[below] = true
				below.supporting[t] = true
				found = true
			}
		}
		if found {
			for _, q := range t.points {
				g.addpoint(q, t)
			}
			return
		}
		for _, q := range t.points {
			q.z = q.z - 1
		}
	}
}

func (t *tetris) is_removable() bool {
	for block := range t.supporting {
		if len(block.supported_by) == 1 {
			return false
		}
	}
	return true
}

func (t *tetris) will_fall(removed map[*tetris]bool) bool {
	for s := range t.supported_by {
		if _, ok := removed[s]; !ok {
			return false
		}
	}
	return true
}

func (t *tetris) remove(removed map[*tetris]bool) {
	removed[t] = true
	for s := range t.supporting {
		// fmt.Println(t, "is supporting", s)
		if s.will_fall(removed) {
			s.remove(removed)
		}
	}
}

func make_tetris(input string) *tetris {
	res := tetris{}
	coords := extract_nums(input)
	if coords[0] != coords[3] {
		// x values are different
		for i := coords[0]; i <= coords[3]; i++ {
			res.points = append(res.points, &threepoint{i, coords[1], coords[2]})
		}
	} else if coords[1] != coords[4] {
		// y values are different
		for i := coords[1]; i <= coords[4]; i++ {
			res.points = append(res.points, &threepoint{coords[0], i, coords[2]})
		}
	} else if coords[2] != coords[5] {
		// z values are different
		for i := coords[2]; i <= coords[5]; i++ {
			res.points = append(res.points, &threepoint{coords[0], coords[1], i})
		}
	} else {
		// single point block
		res.points = append(res.points, &threepoint{coords[0], coords[1], coords[2]})
	}
	res.lowest_z = coords[2]
	res.supported_by = make(map[*tetris]bool)
	res.supporting = make(map[*tetris]bool)
	return &res
}

func d22p1() int {
	defer timeTrack(time.Now(), "d22p1")
	f, err := os.Open(input22)
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var tetrislist []*tetris
	for scanner.Scan() {
		tetrislist = append(tetrislist, make_tetris(scanner.Text()))
	}

	// sort by lowest z
	sort.SliceStable(tetrislist, func(i, j int) bool { return tetrislist[i].lowest_z < tetrislist[j].lowest_z })

	grid := make(tetrisgrid)
	for _, t := range tetrislist {
		t.movedown(grid)
	}

	count := 0
	for _, t := range tetrislist {
		if t.is_removable() {
			count++
		}
	}
	return count
}

func d22p2() int {
	defer timeTrack(time.Now(), "d22p2")
	f, err := os.Open(input22)
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var tetrislist []*tetris
	for scanner.Scan() {
		tetrislist = append(tetrislist, make_tetris(scanner.Text()))
	}

	// sort by lowest z
	sort.SliceStable(tetrislist, func(i, j int) bool { return tetrislist[i].lowest_z < tetrislist[j].lowest_z })

	grid := make(tetrisgrid)
	for _, t := range tetrislist {
		t.movedown(grid)
	}

	sum := 0
	for _, t := range tetrislist {
		removed := make(map[*tetris]bool)
		t.remove(removed)
		// fmt.Println(t, removed)
		sum += len(removed) - 1
	}

	return sum
}
