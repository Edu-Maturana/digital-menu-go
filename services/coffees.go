package services

import (
	"menu/database"
	"menu/models"

	"github.com/rs/xid"
)

func GetCoffees() []models.Coffee {
	var coffees []models.Coffee
	database.Connect().Find(&coffees)
	return coffees
}

func GetCoffee(id int) models.Coffee {
	var coffee models.Coffee
	database.Connect().First(&coffee, id)
	return coffee
}

func CreateCoffee(coffee models.Coffee) models.Coffee {
	coffee.Id = xid.New().String()

	database.Connect().Create(&coffee)
	return coffee
}

func UpdateCoffee(id int, coffee models.Coffee) models.Coffee {
	database.Connect().Model(&coffee).Where("id = ?", id).Updates(coffee)
	return coffee
}

func DeleteCoffee(id int) {
	database.Connect().Delete(&models.Coffee{}, id)
}
