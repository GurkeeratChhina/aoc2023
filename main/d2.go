package main

import (
	"bufio"
	"cmp"
	"os"
	"regexp"
	"slices"
	"strconv"

	"go.arsenm.dev/pcre"
)

const input2 = "../inputs/d2.txt"

func maxrgb(line string) (int, int, int) {
	r := pcre.MustCompile("\\d+(?= red)")
	g := pcre.MustCompile("\\d+(?= green)")
	b := pcre.MustCompile("\\d+(?= blue)")

	redstr := r.FindAllString(line, -1)
	greenstr := g.FindAllString(line, -1)
	bluestr := b.FindAllString(line, -1)

	strcomp := func(a, b string) int {
		x, _ := strconv.Atoi(a)
		y, _ := strconv.Atoi(b)
		return cmp.Compare(x, y)
	}

	// fmt.Println(redstr, greenstr, bluestr)

	maxred, _ := strconv.Atoi(slices.MaxFunc(redstr, strcomp))
	maxgreen, _ := strconv.Atoi(slices.MaxFunc(greenstr, strcomp))
	maxblue, _ := strconv.Atoi(slices.MaxFunc(bluestr, strcomp))

	return maxred, maxgreen, maxblue
}

func compare3(x1, x2, x3, m1, m2, m3 int) bool {
	return x1 <= m1 && x2 <= m2 && x3 <= m3
}

func d2p1() int {
	sum := 0

	file, err := os.Open(input2)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	d, _ := regexp.Compile("\\d+")
	for scanner.Scan() {
		line := scanner.Text()
		lineno, _ := strconv.Atoi(d.FindString(line))
		mr, mg, mb := maxrgb(line)
		if compare3(mr, mg, mb, 12, 13, 14) {
			sum += lineno
		}
	}

	return sum
}

func d2p2() int {
	sum := 0

	file, err := os.Open(input2)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		mr, mg, mb := maxrgb(line)
		sum += mr * mg * mb
	}

	return sum
}
