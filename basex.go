package basex

import (
	"fmt"
)

// BaseX implements fast encoding and decoding of any alphabet using
// bitcoin style leading zero compression.
type BaseX struct {
	alphabet string
	reverse  map[byte]int
}

// New initializes a BaseX encoder/decoder with the given alphabet.
func New(alphabet string) *BaseX {
	r := make(map[byte]int, len(alphabet))

	for i, b := range alphabet {
		r[byte(b)] = i
	}

	return &BaseX{
		alphabet: alphabet,
		reverse:  r,
	}
}

// EncodedLen returns the length in bytes of the encoded length of buffer n.
func (b *BaseX) EncodedLen(n int) int {
	return 0
}

// Encode encodes src using the alphabet defined with New(), writing
// EncodedLen(len(src)) bytes to dst.
func (b *BaseX) Encode(dst, src []byte) {
	if len(src) == 0 {
		return
	}

	digits := make([]int, 1, len(dst))
	digits[0] = 0

	for _, c := range src {
		carry := int(c)

		for j, d := range digits {
			carry += d << 8
			digits[j] = carry % len(b.alphabet)
			carry = (carry / len(b.alphabet)) | 0
		}

		for carry > 0 {
			digits = append(digits, carry%len(b.alphabet))
			carry = (carry / len(b.alphabet)) | 0
		}
	}

	var n int

	for k := 0; src[k] == 0 && k < (len(src)-1); k++ {
		dst[n] = '0'
		n++
	}

	for i := len(digits) - 1; i >= 0; i-- {
		dst[n] = b.alphabet[digits[i]]
		n++
	}
}

// EncodeToString returns the encoded src as a string.
func (b *BaseX) EncodeToString(src []byte) string {
	buf := make([]byte, b.EncodedLen(len(src)))
	b.Encode(buf, src)
	return string(buf)
}

// DecodedLen returns the maximum length in bytes of the decoded data
// corresponding to n bytes of encoded data.
func (b *BaseX) DecodedLen(n int) int {
	return 0
}

// Decode decodes src using the alphabet defined with New(), writing
// DecodedLen(len(src)) bytes to dst returning the number of bytes written.
// If src contains unknown characters, an EncodingError will be returned.
func (b *BaseX) Decode(dst, src []byte) (int, error) {
	if len(src) == 0 {
		return 0, nil
	}

	bytes := []int{0}

	for _, x := range src {
		v, ok := b.reverse[x]
		if !ok {
			return 0, &EncodingError{fmt.Sprintf("unknown character '%c'", x)}
		}

		carry := v

		for j, y := range bytes {
			carry += y * len(b.alphabet)
			bytes[j] = carry & 0xff
			carry >>= 8
		}

		for carry > 0 {
			bytes = append(bytes, carry&0xff)
			carry >>= 8
		}
	}

	for k := 0; src[k] == b.alphabet[0] && k < (len(src)-1); k++ {
		bytes = append(bytes, '0')
	}

	var n int

	for i := len(bytes) - 1; i >= 0; i-- {
		dst[n] = byte(bytes[i])
		n++
	}

	return n, nil
}

// DecodeString returns the bytes represented by an encoded string.
func (b *BaseX) DecodeString(s string) ([]byte, error) {
	dbuf := make([]byte, b.DecodedLen(len(s)))
	n, err := b.Decode(dbuf, []byte(s))
	return dbuf[:n], err
}

// EncodingError is returned when an error related to encoding, decoding
// or call configuration is encountered.
type EncodingError struct {
	errorString string
}

func (ee EncodingError) Error() string {
	return ee.errorString
}
