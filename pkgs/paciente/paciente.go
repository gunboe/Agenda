package paciente

import (
	"Agenda/pkgs/common"
	"errors"
)

type Paciente struct {
	Nome      string
	CPF       string
	NrCelular int
	Email     string
	Endereco  string
	// PlanoSaude []planosaude.PlanoSaude
	secret    string
	Bloqueado bool
}

func VerificarPaciente(pac Paciente) error {
	if pac.Nome == "" || pac.CPF == "" || pac.NrCelular == 0 || pac.secret == "" || !pac.Bloqueado { //|| pac.PlanoSaude == nil {
		return errors.New("Nome, CPF, NrCelular ou Secret está vazio/zerado ou Paciente Bloqueado.")
	} else if !common.CPFvalido(pac.CPF) {
		return errors.New("CPF inválido.")
		// } else {
		// for _, v := range pac.PlanoSaude {
		// err := planosaude.VerificarPlano(v)
		// if err != nil {
		// 	return err
		// }
		// }
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
