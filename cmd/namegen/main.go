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
