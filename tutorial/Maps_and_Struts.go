package main

import (
	"fmt"
	"reflect"
)

type Doctor struct {
	Number     int
	Actorname  string
	Companions []string
}

func main() {
	//statePopulation := make(map[string]int,10)
	statePopulation := map[string]int{
		"A":  1,
		"BC": 1,
		"AD": 1,
	}
	delete(statePopulation, "A")
	fmt.Println(statePopulation)
	//Slices cannot be a key and Array can be used as key
	aDoctor := Doctor{
		Number:    3,
		Actorname: "Ritesh Ghorpade",
		Companions: []string{
			"RMG",
			"RUT",
			"PUT",
		},
	}
	fmt.Println(aDoctor.Companions[1])
	aDoctor2 := struct{ name string }{name: "John Pertwee"}
	fmt.Println(aDoctor2)

	type Animal struct {
		Name   string `required max:"2"`
		Origin string
	}
	type Bird struct {
		Animal
		Speed  float32
		Canfly bool
	}
	b := Bird{
		Animal: Animal{Name: "RMG", Origin: "INDIA"},
		Speed:  23,
		Canfly: false,
	}
	fmt.Println(b)
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)

}
