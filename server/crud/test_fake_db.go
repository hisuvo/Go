package crud

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
	Email string `json:"email"`
}

var users = []Person{
	{
		Id: 1,
		Name: "suvo",
		Age: 24,
		Email: "suvo@gmail.com",
	},
	{
		Id: 2,
		Name: "datta",
		Age: 24,
		Email: "datta@gmail.com",
	},
	{
		Id: 3,
		Name: "sajeeb",
		Age: 23,
		Email: "sajeeb@gmail.com",
	},
}

func userHanlder(w http.ResponseWriter, r *http.Request){
	
	w.Header().Set("Content-Type","application/json")

	// This is not good approce becuse it take memory place
	userByte, err := json.Marshal(users)
	w.Write(userByte)
	
	if err != nil{
		fmt.Println("Error:",err)
	}

}

func createUserHandler(w http.ResponseWriter, r *http.Request){

	defer r.Body.Close()
	
	var newUser Person
	
	err := json.NewDecoder(r.Body).Decode(&newUser)
	
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w,"User infor not valid")
		return
	}
	
	newUser.Id = len(users) + 1

	users = append(users, newUser)

	json.NewEncoder(w).Encode(newUser)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
}

func personHanlder(w http.ResponseWriter, r *http.Request){
	
	w.Header().Set("Content-Type","application/json")
	// This is good way because it provide stream not take memory place
	// err := json.NewEncoder(w).Encode(users)

	encode := json.NewEncoder(w)
	err := encode.Encode(users)
	
	if err != nil{
		fmt.Println("Error:",err)
	}

}

func TestFakeDB() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /users", userHanlder)
	mux.HandleFunc("GET /persons", personHanlder)
	mux.HandleFunc("POST /users", createUserHandler)

	fmt.Println("Server running on port 5000")

	err := http.ListenAndServe(":5000",mux)
	if err != nil{
		fmt.Println("Server error",err)
	}
}