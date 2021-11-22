package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)
	colors = append(colors, "Purple");
	fmt.Println(colors)

	colors = colors[1:];
	colors = colors[:len(colors)-1];
	fmt.Println(colors)

	numbers := make([]int, 5, 5)
	numbers[0] = 1;
	numbers[1] = 2;
	numbers[2] = 4;
	numbers[3] = 8;
	numbers[4] = 16;
	fmt.Println(numbers);

	numbers = append(numbers, 64);
	fmt.Println(numbers);

	//s := []int{5, 2, 6, 3, 1, 4} // unsorted
	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))

	fmt.Println(numbers);

}
