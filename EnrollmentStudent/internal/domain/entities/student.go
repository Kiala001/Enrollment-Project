package entities

import (
	valueobject "Enrollment/internal/domain/value_object"
	"math/rand"
	"time"
)

const chartSet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type Student struct {
	Id string
	Name valueobject.Name
	Email valueobject.Email
}

func NewStudent (id string, name valueobject.Name, email valueobject.Email) Student {
	return Student{
		Id: id,
		Name: name,
		Email: email,
	}
}

func GenerateStudentID(length int) string {
	rand.Seed(time.Now().UnixNano())
	id := make([]byte, length)
	for i := range id{
		id[i] = chartSet[rand.Intn(len(chartSet))]
	}
	return "ST-" + string(id)
}
