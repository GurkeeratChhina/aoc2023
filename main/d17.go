package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"
)

const input17 = "../inputs/d17.txt"

type BFSstate struct {
	x         int
	y         int
	direction int
	steps     int
	value     int
}

type BFSstates [][][]BFSstate

type BFSBinaryTreeQueue struct {
	value *BFSstate
	left  *BFSBinaryTreeQueue
	right *BFSBinaryTreeQueue
}

func (s BFSstate) forward(turn int) (res BFSstate) {
	if turn == 0 {
		res.steps = s.steps + 1
	} else {
		res.steps = 1
	}
	res.direction = (s.direction + turn + 4) % 4
	res.x = s.x + direction_table[res.direction][0]
	res.y = s.y + direction_table[res.direction][1]
	res.value = s.value
	return
}

func (q BFSBinaryTreeQueue) add(t *BFSstate) *BFSBinaryTreeQueue {
	// fmt.Println("adding to", q)
	if q.value == nil {
		q.value = t
	} else if q.value.value < t.value {
		if q.right == nil {
			q.right = &(BFSBinaryTreeQueue{value: t})
		} else {
			q.right = q.right.add(t)
		}
	} else {
		if q.left == nil {
			q.left = &(BFSBinaryTreeQueue{value: t})
		} else {
			q.left = q.left.add(t)
		}

	}
	return &q
}

func (q BFSBinaryTreeQueue) remove_root() *BFSBinaryTreeQueue {
	if q.left == nil && q.right == nil {
		q.value = nil
		return &q
	} else if q.left == nil {
		return q.right
	} else if q.right == nil {
		return q.left
	} else {
		smallest_right, before_smallest_right := q.find_smallest_right()
		q.value = smallest_right.value
		before_smallest_right.left = smallest_right.right
		return &q
	}
}

func (q BFSBinaryTreeQueue) find_smallest_right() (smallest, prior *BFSBinaryTreeQueue) {
	smallest = q.right
	prior = &q
	for {
		if smallest.left == nil {
			return
		} else {
			prior = smallest
			smallest = smallest.left
		}
	}
}

func (q BFSBinaryTreeQueue) pop_smallest() (*BFSBinaryTreeQueue, BFSstate) {
	smallest := &q
	prior := &q
	if q.left == nil {
		result := q.value
		return q.remove_root(), *result
	}

	for {
		if smallest.left == nil {
			prior.left = smallest.right
			return &q, *(smallest.value)
		} else {
			prior = smallest
			smallest = smallest.left
		}
	}
}

func (s BFSstates) add(new BFSstate) bool {
	for _, state := range s[new.x][new.y] {
		if state.direction == new.direction && state.steps == new.steps {
			if state.value <= new.value {
				return false
			} else {
				state.value = new.value
				return true
			}
		}
	}
	s[new.x][new.y] = append(s[new.x][new.y], new)
	return true
}

func BFS(grid [][]int, state_grid BFSstates, minsteps, maxsteps int) int {
	max_x := len(grid) - 1
	max_y := len(grid[0]) - 1

	current := BFSstate{x: 0, y: 0, direction: 1, steps: 0, value: 0}
	var queue BFSBinaryTreeQueue
	q := &queue
	q = (*q).add(&current)

	for {
		// fmt.Println("before pop", q)
		if q.value == nil {
			panic("nothing to pop")
		}
		q, current = q.pop_smallest()
		// fmt.Println("after pop", current, q)
		if current.x == max_x && current.y == max_y && current.steps >= minsteps {
			return current.value
		} else if current.x == max_x && current.y == max_y {
			continue
		} else {
			for i := -1; i < 2; i++ {
				if current.steps < minsteps && i != 0 {
					//fmt.Println("skipping1")
					continue
				}
				next := current.forward(i)
				if next.steps > maxsteps || next.x < 0 || next.y < 0 || next.x > max_x || next.y > max_y {
					//fmt.Println("skipping2")
					continue
				} else {
					//fmt.Println("not skipping")
					next.value += grid[next.x][next.y]
					// fmt.Println("have state to add", next, "for turn", i)
					if state_grid.add(next) {
						q = q.add(&next)
					}
				}
			}
		}
	}
}

func digit_grid(s bufio.Scanner) (grid [][]int, state_grid BFSstates) {
	for s.Scan() {
		var row []int
		var BFSrow [][]BFSstate
		for _, digit := range strings.Split(s.Text(), "") {
			x, _ := strconv.Atoi(digit)
			row = append(row, x)
			BFSrow = append(BFSrow, []BFSstate{})
		}
		grid = append(grid, row)
		state_grid = append(state_grid, BFSrow)
	}
	return
}

func d17(minsteps, maxsteps int) int {
	f, err := os.Open(input17)
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	grid, state_grid := digit_grid(*scanner)

	return BFS(grid, state_grid, minsteps, maxsteps)
}

func d17p1() int {
	defer timeTrack(time.Now(), "d17p1")
	return d17(0, 3)
}

func d17p2() int {
	defer timeTrack(time.Now(), "d17p1")
	return d17(4, 10)
}
