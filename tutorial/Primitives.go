package main

import (
	"fmt"
)

func main() {
	var n bool = false //Boolean
	fmt.Printf("%v %T", n, n)
	fmt.Println()
	var a int8 = 2
	var b int16 = 2777
	var c uint8 = 2
	var d uint8 = 244
	fmt.Printf("%v %T %v %T %v %T %v %T ", a, a, b, b, c, c, d, d)
	m := 3.14
	fmt.Println()
	fmt.Printf(" %v %T", m, m)
	m = 13.7e72

	fmt.Println()
	fmt.Printf(" %v %T", m, m)
	m = 2 / 1E14
	fmt.Println()
	fmt.Printf(" %v %T", m, m)
	fmt.Println()
	var com complex64 = 3 + 4i //or complex(5,12)
	fmt.Printf(" %v %T", com, com)
	fmt.Printf(" %v %T", real(com), com)
	fmt.Printf(" %v %T", imag(com), com)

	s := "this is a string"
	fmt.Printf("%v ,%T\n", s, s)

	by := []byte(s)
	fmt.Printf("%v ,%T\n", by, by)

	r := 'a' //rune is integer 32
	fmt.Printf("%v ,%T\n", r, r)
}
