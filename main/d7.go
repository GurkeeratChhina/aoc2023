package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

const input7 = "../inputs/d7.txt"

func rank_from_count(count map[rune]int) string {
	if len(count) == 5 {
		return "1"
	} else if len(count) == 4 {
		return "2"
	} else if len(count) == 3 {
		twos := 0
		for _, value := range count {
			if value == 2 {
				twos += 1
			}
		}
		if twos == 2 {
			return "3"
		}
		return "4"
	} else if len(count) == 2 {
		for _, value := range count {
			if value == 1 {
				return "6"
			}
		}
		return "5"
	} else if len(count) == 1 {
		return "7"
	} else {
		panic("hand has too many or too few cards")
	}
}

func rank_hand(hand string) string {
	card_count := make(map[rune]int)
	for _, char := range hand {
		card_count[char]++
	}
	return rank_from_count(card_count)
}

func rank_hand_joker(hand string) string {
	card_count := make(map[rune]int)
	for _, char := range hand {
		card_count[char]++
	}
	jokers := card_count['J']
	if jokers < 4 {
		delete(card_count, 'J')
		return rank_from_count(card_count)
	} else {
		return "7"
	}
}

func hand_less(hand1, hand2 []string, card_value map[rune]int) bool {
	a, _ := strconv.Atoi(hand1[2])
	b, _ := strconv.Atoi(hand2[2])
	if a < b {
		return true
	} else if a > b {
		return false
	} else if b == a {
		hand1chars := []rune(hand1[0])
		hand2chars := []rune(hand2[0])
		for i := 0; i < 5; i++ {
			if card_value[hand1chars[i]] < card_value[hand2chars[i]] {
				return true
			} else if card_value[hand1chars[i]] > card_value[hand2chars[i]] {
				return false
			}
		}
	}
	panic("two hands are equal")
}

func eval_hands(card_value map[rune]int, hand_ranking func(s string) string) int {
	file, err := os.Open(input7)
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var hands [][]string
	for scanner.Scan() {
		a := strings.Split(scanner.Text(), " ")
		a = append(a, hand_ranking(a[0]))
		hands = append(hands, a)
	}

	sort.Slice(hands, func(i, j int) bool { return hand_less(hands[i], hands[j], card_value) })
	sum := 0
	for i, hand := range hands {
		bet, _ := strconv.Atoi(hand[1])
		sum += (i + 1) * bet
	}
	return sum
}

func d7p1() int {
	card_values := map[rune]int{'2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9, 'T': 10, 'J': 11, 'Q': 12, 'K': 13, 'A': 14}
	return eval_hands(card_values, rank_hand)
}

func d7p2() int {
	card_values := map[rune]int{'2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9, 'T': 10, 'J': 1, 'Q': 12, 'K': 13, 'A': 14}
	return eval_hands(card_values, rank_hand_joker)
}
