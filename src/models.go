package manager

import (
	"gorm.io/gorm"
)
import "github.com/google/uuid"

type Logger struct {
	Name string
}

type Student struct {
	gorm.Model
	UUID         uuid.UUID `gorm:"unique;not null"`
	Dozl         int
	Dozs         int
	FName        string
	LName        string
	PhoneNumber  string
	FPhoneNumber string
	Adress       string
	number       int
}

type StudentManagerModel struct {
	db     *gorm.DB
	logger *Logger
}

func (S *Student) BeforeCreate(tx *gorm.DB) (err error) {
	S.UUID = uuid.New()
	S.Dozl = 0
	S.Dozs = 0
	return
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
