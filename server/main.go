package main

import (
	"fmt"
	"net/http"

	"example.com/test/crud"
)

func homeRoute (w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w, "Hello SUVO DATTA! Go Server Running...")
}

func aboutRoute (w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w, "about page")
}

func contactRoute (w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w, "conatct page")
}

func server (){
	// ? here use nil listen and server inside
	// http.HandleFunc("/", homeRoute)
	// http.HandleFunc("/about", aboutRoute)
	// http.HandleFunc("/contact", contactRoute)

	// fmt.Println("Server running on port 500")
	// err := http.ListenAndServe(":5000",nil)

	// ? here user mux listen and server inside
	// * 
	mux := http.NewServeMux()

	mux.HandleFunc("/", homeRoute)
	mux.HandleFunc("/about", aboutRoute)
	mux.HandleFunc("/contact", contactRoute)

	fmt.Println("Server running on port 500")
	err := http.ListenAndServe(":5000",mux)

	if err != nil{
		fmt.Println("Server Error", err)
	}
}

func main() {
	// server()
	// crud.TestOne()
	// crud.TestTwo()
	// crud.TestFakeDB()
	// crud.RealDataBase()
	crud.PgxDataBase()
	
}
