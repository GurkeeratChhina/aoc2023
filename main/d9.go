package main

import (
	"bufio"
	"os"
)

const input9 = "../inputs/d9.txt"

func finite_difference(s []int) (r []int) {
	for i := 0; i < len(s)-1; i++ {
		r = append(r, s[i+1]-s[i])
	}
	return
}

func slice_is_zero(s []int) bool {
	for _, val := range s {
		if val != 0 {
			return false
		}
	}
	return true
}

func d9(mode int) int {
	file, err := os.Open(input9)
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		y_values := extract_nums(scanner.Text())
		parity := 1
		for !slice_is_zero(y_values) {
			sum += y_values[mode*(len(y_values)-1)] * parity
			y_values = finite_difference(y_values)
			if mode == 0 {
				parity *= -1
			}
		}
	}
	return sum
}

func d9p1() int {
	return d9(1)
}

func d9p2() int {
	return d9(0)
}
