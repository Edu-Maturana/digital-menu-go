package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"menu/handlers"
)

func main() {
	router := mux.NewRouter()

	// Coffee endpoints
	router.HandleFunc("/coffees", handlers.GetCoffees).Methods("GET")
	router.HandleFunc("/coffees/{id}", handlers.GetCoffee).Methods("GET")
	router.HandleFunc("/coffees", handlers.CreateCoffee).Methods("POST")
	router.HandleFunc("/coffees/{id}", handlers.UpdateCoffee).Methods("PUT")
	router.HandleFunc("/coffees/{id}", handlers.DeleteCoffee).Methods("DELETE")

	fmt.Println("Server running!")
	http.ListenAndServe(":3000", router)
}
