package logic

import (
	"fmt"
	"regexp"
	"strings"
)

/*
	Have the function Calculator(str) take the str parameter being passed
	and evaluate the mathematical expression within in.

	For example, if str were "2+(3-1)*3" the output should be 8.
	Another example: if str were "(2-0)(6/2)" the output should be 6.

	There can be parenthesis within the string, so you must evaluate it
	properly according to the rules of arithmetic.

	The string will contain the operators: +, -, /, *, (, and ).
	If you have a string like this: #/#*# or #+#(#)/#,
	then evaluate from left to right. So divide then multiply,
	and for the second one multiply, divide, then add.

	The evaluations will be such that there will not be any decimal operations,
	so you do not need to account for rounding and whatnot.
*/

func Calculate(str string) (result string) {

	tokens := PreProcessing(str)

	fmt.Println(tokens)

	return
}

/*
Cleaning, parsing and splitting the input string into the array of tokens

1. convert to lowercase
2. replace (a)(b) with (a) * (b)
3. replace a(b) with a * (b)
4. split two bundled characters (but only when the second character is not a digit)
5. split bundled command/operator and a digit. Excl. minus to keep negative numbers intact.
6. remove multiple spaces
7. trim (remove leading and trailing spaces)
8. split to array of strings

param input input string
return array of strings with parsed operators and operands
*/
func PreProcessing(input string) (output []string) {

	// 1.
	text := strings.ToLower(input)

	// 2.
	r := regexp.MustCompile(`(\)\()`)
	text = r.ReplaceAllString(text, ") * (")

	// 3.
	r = regexp.MustCompile(`([\d+])\(`)
	text = r.ReplaceAllString(text, "$1 * (")

	// 4.

	// 5.

	// 6.

	// 7.

	// 8.
	output = strings.Split(text, " ")
	return
}

// https://github.com/ardallie/coderbyte-java/blob/master/src/hard/Calculator.java
