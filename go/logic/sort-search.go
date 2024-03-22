package logic

import (
	"fmt"
	"math"
	"strconv"
)

// https://medium.com/@pvivek4/popular-interview-questions-on-time-and-space-complexity-10606cbefef1

func BubbleSort(array []int) {
	for i := 0; i < len(array); i++ { // O(n)
		for j := 0; j < len(array)-1; j++ { // O(n)
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
}

func BinarySearch(array []int, input int) int {
	var (
		start = 0
		end   = len(array) - 1
	)
	for start <= end {
		mid := int(math.Floor(float64(start+end) / 2))
		if array[mid] == input {
			return mid
		} else if array[mid] < input {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}

func Print3largest(array []int) {
	var (
		i, first, second, third int
	)

	/* There should be atleast three elements */
	if len(array) < 3 {
		fmt.Println(" Invalid Input ")
		return
	}

	first, second, third = math.MinInt, math.MinInt, math.MinInt
	for i = 0; i < len(array); i++ {
		/* If current element is greater than
		   first*/
		if array[i] > first {
			third = second
			second = first
			first = array[i]
		} else if array[i] > second {
			/* If arr[i] is in between first and
			   second then update second  */
			third = second
			second = array[i]
		} else if array[i] > third {
			third = array[i]
		}
	}

	fmt.Println("Three largest elements are " + strconv.Itoa(first) + ", " + strconv.Itoa(second) + ", " + strconv.Itoa(third))
}
