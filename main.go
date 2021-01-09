package main

import (
	"fmt"
	"log"
	"net/http"
	indexRouter "shopify-image-go/router"

	"github.com/gorilla/mux"
)

func Calculate(num int) int {
	return num + 2
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is an Image upload project.")
}

func processRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", indexRouter.RouterTest)
	myRouter.HandleFunc("/images", indexRouter.SelectImage).Methods("GET")
	myRouter.HandleFunc("/image/{id}", indexRouter.SelectImageById).Methods("GET")
	myRouter.HandleFunc("/image/update/{id}", indexRouter.UpdateImageById).Methods("POST")
	myRouter.HandleFunc("/image/create", indexRouter.CreateImage).Methods("POST")
	myRouter.HandleFunc("/image/delete/{id}", indexRouter.DeleteImage).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8001", myRouter))
}

func main() {
	fmt.Println("REST API V1: Shopify image upload!")
	processRequests()
}
