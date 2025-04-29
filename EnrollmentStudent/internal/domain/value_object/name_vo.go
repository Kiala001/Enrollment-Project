package valueobject

import (
    "errors"
    "regexp"
    "strings"
)

type Name struct {
    Valor string
}

var (
    NameRegex = regexp.MustCompile(`^[A-Za-zÀ-ÖØ-öø-ÿ\s]+$`)
)

func NewName(input string) (Name, error) {
    input = strings.TrimSpace(input)

    if input == "" {
        return Name{}, errors.New("Name não pode ser vazio")
    }

    if !NameRegex.MatchString(input) {
        return Name{}, errors.New("Name inválido: não pode conter números ou caracteres especiais")
    }

    return Name{Valor: input}, nil
}

func (n Name) String() string {
    return n.Valor
}
