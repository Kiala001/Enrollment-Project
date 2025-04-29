package cli

import (
	"Enrollment/internal/application/events"
	"Enrollment/internal/application/service"
	"Enrollment/internal/domain/dto"
	"Enrollment/internal/ports"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type CLI struct {
	usecase *service.StudentService
	Event *events.EventBus
	service *service.ClassManagementService
	studentRepo ports.RepositoryStudent
	classRoomRepo ports.RepositoryClassroom
	paymentService *service.PaymentService
}

func NewCLI(usecase *service.StudentService, studentRepo ports.RepositoryStudent, Event *events.EventBus, service *service.ClassManagementService, classRoomRepo ports.RepositoryClassroom, payment *service.PaymentService) *CLI {
	return &CLI{usecase: usecase, studentRepo: studentRepo, Event: Event, service: service, classRoomRepo: classRoomRepo, paymentService: payment}
}

func (cli *CLI) Run(){
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Digite seu nome: ")
	nameInput, _ := reader.ReadString('\n')
	nameInput = strings.TrimSpace(nameInput)

	fmt.Print("Digite seu email: ")
	emailInput, _ := reader.ReadString('\n')
	emailInput = strings.TrimSpace(emailInput)

	dto := dto.StudentDTO{
		Name:  nameInput,
		Email: emailInput,
	}

	cli.Event.Subscribe("Aluno Inscrito", cli.service.HandleStudentRegisteredEvent)
	cli.Event.Subscribe("Turma Criada", cli.paymentService.HandleClassCreatedEvent)

	cli.usecase.HandleRegistrationAttempt(dto.Name, dto.Email)

	fmt.Println("Turmas Criadas: ",cli.classRoomRepo.Count())

	allStudents := cli.studentRepo.GetAllStudents()

	fmt.Println("Estudantes:")
	for _, student := range allStudents{
		fmt.Println(" ID: ", student.Id," Nome: ", student.Name.String()," Email: ", student.Email.String())
	}
}