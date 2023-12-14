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

const input12 = "../inputs/d12.txt"

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

func trim_picross(data *picross) bool {
	data.sequence = strings.Trim(data.sequence, ".")
	if len(data.nums) == 0 {
		return true
	}
	if slice_sum(data.nums)+len(data.nums)-1 > len(strings.Split(data.sequence, "")) {
		return false
	}
	if slice_sum(data.nums)+len(data.nums)-1 == len(strings.Split(data.sequence, "")) {
		return true
	}

	//trim #'s
	res := false
	data.sequence, res = strings.CutPrefix(data.sequence, "#")
	if res == true {
		for i, char := range strings.Split(data.sequence, "") {
			if i >= data.nums[0] {
				break
			} else if i == data.nums[0]-1 && char == "#" {
				return false
			} else if i < data.nums[0]-1 && char == "." {
				return false
			}
		}
		data.sequence = data.sequence[data.nums[0]:]
		data.nums = data.nums[1:]
		if !trim_picross(data) {
			return false
		}
	}
	if len(data.nums) == 0 {
		return true
	}

	//trim first section if its too short to fit first num
	chunks := strings.Split(data.sequence, ".")
	if len(strings.Split(chunks[0], "")) < data.nums[0] {
		for _, c := range strings.Split(chunks[0], "") {
			if c == "#" {
				return false
			}
		}
		data.sequence = strings.Join(chunks[1:], ".")
		if !trim_picross(data) {
			return false
		}
	}
	if len(data.nums) == 0 {
		return true
	}

	chunks = strings.Split(data.sequence, ".")
	first_chunk := chunks[0]
	dist_to_hash := len(strings.Split(strings.Split(first_chunk, "#")[0], ""))

	// grow first hash, and then trim if it's big enough
	if dist_to_hash < data.nums[0]+1 {
		s := strings.Split(first_chunk, "")
		len_of_hashes := 0
		// grow and count length
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
		if len_of_hashes > data.nums[0] {
			return false
		}
		//trim if length is equal to specified amount
		if len_of_hashes == data.nums[0] {
			s = s[dist_to_hash+data.nums[0]:]
			if len(s) > 0 {
				s = s[1:]
			}
			chunks[0] = strings.Join(s, "")
			data.sequence = strings.Join(chunks, ".")
			data.nums = data.nums[1:]
			if !trim_picross(data) {
				return false
			}
		}
	}
	return true
}

func flip_picross(data *picross) {
	data.sequence = Reverse(data.sequence)
	slices.Reverse(data.nums)
}

func reduce_picross(pic *picross) bool {
	if trim_picross(pic) {
		flip_picross(pic)
		if trim_picross(pic) {
			flip_picross(pic)
			return true
		} else {
			flip_picross(pic)
			return false
		}
	}
	return false
}

func picross_recurs_dp(data *picross, saved map[*picross]int) int {
	r := regexp.MustCompile("\\#")
	for key, val := range saved {
		if picross_isEqual(data, key) {
			return val
		}
	}
	if len(data.nums) == 0 {
		for _, char := range strings.Split(data.sequence, "") {
			if char == "#" {
				saved[data] = 0
				return 0
			}
		}
		saved[data] = 1
		return 1
	} else {
		sections := strings.Split(data.sequence, ".")
		first_sec := sections[0]
		len_first := len(strings.Split(sections[0], ""))
		if len(sections) == 1 {
			answer := 0
			if r.MatchString(first_sec) {
				answer = brute_force_picross(first_sec, data.nums)
			} else {
				answer = count_stars_and_bars(len_first, data.nums)
			}
			saved[data] = answer
			return answer
		} else {
			sum := 0
			for i := 0; i <= len(data.nums); i++ {
				beginning := picross{sequence: first_sec, nums: append([]int(nil), data.nums[:i]...)}
				leftover := picross{sequence: strings.Join(sections[1:], "."), nums: append([]int(nil), data.nums[i:]...)}
				if reduce_picross(&beginning) && reduce_picross(&leftover) {
					ans := picross_recurs_dp(&beginning, saved) * picross_recurs_dp(&leftover, saved)
					sum += ans
				}
			}
			saved[data] = sum
			return sum
		}
	}
}

func count_stars_and_bars(length int, sections []int) int {
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

func build_stars_and_bars(combination []int, max int) (res []int) {
	new_c := append([]int{-1}, combination...)
	new_c = append(new_c, max)
	for i := 0; i < len(new_c)-1; i++ {
		res = append(res, new_c[i+1]-new_c[i]-1)
	}
	return
}

func springs_include(seq1, seq2 string) bool {
	chars1 := strings.Split(seq1, "")
	chars2 := strings.Split(seq2, "")
	if len(chars1) != len(chars2) {
		panic("sequences have different lengths!")
	}
	for i := 0; i < len(chars1); i++ {
		if chars1[i] != "?" && chars1[i] != chars2[i] {
			return false
		}
	}
	return true
}

func brute_force_picross(s string, nums []int) int {
	seq_len := len(strings.Split(s, ""))
	leftover := seq_len - slice_sum(nums) - len(nums) + 1
	if leftover < 0 {
		return 0
	} else {
		count := 0
		for _, combin := range combin.Combinations(leftover+len(nums), len(nums)) {
			buckets := build_stars_and_bars(combin, leftover+len(nums))
			var newseq string

			for k := 0; k < len(nums); k++ {
				dots := buckets[k]
				if k > 0 {
					dots++
				}
				newseq = newseq + strings.Repeat(".", dots) + strings.Repeat("#", nums[k])
			}
			newseq = newseq + strings.Repeat(".", buckets[len(nums)])

			if springs_include(s, newseq) {
				count++
			}
		}
		return count
	}
}

func get_seq_counts(line string, copies int) picross {
	a := strings.Split(line, " ")
	r := regexp.MustCompile("\\.+")
	seq := r.ReplaceAllString(a[0], ".")
	counts := extract_nums(a[1])
	for i := 1; i < copies; i++ {
		seq = seq + "?" + r.ReplaceAllString(a[0], ".")
		counts = append(counts, extract_nums(a[1])...)
	}
	return picross{sequence: seq, nums: counts}
}

func d12(amount int) int {
	f, err := os.Open(inputtest)
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	sum := 0
	for scanner.Scan() {
		answer_cache := make(map[*picross]int)
		data := get_seq_counts(scanner.Text(), amount)

		// fmt.Println(data)
		reduce_picross(&data)
		// fmt.Println(data)
		ans := picross_recurs_dp(&data, answer_cache)
		fmt.Println(ans)
		sum += ans
	}
	return sum
}

func d12p1() int {
	return d12(1)
}

func d12p2() int {
	return d12(5)
}
