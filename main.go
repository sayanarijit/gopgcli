package main

import (
	"flag"
	"fmt"

	"github.com/sayanarijit/gopassgen"
)

func main() {
	p := gopassgen.NewPolicy()

	flag.IntVar(&p.MaxLength, "max-length", 16, "Maximum length of password")
	flag.IntVar(&p.MinLength, "min-length", 6, "Minimum length of password")

	flag.IntVar(&p.MinCapsAlpha, "min-caps-alpha", 0, "Minimun length of capital letters")
	flag.IntVar(&p.MinSmallAlpha, "min-small-alpha", 0, "Minimun length of small letters")
	flag.IntVar(&p.MinDigits, "min-digits", 0, "Minimun length of digits")
	flag.IntVar(&p.MinSpclChars, "min-spcl-chars", 0, "Minimun length of spcial characters")

	flag.StringVar(&p.CapsAlphaPool,
		"caps-alpha-pool",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"Permitted capital letters",
	)
	flag.StringVar(&p.SmallAlphaPool,
		"small-alpha-pool",
		"abcdefghijklmnopqrstuvwxyz",
		"Permitted small letters",
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

	fmt.Println(gopassgen.Generate(p))
}
