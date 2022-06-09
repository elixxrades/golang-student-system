package manager

import (
	"gorm.io/gorm"
	"strconv"
)

func StudentManager(database *gorm.DB, logger *Logger) *StudentManagerModel {
	w := &StudentManagerModel{
		db:     database,
		logger: logger,
	}
	return w
}

func (m *StudentManagerModel) CreateUser(student *Student) *Student {
	m.db.Create(student)
	return student
}

func (m *StudentManagerModel) UpdateUser(student *Student) {

	m.RemoveUserWithUUID(student.UUID.String())
	m.CreateUser(student)
}

func (m *StudentManagerModel) GetStudentWithUUID(data string) Student {
	var result Student
	m.db.First(&result, "UUID = ?", data)

	return result
}

func (m *StudentManagerModel) GetStudentWithID(data string) Student {
	var result Student
	m.db.First(&result, "ID = ?", data)

	return result
}

func (m *StudentManagerModel) RemoveUserWithID(data uint) {
	m.db.Delete(&Student{}, data)
}

func (m *StudentManagerModel) RemoveUserWithUUID(str string) {
	m.db.Delete(&Student{}, "uuid = ?", str)
}

func (S *Student) GetStudentName() string {
	return S.FName + " " + S.LName
}
func (S *Student) GetDevamsizlikSTR() string {
	return strconv.Itoa(int(S.Dozs + S.Dozl))
}
func (S *Student) DevamsizlikEkle(t string) Student {
	db := GetDatabase()

	var result Student
	if t == "ozursuz" {
		db.First(&result, "UUID = ?", S.UUID).Update("dozs", S.Dozs+1)
	} else if t == "ozurlu" {
		db.First(&result, "UUID = ?", S.UUID).Update("dozl", S.Dozl+1)
	}
	return result
}
