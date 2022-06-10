package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/rb-go/namegen"
)

func main() {
	var retriesCount int

	flag.IntVar(&retriesCount, "retries", 0, "If retry is non-zero, a random integer between 0 and 10 will be added to the end of the name, e.g 'focused_turing3'")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())
	fmt.Println(namegen.GetName(retriesCount))
}
