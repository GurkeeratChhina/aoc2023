package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const input18 = "../inputs/d18.txt"

type point struct {
	x, y int
}

func shoelace(vertex_set []point) int {
	sum := 0
	for i := 0; i < len(vertex_set)-1; i++ {
		sum += (vertex_set[i].y + vertex_set[i+1].y) * (vertex_set[i].x - vertex_set[i+1].x)
	}
	return sum / 2
}

func dirAtoi(char rune) int {
	if char == 'U' {
		return 0
	} else if char == 'R' {
		return 1
	} else if char == 'D' {
		return 2
	} else if char == 'L' {
		return 3
	} else {
		fmt.Println(char)
		panic("invalid letter")
	}
}

func d18p1() int {
	defer timeTrack(time.Now(), "d18p1")
	f, err := os.Open(input18)
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	perimeter := 0
	current_vertex := point{x: 0, y: 0}
	var vertices []point
	for scanner.Scan() {
		direction := dirAtoi([]rune(scanner.Text())[0])
		amount := extract_nums(scanner.Text())[0]
		vertices = append(vertices, current_vertex)
		current_vertex.y += direction_table[direction][0] * amount
		current_vertex.x += direction_table[direction][1] * amount
		perimeter += amount
	}

	return shoelace(vertices) + perimeter/2 + 1
}

func d18p2() int {
	defer timeTrack(time.Now(), "d18p2")
	f, err := os.Open(input18)
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	perimeter := 0
	current_vertex := point{x: 0, y: 0}
	var vertices []point
	for scanner.Scan() {
		hex := strings.Split(scanner.Text(), "#")[1]
		direction := int(hex[5]) - int('0')
		amount, err := strconv.ParseInt(hex[:5], 16, 0)
		if err != nil {
			panic(err)
		}
		vertices = append(vertices, current_vertex)
		current_vertex.y += direction_table[direction][0] * int(amount)
		current_vertex.x += direction_table[direction][1] * int(amount)
		perimeter += int(amount)
	}

	return shoelace(vertices) + perimeter/2 + 1
}
