package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/ia-Simon/colorsAPI/models"
)

var (
	//DBConn represents the connection with the database via Gorm.
	DBConn *gorm.DB
)

func connectToDB() (err error) {
	DBConn, err = gorm.Open("sqlite3", "db.sqlite3")
	return
}

func autoMigrate(gormModels ...interface{}) {
	DBConn.AutoMigrate(gormModels...)
}

//InitDBConn connects to the db and executes an AutoMigrate on all models passed
func InitDBConn() {
	if err := connectToDB(); err != nil {
		panic("Couldn't connect to the database")
	}
	fmt.Println("Successfully connected to the database")
	autoMigrate(&models.Color{})
	fmt.Println("Database migrated")
}
