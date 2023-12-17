package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"
)

type BFSstate struct {
	x         int
	y         int
	direction int
	steps     int
	value     int
}

func (s BFSstate) forward(turn, value int) (res BFSstate) {
	if turn == 0 {
		res.steps = s.steps + 1
	} else {
		res.steps = 0
	}
	res.direction = s.direction + turn
	res.x = s.x + direction_table[res.direction][0]
	res.y = s.y + direction_table[res.direction][1]
	res.value = s.value + value
	return
}

type GridVisited [][][]BFSstate

type BFSBinaryTreeQueue struct {
	value *BFSstate
	left  *BFSBinaryTreeQueue
	right *BFSBinaryTreeQueue
}

func (q BFSBinaryTreeQueue) add(t BFSstate) {
	if q.value == nil {
		q.value = &t
	} else if q.value.value < t.value {
		q.right.add(t)
	} else {
		q.left.add(t)
	}
}

func (q BFSBinaryTreeQueue) remove_root() *BFSBinaryTreeQueue {
	if q.left == nil {
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

func (q BFSBinaryTreeQueue) pop_smallest() *BFSstate {
	smallest := q.left
	prior := &q
	for {
		if smallest.left == nil {
			prior.left = smallest.right
			return smallest.value
		} else {
			prior = smallest
			smallest = smallest.left
		}
	}
}

func digit_grid(s bufio.Scanner) (grid [][]int) {
	for s.Scan() {
		var row []int
		for _, digit := range strings.Split(s.Text(), "") {
			x, _ := strconv.Atoi(digit)
			row = append(row, x)
		}
		grid = append(grid, row)
	}
	return
}

func d17p1() int {
	defer timeTrack(time.Now(), "d16p2")
	f, err := os.Open(input16)
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	return 0
}
