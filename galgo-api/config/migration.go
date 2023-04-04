package config

import (
	"os"
	"strconv"

	BookModel "github.com/marceloamoreno87/galgo-api/api/books/models"
	UserModel "github.com/marceloamoreno87/galgo-api/api/users/models"
	"gorm.io/gorm"
)

func RunMigration(database *gorm.DB) {
	databaseSync := os.Getenv("DATABASE_SYNC")
	databaseSyncBool, _ := strconv.ParseBool(databaseSync)
	if databaseSyncBool {
		database.AutoMigrate(DefineStructs()...)
	}

}

func DefineStructs() []interface{} {
	structs := []interface{}{
		&BookModel.Book{},
		&UserModel.User{},
	}
	return structs
}
