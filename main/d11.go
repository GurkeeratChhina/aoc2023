package main

import (
	"bufio"
	"os"
	"sort"
	"strings"
)

const input11 = "../inputs/d11.txt"

// finds sum of pairwise dist from sorted array
func pairwise_dist(xs []int) (sum int) {
	for i := 0; i < len(xs); i++ {
		sum += (-len(xs) + 1 + 2*i) * xs[i]
	}
	return
}

// expands a sorted array
func expand(xs []int, amount int) {
	next := xs[len(xs)-1]
	for i := len(xs) - 2; i >= 0; i-- {
		if next-xs[i] > 1 {
			for j := i + 1; j < len(xs); j++ {
				xs[j] += amount * (next - xs[i] - 1)
			}
		}
		next = xs[i]
	}
}

func star_locs(file string) (xs, ys []int) {
	f, err := os.Open(file)
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for row := 0; scanner.Scan(); row++ {
		for col, symb := range strings.Split(scanner.Text(), "") {
			if symb == "#" {
				xs = append(xs, row)
				ys = append(ys, col)
			}
		}
	}
	return
}

func d11p1() int {
	xs, ys := star_locs(input11)
	sort.Ints(xs)
	sort.Ints(ys)
	expand(xs, 1)
	expand(ys, 1)
	return pairwise_dist(xs) + pairwise_dist(ys)
}

func d11p2() int {
	xs, ys := star_locs(input11)
	sort.Ints(xs)
	sort.Ints(ys)
	expand(xs, 1000000-1)
	expand(ys, 1000000-1)
	return pairwise_dist(xs) + pairwise_dist(ys)
}
