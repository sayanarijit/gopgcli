# gopgcli

[![Go Report Card](https://goreportcard.com/badge/github.com/sayanarijit/gopgcli)](https://goreportcard.com/report/github.com/sayanarijit/gopgcli)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/sayanarijit/gopgcli/blob/master/LICENSE)

This tool uses [gopassgen](https://github.com/sayanarijit/gopassgen) library to generate password based on given policy.

### Download binary

[latest release](https://github.com/sayanarijit/gopgcli/releases/latest)

[all releases](https://github.com/sayanarijit/gopgcli/releases)

### Run

```bash
gopgcli
```

### Command-line help menu

```bash
gopgcli -h
```

```bash
Usage of ./gopgcli:
  -digit-pool string
        Permitted digits (default "0123456789")
  -lower-pool string
        Permitted lower case letters (default "abcdefghijklmnopqrstuvwxyz")
  -max-length int
        Maximum length of password (default 16)
  -min-digits int
        Minimun length of digits
  -min-length int
        Minimum length of password (default 6)
  -min-lowers int
        Minimun length of lower case letters
  -min-spcl-chars int
        Minimun length of spcial characters
  -min-uppers int
        Minimun length of upper case letters
  -spcl-char-pool string
        Permitted special characters (default "!@#$%^&*()-_=+,.?/:;{}[]~")
  -upper-pool string
        Permitted upper case letters (default "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
```
