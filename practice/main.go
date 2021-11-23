package main

import (
	"fmt"
)

func main() {
	fmt.Println("Structs")

	poodle := Dog{
		Breed:  "Poodle",
		Weight: 10,
	}

	fmt.Println(poodle);
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("Breed : %v\nWeight : %v\n", poodle.Breed, poodle.Weight);
	poodle.Weight = 9;
	fmt.Printf("Breed : %v\nWeight : %v\n", poodle.Breed, poodle.Weight);

}

// dog is a struct
type Dog struct {
	Breed string;
	Weight int;
}
