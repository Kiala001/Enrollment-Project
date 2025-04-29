package db

import (
	"Enrollment/internal/config"
	"Enrollment/internal/domain/entities"
	"log"
)

type SQLitePaymentRepository struct {
	*config.Database
}

func NewSQLitePaymentRepository (db *config.Database) (*SQLitePaymentRepository) {
	repo := &SQLitePaymentRepository{Database: db}
	return repo
}

func (r *SQLitePaymentRepository) Save(paymentNote entities.PaymentNote) {
	_, err := r.Db.Exec(`
	INSERT INTO payment_note (student_id, price, date) VALUES (?, ?, ?) 
	`, paymentNote.StudentId, paymentNote.Price, paymentNote.Date)	

	if err != nil {
		log.Fatal(err)
	}
}

func (r *SQLitePaymentRepository) Count() int {
	var count int
	err := r.Db.QueryRow(`SELECT COUNT(*) FROM payment_note`).Scan(&count)

	if err != nil {
		log.Fatal(err)
	}

	return count
}