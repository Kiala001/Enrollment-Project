package service

import (
	"Enrollment/internal/application/events"
	"Enrollment/internal/domain/entities"
	valueobject "Enrollment/internal/domain/value_object"
	"Enrollment/internal/ports"
	"fmt"
)

type StudentService struct {
	repository    ports.RepositoryStudent
	event   *events.EventBus
}

func NewStudentService(repository ports.RepositoryStudent, event *events.EventBus) *StudentService {
	return &StudentService{repository: repository, event: event}
}

func (s *StudentService) RegisterStudent(student entities.Student) {

	s.repository.Save(student)
	fmt.Printf("Estudante %s inscrito com sucesso com ID: %s.\n", student.Name.String(), student.Id)
	s.event.Publish("Aluno Inscrito", student)
	
}

func (s *StudentService) HandleRegistrationAttempt(nameInput string, emailInput string) (entities.Student, error) {

	name, errName := valueobject.NewName(nameInput)
	if errName != nil {
		fmt.Printf("Falha ao validar o nome para input '%s': %v\n", nameInput, errName)
		return entities.Student{}, fmt.Errorf("nome inválido '%s': %w", nameInput, errName)
	}

	email, errEmail := valueobject.NewEmail(emailInput)
	if errEmail != nil {
		fmt.Printf("Falha ao validar o email para input '%s': %v\n", emailInput, errEmail)
		return entities.Student{}, fmt.Errorf("email inválido '%s': %w", emailInput, errEmail)
	}

	StudentId := s.GenerateUniqueIdStudent()

	student := entities.NewStudent(StudentId, name, email)	

	s.RegisterStudent(student)
	return student, nil
}

func (s *StudentService) GenerateUniqueIdStudent() string {
	var nextStudentID = entities.GenerateStudentID(8)
	exists := s.repository.StudentExists(nextStudentID)

	if exists {
		nextStudentID = entities.GenerateStudentID(8)
	}

	return nextStudentID
}

