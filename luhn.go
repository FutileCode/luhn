package luhn

import (
	"fmt"
	"strconv"
	"strings"
)

// checks a cc number against the luhn algorithm
func Valid(num string) bool {

	// remove whitespace
	trimmed := strings.ReplaceAll(num, " ", "")

	// ensure valid length
	if len(trimmed) <= 1 {
		return false
	}

	var sum int
	var rev bool

	if len(trimmed)%2 == 0 {
		rev = true
	} else {
		rev = false
	}
	// iterate over each integer in the cc number
	for i := len(trimmed) - 1; i >= 0; i-- {

		// ensure it is an integer
		integer, err := strconv.Atoi(string(trimmed[i]))
		//fmt.Println(integer)
		if err != nil {
			return false
		}

		// add if even index (from right) and double for odd index, subtract 9 if doubled > 9
		switch i % 2 {
		case 0:
			if rev {
				integer *= 2
				if integer > 9 {
					integer -= 9
				}
				sum += integer
			} else {
				sum += integer
			}

		default:
			if rev {
				sum += integer
			} else {
				integer *= 2
				if integer > 9 {
					integer -= 9
				}
				sum += integer
			}

		}
	}
	fmt.Println(sum)
	// check if sum is divisible by 10
	return sum%10 == 0
}
