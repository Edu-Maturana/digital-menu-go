package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"menu/models"
	"menu/services"
)

func CreateEmployee(res http.ResponseWriter, req *http.Request) {
	var employee models.Employee
	_ = json.NewDecoder(req.Body).Decode(&employee)
	employee = services.CreateEmployee(employee)
	json.NewEncoder(res).Encode(employee)
}

func UpdateEmployee(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, _ := params["id"]
	var employee models.Employee
	_ = json.NewDecoder(req.Body).Decode(&employee)
	employee = services.UpdateEmployee(id, employee)
	fmt.Fprintf(res, "Employee with ID %v has been updated.", id)
}

func DeleteEmployee(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, _ := params["id"]
	services.DeleteEmployee(id)
	fmt.Fprintf(res, "Employee with ID %v has been deleted.", id)
}
