package service

import (
	"Enrollment/internal/application/events"
	"Enrollment/internal/domain/entities"
	"Enrollment/internal/ports"
	"fmt"
)

const Amount = 10000

type PaymentService struct {
	payment_repository  ports.PaymentRepository
	event             *events.EventBus
}

func NewPaymentService(event *events.EventBus, repositoryPayment ports.PaymentRepository) *PaymentService {
	return &PaymentService{event: event, payment_repository: repositoryPayment}
}

func (c *PaymentService) GeneratePaymentNotes(students [] entities.Student){
	for _, student := range students{
		note := entities.NewPaymentNote(student.Name.String(), student.Id, Amount)
		c.SavePaymentOrder(note)
		fmt.Printf("Nota de pagamento para %s gerada com id %s, valor %d e data %s.\n", note.Name, note.StudentId, note.Price, note.Date)
	}
}

func (c *PaymentService) SavePaymentOrder(note entities.PaymentNote){
	c.payment_repository.Save(note)
}

func (c *PaymentService) NotifyStudents(students []entities.Student){
	for _, student:= range students{
		fmt.Println("Email enviado para " + student.Email.String())
	}
	c.event.Publish("Alunos Notificados", map[string]any{
		"students": students,
	})
}

func (c *PaymentService) HandleClassCreatedEvent(payload any) {

	ClassRoom, Error := payload.(entities.ClassRoom)

	if !Error {
		fmt.Println("Payload inesperado para o evento Turma Criada.")
		return
	}

	c.GeneratePaymentNotes(ClassRoom.Students)
	c.NotifyStudents(ClassRoom.Students)
}