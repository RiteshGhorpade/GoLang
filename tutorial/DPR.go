package main

import (
	"fmt"
	// "io/ioutil"
	// "log"
	// "net/http"
)

func main() {
	// res, err := http.Get("http://www.google.com/robots.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// robots, err := ioutil.ReadAll(res.Body)
	// defer res.Body.Close()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s", robots)
	//DEFER does at the last of function before returning in mutiple defer works in LIFO manner
	//Panic
	// a, b := 1, 0
	// ans := a / b
	// fmt.Println(ans)

	fmt.Println("start")
	defer fmt.Println("This is deffered")
	panic("Something bad Happened")
	fmt.Println("end")

}
