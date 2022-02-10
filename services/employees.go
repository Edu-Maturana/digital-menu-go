package services

import (
	"menu/database"
	"menu/models"
)

func CreateEmployee(emp models.Employee, user models.Employee) models.Employee {
	var employee models.Employee
	database.Connect().Create(&employee)
	return employee
}

func UpdateEmployee(id int, emp models.Employee, user models.Employee) models.Employee {
	var employee models.Employee
	database.Connect().Model(&employee).Where("id = ?", id).Updates(emp)
	return employee
}

func DeleteEmployee(id int, user models.Employee) {
	var employee models.Employee
	database.Connect().Delete(&employee, id)
}
