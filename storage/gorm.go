package storage

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"fmt"

)

func GetGormDB(dbName string) *gorm.DB {
	db, err := gorm.Open("sqlite3", dbName)
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}
	return db
}

func MigrateAll(db *gorm.DB ) {
	db.AutoMigrate(&Session{})
	//db.AutoMigrate(&Address{})
	db.AutoMigrate(&Phone{})

	// phone4 :=  &Phone{}
	// phone4.ChatID = 0
	// phone4.Number = "+79131964105"
	// db.Create(&phone4)


}
