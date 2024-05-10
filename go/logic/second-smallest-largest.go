package logic

import (
	"sort"
	"strconv"
)

func SecondSmallestLargest(numbers []int) (result string) {
	if len(numbers) > 1 {
		sort.Ints(numbers)
		result = strconv.Itoa(numbers[1]) + " " + strconv.Itoa(numbers[len(numbers)-2])
	}
	return
}
