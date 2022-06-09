package manager

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDatabase(logger *Logger) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	logger.Info("Database loaded.", nil)
	err = db.AutoMigrate(&Student{})
	if err != nil {
		panic(err)
	}
	return db
}

func GetDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db

}
