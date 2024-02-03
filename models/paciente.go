package models

import (
	"Agenda/common"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Paciente struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nome         string             `bson:"nome,omitempty" json:"nome" binding:"required"`
	CPF          string             `bson:"cpf,omitempty" json:"cpf" binding:"required,valcpf"`
	NrCelular    int64              `bson:"nr_celular,omitempty" json:"nr_celular" binding:"required,valnrcelular"`
	Email        string             `bson:"email,omitempty" json:"email" binding:"email"`
	Endereco     string             `bson:"endereco,omitempty" json:"endereco"`
	PlanosPgts   []PlanoPgto        `bson:"planospgts,omitempty" json:"planospgts"`
	CanaisPrefer []string           `bson:"canais_prefer,omitempty" json:"canais_prefer"`
	// Campo "bloqueado": Default = FALSE(NIL) -> Não-Bloqueado
	Bloqueado bool     `bson:"bloqueado" json:"bloqueado"`
	secret    string   `bson:"secret,omitempty" json:"secret"`
	obs       []string `bson:"obs,omitempty" json:"obs"`
}

// Função de Validação do Paciente
func ChecarPaciente(pac Paciente) error {
	if pac.Bloqueado {
		return errors.New("Paciente (" + pac.Nome + ") está bloqueado, consulte as observações do Paciente")
	}
	if pac.Nome == "" || pac.CPF == "" || pac.NrCelular == 0 {
		return errors.New("Nome, CPF ou NrCelular está vazio/zerado")
	} else if _, ok := common.CPFvalido(pac.CPF); !ok {
		return errors.New("CPF inválido")
	} else if !common.NrCelValido(pac.NrCelular) {
		return errors.New("NrCelular no formato inválido")
	} else if pac.Email != "" && !common.EmailValido(pac.Email) {
		return errors.New("Email no formato inválido")
	}
	// Verificar os PlanosPgto
	for i, v := range pac.PlanosPgts {
		err := ChecarPlanoPgto(v)
		if err != nil {
			return err
		}
		// Checar PlanoPgto duplicados neste objeto Paciente (obs: Não checa os já adicionados no banco!)
		for j := i + 1; j < len(pac.PlanosPgts); j++ {
			if pac.PlanosPgts[i].ConvenioId == pac.PlanosPgts[j].ConvenioId && pac.PlanosPgts[i].NrPlano == pac.PlanosPgts[j].NrPlano {
				return errors.New("PlanoPgto Nr:" + pac.PlanosPgts[i].NrPlano + " já existe, está duplicado")
			}
		}
	}
	return nil
}

// Função privada manippulação da Secret
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

// Função privada de manipulação de Obs (Observações)
func (p *Paciente) SetObs(obs string) error {
	if obs == "" {
		return errors.New("\"SetObs\": Nula ou vazia!")
	} else {
		p.obs = append(p.obs, obs)
		return nil
	}
}
func (p *Paciente) GetObs() []string {
	return p.obs
}
