package tests

import (
	"Enrollment/internal/adapters/inmemory"
	"Enrollment/internal/application/events"
	"Enrollment/internal/application/service"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestDeveSalvarNotaDePagamentoNoRepositorio(t *testing.T) {

	student_repository := inmemory.NewStudentRepositoryInMemory()
	classroom_repository  := inmemory.NewClassroomRepositoryInMemory()
	payment_repository  := inmemory.NewRepositoryPaymentInMemory()

	event := events.NewEventBus()

	event.On("Publish", mock.Anything, mock.Anything).Return()

	student_service := service.NewStudentService(student_repository, event)
	enroll_service := service.NewClassManagementService(classroom_repository ,event, student_repository)
	payment_service := service.NewPaymentService(event, payment_repository)

	event.Subscribe("Aluno Inscrito", enroll_service.HandleStudentRegisteredEvent)
	event.Subscribe("Turma Criada", payment_service.HandleClassCreatedEvent)

	student_service.HandleRegistrationAttempt("Rui", "rui@gmail.com")
	student_service.HandleRegistrationAttempt("Kiala", "kiala@gmail.com")

	lenght_payment_repo := payment_repository.Count()
	expected_lenght := 2

	if lenght_payment_repo != expected_lenght{
		t.Errorf("Esperava %d notas de pagamento, mas obteve %d", expected_lenght, lenght_payment_repo)
	}
}

func TestDeveDispararEventoAlunoNotificado(t *testing.T) { 

	student_repository := inmemory.NewStudentRepositoryInMemory()
	classroom_repository  := inmemory.NewClassroomRepositoryInMemory()
	payment_repository := inmemory.NewRepositoryPaymentInMemory()

	event := events.NewEventBus()

	event.On("Publish", mock.Anything, mock.Anything).Return()

	student_service := service.NewStudentService(student_repository, event)
	enroll_service := service.NewClassManagementService(classroom_repository ,event, student_repository)
	payment_service := service.NewPaymentService(event, payment_repository)

	event.Subscribe("Aluno Inscrito", enroll_service.HandleStudentRegisteredEvent)
	event.Subscribe("Turma Criada", payment_service.HandleClassCreatedEvent)

	student_service.HandleRegistrationAttempt("Rui", "rui@gmail.com")
	student_service.HandleRegistrationAttempt("Kiala", "kiala@gmail.com")

	event.AssertCalled(t, "Publish", "Alunos Notificados")
}