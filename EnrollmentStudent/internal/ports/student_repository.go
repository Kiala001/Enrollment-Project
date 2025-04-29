package ports

import (
	"Enrollment/internal/domain/entities"
)


type RepositoryStudent interface {
	Save(student entities.Student)
	Count() int
	CountStudentWithoutClass() int
	StudentExists(id string) bool
	GetAllStudents() []entities.Student
	UpdateStudentWithoutClass(classroomID string)
}
