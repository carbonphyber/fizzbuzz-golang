package fizzbuzzapp

import (
	"errors"
	"fmt"
)

func Fizzbuzz(input int) (*string, error) {
	if input < 0 || input > 9999 {
		return nil, errors.New("")
	}
	ret := ""
	if input%15 == 0 {
		ret = "fizzbuzz"
	} else if input%3 == 0 {
		ret = "fizz"
	} else if input%5 == 0 {
		ret = "buzz"
	} else {
		ret = fmt.Sprintf("%d", input)
	}
	return &ret, nil
}
