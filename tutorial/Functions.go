package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello,playground")
	get := sayMessage("Hello", 2)
	fmt.Println(get)

	func() {
		fmt.Println("Anonymous")
	}()
}

func sayMessage(msg string, idx int) string {
	fmt.Println(msg)
	fmt.Println(idx)
	return "This"
}

func divide(a, b float64) (float64, error) {

	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide")
	}
	return a / b, nil
}
