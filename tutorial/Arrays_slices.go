package main

import (
	"fmt"
)

func main() {
	grades := [3]int{97, 85, 93}
	//grades :=[...]int{97,85,93} //Alternate syntax
	fmt.Printf("Grades: %v", grades)
	var students [3]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "Lisa"
	students[1] = "RMG"
	students[2] = "TEST"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Number of Students: %v\n", len(students))
	var identityMtrix [3][3]int
	identityMtrix[0] = [3]int{1, 0, 1}
	identityMtrix[1] = [3]int{0, 0, 1}
	identityMtrix[2] = [3]int{1, 0, 0}
	fmt.Printf("Matrix: %v\n", identityMtrix)
	a := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("Length:%v\n", len(a))
	fmt.Printf("Capacity:%v\n", cap(a))

	//Various Ways of creating Slices
	b := []int{1, 2, 3, 4, 5, 67, 89, 9}
	c := b[:]
	e := b[3:]
	d := b[:4]
	f := b[2:5]
	fmt.Println(c)
	fmt.Println(e)
	fmt.Println(d)
	fmt.Println(f)

	m := make([]int, 3, 100)
	fmt.Printf("Length:%v\n", len(m))
	fmt.Printf("Capacity:%v\n", cap(m))

	n := []int{}
	fmt.Printf("Length:%v\n", len(n))
	fmt.Printf("Capacity:%v\n", cap(n))
	n = append(n, 1)
	fmt.Println(n)
	fmt.Printf("Length:%v\n", len(n))
	fmt.Printf("Capacity:%v\n", cap(n))
	fmt.Println(n)
	n = append(n, []int{198, 34, 23}...) //the spread opertor speads it
	fmt.Println(n)
	fmt.Printf("Length:%v\n", len(n))
	fmt.Printf("Capacity:%v\n", cap(n))

}

//a:=[...]int{1,2,3}
//b:=a //this will generate copy of a and will not refernce to same as a
// But it is different for slices
// grades := []int{97, 85, 93} this is a slice not an array the only differenece is we dont specify ... or size
// and b =grades will point to same data
