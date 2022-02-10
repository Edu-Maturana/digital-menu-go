package services

import (
	"menu/database"
	"menu/models"

	"github.com/rs/xid"
	"golang.org/x/crypto/bcrypt"
)

func CreateEmployee(models.Employee) models.Employee {
	employee := models.Employee{}

	employee.Id = xid.New().String()

	// hash password
	hashed, _ := bcrypt.GenerateFromPassword([]byte(employee.Password), 10)
	employee.Password = string(hashed)

	database.Connect().Create(&employee)
	return employee
}

func UpdateEmployee(id int, employee models.Employee) models.Employee {
	database.Connect().Model(&employee).Where("id = ?", id).Updates(employee)
	return employee
}

func DeleteEmployee(id int) {
	database.Connect().Delete(&models.Employee{}, id)
}
