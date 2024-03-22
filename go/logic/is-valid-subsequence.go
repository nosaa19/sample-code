package logic

func IsValidSubsequence(array []int, sequence []int) bool {

	var (
		seqIndex int  = 0
		isValid  bool = false
	)

	for _, num := range array {
		if num == sequence[seqIndex] {
			seqIndex++
		}
	}

	if seqIndex == len(sequence) {
		isValid = true
	}

	return isValid
}
