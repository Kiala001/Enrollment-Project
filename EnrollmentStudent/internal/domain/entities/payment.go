package entities

import "time"

type PaymentNote struct {
	Name string
	StudentId string
	Price int
	Date string
}

func NewPaymentNote (name string , studentId string, price int) PaymentNote {
	return PaymentNote{
		Name: name,
		StudentId: studentId,
		Price: price,
		Date: time.Now().Format("02/01/2006 15:04:05"),
	}
}