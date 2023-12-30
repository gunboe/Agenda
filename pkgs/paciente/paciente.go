package paciente

import (
	"Agenda/pkgs/common"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Paciente struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Nome      string             `bson:"nome,omitempty"`
	CPF       string             `bson:"cpf,omitempty"`
	NrCelular int                `bson:"nr_celular,omitempty"`
	Email     string             `bson:"email,omitempty"`
	Endereco  string             `bson:"endereco,omitempty"`
	secret    string             `bson:"secret,omitempty"`
	Bloqueado bool               `bson:"bloqueado,omitempty"`
	// PlanoSaude []planosaude.PlanoSaude
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
