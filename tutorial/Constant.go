package main

import (
	"fmt"
)

const (
	a = iota
	b = iota
	c = iota
	d //Comipler sees it and automatically Assigns the value of iota
	e
)

func main() {
	const myConst int = 42
	fmt.Printf("%v ,%T\n", myConst, myConst)
	fmt.Printf(" %v ,%T\n", a, a)
	fmt.Printf(" %v ,%T\n", b, b)
	fmt.Printf(" %v ,%T\n", c, c)
	fmt.Printf(" %v ,%T\n", d, d)
	fmt.Printf(" %v ,%T\n", e, e)

}
