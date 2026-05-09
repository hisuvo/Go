package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello world")
	http.HandleFunc("/",handler)
	fmt.Println("Hello Bangladesh")
	log.Fatal(http.ListenAndServe("localhost:8000",nil))
}

type Person struct{
	name string
	age int
}


func handler(w http.ResponseWriter, r *http.Request){
	per1 := Person{
		name: "suvo data",
		age: 24,
	}
	fmt.Fprintf(w, "URL.path ---> = %q\n",per1)
}