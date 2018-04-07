# BaseX

[![](https://godoc.org/github.com/jamescun/basex?status.svg)](http://godoc.org/github.com/jamescun/basex)

BaseX is a Go library and command line utility for fast base encoding and decoding of any alphabet using bitcoin style leading zero compression.

```sh
go get -u github.com/jamescun/basex/cmd
```

## Alphabets

While BaseX supports custom definition of alphabets, BaseX embeds and exports some commonly used alphabets by their respective base.

| Base | Alphabet                                                              |
|------|-----------------------------------------------------------------------|
| 2    | `01`                                                                  |
| 8    | `01234567`                                                            |
| 11   | `0123456789a`                                                         |
| 16   | `0123456789abcdef`                                                    |
| 32   | `0123456789ABCDEFGHJKMNPQRSTVWXYZ`                                    |
| 36   | `0123456789abcdefghijklmnopqrstuvwxyz`                                |
| 58   | `123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz`          |
| 62   | `0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz`      |
| 64   | `ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/`    |
| 66   | `ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_.!~` |


## Example

### Command Line

Encode a string as base 32:

```sh
$ echo "hello world" | basex --encode --base 32
38CNP6RVS0EXQQ4V34
```

Decode a string from base 32:

```sh
$ echo "38CNP6RVS0EXQQ4V34" | basex --decode --base 32
hello world
```

Encode with a custom alphabet (reverse base 8):

```sh
$ echo "hello world" | basex --encode --alphabet 76543210
13632447116206770422043311633
```

Decode with a custom alphabet (reverse base 8):

```sh
$ echo "13632447116206770422043311633" | basex --decode --alphabet 76543210
hello world
```

### Go Library

Encode a string as base 32:

```go
str := basex.Base32.EncodeToString("hello world")
// str => "38CNP6RVS0EXQQ4V34"
```

Decode a string from base32:

```go
bytes, _ := basex.Base32.DecodeString("38CNP6RVS0EXQQ4V34")
// bytes => []byte("hello world")
```

Encode with a custom alphabet (reverse base 8):

```go
reverseBase8 := New("76543210")

str := reverseBase8.EncodeToString("hello world")
// str => "13632447116206770422043311633"
```

Decode with a custom alphabet (reverse base 8):

```go
reverseBase8 = New("76543210")

bytes, _ := reverseBase8.DecodeString("13632447116206770422043311633")
// bytes => []byte("hello world")
```
