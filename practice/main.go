package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

const url = "http://services.explorecalifornia.org/json/tours.php"

func main() {
	fmt.Println("Network requests")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response type: %T\n", response)

	defer response.Body.Close()

	bytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(bytes)
	// fmt.Print(content)

	tours := toursFromJson(content);
	for _, tour := range tours {
		price, err := strconv.ParseFloat(tour.Price, 64)
		if err != nil {
			panic(err);
		}
		fmt.Printf("%v - %v\n", tour.Name, price )
	}

}

func toursFromJson(content string) []Tour {
	tours := make([]Tour, 0, 20)

	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	if err != nil {
		panic(err)
	}

	var tour Tour
	for decoder.More() {
		err := decoder.Decode(&tour)
		if err != nil {
			panic(err)
		}
		tours = append(tours, tour);
	}

	return tours;
}

type Tour struct {
	Name string
	Price string
}
