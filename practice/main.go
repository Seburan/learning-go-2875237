package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays")

	var colors [3]string;
	colors[0] = "Blue";
	colors[1] = "White";
	colors[2] = "Red";
	fmt.Println(colors);
	fmt.Println(colors[0]);

	var numbers = [5]int{1,2,3,4,5}
	fmt.Println(numbers)

	fmt.Println("Numbers of colors : ", len(colors));
	fmt.Println("Numbers of numbers : ", len(numbers));
}
