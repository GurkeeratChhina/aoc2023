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
	x, m, a, s, total int
}

type part_condition struct {
	ineq string
	res  string
}

type pipeline struct {
	conditions []part_condition
}

// "half" of parsing input
func make_part(input_line string) part {
	nums := extract_nums(input_line)
	return part{x: nums[0], m: nums[1], a: nums[2], s: nums[3], total: slice_sum(nums)}
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
	return pipeline{conditions: conds}
}

// actually doing the inequality somehow???
func (p part) check_condition(ineq string) bool {
	r := regexp.MustCompile("<")
	if ineq == "true" {
		return true
	} else if r.MatchString(ineq) {
		if p.get_attr([]rune(ineq)[0]) < extract_nums(ineq)[0] {
			return true
		}
		return false
	} else {
		if p.get_attr([]rune(ineq)[0]) > extract_nums(ineq)[0] {
			return true
		}
		return false
	}
}

func (p part) get_attr(c rune) int {
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

// how to deal with the result of inequality, using recursion
func (p part) process_pipeline(pl_key string, pipeline_map map[string]pipeline) int {
	for _, cond := range pipeline_map[pl_key].conditions {
		if p.check_condition(cond.ineq) {
			if cond.res == "A" {
				return p.total
			} else if cond.res == "R" {
				return 0
			} else {
				return p.process_pipeline(cond.res, pipeline_map)
			}
		}
	}
	panic("reached end of conditions somehow")
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

	// summing results for each part
	sum := 0
	for _, p := range part_list {
		sum += p.process_pipeline("in", pipeline_map)
	}
	return sum
}
