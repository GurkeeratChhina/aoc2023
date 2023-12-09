package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
)

const input = "../inputs/d1.txt"

func d1p1() int {
	sum := 0

	file, err := os.Open(input)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		r, _ := regexp.Compile("[^\\d]")
		digits := r.ReplaceAllString(scanner.Text(), "")
		sum += 10*int(digits[0]-'0') + int(digits[len(digits)-1]-'0')
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return sum
}

func d1p2() int {
	return 0
}
