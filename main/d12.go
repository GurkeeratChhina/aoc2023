package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"

	"gonum.org/v1/gonum/stat/combin"
)

type picross struct {
	sequence string
	nums     []int
}

func picross_isEqual(data1, data2 *picross) bool {
	if data1.sequence == data2.sequence && len(data1.nums) == len(data2.nums) {
		for i, _ := range data1.nums {
			if data1.nums[i] != data2.nums[i] {
				return false
			}
		}
		return true
	}
	return false
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

func picross_recurs_dp(data *picross, saved map[*picross]int) int {
	r := regexp.MustCompile("\\#")
	for key, val := range saved {
		if picross_isEqual(data, key) {
			return val
		}
	}
	if slice_sum(data.nums)+len(data.nums)-1 == len(strings.Split(data.sequence, "")) {
		return 1
	} else {
		sections := strings.Split(data.sequence, ".")
		first_sec := sections[0]
		len_first := len(strings.Split(sections[0], ""))
		if !r.MatchString(first_sec) {
			if len(sections) == 1 {
				answer := stars_and_bars(len_first, data.nums)
				return answer
			} else {
				sum := 0
				for i := 0; i < len(data.nums); i++ {
					leftover := picross{sequence: strings.Join(sections[1:], ""), nums: data.nums[i:]}
					sum += stars_and_bars(len_first, data.nums[:i]) * picross_recurs_dp(&leftover, saved)
				}
				saved[data] = sum
				return sum
			}
		} else {
			panic("case not implemented yet")
		}
	}
}

func picross_recurs_dp2(data *picross, saved map[*picross]int) int {
	r := regexp.MustCompile("\\#")
	for key, val := range saved {
		if picross_isEqual(data, key) {
			return val
		}
	}
	if slice_sum(data.nums)+len(data.nums)-1 == len(strings.Split(data.sequence, "")) {
		return 1
	} else if !r.MatchString(data.sequence) {
		sections := strings.Split(data.sequence, ".")
		len_first := len(strings.Split(sections[0], ""))
		if len(sections) == 1 {
			answer := stars_and_bars(len_first, data.nums)
			return answer
		} else {
			sum := 0
			for i := 0; i < len(data.nums); i++ {
				leftover := picross{sequence: strings.Join(sections[1:], ""), nums: data.nums[i:]}
				sum += stars_and_bars(len_first, data.nums[:i]) * picross_recurs_dp(&leftover, saved)
			}
			saved[data] = sum
			return sum
		}
	} else {
		panic("case not implemented yet")
	}
}

func stars_and_bars(length int, sections []int) int {
	stars := length - slice_sum(sections) - len(sections) + 1
	bars := len(sections)
	if stars == 0 {
		return 1
	} else if stars > 0 && bars == 0 {
		return 1
	} else if stars < 0 {
		return 0
	} else {
		return combin.Binomial(stars+bars, bars)
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
	sum := 0
	answer_cache := make(map[*picross]int)

	for scanner.Scan() {
		data := get_seq_counts(scanner.Text())

		fmt.Println(data)
		reduce_picross(&data)
		fmt.Println(data)
		ans := picross_recurs_dp(&data, answer_cache)
		fmt.Println(ans)
		sum += ans
	}

	return sum
}
