package paciente

import (
	"Agenda/pkgs/common"
	"Agenda/pkgs/planopgto"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Paciente struct {
	ID           primitive.ObjectID    `bson:"_id,omitempty"`
	Nome         string                `bson:"nome,omitempty"`
	CPF          string                `bson:"cpf,omitempty"`
	NrCelular    int64                 `bson:"nr_celular,omitempty"`
	Email        string                `bson:"email,omitempty"`
	Endereco     string                `bson:"endereco,omitempty"`
	PlanosPgts   []planopgto.PlanoPgto `bson:"planospgts,omitempty"`
	CanaisPrefer []string              `bson:"canais_prefer,omitempty"`
	Bloqueado    bool                  `bson:"bloqueado"` // Default: Bloqueado=FALSE(NIL) -> Não-Bloqueado
	secret       string                `bson:"secret,omitempty"`
	obs          []string              `bson:"obs,omitempty"`
}

// Função de Validação do Paciente
func ChecarPaciente(pac Paciente) error {
	if pac.Bloqueado {
		return errors.New("Paciente (" + pac.Nome + ") está Bloqueado, consulte as observações do Paciente.")
	}
	if pac.Nome == "" || pac.CPF == "" || pac.NrCelular == 0 {
		return errors.New("Nome, CPF ou NrCelular  está vazio/zerado ou Paciente Bloqueado.")
	} else if !common.CPFvalido(pac.CPF) {
		return errors.New("CPF inválido.")
	} else if !common.NrCelValido(pac.NrCelular) {
		return errors.New("NrCelular no formato inválido.")
	} else if pac.Email != "" && !common.EmailValido(pac.Email) {
		return errors.New("Email no formato inválido.")
	}
	// TODO (TESTAR) Verificar os PlanosPgto dele
	for _, v := range pac.PlanosPgts {
		err := planopgto.ChecarPlanoPgto(v)
		if err != nil {
			return err
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
