package main

import (
	"bufio"
	"fmt"
	"math"
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

func trim_picross(data *picross) {
	if slice_sum(data.nums)+len(data.nums)-1 >= len(strings.Split(data.sequence, "")) {
		return
	}
	if len(data.nums) == 0 {
		return
	}
	fmt.Println("trimming", data)
	res := false
	data.sequence = strings.Trim(data.sequence, ".")

	//trim #'s and numbers
	data.sequence, res = strings.CutPrefix(data.sequence, "#")
	if res == true {
		data.sequence = data.sequence[data.nums[0]:]
		data.nums = data.nums[1:]
		fmt.Println("trimmed #")
		trim_picross(data)
	}
	if len(data.nums) == 0 {
		return
	}
	//trim first section if its too short to fit first num
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
	flip_picross(data)
}

func picross_recurs_dp(data *picross, saved map[*picross]int) int {
	fmt.Println("recurs call", data.sequence, data.nums)
	r := regexp.MustCompile("\\#")
	for key, val := range saved {
		// fmt.Println("found saved")
		if picross_isEqual(data, key) {
			return val
		}
	}
	if len(data.nums) == 0 {
		fmt.Println("no nums")
		count := 0
		for _, char := range strings.Split(data.sequence, "") {
			if char == "?" {
				count++
			} else if char == "#" {
				saved[data] = 0
				return 0
			}
		}
		ans := int(math.Pow(2, float64(count)))
		saved[data] = ans
		return ans
	} else if slice_sum(data.nums)+len(data.nums)-1 == len(strings.Split(data.sequence, "")) {
		fmt.Println("exact length")
		saved[data] = 1
		return 1
	} else {
		sections := strings.Split(data.sequence, ".")
		first_sec := sections[0]
		len_first := len(strings.Split(sections[0], ""))
		if len(sections) == 1 {
			fmt.Println("one section")
			answer := 0
			if r.MatchString(first_sec) {
				fmt.Println("bruteforcing")
				answer = brute_force_picross(first_sec, data.nums)
			} else {
				fmt.Println("combinatorics")
				answer = stars_and_bars(len_first, data.nums)
			}
			saved[data] = answer
			return answer
		} else {
			fmt.Println("multiple sections", sections)
			sum := 0
			for i := 0; i < len(data.nums); i++ {
				beginning := picross{sequence: first_sec, nums: data.nums[:i]}
				leftover := picross{sequence: strings.Join(sections[1:], "."), nums: data.nums[i:]}
				fmt.Println("beginning", beginning, "leftover", leftover)
				reduce_picross(&beginning)
				reduce_picross(&leftover)
				ans := picross_recurs_dp(&beginning, saved) * picross_recurs_dp(&leftover, saved)
				fmt.Println(beginning, leftover, "ans", ans)
				sum += ans
			}
			saved[data] = sum
			return sum
		}
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

func springs_include(seq1, seq2 string) bool {
	chars1 := strings.Split(seq1, "")
	chars2 := strings.Split(seq2, "")
	if len(chars1) != len(chars2) {
		fmt.Println(seq1, seq2)
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
	} else if leftover == 0 {
		return 1
	} else {
		count := 0
		buckets := []int{leftover}
		for i := 0; i < len(nums); i++ {
			buckets = append(buckets, 0)
		}
		// fmt.Println(s, nums, buckets)
		i, j := 0, 0
		for {
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
			if buckets[len(nums)] == leftover {
				break
			} else {
				// advance loop
				if j == len(nums) {
					buckets[j]--
					buckets[i+1]++
					j = i
				}
				buckets[j]--
				buckets[j+1]++
				if buckets[j] == 0 {
					i++
				}
				j++
			}
		}
		return count
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
	f, err := os.Open(input12)
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	sum := 0
	answer_cache := make(map[*picross]int)

	for scanner.Scan() {
		data := get_seq_counts(scanner.Text())

		// fmt.Println(data)
		reduce_picross(&data)
		// fmt.Println(data)
		ans := picross_recurs_dp(&data, answer_cache)
		fmt.Println(ans)
		sum += ans
	}

	return sum
}
