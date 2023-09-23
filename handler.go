package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

func getStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range students {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			break
		}
		return
	}
	json.NewEncoder(w).Encode(&Student{})
}

func includeStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var newStudent Student
	_ = json.NewDecoder(r.Body).Decode(&newStudent)

	students = append(students, newStudent)
	json.NewEncoder(w).Encode(&newStudent)
}

func updateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	studentMarks, err := strconv.Atoi(params["marks"])
	if err != nil {
		http.Error(w, "Invalid marks", http.StatusBadRequest)
		return
	}
	for index, item := range students {
		if item.ID == params["id"] {
			students = append(students[:index], students[index+1:]...)

			var student Student
			_ = json.NewDecoder(r.Body).Decode(&student)
			student.ID = params["id"]
			student.Name = params["name"]
			student.Marks = studentMarks
			students = append(students, student)
			json.NewEncoder(w).Encode(&student)

			return
		}
	}
	json.NewEncoder(w).Encode(&students)
}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range students {
		if item.ID == params["id"] {
			students = append(students[:index], students[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(students)
}
