package services

import (
	"menu/database"
	"menu/models"

	"github.com/rs/xid"
)

var db = database.Connect()

func GetCoffees() []models.Coffee {
	var coffees []models.Coffee
	db.Find(&coffees)
	return coffees
}

func CreateCoffee(coffee models.Coffee) models.Coffee {
	coffee.Id = xid.New().String()
	db.Create(&coffee)
	return coffee
}

func UpdateCoffee(id string, coffee models.Coffee) models.Coffee {
	db.First(&coffee, id)
	db.Model(&coffee).Updates(coffee)
	return coffee
}

func DeleteCoffee(id string) {
	var coffee models.Coffee
	db.First(&coffee, "id = ?", id)
	db.Delete(&coffee)
}
