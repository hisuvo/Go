package main

import (
	"context"
	"fmt"
	"learn_go_project/db"
	"log"
	"net/http"

	"github.com/lpernett/godotenv"
)

func main() {
	var err error

	err = godotenv.Load()
	if err != nil {
		panic(".env file not found!")
	}

	db.ConnectDB()
	defer db.DB.Close(context.Background())

	mux := http.NewServeMux()

	mux.HandleFunc("GET /teachers", getTeacher)
	mux.HandleFunc("GET /teachers/{id}", getSingleTeacher)
	mux.HandleFunc("POST /teachers", createTeacher)
	mux.HandleFunc("PUT /teachers/{id}", updateTeacher)
	mux.HandleFunc("DELETE /teachers/{id}", deleteTeacher)

	fmt.Println("Server run on the port 5000")

	err = http.ListenAndServe(":5000", mux)

	if err != nil {
		log.Fatal("Error", err)
	}
}