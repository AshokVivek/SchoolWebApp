package main

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Student struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Marks int    `json:"marks"`
}

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "root:Pendrive@123@tcp(127.0.0.1:3306)/schoolwebapp")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	router := mux.NewRouter()

	router.HandleFunc("/students", getStudents).Methods("GET")
	router.HandleFunc("/students", includeStudent).Methods("POST")
	router.HandleFunc("/students/topper", fetchTopper).Methods("GET")
	router.HandleFunc("/students/{id}", getStudent).Methods("GET")
	router.HandleFunc("/students/{id}", updateStudent).Methods("PUT")
	router.HandleFunc("/students/{id}", deleteStudent).Methods("DELETE")

	http.ListenAndServe(":8000", router)
}
