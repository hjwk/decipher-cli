[![Go Report Card](https://goreportcard.com/badge/github.com/hjwk/decipher)](https://goreportcard.com/report/github.com/hjwk/decipher)
[![codecov](https://codecov.io/gh/hjwk/decipher/branch/main/graph/badge.svg?token=NTYTJ9FMH4)](https://codecov.io/gh/hjwk/decipher)
# Decipher

Decipher-CLI allows you to play around with various ciphers.

Supported ciphers:
- caeasar (enciphering and deciphering)
- scytale (enciphering)

## How to use

To encipher a text using the caeasar cipher (and the standard shift of 3)
```
./decipher cipher -c caesar -t "Attack at once"
```

To attempt an automatic decipher of a suspected caesar ciphertext
```
./decipher decipher -c caesar -t "dwwdfndwrqfh"
```
