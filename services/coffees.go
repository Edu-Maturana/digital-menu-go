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

func GetCoffee(id string) models.Coffee {
	var coffee models.Coffee
	db.First(&coffee, "id = ?", id)
	return coffee
}

func CreateCoffee(coffee models.Coffee) models.Coffee {
	coffee.Id = xid.New().String()
	db.Create(&coffee)
	return coffee
}

func UpdateCoffee(id string, coffee models.Coffee) models.Coffee {
	db.Model(&coffee).Where("id = ?", id).Updates(coffee)
	return coffee
}

func DeleteCoffee(id string) {
	var coffee models.Coffee
	db.First(&coffee, "id = ?", id)
	db.Delete(&coffee)
}
