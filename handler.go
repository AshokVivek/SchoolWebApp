package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var students []Student
	result, err := db.Query("SELECT id, name, marks from students")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		var student Student
		err := result.Scan(&student.ID, &student.Name, &student.Marks)
		if err != nil {
			panic(err.Error())
		}
		students = append(students, student)
	}

	json.NewEncoder(w).Encode(students)
}

func getStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	studentID := params["id"]

	result, err := db.Query("SELECT id, name, marks from students where id=?", studentID)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	var student Student
	if result.Next() {
		err = result.Scan(&student.ID, &student.Name, &student.Marks)
		if err != nil {
			panic(err.Error())
		}

		json.NewEncoder(w).Encode(student)
		fmt.Fprintf(w, "Student with ID %s fetched successfully", studentID)
	} else {
		fmt.Fprintf(w, "There is no student with ID %s", studentID)
	}
}

func fetchTopper(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	result, err := db.Query("SELECT ID, Name, Marks FROM students ORDER BY Marks DESC LIMIT 1")
	if err != nil {
		panic(err.Error())
	}

	var topperStudent Student
	if result.Next() {
		err = result.Scan(&topperStudent.ID, &topperStudent.Name, &topperStudent.Marks)
		if err != nil {
			panic(err.Error())
		}

		json.NewEncoder(w).Encode(topperStudent)
		fmt.Fprintf(w, "Topper student ID is %s", topperStudent.ID)
	} else {
		fmt.Fprintf(w, "There is no students")
	}

	defer result.Close()
}

func includeStudent(w http.ResponseWriter, r *http.Request) {
	var newStudent Student
	err := json.NewDecoder(r.Body).Decode(&newStudent)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	insertStatement, err := db.Prepare("INSERT INTO students (id, name, marks) VALUES (?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	_, err = insertStatement.Exec(newStudent.ID, newStudent.Name, newStudent.Marks)
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprint(w, "Student included successfully")
}

func updateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	studentID := params["id"]

	var updatedStudent Student
	err := json.NewDecoder(r.Body).Decode(&updatedStudent)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	_, err = db.Exec(
		"UPDATE students SET id=?, name=?, marks=? where id=?",
		updatedStudent.ID,
		updatedStudent.Name,
		updatedStudent.Marks,
		studentID,
	)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to update student detail", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Student with ID %s updated successfully", updatedStudent.ID)
}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	studentID := params["id"]

	_, err := db.Exec("DELETE FROM students WHERE id=?", studentID)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to delete student", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Student with ID %s deleted successfully", studentID)
}
