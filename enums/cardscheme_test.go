package enums_test

import (
	"testing"

	"github.com/hamedhaghi/card-scheme/enums"
	"github.com/stretchr/testify/assert"
)

func TestCardSchemeString(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		cardScheme enums.CardScheme
		expected   string
	}{
		{
			cardScheme: enums.AmericanExpress,
			expected:   "American Express",
		},
		{
			cardScheme: enums.JCB,
			expected:   "JCB",
		},
		{
			cardScheme: enums.Maestro,
			expected:   "Maestro",
		},
		{
			cardScheme: enums.Visa,
			expected:   "Visa",
		},
		{
			cardScheme: enums.MasterCard,
			expected:   "Master Card",
		},
		{
			cardScheme: "InvalidScheme",
			expected:   "",
		},
		{
			cardScheme: "Hello world",
			expected:   "",
		},
	}

	for _, tc := range testCases {
		result := tc.cardScheme.String()
		assert.Equal(t, tc.expected, result)
	}
}
