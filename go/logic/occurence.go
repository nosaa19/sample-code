package logic

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func ShowStatistics(array []int) {
	stats := make(map[int]int)
	keys := make([]string, 0)

	for _, number := range array {
		if count, found := stats[number]; found {
			count++
			stats[number] = count
		} else {
			count = 1
			stats[number] = count
			keys = append(keys, strconv.Itoa(number))
		}
	}

	sort.Strings(keys)

	fmt.Println("The statistics of", array)
	for _, keyStr := range keys {
		key, _ := strconv.Atoi(keyStr)
		fmt.Println(key, ":", stats[key])
	}
}

func MaxOccurenceCharOfSubstring(str string) (result string) {

	/*
		input: hello world from France
		output: hello, because 2 l's

		input: strong and weak
		output: -1, all subtring contain unique char
	*/

	chunk := strings.Split(str, " ")

	result = "-1"
	maxCount := 0

	for i := 0; i < len(chunk); i++ {
		currentMax := MaxFrequency(chunk[i])
		fmt.Println("chunk[i]", currentMax)
		if currentMax > 1 && currentMax > maxCount {
			maxCount = currentMax
			result = chunk[i]
		}
	}
	return
}

func MaxFrequency(str string) (result int) {
	stats := make(map[rune]int, len(str))
	maxCount := 0

	for _, char := range str {
		if count, found := stats[char]; !found {
			stats[char], maxCount = 1, 1
		} else {
			count += 1
			stats[char] = count
			if count > maxCount {
				maxCount = count
			}
		}
	}
	result = maxCount
	return
}
