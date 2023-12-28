package paciente

import (
	"Agenda/lib"
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

func VerificarPaciente(pac Paciente) error {
	if pac.Nome == "" || pac.CPF == "" || pac.NrCelular == 0 || pac.Endereco == "" || pac.secret == "" || pac.PlanoSaude.Convenio.NomeConv == "" {
		return errors.New("Nome, CPF, NrCelular, Plano de Saúde ou Secret está vazio/zerado.")
	} else if !lib.CPFvalido(pac.CPF) {
		return errors.New("CPF inválido.")
	} else {
		err := planosaude.VerificarPlano(pac.PlanoSaude)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Paciente) SetSecret(s string) error {
	if s == "" {
		return errors.New("\"GetSecret\": Segredo Nulo ou vazio!")
	} else {
		p.secret = s
		return nil
	}
}
func (p *Paciente) GetSecret() (string, error) {
	if p.secret == "" {
		return "", errors.New("\"GetSecret\": Nulo ou vazio!")
	} else {
		return p.secret, nil
	}
}
