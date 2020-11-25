package main

import (
	"fmt"
)

func main() {
	grades := [3]int{97, 85, 93}
	//grades :=[...]int{97,85,93}
	fmt.Printf("Grades: %v", grades)
	var students [3]string
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Number of Students: %v\n", len(students))

}

//a:=[...]int{1,2,3}
//b:=a //this will generate copy of a and will not refernce to same as a
// But it is different for slices
// grades := []int{97, 85, 93} this is a slice not an array
// and b =grades will point to same data
