package paciente

import (
	"Agenda/planosaude"
	"errors"
)

type Paciente struct {
	Nome       string
	CPF        string
	NrCelular  int
	Email      string
	Endereco   string
	PlanoSaude planosaude.PlanoSaude
	secret     string
}

func (p *Paciente) SetSecret(s string) error {
	if s == "" {
		return errors.New("[SetSecret] Segredo Nulo ou vazio!")
	} else {
		p.secret = s
		return nil
	}
}
func (p *Paciente) GetSecret() (string, error) {
	if p.secret == "" {
		return "", errors.New("[GetSecret] Segredo Nulo ou vazio!")
	} else {
		return p.secret, nil
	}
}
