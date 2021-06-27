package decipher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCaesar(t *testing.T) {
	testCases := map[string]struct {
		in       string
		shift    int
		expected string
	}{
		"Shift of 0": {
			"test",
			0,
			"test",
		},
		"Shift of 1": {
			"ABCDEF",
			1,
			"BCDEFG",
		},
		"Shift of -1": {
			"ABCDEF",
			-1,
			"ZABCDE",
		},
		"Shift of 25": {
			"ABCDEF",
			25,
			"ZABCDE",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := caesar(tc.in, tc.shift)
			assert.Equal(t, tc.expected, got)
		})
	}
}
