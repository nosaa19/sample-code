package logic

import (
	"fmt"
	"sort"
	"strconv"
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
