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

func TestCeaserDecipher(t *testing.T) {
	testCases := map[string]struct {
		in            string
		lang          string
		expectedShift int
		expectedText  string
	}{
		"Not encoded": {
			"this string is not encoded, and is quite short",
			"eng",
			0,
			"this string is not encoded, and is quite short",
		},
		"Basic": {
			"wklv lv d whvw",
			"eng",
			3,
			"this is a test",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			shift, msg := breakCaesar(tc.in, "eng")
			assert.Equal(t, tc.expectedShift, shift)
			assert.Equal(t, tc.expectedText, msg)
		})
	}
}
