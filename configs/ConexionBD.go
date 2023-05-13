package configs

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DSN = VariablesEnv("DSN")

func ConectarBD() *gorm.DB {
	DB, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatalf(err.Error())
	} else {
		log.Println("BD conectada")
	}
	return DB
}
