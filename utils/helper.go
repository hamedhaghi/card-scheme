package utils

import (
	"slices"
	"strconv"
	"strings"

	"github.com/hamedhaghi/card-scheme/enums"
)

func CleanCardNumber(cardNumber string) string {
	return strings.ReplaceAll(strings.TrimSpace(cardNumber), " ", "")
}

func IsCardNumberValid(cardNumber string) bool {
	cardNumber = CleanCardNumber(cardNumber)
	length := len(cardNumber)

	if length == 0 {
		return false
	}

	sum := 0
	for i := length - 1; i >= 0; i-- {
		isSecondDigit := i%2 == 0
		num, err := strconv.Atoi(string(cardNumber[i]))
		if err != nil {
			return false
		}

		if !isSecondDigit {
			sum += num
		} else {
			doubled := num * 2
			if doubled > 9 {
				doubled = doubled%10 + doubled/10
			}
			sum += doubled
		}
	}

	return sum%10 == 0
}

func GetCardScheme(cardNumber string) enums.CardScheme {
	cardNumber = CleanCardNumber(cardNumber)
	length := len(cardNumber)

	if length == 0 {
		return ""
	}

	if length == 15 {
		twoDigit, err := strconv.Atoi(cardNumber[:2])
		if err == nil && slices.Contains([]int{34, 37}, twoDigit) {
			return enums.AmericanExpress
		}
	}

	if length >= 16 && length <= 19 {
		fourDigit, err := strconv.Atoi(cardNumber[:4])
		if err == nil && fourDigit >= 3528 && fourDigit <= 3589 {
			return enums.JCB
		}
	}

	if length >= 12 && length <= 19 {
		validNumbers := []int{50, 6}
		oneDigit, err1 := strconv.Atoi(cardNumber[:1])
		twoDigit, err2 := strconv.Atoi(cardNumber[:2])
		if (err1 == nil && slices.Contains(validNumbers, oneDigit)) ||
			(err2 == nil && slices.Contains(validNumbers, twoDigit)) ||
			(err2 == nil && twoDigit >= 56 && twoDigit <= 58) {
			return enums.Maestro
		}
	}

	if length == 13 || length == 16 || length == 19 {
		digit, err := strconv.Atoi(cardNumber[:1])
		if err == nil && digit == 4 {
			return enums.Visa
		}
	}

	if length == 16 {
		twoDigit, err1 := strconv.Atoi(cardNumber[:2])
		fourDigit, err2 := strconv.Atoi(cardNumber[:4])
		if (err1 == nil && twoDigit >= 51 && twoDigit <= 55) ||
			(err2 == nil && fourDigit >= 2221 && fourDigit <= 2720) {
			return enums.MasterCard
		}
	}

	return ""
}
