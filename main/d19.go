package main

import (
	"bufio"
	"os"
	"regexp"
	"strings"
	"time"
)

const input19 = "../inputs/d19.txt"

type part struct {
	x, m, a, s []int
}

type part_condition struct {
	ineq string
	res  string
}

type pipeline struct {
	criteria []part_condition
}

// "half" of parsing input
func make_part(input_line string) part {
	nums := extract_nums(input_line)
	return part{
		x: []int{nums[0], nums[0] + 1},
		m: []int{nums[1], nums[1] + 1},
		a: []int{nums[2], nums[2] + 1},
		s: []int{nums[3], nums[3] + 1},
	}
}

// other "half" of parsing input
func make_pipeline(input string) pipeline {
	r := regexp.MustCompile(":")
	var conds []part_condition

	for _, condition := range strings.Split(input, ",") {
		// could and should use regex here
		if r.MatchString(condition) {
			// has a :
			ineq := strings.Split(condition, ":")[0]
			res := strings.Split(condition, ":")[1]
			conds = append(conds, part_condition{ineq: ineq, res: res})
		} else {
			// does not have a :
			ineq := "true"
			res := condition[:len(condition)-1]
			conds = append(conds, part_condition{ineq: ineq, res: res})
		}
	}
	return pipeline{criteria: conds}
}

func (p part) get_attr(c rune) []int {
	switch c {
	case 'x':
		return p.x
	case 'm':
		return p.m
	case 'a':
		return p.a
	case 's':
		return p.s
	}
	panic("trying to get attribute not one of xmas")
}

func (p *part) set_attr(c rune, val []int) {
	switch c {
	case 'x':
		p.x = val
		return
	case 'm':
		p.m = val
		return
	case 'a':
		p.a = val
		return
	case 's':
		p.s = val
		return
	}
	panic("trying to set attribute not one of xmas")
}

// returns a range of parts that match the inequality condition, and set p to the leftovers.
func (p *part) intersect_condition(ineq string) part {
	r := regexp.MustCompile("<")
	var ineq_range []int
	if ineq == "true" {
		return *p
	} else if r.MatchString(ineq) {
		ineq_range = []int{1, extract_nums(ineq)[0]}
	} else {
		ineq_range = []int{extract_nums(ineq)[0] + 1, 4001}
	}

	match := part{x: p.x, m: p.m, a: p.a, s: p.s}
	char := []rune(ineq)[0]
	new_range := intersect_intervals(p.get_attr(char), ineq_range)
	leftover_ranges := interval_subtract(p.get_attr(char), new_range)

	match.set_attr(char, new_range)
	if len(leftover_ranges) == 0 {
		p.set_attr(char, []int{})
	} else {
		p.set_attr(char, leftover_ranges[0])
	}
	return match
}

func (p part) is_valid() bool {
	if len(p.x) < 2 || len(p.m) < 2 || len(p.a) < 2 || len(p.s) < 2 {
		return false
	}
	if p.x[0] >= p.x[1] || p.m[0] >= p.m[1] || p.a[0] >= p.a[1] || p.s[0] >= p.s[1] {
		return false
	}
	return true
}

// how to deal with the result of inequality, using recursion
func (p part) process_pipeline(pl_key string, pipeline_map map[string]pipeline) (accepted_parts []part) {
	for _, cond := range pipeline_map[pl_key].criteria {
		// produces a part that matches the criteria, and sets p to the remainder which doesn't match the criteria
		match := p.intersect_condition(cond.ineq)
		if match.is_valid() {
			if cond.res == "A" {
				accepted_parts = append(accepted_parts, match)
			} else if cond.res == "R" {
			} else {
				accepted_parts = append(accepted_parts, match.process_pipeline(cond.res, pipeline_map)...)
			}
		}
		if !p.is_valid() {
			break
		}
	}
	return
}

func d19p1() int {
	defer timeTrack(time.Now(), "d19p1")
	f, err := os.Open(input19)
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	// make pipelines
	pipeline_map := make(map[string]pipeline)
	for scanner.Scan() {
		if scanner.Text() == "" { // reached the empty line
			break
		}
		pipeline_key := strings.Split(scanner.Text(), "{")[0]
		pipeline_data := strings.Split(scanner.Text(), "{")[1]
		pipeline_map[pipeline_key] = make_pipeline(pipeline_data)
	}

	// make parts
	var part_list []part
	for scanner.Scan() {
		part_list = append(part_list, make_part(scanner.Text()))

	}

	// finding accepted parts
	var accepted []part
	for _, p := range part_list {
		accepted = append(accepted, p.process_pipeline("in", pipeline_map)...)
	}

	sum := 0
	for _, p := range accepted {
		sum += p.x[0] + p.m[0] + p.a[0] + p.s[0]
	}
	return sum
}

func d19p2() int {
	defer timeTrack(time.Now(), "d19p2")
	f, err := os.Open(input19)
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	// make pipelines
	pipeline_map := make(map[string]pipeline)
	for scanner.Scan() {
		if scanner.Text() == "" { // reached the empty line
			break
		}
		pipeline_key := strings.Split(scanner.Text(), "{")[0]
		pipeline_data := strings.Split(scanner.Text(), "{")[1]
		pipeline_map[pipeline_key] = make_pipeline(pipeline_data)
	}

	max_part := part{x: []int{1, 4001}, m: []int{1, 4001}, a: []int{1, 4001}, s: []int{1, 4001}}
	accepted := max_part.process_pipeline("in", pipeline_map)

	sum := 0
	for _, p := range accepted {
		sum += (p.x[1] - p.x[0]) * (p.m[1] - p.m[0]) * (p.a[1] - p.a[0]) * (p.s[1] - p.s[0])
	}
	return sum
}
