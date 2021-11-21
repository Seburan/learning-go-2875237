package main

import (
	"fmt"
)

func main() {
	fmt.Println("Pointers")

	anInt := 42;
	var p = &anInt;
	fmt.Println("Value of p = ", *p);

	value1 := 42.13
	pointer1 := &value1;
	fmt.Println("Value of value1 = ", *pointer1)

	*pointer1 = *pointer1 / 31;
	fmt.Println("Value of *pointer1 = ", *pointer1);
	fmt.Println("Value of value1 = ", value1)

}
