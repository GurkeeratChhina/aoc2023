package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"
)

type picross struct {
	sequence string
	nums     []int
}

func trim_picross(data *picross) {
	if slice_sum(data.nums)+len(data.nums)-1 == len(strings.Split(data.sequence, "")) {
		return
	}
	res := false
	data.sequence = strings.Trim(data.sequence, ".")
	data.sequence, res = strings.CutPrefix(data.sequence, "#")
	if res == true {
		data.sequence = data.sequence[data.nums[0]:]
		data.nums = data.nums[1:]
		trim_picross(data)
	}
	chunks := strings.Split(data.sequence, ".")
	if len(strings.Split(chunks[0], "")) < data.nums[0] {
		data.sequence = strings.Join(chunks[1:], "")
		trim_picross(data)
	}
	chunks = strings.Split(data.sequence, ".")
	first_chunk := chunks[0]
	dist_to_hash := len(strings.Split(strings.Split(first_chunk, "#")[0], ""))
	if dist_to_hash < data.nums[0]+1 {
		s := strings.Split(first_chunk, "")
		len_of_hashes := 0
		for i := dist_to_hash; i < len(s); i++ {
			if i < data.nums[0] {
				s[i] = "#"
			}
			if s[i] == "#" {
				len_of_hashes++
			} else {
				break
			}
		}
		if len_of_hashes == data.nums[0] {
			s = s[dist_to_hash+data.nums[0]:]
			if len(s) > 0 {
				s = s[1:]
			}
			chunks[0] = strings.Join(s, "")
			data.sequence = strings.Join(chunks, "")
			data.nums = data.nums[1:]
			trim_picross(data)
		}
	}
}

func flip_picross(data *picross) {
	data.sequence = Reverse(data.sequence)
	slices.Reverse(data.nums)
}

func reduce_picross(data *picross) {
	trim_picross(data)
	flip_picross(data)
	trim_picross(data)
}

func brute_force_picross(data *picross) int {
	reduce_picross(data)
	r := regexp.MustCompile("\\#")
	if slice_sum(data.nums)+len(data.nums)-1 == len(strings.Split(data.sequence, "")) {
		return 1
	} else if !r.MatchString(data.sequence) {
		return 0
	} else {
		return 0
	}
}

func get_seq_counts(line string) picross {
	a := strings.Split(line, " ")
	r := regexp.MustCompile("\\.+")
	seq := r.ReplaceAllString(a[0], ".")
	counts := extract_nums(a[1])
	return picross{sequence: seq, nums: counts}
}

func d12p1() int {
	f, err := os.Open(inputtest)
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		data := get_seq_counts(scanner.Text())
		fmt.Println(data)
		reduce_picross(&data)
		fmt.Println(data)
	}

	return 0
}
