package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"unicode/utf8"
)

const inputtest = "../inputs/test.txt"

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

func extract_nums(s string) []int {
	r := regexp.MustCompile("-?\\d+")
	return sliceAtoI(r.FindAllString(s, -1))
}

func intersect_intervals(ab, cd []int) []int {
	if ab[0] >= cd[1] || ab[1] <= cd[0] {
		return nil
	} else {
		return []int{max(ab[0], cd[0]), min(ab[1], cd[1])}
	}
}

func interval_subtract(ab, cd []int) [][]int {
	if ab[0] < cd[0] && ab[1] > cd[1] {
		//second interval is inside first
		return [][]int{{ab[0], cd[0]}, {cd[1], ab[1]}}
	} else if ab[0] >= cd[0] && ab[1] <= cd[1] {
		// second interval is superset of first
		return [][]int{}
	} else if ab[0] >= cd[0] && ab[1] > cd[1] {
		// second interval is to the left of first
		return [][]int{{cd[1], ab[1]}}
	} else if ab[0] < cd[0] && ab[1] <= cd[1] {
		//second interval is to the right of first
		return [][]int{{ab[0], cd[0]}}
	} else {
		panic("interval_subtract invalid intervals")
	}
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

func LCM_slice(s []int) int {
	l := s[0]
	for _, i := range s {
		l = LCM(l, i)
	}
	return l
}

func Reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}

func slice_sum(s []int) (r int) {
	for _, i := range s {
		r += i
	}
	return
}

func main() {
	fmt.Println("Part 1: ", d14p1())
	fmt.Println("Part 2: ", d14p2())
}
