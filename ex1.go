package main

import (
	"fmt"
)

/**/
func Names() (string, string) {
	return "Foo", "Bar"
}

func TwoNames() (first string, second string) {
	first = "one"
	second = "two"
	return
}

func main() {
	n1, n2 := Names()
	fmt.Println(n1)
	fmt.Println(n2)

	n3, _ := TwoNames()
	fmt.Println(n3)
}
