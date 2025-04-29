package inmemory

import (
	"Enrollment/internal/domain/entities"
)

type RepositoryPaymentInMemory struct {
	payment map[string] entities.PaymentNote
}

func NewRepositoryPaymentInMemory() *RepositoryPaymentInMemory {
	return &RepositoryPaymentInMemory{payment: make(map[string]entities.PaymentNote)}
}

func (r *RepositoryPaymentInMemory) Save(PaymentNote entities.PaymentNote) {
	r.payment[PaymentNote.StudentId] = PaymentNote 
}

func (r *RepositoryPaymentInMemory) Count() int {
	return len(r.payment)
}
