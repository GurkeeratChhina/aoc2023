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
	fmt.Println("Part 1: ", d2p1())
	fmt.Println("Part 2: ", d2p2())
}
