package crud

import (
	"fmt"
	"net/http"
)

func faqHandler(w http.ResponseWriter, r * http.Request){
	fmt.Fprintln(w,"this is FAQ route. check you quetion")
}

func registerHandler(w http.ResponseWriter, r * http.Request){
	fmt.Fprintln(w,"This is register route. create account please")
}


func TestTwo() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /faq", faqHandler)
	mux.HandleFunc("POST /register", registerHandler)

	fmt.Println("Server running on port 5000")
	
	err := http.ListenAndServe(":5000", mux)

	if err != nil {
		fmt.Println("Server error",err)
	}
}