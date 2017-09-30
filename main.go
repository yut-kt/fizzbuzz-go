package main

import (
	"fmt"
	"flag"
)

var n *int = flag.Int("n", 100, "Setting last position. (default is 100)")

func main() {
	flag.Parse()
	for i := 1; i <= *n; i++ {
		fmt.Println(i, FizzBuzz(i))
	}
}
