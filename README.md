# gopgcli

[![Go Report Card](https://goreportcard.com/badge/github.com/sayanarijit/gopgcli)](https://goreportcard.com/report/github.com/sayanarijit/gopgcli)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/sayanarijit/gopgcli/blob/master/LICENSE)

This tool uses [gopassgen](https://github.com/sayanarijit/gopassgen) library to generate password based on given policy.

### Install

```bash
go get -u github.com/sayanarijit/gopgcli
```

### Run

```bash
gopgcli
```

### Command-line help menu

```bash
gopgcli -h
```

```bash
Usage of gopgcli:
  -caps-alpha-pool string
    	Permitted capital letters (default "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
  -digit-pool string
    	Permitted capital letters (default "0123456789")
  -max-length int
    	Maximum length of password (default 16)
  -min-caps-alpha int
    	Minimun length of capital letters
  -min-digits int
    	Minimun length of digits
  -min-length int
    	Minimum length of password (default 6)
  -min-small-alpha int
    	Minimun length of small letters
  -min-spcl-chars int
    	Minimun length of spcial characters
  -small-alpha-pool string
    	Permitted capital letters (default "abcdefghijklmnopqrstuvwxyz")
  -spcl-char-pool string
    	Permitted capital letters (default "!@#$%^&*()-_=+,.?/:;{}[]~")
```
