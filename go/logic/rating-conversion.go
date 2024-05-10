package logic

import (
	"strconv"
	"strings"
)

const (
	fullStr  = "full"
	halfStr  = "half"
	emptyStr = "empty"
)

func RatingConversion(str string) (result string) {

	/*
		3.48; integer digit = 3, fractional digit = 48
	*/
	integerDigit := 0
	fractionalDigit := 0
	onlyOnce := false

	number := strings.Split(str, ".")

	switch len(number) {
	case 1:
		integerDigit, _ = strconv.Atoi(number[0])
	case 2:
		integerDigit, _ = strconv.Atoi(number[0])
		switch len(number[1]) {
		case 0:
			break
		case 1:
			fractionalDigit, _ = strconv.Atoi(number[1] + "0")
		case 2:
			fractionalDigit, _ = strconv.Atoi(number[1])
		default:
			fractionalDigit, _ = strconv.Atoi(number[1][:2])
		}
	default:
		break
	}

	defer func() {
		/*
			example:
			- input: 3.5
			- output: full full full half empty

		*/

		for star := 1; star <= 5; star++ {
			if star <= integerDigit {
				result = result + fullStr
			} else if !onlyOnce {
				if fractionalDigit > 50 {
					result = result + fullStr
				} else if fractionalDigit > 0 {
					result = result + halfStr
				} else {
					result = result + emptyStr
				}
				onlyOnce = true
			} else {
				result = result + emptyStr
			}

			if star < 5 {
				result = result + " "
			}
		}
	}()
	return
}
