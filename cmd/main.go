package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/jamescun/basex"
)

// command line options
var (
	Encode = flag.Bool("encode", false, "encode data from stdin to stdout")
	Decode = flag.Bool("decode", false, "decode data from stdin to stdout")

	Base     = flag.Int("base", 0, "select a pre-defined base to encode/decode")
	Alphabet = flag.String("alphabet", "", "define a custom alphabet to encode/decode")
)

func main() {
	flag.Parse()

	if (!*Encode && !*Decode) || (*Encode && *Decode) {
		exitWithConfigError("must select either --encode or --decode")
	} else if *Base == 0 && *Alphabet == "" {
		exitWithConfigError("must select a --base or specify a --alphabet")
	}

	var base *basex.BaseX
	if *Base != 0 {
		base = selectBase(*Base)
		if base == nil {
			exitWithConfigError("base %d is unknown, use --alphabet", *Base)
		}
	} else {
		base = basex.New(*Alphabet)
	}

	in := bufio.NewScanner(os.Stdin)
	out := make([]byte, 4*1024)

	for in.Scan() {
		var n int
		var err error

		if *Encode {
			b := in.Bytes()
			base.Encode(out, b)
			n = base.EncodedLen(len(b))
		} else {
			n, err = base.Decode(out, in.Bytes())
			if err != nil {
				exitWithRuntimeError(err.Error())
			}
		}

		out[n] = '\n'

		_, err = os.Stdout.Write(out[:n+1])
		if err != nil {
			exitWithRuntimeError(err.Error())
		}
	}
}

func selectBase(base int) *basex.BaseX {
	switch base {
	case 2:
		return basex.Base2
	case 8:
		return basex.Base8
	case 11:
		return basex.Base11
	case 16:
		return basex.Base16
	case 32:
		return basex.Base32
	case 36:
		return basex.Base36
	case 58:
		return basex.Base58
	case 62:
		return basex.Base62
	case 64:
		return basex.Base64
	case 66:
		return basex.Base66
	default:
		return nil
	}
}

func exitWithConfigError(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "config error: "+format+"\n", args...)
	os.Exit(2)
}

func exitWithRuntimeError(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "runtime error: "+format+"\n", args...)
	os.Exit(1)
}
