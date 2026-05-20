package main

import (
	"encoding/json"
	"fmt"
)

type Address struct{
	Village string `json:"village"`
	Post string `json:"postOffice"`
	Thana string `json:"thana"`
	Zila string `json:"zila"`
	Country string `json:"country"`
}

type Person struct {
	Name    string `json:"name"`
	Age     int	`json:"age"`
	Address Address `json:"adress"`
}

func main() {

	// Struct to JSON
	suvo := Person{
		Name:    "suvo datta",
		Age:     24,
		Address: Address{
			Village: "Baulia",
			Post: "Bottla Bazar",
			Thana: "Shalikha",
			Zila: "Magura",
			Country: "Bangladesh",
		},
		
	}

	jsonByte,error := json.Marshal(suvo)

	if error != nil{
		fmt.Println("Error:",error)
	}

	fmt.Println("Json Data",string(jsonByte))

	// JSON to struct
	var rajib Person
	jsonData := `{"name":"suvo datta","age":24,"adress":{"village":"Baulia","postOffice":"Bottla","thana":"Shalikha","zila":"Magura","country":"Bangladeshi"}}`

	err := json.Unmarshal([]byte(jsonData),&rajib)
	
	if err != nil{
		fmt.Println("Error",err)
	}

	fmt.Println("Rajib josn data", rajib)
}