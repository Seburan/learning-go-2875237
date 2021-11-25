package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	float1 := inputValue("Value 1")
	float2 := inputValue("Value 2")
	var result float64;

	fmt.Print("Select an operation (+ - * /) : ")
	inputString, _ := reader.ReadString('\n')
	operationString := strings.TrimSpace(inputString)

	switch operationString {
	case "+":
		result = calcAddition(float1, float2)
		fmt.Printf("The sum of %v and %v is %v\n\n", float1, float2, result)
	case "-":
		result = calcSubstraction(float1, float2)
		fmt.Printf("The substraction of %v and %v is %v\n\n", float1, float2, result)
	case "*":
		result = calcMultiplication(float1, float2)
		fmt.Printf("The multiplication of %v by %v is %v\n\n", float1, float2, result)
	case "/":
		result = calcDivision(float1, float2)
		fmt.Printf("The division of %v by %v is %v\n\n", float1, float2, result)
	default:
		panic("Incorrect Operation")
	}

	// result := calcAddition(float1, float2)
	// sum = math.Round(sum*100) / 100
	// fmt.Printf("The sum of %v and %v is %v\n\n", float1, float2, result)

}

// obtain user input of a number
func inputValue(valueName string) float64 {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("%v : ", valueName)
	inputString, _ := reader.ReadString('\n')
	inputFloat, err := strconv.ParseFloat(strings.TrimSpace(inputString), 64)
	if err != nil {
		fmt.Println(err)
		panic(fmt.Sprintf("%v must be a number", valueName))
	}

	return inputFloat
}

// sum of v1 and v2
func calcAddition(v1, v2 float64) float64 {
	return v1 + v2
}

// substraction of v1 and v2
func calcSubstraction(v1, v2 float64) float64 {
	return v1 - v2
}

// Multiplication of v1 and v2
func calcMultiplication(v1, v2 float64) float64 {
	return v1 * v2
}

// substraction of v1 and v2
func calcDivision(v1, v2 float64) float64 {
	return v1 / v2
}
