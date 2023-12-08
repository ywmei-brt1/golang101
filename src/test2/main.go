package main

import (
	"fmt"
)

type A struct {
	a int
}

func main() {
	a := 100
	a, b := foobar()
	fmt.Println(a, b)

	objA := A{a: 123}
	objA.a, c := foobar()
	fmt.Println(objA, c)
}

func foobar() (int, int) {
	return 1, 2
}
