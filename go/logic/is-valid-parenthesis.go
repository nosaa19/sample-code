package logic

func IsValidParenthesis(input string) bool {

	if len(input)%2 != 0 {
		return false
	}

	stack := make([]rune, 0)
	mapPair := map[rune]rune{
		'[': ']',
		'{': '}',
		'(': ')',
	}

	for _, char := range input {
		if closing, ok := mapPair[char]; ok {
			stack = append(stack, closing)
			continue
		}

		lastIndex := len(stack) - 1

		// check current char same with latest char on the stack
		if lastIndex < 0 || stack[lastIndex] != char {
			return false
		}

		// remove the latest element, meaning LIFO
		stack = stack[:lastIndex]
	}

	// last step, check if stack empty
	return len(stack) == 0
}
