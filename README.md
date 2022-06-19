# namegen

[![Go Reference](https://pkg.go.dev/badge/github.com/rb-go/namegen.svg)](https://pkg.go.dev/github.com/rb-go/namegen)
[![Go Report Card](https://goreportcard.com/badge/github.com/rb-go/namegen)](https://goreportcard.com/report/github.com/rb-go/namegen)
[![Coverage Status](https://coveralls.io/repos/github/rb-go/namegen/badge.svg)](https://coveralls.io/github/rb-go/namegen)
[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/rb-go/namegen)](https://github.com/rb-go/namegen/releases)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/rb-go/namegen)](https://github.com/rb-go/namegen)
[![license](https://img.shields.io/github/license/rb-go/namegen.svg)](LICENSE)

Funny name generator based on docker namesgenerator but with some additions

## About 

Package namegen is a name generator util for golang that generates funny names like `hungry_chaplygin` with cmd command support (see installation section)

## Usage

```go
package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/rb-go/namegen"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(namegen.GetName(0))
}
```

## Use binary command

You can install binary `namegen` to get names in command line

### Example Usage

Example:

```
namegen -retries=1
```

or just call `namegen` binary to get new name.

if `retries` flag is non-zero, a random integer between 0 and 10 will be added to the end of the name, e.g 'focused_turing3'

### if golang installed

```
go install github.com/rb-go/namegen/cmd/namegen@latest
```

### from binary releases

Check [releases page](https://github.com/rb-go/namegen/releases) to get latest binary for many platforms

## Contribute

Pull request are welcome :) 

But follow with rules:

- `left` - must be an adjective
- `right` - must be the last name of a famous scientist or person associated with IT. Political figures are not allowed!

## Credits

- Based on moby/moby namesgenerator source code
