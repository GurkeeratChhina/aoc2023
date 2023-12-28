package main

import (
	"bufio"
	"os"
	"slices"
	"time"
)

const input23 = "../inputs/d23.txt"

type node struct {
	value          int
	incoming_edges []*edge
	outgoing_edges []*edge
}

type edge struct {
	value  int
	source *node
	dest   *node
}

type mygraph struct {
	nodes []*node
	edges []*edge
}

type node_locs map[int]map[int]*node

type maze [][]rune

var AtoP = map[rune]point{'^': {-1, 0}, '>': {0, 1}, 'v': {1, 0}, '<': {0, -1}}

func (m maze) walk(current, direction point) (int, point) {
	steps := 1
	for {
		current.x += direction.x
		current.y += direction.y
		steps++
		if m[current.x][current.y] != '.' {
			break
		}
		// reach end of maze
		if current.x == len(m)-1 {
			return steps, current
		}
		// check if need to change direction
		if m[current.x+direction.x][current.y+direction.y] == '#' {
			if direction.x == 0 { // going horizontal
				if m[current.x+1][current.y] == '#' { // wall below
					direction.x = -1
				} else {
					direction.x = 1
				}
				direction.y = 0
			} else { // going vertical
				if m[current.x][current.y+1] == '#' { // wall to the right
					direction.y = -1
				} else {
					direction.y = 1
				}
				direction.x = 0
			}
		}

	}
	arrow := m[current.x][current.y]
	current.x += AtoP[arrow].x
	current.y += AtoP[arrow].y
	return steps + 1, current
}

func (g *mygraph) max_path(end *node) int {
	for i := 0; i < len(g.nodes); i++ {
		for _, e := range g.edges {
			e.dest.value = max(e.dest.value, e.source.value+e.value)
		}
	}
	return end.value
}

func create_node(val int, g *mygraph, locs node_locs, location point) {
	if _, ok := locs[location.x][location.y]; ok { // node already exists at point
		return
	}
	if _, ok := locs[location.x]; !ok {
		locs[location.x] = make(map[int]*node)
	}
	n := &node{value: val}
	locs[location.x][location.y] = n
	g.nodes = append(g.nodes, n)
}

func (g *mygraph) create_nodes(locs node_locs, m maze) {
	for i, row := range m {
		for j, symbol := range row {
			if symbol == 'v' || symbol == '^' {
				if m[i+1][j+1] == '<' || m[i+1][j+1] == '>' || m[i+1][j-1] == '<' || m[i+1][j-1] == '>' {
					create_node(1, g, locs, point{i + 1, j})
				} else if m[i-1][j+1] == '<' || m[i-1][j+1] == '>' || m[i-1][j-1] == '<' || m[i-1][j-1] == '>' {
					create_node(1, g, locs, point{i - 1, j})
				}
			}
		}
	}
	create_node(0, g, locs, point{0, 1})
	create_node(1, g, locs, point{len(m) - 1, len(m[0]) - 2})
}

func create_edge(src *node, g *mygraph, m maze, locs node_locs, start point, dir point) {
	steps, end_point := m.walk(start, dir)
	end_node := locs[end_point.x][end_point.y]
	newedge := &edge{value: steps, source: src, dest: end_node}
	src.outgoing_edges = append(src.outgoing_edges, newedge)
	end_node.incoming_edges = append(end_node.incoming_edges, newedge)
	g.edges = append(g.edges, newedge)
}

func (g *mygraph) create_edges(locs node_locs, m maze) {
	for i, val1 := range locs {
		if i == 0 { // start node
			create_edge(val1[1], g, m, locs, point{0, 1}, point{1, 0})
		} else if i == len(m)-1 { // end node, don't make any edges starting at it
		} else {
			for j, n := range val1 {
				if m[i-1][j] == '^' {
					create_edge(n, g, m, locs, point{i - 1, j}, AtoP[m[i-1][j]])
				}
				if m[i+1][j] == 'v' {
					create_edge(n, g, m, locs, point{i + 1, j}, AtoP[m[i+1][j]])
				}
				if m[i][j-1] == '<' {
					create_edge(n, g, m, locs, point{i, j - 1}, AtoP[m[i][j-1]])
				}
				if m[i][j+1] == '>' {
					create_edge(n, g, m, locs, point{i, j + 1}, AtoP[m[i][j+1]])
				}
			}
		}

	}
}

func (g *mygraph) max_undirected_simple(current *node, end *node, visited []*node) (bool, int) {
	max_so_far := current.value
	if current == end {
		// fmt.Println("found end")
		return true, max_so_far
	}
	for _, e := range current.incoming_edges {
		if !slices.Contains(visited, e.source) {
			valid, max_len := g.max_undirected_simple(e.source, end, append(visited, current))
			if valid {
				// fmt.Println("src", current.value, max_len, e.value)
				max_so_far = max(max_so_far, max_len+e.value)
			}
		}
	}
	for _, e := range current.outgoing_edges {
		if !slices.Contains(visited, e.dest) {
			valid, max_len := g.max_undirected_simple(e.dest, end, append(visited, current))
			if valid {
				// fmt.Println("dst", current.value, max_len, e.value, max_so_far)
				max_so_far = max(max_so_far, max_len+e.value)
			}
		}
	}
	if max_so_far > current.value {
		// fmt.Println("max that is returned", max_so_far)
		return true, max_so_far
	} else {
		return false, 0
	}
}

func d23p1() int {
	defer timeTrack(time.Now(), "d23p1")
	f, err := os.Open(input23)
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var m maze
	for scanner.Scan() {
		m = append(m, []rune(scanner.Text()))
	}

	locs := make(node_locs)
	g := mygraph{}
	g.create_nodes(locs, m)
	g.create_edges(locs, m)

	g.max_path(locs[len(m)-1][len(m[0])-2])
	// fmt.Println("nodes of g")
	// for _, n := range g.nodes {
	// 	fmt.Println(n)
	// }
	// fmt.Println(locs)

	return g.max_path(locs[len(m)-1][len(m[0])-2]) - 1
}

func d23p2() int {
	defer timeTrack(time.Now(), "d23p2")
	f, err := os.Open(input23)
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var m maze
	for scanner.Scan() {
		m = append(m, []rune(scanner.Text()))
	}

	locs := make(node_locs)
	g := mygraph{}
	g.create_nodes(locs, m)
	g.create_edges(locs, m)

	_, ans := g.max_undirected_simple(locs[0][1], locs[len(m)-1][len(m[0])-2], []*node{})
	return ans - 2
}
