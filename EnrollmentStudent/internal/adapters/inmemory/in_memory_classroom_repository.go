package inmemory

import (
	"Enrollment/internal/domain/entities"
)

type RepositoryClassroomInMemory struct {
	classes map[string]entities.ClassRoom
}

func NewClassroomRepositoryInMemory() *RepositoryClassroomInMemory {
	return &RepositoryClassroomInMemory{classes: make(map[string]entities.ClassRoom)}
}

func (r *RepositoryClassroomInMemory) Save(ClassRoom entities.ClassRoom) {
	r.classes[ClassRoom.ClassId] = ClassRoom
}

func (r *RepositoryClassroomInMemory) Count() int {
	return len(r.classes)
} 
