package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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
	id, _ := strconv.Atoi(params["id"])
	var employee models.Employee
	_ = json.NewDecoder(req.Body).Decode(&employee)
	employee = services.UpdateEmployee(id, employee)
	json.NewEncoder(res).Encode(employee)
}

func DeleteEmployee(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, _ := strconv.Atoi(params["id"])
	services.DeleteEmployee(id)
	fmt.Fprintf(res, "Employee with ID %v has been deleted.", id)
}
