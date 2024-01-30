package models

import (
	"Agenda/common"
	"errors"
)

type Agente struct {
	Nome           string
	CPF            string
	NrCelular      int
	Especialidades []string
	secret         string
	CanaisPrefer   []string
}

func (ag *Agente) SetSecret(s string) error {
	if s == "" {
		return errors.New("[SetSecret] Segredo Nulo ou vazio!")
	} else {
		ag.secret = s
		return nil
	}
}

func (ag *Agente) GetSecret() (string, error) {
	if ag.secret == "" {
		return "", errors.New("[GetSecret] Segredo Nulo ou vazio!")
	} else {
		return ag.secret, nil
	}
}

// Agente
func VerificarAgente(ag Agente) error {
	agsecret, err := ag.GetSecret()
	if err != nil {
		return err
	}
	if ag.Nome == "" || ag.CPF == "" || ag.NrCelular == 0 || ag.Especialidades == nil || agsecret == "" {
		return errors.New("Nome, CPF, NrCelular, Especialidades ou Secret está vazio/zerado.")
	} else if _, ok := common.CPFvalido(ag.CPF); !ok {
		return errors.New("CPF inválido.")
	}
	return nil
}
