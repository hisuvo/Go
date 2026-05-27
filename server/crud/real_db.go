package crud

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Student struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Class string `json:"class"`
	Email string `json:"email"`
}

var students = []Student{
	{
		Id: 1,
		Name: "Anupom",
		Class: "M.Sc.",
		Email: "anupom@gmail.com",
	},
	{
		Id: 2,
		Name: "Puja",
		Class: "Eight",
		Email: "puja@gmail.com",
	},
	{
		Id: 3,
		Name: "Arohi",
		Class: "seven",
		Email: "arohi@gmail.com",
	},
}


func getStudent(w http.ResponseWriter, r *http.Request){

	w.Header().Set("content-type","application/json")
	err := json.NewEncoder(w).Encode(students)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w,err)
	}
}

func getSingleStudent(w http.ResponseWriter, r *http.Request){
	paramId := r.PathValue("id") // get route params by r.PathValue() method
	
	// string to int conversion standerd way 
	id, err := strconv.Atoi(paramId)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid student id")
        return
	}

	for _, student := range students{
		if student.Id == id{
			w.Header().Set("content-type","application/json")
			json.NewEncoder(w).Encode(student)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w,"Student not found!")
}

func createStudent(w http.ResponseWriter, r *http.Request){
	var students Student

	err := json.NewDecoder(r.Body).Decode(&students)

	if err != nil {
		fmt.Println("Decod error",err)
	}


	w.Header().Set("Content-type","application/json")
	w.WriteHeader(http.StatusCreated)
	encodErr := json.NewEncoder(w).Encode(students)

	if encodErr != nil {
		fmt.Println("Encod error",err)
	}
}

func updateStudent(w http.ResponseWriter, r *http.Request){
	
	paramId := r.PathValue("id")
	
	id, err := strconv.Atoi(paramId)
	
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w,"Invalied student id")
		return
	}
	
	var updateStudent Student

	err = json.NewDecoder(r.Body).Decode(&updateStudent)

	if err!= nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w,"Invaild request body")
		return
	}

	for idx, student := range students{
		if student.Id == id {
			updateStudent.Id = id
			students[idx] = updateStudent

			w.Header().Set("Content-Type","application/json")
			json.NewEncoder(w).Encode(updateStudent)
			return
		}
	}

	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintln(w,"Student not found!")
}

func deleteStudent(w http.ResponseWriter, r *http.Request){
	paramId := r.PathValue("id")

	id, err := strconv.Atoi(paramId)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w,"Invalide student id")
		return
	}

	for idx, student := range students {
		if student.Id == id{
			students = append(students[:idx], students[idx+1:]... )
			json.NewEncoder(w).Encode(student)
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
	fmt.Fprintln(w,"Student not found")

}

func RealDataBase() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /students", getStudent)
	mux.HandleFunc("GET /students/{id}", getSingleStudent)
	mux.HandleFunc("POST /students", createStudent)
	mux.HandleFunc("PUT /students/{id}",updateStudent)
	mux.HandleFunc("DELETE /students/{id}",deleteStudent)

	fmt.Println("Server run on the port 5000")

	err := http.ListenAndServe(":5000", mux)

	if err != nil {
		log.Fatal("Error",err)
	}

}