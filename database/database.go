package database

import (
	"fmt"
	"menu/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var url = utils.GetEnv("DB_URL")
var Connect = func() *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to database")
		panic(err)
	} else {
		fmt.Println("Database connected")

		return db
	}
}
