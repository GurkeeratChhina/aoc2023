package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const input20 = "../inputs/d20.txt"

type pulse struct {
	signal bool
	src    string
	dest   string
}

type flipflip struct {
	state   bool
	outputs []string
}

type conjunction struct {
	inputs  map[string]bool
	outputs []string
}

type broadcaster struct {
	outputs []string
}

type module interface {
	get_outputs() []string
	process(input string, signal bool) string
	add_input(inp string)
}

type pulse_queue []pulse

func (t *flipflip) process(input string, signal bool) string {
	if !signal {
		t.state = !t.state
		if t.state {
			return "true"
		} else {
			return "false"
		}
	} else {
		return "none"
	}
}

func (t *flipflip) get_outputs() []string {
	return t.outputs
}

func (t *flipflip) add_input(inp string) {
}

func (c *conjunction) process(input string, signal bool) string {
	c.inputs[input] = signal
	for _, val := range c.inputs {
		if !val {
			return "true"
		}
	}
	return "false"
}

func (c *conjunction) get_outputs() []string {
	return c.outputs
}

func (c *conjunction) add_input(inp string) {
	c.inputs[inp] = false
}

func (b *broadcaster) process(input string, signal bool) string {
	if signal {
		return "true"
	} else {
		return "false"
	}
}

func (b *broadcaster) get_outputs() []string {
	return b.outputs
}

func (b *broadcaster) add_input(inp string) {
}

func str_to_bool(s string) bool {
	if s == "true" {
		return true
	} else {
		return false
	}
}

func process_signal(m module, p pulse, q pulse_queue) pulse_queue {
	res := m.process(p.src, p.signal)
	if res == "none" {
		return q
	}
	for _, output := range m.get_outputs() {
		q = append(q, pulse{signal: str_to_bool(res), src: p.dest, dest: output})
	}
	return q
}

// parse input
func build_module_network(s bufio.Scanner) (network map[string]module) {
	network = make(map[string]module)
	for s.Scan() {
		first_char := []rune(s.Text())[0]
		outs := strings.Split(strings.Split(s.Text(), "-> ")[1], ", ")
		if first_char == '%' {
			// flipflop
			network[s.Text()[1:3]] = &flipflip{state: false, outputs: outs}
		} else if first_char == '&' {
			// conjunction
			network[s.Text()[1:3]] = &conjunction{inputs: make(map[string]bool), outputs: outs}
		} else {
			//broadcaster
			network["broadcaster"] = &broadcaster{outputs: outs}
		}
	}

	for src, mod := range network {
		for _, dest := range mod.get_outputs() {
			val, ok := network[dest]
			if ok {
				val.add_input(src)
			}

		}
	}

	return
}

func d20p1() int {
	defer timeTrack(time.Now(), "d20p1")
	f, err := os.Open(input20)
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	myNetwork := build_module_network(*scanner)
	var q pulse_queue

	lowcount := 0
	highcount := 0
	for i := 0; i < 1000; i++ {
		q = append(q, pulse{signal: false, src: "button", dest: "broadcaster"})
		for {
			if len(q) == 0 {
				break
			}
			current_signal := q[0]
			// fmt.Println(current_signal)
			q = q[1:]
			if current_signal.signal {
				highcount++
			} else {
				lowcount++
			}
			val, ok := myNetwork[current_signal.dest]
			if ok {
				q = process_signal(val, current_signal, q)
			}
		}
	}
	q = append(q, pulse{signal: false, src: "button", dest: "broadcaster"})
	fmt.Println(lowcount, highcount)
	return lowcount * highcount
}

func d20p2() int {
	defer timeTrack(time.Now(), "d20p2")
	f, err := os.Open(input20)
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	myNetwork := build_module_network(*scanner)

	var cycles []int
	for _, start := range myNetwork["broadcaster"].get_outputs() {
		var conj_key string
		for _, output := range myNetwork[start].get_outputs() {
			if string_slice_contains(myNetwork[output].get_outputs(), start) {
				conj_key = output
				break
			}
		}

		multiplier := 1
		cycle_length := 0
		for current_key := start; ; {
			// fmt.Println(current_key)
			outs := myNetwork[current_key].get_outputs()
			if string_slice_contains(outs, conj_key) {
				cycle_length += multiplier
			}
			multiplier *= 2
			for _, out := range outs {
				if out != conj_key {
					current_key = out
					goto end
				}
			}
			break
		end:
		}
		cycles = append(cycles, cycle_length)
	}
	fmt.Println(cycles)

	return LCM_slice(cycles)
}
