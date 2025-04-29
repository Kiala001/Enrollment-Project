package tests

import (
	"Enrollment/internal/adapters/inmemory"
	"Enrollment/internal/application/events"
	"Enrollment/internal/application/service"
	"Enrollment/internal/domain/entities"
	valueobject "Enrollment/internal/domain/value_object"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestDeveVerificarSeOEventoTurmaCriadaFoiDisparado(t *testing.T) {

	event := events.NewEventBus()
	event.On("Publish", "Turma Criada", mock.Anything).Return()

	classroom_repository := inmemory.NewClassroomRepositoryInMemory()
	student_repository := inmemory.NewStudentRepositoryInMemory()
	enroll_service := service.NewClassManagementService(classroom_repository , event, student_repository)

	aluno1 := entities.Student{Id: "ST0001", Name: valueobject.Name{Valor: "Rui"}, Email: valueobject.Email{Valor: "rui@gmail.com"}}
	aluno2 := entities.Student{Id: "ST0002", Name: valueobject.Name{Valor: "Kiala"}, Email: valueobject.Email{Valor: "kiala@gmail.com"}}

	student_repository.Save(aluno1)
	student_repository.Save(aluno2)

	enroll_service.HandleStudentRegisteredEvent(aluno2)

	event.AssertCalled(t, "Publish", "Turma Criada", mock.Anything)
} 

func TestDeveSalvarTurmaNoRepositorioQuandoNumeroIdealForAtingido(t *testing.T) {

	student_repository := inmemory.NewStudentRepositoryInMemory()
	classroom_repository := inmemory.NewClassroomRepositoryInMemory()
	payment_repository := inmemory.NewRepositoryPaymentInMemory()

	eventBus := events.NewEventBus()
	eventBus.On("Publish", mock.Anything, mock.Anything).Return()

	student_service:= service.NewStudentService(student_repository, eventBus)
	enroll_service := service.NewClassManagementService(classroom_repository , eventBus, student_repository)
	payment_service := service.NewPaymentService(eventBus, payment_repository)

	eventBus.Subscribe("Aluno Inscrito", enroll_service.HandleStudentRegisteredEvent)
	eventBus.Subscribe("Turma Criada", payment_service.HandleClassCreatedEvent)
		
	studentName1 := "Rui"
	studentEmail1 := "rui@gmail.com"

	studentName2 := "Kiala"
	studentEmail2 := "kiala@gmail.com"

	_, err1 := student_service.HandleRegistrationAttempt(studentName1, studentEmail1)
	if err1 != nil{
		t.Fatal("Falha ao registrar ",studentName1)
	}

	_, erro2 := student_service.HandleRegistrationAttempt(studentName2, studentEmail2)
	if erro2 != nil{
		t.Fatal("Falha ao registrar ",studentName2)
	}

	repolenght := classroom_repository.Count()
	expectedLenght := 1

	if repolenght != expectedLenght{
		t.Errorf("Esperava %d, mas obteve %d", expectedLenght, repolenght)
	}	
}

func TestNaoDeveCriarTurmaSeNumeroDeEstudantesForInsuficiente(t *testing.T){

	student_repository := inmemory.NewStudentRepositoryInMemory()
	classroom_repository  := inmemory.NewClassroomRepositoryInMemory()
	payment_repository := inmemory.NewRepositoryPaymentInMemory()

	eventBus := events.NewEventBus()

	eventBus.On("Publish", mock.Anything, mock.Anything).Return()

	student_service := service.NewStudentService(student_repository, eventBus)

	student_service.HandleRegistrationAttempt("Rui", "Rui@gmail.com")

	enroll_service := service.NewClassManagementService(classroom_repository , eventBus, student_repository)
	payment_service := service.NewPaymentService(eventBus, payment_repository)

	eventBus.Subscribe("Aluno Inscrito", enroll_service.HandleStudentRegisteredEvent)
	eventBus.Subscribe("Turma Criada", payment_service.HandleClassCreatedEvent)

	eventBus.AssertNotCalled(t,"Publish", "Turma Criada")

	count_classroom_repo := classroom_repository.Count()
	expected_count := 0

	if count_classroom_repo != expected_count{
		t.Errorf("Esperava %d turmas, mas obteve %d", expected_count, count_classroom_repo)
	}
}


