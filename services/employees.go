package services

import (
	"menu/models"

	"github.com/rs/xid"
	"golang.org/x/crypto/bcrypt"
)

func CreateEmployee(employee models.Employee) models.Employee {
	employee.Id = xid.New().String()
	hashed, _ := bcrypt.GenerateFromPassword([]byte(employee.Password), 10)
	employee.Password = string(hashed)
	db.Create(&employee)
	return employee
}

func UpdateEmployee(id string, employee models.Employee) models.Employee {
	db.Model(&employee).Where("id = ?", id).Updates(employee)
	return employee
}

func DeleteEmployee(id string) {
	var employee models.Employee
	db.First(&employee, "id = ?", id)
	db.Delete(&employee)
}
