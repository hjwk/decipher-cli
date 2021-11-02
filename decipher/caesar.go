package decipher

import (
	"sort"
	"strings"
)

var freqs_eng = "etaoinshrd"

func caesar(in string, shift int) string {
	shift = (shift%26 + 26) % 26 // [0, 25]
	bytes := make([]byte, len(in))
	for i := 0; i < len(in); i++ {
		c := in[i]
		var a int
		switch {
		case 'a' <= c && c <= 'z':
			a = 'a'
		case 'A' <= c && c <= 'Z':
			a = 'A'
		default:
			bytes[i] = c
			continue
		}
		bytes[i] = byte(a + ((int(c)-a)+shift)%26)
	}

	return string(bytes)
}

func breakCaesar(in, lang string) []string {
	letters := countFrequencies(in)

	var freqs_ref string
	switch lang {
	case "eng":
		freqs_ref = freqs_eng
	default:
		freqs_ref = freqs_eng
	}

	possibilities := make([]string, 5)
	for i := 0; i < 5; i++ {
		shift := int(freqs_ref[i]) - int(letters[i])
		possibilities[i] = caesar(in, shift)
	}

	return possibilities
}

func countFrequencies(in string) string {
	counts := make(map[string]int)
	for _, r := range in {
		if r != ' ' && r != ',' {
			counts[string(r)]++
		}
	}

	letters := make([]string, 0, len(counts))
	for l := range counts {
		letters = append(letters, l)
	}

	sort.Slice(letters, func(i, j int) bool {
		return counts[letters[i]] > counts[letters[j]]
	})

	return strings.Join(letters, "")
}
