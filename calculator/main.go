package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("A simple calculator")

	// var value1 float64;
	//var value2 float64;
	// var valueSum float64;


	reader := bufio.NewReader(os.Stdin);
	fmt.Print("Value 1 : ")
	input1, errReader := reader.ReadString('\n')
	if errReader != nil {
		fmt.Println(errReader)
		panic("Value 1 must be a number")
	}
	value1, errConversion := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	if errConversion != nil {
		fmt.Println(errConversion)
		panic("Could not convert your input to a float64")
	}

	fmt.Print("Value 2 : ")
	input2, errReader2 := reader.ReadString('\n')
	if errReader2 != nil {
		fmt.Println(errReader2)
		panic("Value 2 must be a number")
	}
	value2, errConversion2 := strconv.ParseFloat(strings.TrimSpace(input2), 64)
	if errConversion2 != nil {
		fmt.Println(errConversion2)
		panic("Could not convert your input to a float64")
	}

	valueSum := value1 + value2;
	valueSum = math.Round(valueSum*10000)/10000;
	fmt.Println("The sum of ",value1, "and ", value2, " is ", valueSum);

}
