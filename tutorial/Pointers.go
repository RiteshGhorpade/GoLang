package main

import (
	"fmt"
)

func main() {
	var a int = 42
	var b *int = &a
	fmt.Println(&a, *b)

	var ms *myStruct   //pointer to myStruct
	ms = new(myStruct) //intialize with zero value
	(*ms).foo = 42     //Dereferencing and assigning value brackets imp as . operator has more Precedence than pointer
	fmt.Println((*ms))
	fmt.Println(ms.foo) //This works beacuse compiler understands it ,is syntactic sugar
	//Slices and Maps Actually point to underline data , so be careful to use them
}

type myStruct struct {
	foo int
}
