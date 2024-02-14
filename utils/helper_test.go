package utils_test

import (
	"testing"

	"github.com/hamedhaghi/card-scheme/enums"
	"github.com/hamedhaghi/card-scheme/utils"
	"github.com/stretchr/testify/assert"
)

func TestCleanCardNumber(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		cardNumber string
		expected   string
	}{
		{
			cardNumber: "5237 2516 2477 8133",
			expected:   "5237251624778133",
		},
		{
			cardNumber: " 5237251624778133",
			expected:   "5237251624778133",
		},
		{
			cardNumber: "5237251624778133 ",
			expected:   "5237251624778133",
		},
		{
			cardNumber: "523725162477 8133",
			expected:   "5237251624778133",
		},
		{
			cardNumber: " 5237251624778133 ",
			expected:   "5237251624778133",
		},
		{
			cardNumber: "52372516 24778133",
			expected:   "5237251624778133",
		},
		{
			cardNumber: "  5 2 3 7 2 5 1 6 2 4 7 7 8 1 3 3 ",
			expected:   "5237251624778133",
		},
		{
			cardNumber: "  523 7251 62477 8133",
			expected:   "5237251624778133",
		},
		{
			cardNumber: "  5237251",
			expected:   "5237251",
		},
	}

	for _, tc := range testCases {
		result := utils.CleanCardNumber(tc.cardNumber)
		assert.Equal(t, tc.expected, result)
	}
}

func TestIsCardNumberValid(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		cardNumber string
		expected   bool
	}{
		{
			cardNumber: "5237 2516 2477 8133",
			expected:   true,
		},
		{
			cardNumber: "523725162477 8133",
			expected:   true,
		},
		{
			cardNumber: " 5237251624778133 ",
			expected:   true,
		},
		{
			cardNumber: "5237251624778133",
			expected:   true,
		},
		{
			cardNumber: "",
			expected:   false,
		},
		{
			cardNumber: "11",
			expected:   false,
		},
		{
			cardNumber: "523725162477833",
			expected:   false,
		},
		{
			cardNumber: "5217 2576 2477 8133",
			expected:   false,
		},
		{
			cardNumber: "5217 2576 2477 8133",
			expected:   false,
		},
		{
			cardNumber: "xxx xxx xxx xxx xx",
			expected:   false,
		},
	}

	for _, tc := range testCases {
		result := utils.IsCardNumberValid(tc.cardNumber)
		assert.Equal(t, tc.expected, result)
	}
}

func TestGetCardScheme(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		cardNumber string
		expected   enums.CardScheme
	}{
		{
			cardNumber: "378282246310005",
			expected:   enums.AmericanExpress,
		},
		{
			cardNumber: "3530111333300000",
			expected:   enums.JCB,
		},
		{
			cardNumber: "6759649826438453",
			expected:   enums.Maestro,
		},
		{
			cardNumber: "4012888888881881",
			expected:   enums.Visa,
		},
		{
			cardNumber: "5105105105105100",
			expected:   enums.MasterCard,
		},
		{
			cardNumber: "",
			expected:   "",
		},
		{
			cardNumber: "13105105105105100",
			expected:   "",
		},
		{
			cardNumber: "1",
			expected:   "",
		},
		{
			cardNumber: "510510510510510012345",
			expected:   "",
		},
	}

	for _, tc := range testCases {
		result := utils.GetCardScheme(tc.cardNumber)
		assert.Equal(t, tc.expected, result)
	}
}
