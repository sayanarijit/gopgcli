package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sayanarijit/gopassgen"
)

func main() {
	p := gopassgen.NewPolicy()

	flag.IntVar(&p.MaxLength, "max-length", 16, "Maximum length of password")
	flag.IntVar(&p.MinLength, "min-length", 6, "Minimum length of password")

	flag.IntVar(&p.MinUppers, "min-uppers", 0, "Minimun length of upper case letters")
	flag.IntVar(&p.MinUppers, "min-lowers", 0, "Minimun length of lower case letters")
	flag.IntVar(&p.MinDigits, "min-digits", 0, "Minimun length of digits")
	flag.IntVar(&p.MinSpclChars, "min-spcl-chars", 0, "Minimun length of spcial characters")

	flag.StringVar(&p.UpperPool,
		"upper-pool",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"Permitted upper case letters",
	)
	flag.StringVar(&p.LowerPool,
		"lower-pool",
		"abcdefghijklmnopqrstuvwxyz",
		"Permitted lower case letters",
	)
	flag.StringVar(&p.DigitPool,
		"digit-pool",
		"0123456789",
		"Permitted digits",
	)
	flag.StringVar(&p.SpclCharPool,
		"spcl-char-pool",
		"!@#$%^&*()-_=+,.?/:;{}[]~",
		"Permitted special characters",
	)

	flag.Parse()

	r, err := gopassgen.Generate(p)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	fmt.Println(r)
}
