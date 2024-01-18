package main

import (
	"crud/src/config"
	"crud/src/controllers/categorycontroller"
	"crud/src/controllers/homecontroller"
	"log"
	"net/http"
)

func main() {

	config.ConnectDB()

	// 1. Home Page

	http.HandleFunc("/", homecontroller.Welcome)

	// 2. Category

	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	log.Println("Server running at port 8080")
	http.ListenAndServe(":8080", nil)
}
