package ports

import (
	"Enrollment/internal/domain/entities"
)

type PaymentRepository interface {
	Save(PaymentNote entities.PaymentNote)
	Count() int
}