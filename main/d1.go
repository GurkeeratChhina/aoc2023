package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
)

const input = "../inputs/d1.txt"

func digit_word_to_int(x string) int {
	switch {
	case len(x) == 1:
		return int(x[0] - '0')
	case x == "one":
		return 1
	case x == "two":
		return 2
	case x == "three":
		return 3
	case x == "four":
		return 4
	case x == "five":
		return 5
	case x == "six":
		return 6
	case x == "seven":
		return 7
	case x == "eight":
		return 8
	case x == "nine":
		return 9
	}
	log.Fatal("No Match")
	return 0
}

func d1(regexinput string) int {
	sum := 0

	file, err := os.Open(input)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		r, _ := regexp.Compile(regexinput)
		first := r.FindString(line)
		i := len(line) - 1
		for !r.MatchString(line[i:]) {
			i--
		}
		last := r.FindString(line[i:])
		sum += 10*digit_word_to_int(first) + digit_word_to_int(last)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return sum
}

func d1p1() int {
	return d1("\\d")
}

func d1p2() int {
	return d1("\\d|one|two|three|four|five|six|seven|eight|nine")
}
