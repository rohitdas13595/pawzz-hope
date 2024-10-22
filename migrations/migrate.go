package migrations

import (
	"github.com/rohitdas13595/pawzz-hope/app/admin"
	"github.com/rohitdas13595/pawzz-hope/db"
)

func AutoMigrate() {

	db.DB.AutoMigrate(&admin.Admin{})
}
