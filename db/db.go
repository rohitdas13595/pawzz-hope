package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB = InitDB()

func InitDB() *gorm.DB {
	var dsn = "host=localhost user=postgres password=mysecretpassword dbname=pawzz port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return db
}

// https://github.com/go-gorm/postgres
// var db, err = gorm.Open(postgres.New(postgres.Config{
// 	DSN:                  "user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai",
// 	PreferSimpleProtocol: true, // disables implicit prepared statement usage
// }), &gorm.Config{})
