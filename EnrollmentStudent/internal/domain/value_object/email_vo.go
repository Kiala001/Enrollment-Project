package valueobject

import (
    "errors"
    "regexp"
	"net/mail"
    "strings"
)

type Email struct {
    Valor string
}

var (
    EmailRegex = regexp.MustCompile(`^[A-Za-z0-9._]+@[a-z0-9./-]+\.[a-z]{2,}$`)
)

func NewEmail(input string) (Email, error) {
    input = strings.TrimSpace(input)

    if input == "" {
        return Email{}, errors.New("email não pode ser vazio")
        
    }

    if !EmailRegex.MatchString(input) {
        return Email{}, errors.New("o formato do email está incorrecto")
    }     

    _, err := mail.ParseAddress(input)
    if err != nil {
        return Email{}, errors.New("email inválido")
    }

    return Email{Valor: input}, nil
}

func (e Email) String() string {
    return e.Valor
}