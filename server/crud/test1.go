package crud

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r * http.Request){
	fmt.Fprintln(w,"this is home route")
}

// basic post method use case here
func loginHandler(w http.ResponseWriter, r * http.Request){
	if r.Method != "POST"{
		// w.WriteHeader(405)
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w,"Method not allowed")
		return
	}

	fmt.Fprintln(w,"login in please this is login route")
}


func TestOne() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/login", loginHandler)

	fmt.Println("Server running on port 5000")
	
	err := http.ListenAndServe(":5000", mux)

	if err != nil {
		fmt.Println("Server error",err)
	}
}