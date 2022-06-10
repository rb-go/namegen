
[![license](https://img.shields.io/github/license/rb-go/namegen.svg)](LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/rb-go/namegen.svg)](https://pkg.go.dev/github.com/rb-go/namegen)
[![Go Report Card](https://goreportcard.com/badge/github.com/rb-go/namegen)](https://goreportcard.com/report/github.com/rb-go/namegen)
[![Coverage Status](https://coveralls.io/repos/github/rb-go/namegen/badge.svg)](https://coveralls.io/github/rb-go/namegen)

# namegen

Funny name generator based on docker namesgenerator

## About 

Package namegen is a name generator util for golang that generates funny names like `hungry_chaplygin`

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

## Install binary

You can install binary with `namegen` name to get names in command line

### if golang installed

```
go install github.com/rb-go/namegen/cmd/namegen@latest
```

### from binary releases

todo

## Credits

- Based on moby/moby namesgenerator source code
