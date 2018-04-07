package basex

// commonly used alphabets defined by their respective base
var (
	Base2  = New("01")
	Base8  = New("01234567")
	Base11 = New("0123456789a")
	Base16 = New("0123456789abcdef")
	Base32 = New("0123456789ABCDEFGHJKMNPQRSTVWXYZ")
	Base36 = New("0123456789abcdefghijklmnopqrstuvwxyz")
	Base58 = New("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")
	Base62 = New("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	Base64 = New("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
	Base66 = New("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_.!~")
)
