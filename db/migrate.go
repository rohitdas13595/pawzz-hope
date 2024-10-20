package db

import "github.com/rohitdas13595/pawzz-hope/models"

func AutoMigrate() {

	DB.AutoMigrate(&models.Admin{})
}
