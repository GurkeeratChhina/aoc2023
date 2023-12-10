package main

import (
	"fmt"
	"log"
	"strconv"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func intersect_ints(s1, s2 []int) (intersection []int) {
	hash := make(map[int]bool)
	for _, e := range s1 {
		hash[e] = true
	}
	for _, e := range s2 {
		// If elements present in the hashmap then append intersection list.
		if hash[e] {
			intersection = append(intersection, e)
			hash[e] = false
		}
	}
	return
}

func sliceAtoI(s []string) (r []int) {
	for _, e := range s {
		x, _ := strconv.Atoi(e)
		r = append(r, x)
	}
	return
}

func main() {
	fmt.Println("Part 1: ", d4p1())
	fmt.Println("Part 2: ", d4p2())
}
