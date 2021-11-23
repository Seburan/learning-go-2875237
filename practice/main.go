package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	// fmt.Println("Day", dow)

	var result string
	switch dow := rand.Intn(7) + 1; dow {
	case 1:
		result = "It's Monday !"
		// fallthrough;
	case 2:
		result = "It's Tuesday !"
		// fallthrough;
	default:
		result = "It's another day !"
	}
	fmt.Println(result)

}
