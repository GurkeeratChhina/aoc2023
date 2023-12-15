package main

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

const input15 = "../inputs/d15.txt"

type Lens struct {
	label string
	focus int
}

type HASHMAP map[int][]Lens

func HASH(word string) (sum int) {
	for _, letter := range []rune(word) {
		sum = (sum + int(letter)) % 256
		sum = (sum * 17) % 256
	}
	return
}

func (m HASHMAP) remove(label string) {
	bin := m[HASH(label)]
	for i, lens := range bin {
		if lens.label == label {
			m[HASH(label)] = append(bin[:i], bin[i+1:]...)
			return
		}
	}
}

func (m HASHMAP) add(new Lens) {
	bin := m[HASH(new.label)]
	for i, lens := range bin {
		if lens.label == new.label {
			bin[i] = new
			m[HASH(new.label)] = bin
			return
		}
	}
	m[HASH(new.label)] = append(bin, new)
}

func d15p1() int {
	f, err := os.Open(input15)
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	sum := 0
	scanner.Scan()
	for _, word := range strings.Split(scanner.Text(), ",") {
		sum += HASH(word)
	}
	return sum
}

func d15p2() int {
	f, err := os.Open(input15)
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	letters := regexp.MustCompile("[a-z]+")
	hashmap := make(HASHMAP)
	scanner.Scan()
	for _, instruction := range strings.Split(scanner.Text(), ",") {
		label := letters.FindString(instruction)
		focus := extract_nums(instruction)
		if len(focus) == 0 {
			// instruction is a -
			hashmap.remove(label)
		} else if len(focus) == 1 {
			// instruction is a =
			hashmap.add(Lens{label: label, focus: focus[0]})
		} else {
			panic("bad instruction")
		}
	}
	// fmt.Println(hashmap)

	sum := 0
	for key, lenses := range hashmap {
		for slot, lens := range lenses {
			sum += (key + 1) * (slot + 1) * lens.focus
		}
	}
	return sum
}
