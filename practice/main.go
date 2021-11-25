package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "http://services.explorecalifornia.org/json/tours.php"

func main() {
	fmt.Println("Network requests")

	response,err := http.Get(url);
	if err != nil {
		panic(err);
	}

	fmt.Printf("Response type: %T\n", response);

	defer response.Body.Close();

	bytes, err := io.ReadAll(response.Body);

	if err != nil {
		panic(err);
	}

	content := string(bytes);
	fmt.Print(content);

}
