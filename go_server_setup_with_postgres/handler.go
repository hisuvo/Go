package main

import (
	"context"
	"encoding/json"
	"fmt"
	"learn_go_project/db"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5"
)

func getTeacher(w http.ResponseWriter, r *http.Request) {

	query := `
	SELECT id, name, class, email
	FROM teachers
	`

	rows, err := db.DB.Query(context.Background(), query)
	if err != nil {
		fmt.Println("error", err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Could not get users")
		return
	}

	var teachers []Teacher

	for rows.Next() {
		var teacher Teacher
		err := rows.Scan(&teacher.Id, &teacher.Name, &teacher.Class, &teacher.Email)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "Could not get users")
			return
		}

		teachers = append(teachers, teacher)
	}

	err = rows.Err()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Could not get users")
		return
	}

	w.Header().Set("content-type", "application/json")
	err = json.NewEncoder(w).Encode(teachers)

}

func getSingleTeacher(w http.ResponseWriter, r *http.Request) {
	paramId := r.PathValue("id") // get route params by r.PathValue() method

	// string to int conversion standerd way
	id, err := strconv.Atoi(paramId)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid student id")
		return
	}

	var teacher Teacher

	query := `
		SELECT * FROM teachers
		WHERE id=$1
	`
	err = db.DB.QueryRow(context.Background(), query, id).Scan(&teacher.Id, &teacher.Name, &teacher.Class, &teacher.Email)

	if err == pgx.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "User not found!")
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Users not found!")
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(teacher)
}

func createTeacher(w http.ResponseWriter, r *http.Request) {
	var teacher Teacher

	err := json.NewDecoder(r.Body).Decode(&teacher)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid request body")
		return
	}

	// 2. insert query
	query := `
		INSERT INTO teachers(name, class, email)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	// 3. execute query
	err = db.DB.QueryRow(context.Background(), query, teacher.Name, teacher.Class, teacher.Email).Scan(&teacher.Id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Insert failed:", err)
		return
	}

	// 4. response
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(teacher)

}

func updateTeacher(w http.ResponseWriter, r *http.Request) {

	paramId := r.PathValue("id")

	id, err := strconv.Atoi(paramId)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalied users id")
		return
	}

	var updateTeacher Teacher

	err = json.NewDecoder(r.Body).Decode(&updateTeacher)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invaild request body")
		return
	}

	query := `
		UPDATE teachers
		SET name=$1, class=$2, email=$3
		WHERE id=$4
		RETURNING id, name, class, email
	`

	err = db.DB.QueryRow(context.Background(), query, updateTeacher.Name, updateTeacher.Class, updateTeacher.Email, id).Scan(&updateTeacher.Id, &updateTeacher.Name, &updateTeacher.Class, &updateTeacher.Email)

	if err == pgx.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "User not found!")
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Could not update users")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updateTeacher)
}

func deleteTeacher(w http.ResponseWriter, r *http.Request) {

	var deleteTeacher Teacher

	paramId := r.PathValue("id")

	id, err := strconv.Atoi(paramId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalide student id")
		return
	}

	query := `
		DELETE FROM teachers
		WHERE id=$1
		RETURNING id, name, class, email
	`
	err = db.DB.QueryRow(context.Background(), query, id).Scan(&deleteTeacher.Id, &deleteTeacher.Name, &deleteTeacher.Class, &deleteTeacher.Email)

	if err == pgx.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "User not found!")
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Could not delete user")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(deleteTeacher)

}