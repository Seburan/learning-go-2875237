package main

import (
	"fmt"
	"math"
)

func main() {

	i1, i2, i3 := 12, 45, 68;
	intSum := i1 +i2 + i3;
	fmt.Println("Integer Sum = ",intSum)

	f1, f2, f3 := 12.5, 45.1, 68.3;
	floatSum := f1 +f2 + f3;
	fmt.Println("Float Sum = ",floatSum)

	floatSum = math.Round(floatSum*1000)/1000
	fmt.Println("Float Sum is now = ",floatSum)

	circleRadius := 15.5
	circumference := circleRadius * 2 * math.Pi;
	fmt.Printf("Circumference is %.5f \n", circumference)
}
