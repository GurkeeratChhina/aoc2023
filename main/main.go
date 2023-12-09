package main

import (
	"fmt"
	"log"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	fmt.Println("Part 1: ", d1p1())
	fmt.Println("Part 2: ", d1p2())
}
