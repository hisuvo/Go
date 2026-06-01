package routes

import (
	"fmt"
	"go_postgresql/controller"
	"net/http"
)

func Routes() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /users", controller.GetUser)
	fmt.Println("Server is rounning on port 5000")

	http.ListenAndServe(":5000", mux)
}