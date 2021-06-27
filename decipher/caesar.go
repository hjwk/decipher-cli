package decipher

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
