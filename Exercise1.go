package main

import (
	"fmt"
	"reflect"
)

var a int = 42
var b string = "James Bond"
var c bool = true
type hotdog int
var h hotdog
var d int

func main() {

	x := 42
	y := "James Bond"
	z := true
	fmt.Println("Execise 1")
	fmt.Println(x, y, z)
	fmt.Println()
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Println("Exercise 2")
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(c))
	fmt.Println()

	fmt.Println("Exercise 3")
	s := fmt.Sprintf("%v\t%v\t%v\t", a,b,c)
	fmt.Println(s)

	fmt.Println()
	fmt.Println("Exercise 4")
	fmt.Printf("%T\n", h)
	h = 42
	fmt.Println(h)

	fmt.Println()
	fmt.Println("Exercise 5")
	d = int(h)
	fmt.Println(d)
	fmt.Printf("%T\n", d)
}