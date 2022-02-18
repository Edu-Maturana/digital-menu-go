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

func CreateCoffee(res http.ResponseWriter, req *http.Request) {
	var coffee models.Coffee
	_ = json.NewDecoder(req.Body).Decode(&coffee)
	coffee, err := services.CreateCoffee(coffee)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
	} else {
		json.NewEncoder(res).Encode(coffee)
	}
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
