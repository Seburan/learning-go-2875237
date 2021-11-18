package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var aString string = "This is Go !"
	fmt.Println(aString)

	var anInteger int = 64;
	fmt.Printf("%v is a %T \n", anInteger, anInteger);

	myString := "This is also a string"
	fmt.Println(myString)
	fmt.Printf("myString type is %T \n", myString)

	reader := bufio.NewReader(os.Stdin);
	fmt.Print("Enter your name : ")
	input, _ := reader.ReadString('\n');
	fmt.Println("Hi ", strings.TrimSpace(input), " !");

	fmt.Print("Enter your age : ")
	input, _ = reader.ReadString('\n');
	age, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err);
	} else {
		fmt.Printf("Your age %v is a %T \n", age, age);
	}

}
