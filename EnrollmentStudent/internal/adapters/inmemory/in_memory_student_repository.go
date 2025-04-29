package inmemory

import "Enrollment/internal/domain/entities"

type RepositoryStudentInMemory struct {
	students map[string]entities.Student
}

func NewStudentRepositoryInMemory() *RepositoryStudentInMemory {
	return &RepositoryStudentInMemory{students: make(map[string]entities.Student)}
}

func (r *RepositoryStudentInMemory) Save(student entities.Student) {
	r.students[student.Id] = student
}

func (r *RepositoryStudentInMemory) Count() int {
	return len(r.students)
}

func (r *RepositoryStudentInMemory) StudentExists(id string) bool{
	_, exists := r.students[id]
    return exists
}

func (r *RepositoryStudentInMemory) GetAllStudents() []entities.Student{
	allStudents:= make([]entities.Student, 0, len(r.students))

	for _, student := range r.students{
		allStudents = append(allStudents, student)
	}
	return allStudents
}

func (r *RepositoryStudentInMemory) CountStudentWithoutClass() int{return len(r.students)}

func (r *RepositoryStudentInMemory) UpdateStudentWithoutClass(classroomID string) {}