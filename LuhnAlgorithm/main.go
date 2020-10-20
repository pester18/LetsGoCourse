package main

import (
	"fmt"
	"strconv"
)

func checkLuhn(digits string) bool {
	sum := 0
	nDigits := len(digits)
	parity := nDigits % 2
	i := 0
	for i < nDigits {
		digit, _ := strconv.Atoi(digits[i : i+1])

		if i%2 == parity {
			digit = digit * 2
			if digit > 9 {
				digit = digit - 9
			}
		}
		sum = sum + digit
		i++
	}
	return (sum % 10) == 0
}

func main() {
	validNum := "4561261212345467"
	res := checkLuhn(validNum)
	fmt.Println(validNum, res)
}
