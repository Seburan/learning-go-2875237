package main

import (
	"fmt"
)

func main() {
	fmt.Println("Functions")
	doSomethings();

	sum := addValues(5, 8);
	fmt.Println("The sum is ", sum)

	multiSum, multiCount := addAllValues(4,7,9,10);
	fmt.Println("The sum is ", multiSum)
	fmt.Println("Count of items", multiCount)
}

func doSomethings() {
	fmt.Println("Doing Something")
}

func addValues(value1, value2 int) int {
	return value1 + value2;
}

func addAllValues(values ...int) (int, int) {
	total := 0;
	for _, v := range values {
		total += v;
	}

	return total, len(values);

}
