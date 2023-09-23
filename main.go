package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Student struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Marks int    `json:"marks"`
}

var students []Student

func main() {
	router := mux.NewRouter()

	students = append(students, Student{ID: "1001", Name: "Soham Sinha", Marks: 476})
	students = append(students, Student{ID: "1002", Name: "Rahul Majumdar", Marks: 472})

	router.HandleFunc("/students", getStudents).Methods("GET")
	router.HandleFunc("/students", includeStudent).Methods("POST")
	router.HandleFunc("/students/{id}", getStudent).Methods("GET")
	router.HandleFunc("/students/{id}", updateStudent).Methods("PUT")
	router.HandleFunc("/students/{id}", deleteStudent).Methods("DELETE")

	http.ListenAndServe(":8000", router)
}
