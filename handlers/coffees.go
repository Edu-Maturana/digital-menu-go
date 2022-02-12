package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"menu/models"
	"menu/services"
)

func GetCoffees(res http.ResponseWriter, req *http.Request) {
	coffees := services.GetCoffees()
	json.NewEncoder(res).Encode(coffees)
}

func GetCoffee(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]
	coffee := services.GetCoffee(id)
	json.NewEncoder(res).Encode(coffee)
}

func CreateCoffee(res http.ResponseWriter, req *http.Request) {
	var coffee models.Coffee
	_ = json.NewDecoder(req.Body).Decode(&coffee)
	coffee = services.CreateCoffee(coffee)
	json.NewEncoder(res).Encode(coffee)
}

func UpdateCoffee(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]
	var coffee models.Coffee
	_ = json.NewDecoder(req.Body).Decode(&coffee)
	coffee = services.UpdateCoffee(id, coffee)
	fmt.Fprintf(res, "Coffee with id %s updated", id)
}

func DeleteCoffee(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]
	services.DeleteCoffee(id)
	fmt.Fprintf(res, "Coffee with id %s deleted", id)
}
