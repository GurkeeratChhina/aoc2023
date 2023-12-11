package main

import (
	"bufio"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const input6 = "../inputs/d6.txt"

func quadratic_formula(a, b, c float64) (float64, float64) {
	disc := math.Sqrt(math.Pow(b, 2) - 4*a*c)
	if a < 0 {
		return 0.5 * (-b + disc) / a, 0.5 * (-b - disc) / a
	} else {
		return 0.5 * (-b - disc) / a, 0.5 * (-b + disc) / a
	}
}

func d6p1() int {
	file, err := os.Open(input6)
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	t := extract_nums(scanner.Text())
	scanner.Scan()
	d := extract_nums(scanner.Text())
	prod := 1

	for i := 0; i < len(t); i++ {
		r1, r2 := quadratic_formula(-1, float64(t[i]), float64(-d[i]))
		prod *= int(math.Ceil(r2-1) - math.Floor(r1+1) + 1)
	}
	return prod
}

func d6p2() int {
	file, err := os.Open(input6)
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	r := regexp.MustCompile("\\d+")

	scanner.Scan()
	t, _ := strconv.Atoi(strings.Join(r.FindAllString(scanner.Text(), -1), ""))
	scanner.Scan()
	d, _ := strconv.Atoi(strings.Join(r.FindAllString(scanner.Text(), -1), ""))

	r1, r2 := quadratic_formula(-1, float64(t), float64(-d))
	return int(math.Ceil(r2-1) - math.Floor(r1+1) + 1)
}
