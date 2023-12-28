package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/twmb/algoimpl/go/graph"
)

const input25 = "../inputs/d25.txt"

func d25p1() int {
	defer timeTrack(time.Now(), "d25p1")
	f, err := os.Open(input25)
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	g := graph.New(graph.Undirected)
	nodes := make(map[string]graph.Node)

	// build graph
	for scanner.Scan() {
		src := strings.Split(scanner.Text(), ": ")[0]
		dsts := strings.Split(strings.Split(scanner.Text(), ": ")[1], " ")
		if _, ok := nodes[src]; !ok {
			nodes[src] = g.MakeNode()
		}
		for _, dst := range dsts {
			if _, ok := nodes[dst]; !ok {
				nodes[dst] = g.MakeNode()
			}
			g.MakeEdge(nodes[src], nodes[dst])
		}
	}

	var mincut []graph.Edge
	for len(mincut) != 3 {
		mincut = g.RandMinimumCut(1, 1)
	}

	for _, e := range mincut {
		g.RemoveEdge(e.Start, e.End)
	}

	comps := g.StronglyConnectedComponents()

	return len(comps[0]) * len(comps[1])
}

func d25p2() int {
	fmt.Println("All done!")
	return 0
}
