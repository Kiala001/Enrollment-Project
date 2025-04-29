package service

import (
	"Enrollment/internal/application/events"
	"Enrollment/internal/domain/entities"
	"Enrollment/internal/ports"
	"fmt"
)

const IdealNumberOfStudents = 2

type ClassManagementService struct {
	classroom_repository   ports.RepositoryClassroom
	event             *events.EventBus
	student_repository ports.RepositoryStudent
}

func NewClassManagementService(repository ports.RepositoryClassroom, event *events.EventBus, repositoryStudent ports.RepositoryStudent) *ClassManagementService {
	return &ClassManagementService{classroom_repository: repository, event: event, student_repository: repositoryStudent}
}

func (c *ClassManagementService) performClassCreation(students []entities.Student) string{
	nextClassId := c.GenerateClassroomId()

	newClass := entities.NewClassRoom(nextClassId, students)
	
	c.classroom_repository.Save(newClass)

	fmt.Printf("Turma Criada com sucesso!\n")
	c.event.Publish("Turma Criada", newClass)

	return nextClassId
}

func (c *ClassManagementService) GenerateClassroomId() string {
	currentNumberOfClasses := c.classroom_repository.Count()
	nextClassId := fmt.Sprintf("TM%04d", currentNumberOfClasses + 1)

	return nextClassId
}

func (c *ClassManagementService) HandleStudentRegisteredEvent(payload any) {
	currentStudentCount:= c.student_repository.CountStudentWithoutClass()

	if currentStudentCount == IdealNumberOfStudents{
		fmt.Printf("Número ideal de alunos atingido. Criando turma...\n")
		
		allStudents := c.student_repository.GetAllStudents()
		classroom_id := c.performClassCreation(allStudents)
		c.student_repository.UpdateStudentWithoutClass(classroom_id)
	}
}





