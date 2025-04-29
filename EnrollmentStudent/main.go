package main

import (
	"Enrollment/internal/adapters/cli"
	"Enrollment/internal/adapters/db"
	"Enrollment/internal/application/events"
	"Enrollment/internal/application/service"
	"Enrollment/internal/config"
	"Enrollment/internal/ports"

	"github.com/stretchr/testify/mock"
)

func main() {

	event := events.NewEventBus()
	event.On("Publish", mock.Anything, mock.Anything).Return()

	database := config.NewDatabase()

	studentRepo := db.NewSQLiteStudentRepository(database)
	usecase := service.NewStudentService(studentRepo, event)

	var ClassRepo ports.RepositoryClassroom = db.NewSQLiteClassroomRepository(database)
	var paymentRepo ports.PaymentRepository = db.NewSQLitePaymentRepository(database)
	classroom_service := service.NewClassManagementService(ClassRepo, event, studentRepo)
	payment_service := service.NewPaymentService(event, paymentRepo)

	CLI := cli.NewCLI(usecase, studentRepo, event, classroom_service, ClassRepo, payment_service)

	CLI.Run()
}