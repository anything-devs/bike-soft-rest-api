package configs

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DSN = VariablesEnv("DSN")
var BD *gorm.DB

func ConectarBD() {
	var err error

	BD, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatalf(err.Error())
	} else {
		log.Println("BD conectada")
	}

}
