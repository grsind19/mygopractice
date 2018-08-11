package main

import (
	"flag"
	"fmt"
)

var name = flag.String("name", "world", "A way to say hello")

var spanish bool

func init() {
	flag.BoolVar(&spanish, "spanish", false, "Use spanish")
	flag.BoolVar(&spanish, "s", false, "Use spanish")
}

func main() {
	flag.Parse()

	if spanish == true {
		fmt.Printf("Hola %s!\n", *name)
	} else {
		fmt.Printf("Hello %s!\n", *name)
	}
}
