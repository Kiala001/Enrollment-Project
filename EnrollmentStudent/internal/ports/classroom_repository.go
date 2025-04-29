package ports

import "Enrollment/internal/domain/entities"


type RepositoryClassroom interface {
	Save(ClassRoom entities.ClassRoom)
	Count() int
}