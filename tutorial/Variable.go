package main

import (
	"fmt"
	"strconv"
)

// var i int = 42 //The Third Method wont work here at paxkage level

var (
	actorname    string = "Ritesh" //variabeles can be declared in separate Blocks
	actorSurname string = "Ghorpade"
)
var (
	noOfFootbalteam string = "5"
	noOfleagues     string = "20"
)

func main() {
	// var i int //First Method
	// i = 42
	// var i int = 42 //Second Method
	i := 42 //Third Method
	fmt.Println(i)
	fmt.Println(noOfFootbalteam)
	var noOfFootbalteam int = 25 //Shadowing the above variable declartion
	fmt.Println(noOfFootbalteam)

	var k int = 32
	var j = float32(k)
	fmt.Printf("%v, %T", j, j)
	fmt.Println()
	var m string = strconv.Itoa(k)
	fmt.Printf("%v, %T", m, m)
}
