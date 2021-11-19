package main

import (
	"fmt"
	"time"
)

func main() {
	n := time.Now();
	fmt.Println("Current Time is ", n);

	france, err := time.LoadLocation("Europe/Paris");
	if err != nil {
		panic(err)
	}


	t := time.Date(1986, time.July, 17, 11,0,0,0, france);
	fmt.Println("I was born on ", t);
	fmt.Println(t.Format(time.ANSIC))

	parsedTime, err := time.Parse(time.ANSIC, "Thu Jul 17 11:00:00 1986")
	if err != nil {
		panic(err)
	}
	fmt.Printf("The type of parsedTime is %T \n", parsedTime);
}
