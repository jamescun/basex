package basex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		Name string
		Base *BaseX

		Src, Dst []byte
	}{
		{"Base2", Base2, []byte("foo"), []byte("11001100110111101101111")},
		{"Base8", Base8, []byte("foo"), []byte("31467557")},
		{"Base11", Base11, []byte("foo"), []byte("38757a9")},
		{"Base16", Base16, []byte("foo"), []byte("666f6f")},
		{"Base32", Base32, []byte("foo"), []byte("6CVVF")},
		{"Base36", Base36, []byte("foo"), []byte("3zvxr")},
		{"Base58", Base58, []byte("foo"), []byte("bQbp")},
		{"Base62", Base62, []byte("foo"), []byte("SAPP")},
		{"Base64", Base64, []byte("foo"), []byte("Zm9v")},
		{"Base66", Base66, []byte("foo"), []byte("WVgA")},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			dst := make([]byte, len(test.Dst))
			test.Base.Encode(dst, test.Src)

			assert.Equal(t, test.Dst, dst)
		})
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		Name string
		Base *BaseX

		Dst, Src []byte
	}{
		{"Base2", Base2, []byte("foo"), []byte("11001100110111101101111")},
		{"Base8", Base8, []byte("foo"), []byte("31467557")},
		{"Base11", Base11, []byte("foo"), []byte("38757a9")},
		{"Base16", Base16, []byte("foo"), []byte("666f6f")},
		{"Base32", Base32, []byte("foo"), []byte("6CVVF")},
		{"Base36", Base36, []byte("foo"), []byte("3zvxr")},
		{"Base58", Base58, []byte("foo"), []byte("bQbp")},
		{"Base62", Base62, []byte("foo"), []byte("SAPP")},
		{"Base64", Base64, []byte("foo"), []byte("Zm9v")},
		{"Base66", Base66, []byte("foo"), []byte("WVgA")},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			dst := make([]byte, len(test.Dst))
			n, _ := test.Base.Decode(dst, test.Src)

			if assert.Equal(t, len(dst), n) {
				assert.Equal(t, test.Dst, dst)
			}
		})
	}
}

func BenchmarkEncodeBase16(b *testing.B) {
	dst := make([]byte, 22)
	src := []byte("hello world")

	b.SetBytes(int64(len(src)))

	for i := 0; i < b.N; i++ {
		Base16.Encode(dst, src)
	}
}

func BenchmarkEncodeBase32(b *testing.B) {
	dst := make([]byte, 18)
	src := []byte("hello world")

	b.SetBytes(int64(len(src)))

	for i := 0; i < b.N; i++ {
		Base32.Encode(dst, src)
	}
}

func BenchmarkEncodeBase62(b *testing.B) {
	dst := make([]byte, 15)
	src := []byte("hello world")

	b.SetBytes(int64(len(src)))

	for i := 0; i < b.N; i++ {
		Base62.Encode(dst, src)
	}
}

func BenchmarkDecodeBase16(b *testing.B) {
	dst := make([]byte, 12)
	src := []byte("68656c6c6f20776f726c64")

	b.SetBytes(int64(len(src)))

	for i := 0; i < b.N; i++ {
		Base16.Decode(dst, src)
	}
}

func BenchmarkDecodeBase32(b *testing.B) {
	dst := make([]byte, 12)
	src := []byte("38CNP6RVS0EXQQ4V34")

	b.SetBytes(int64(len(src)))

	for i := 0; i < b.N; i++ {
		Base32.Decode(dst, src)
	}
}

func BenchmarkDecodeBase62(b *testing.B) {
	dst := make([]byte, 12)
	src := []byte("aaWF93RVY4AwqvW")

	b.SetBytes(int64(len(src)))

	for i := 0; i < b.N; i++ {
		Base62.Decode(dst, src)
	}
}
