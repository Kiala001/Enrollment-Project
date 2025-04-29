package tests

import (
	"Enrollment/internal/adapters/inmemory"
	"Enrollment/internal/application/events"
	"Enrollment/internal/application/service"
	"testing"
)

func TestDeveSalvarAlunoNoRepositorio(t *testing.T) {
	name := "Rui"
	email := "ruimanuel12@gmail.com"

	event := new(events.EventBus)
	event.On("Publish", "Aluno Inscrito").Return()

	student_repository := inmemory.NewStudentRepositoryInMemory()
	student_service := service.NewStudentService(student_repository, event)
	student_service.HandleRegistrationAttempt(name, email)

	length := student_repository.Count()
	if length != 1 {
		t.Errorf("Esperava %d, mas obteve %d", 1, length)
	}
}

func TestNaoDeveSalvarAlunoNoRepositorioSeOFormatoEstiverErrado(t *testing.T) {
	name := "Rui"
	email := "ruimanuel12gmail.com"

	event := new(events.EventBus)
	event.On("Publish", "Aluno Inscrito").Return()

	student_repository := inmemory.NewStudentRepositoryInMemory()
	student_service := service.NewStudentService(student_repository, event)
	student_service.HandleRegistrationAttempt( name, email)

	length := student_repository.Count()
	if length != 0 {
		t.Errorf("Esperava %d, mas obteve %d", 0, length)
	}
}
func TestDeveDispararEventoAlunoInscrito(t *testing.T) {
	name := "Rui"
	email :="rui@gmail.com"

	event := new(events.EventBus)

	event.On("Publish", "Aluno Inscrito").Return()

	student_repository := inmemory.NewStudentRepositoryInMemory()
	student_service := service.NewStudentService(student_repository, event)
	student_service.HandleRegistrationAttempt( name, email)

	event.AssertCalled(t, "Publish", "Aluno Inscrito")
}
