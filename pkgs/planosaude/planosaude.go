// planosaunde.go
package planosaude

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PlanoSaude struct {
	Convenio     Convenios `bson:"nome_conv"`
	NrPlano      string    `bson:"nr_plano"`
	DataValidade time.Time `bson:"data_validade_plano"`
}

type Convenios struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	NomeConv         string             `bson:"plano,omitempty"`
	Endereco         string             `bson:"endereco,omitempty"`
	DataContratoConv time.Time          `bson:"data_contrato_conv,omitempty"`
	Disponivel       bool               `bson:"disponivel,omitempty"`
}

// Regras de Negocio dos Planos e Convenios de Saude
// Convenios
func VerificarConvenio(c Convenios) error {
	if c.NomeConv == "" || c.Endereco == "" || c.DataContratoConv.IsZero() {
		return errors.New("Nome, Número ou Data de validade vazia.")
	} else if !c.Disponivel {
		return errors.New("Convenio não está mais disponível.")
	} else if c.DataContratoConv.Before(time.Now()) {
		return errors.New("Convênio está com a data de contrato vencida.")
	} else {
		return nil
	}
}

// Checar Planos em Convênios
func VerificarPlano(plano PlanoSaude) error {
	// Verifica os campos
	if plano.Convenio.NomeConv == "" || plano.NrPlano == "" || plano.DataValidade.IsZero() {
		return errors.New("Nome, Número ou Data de validade vazio.")
		// Verifica a validado do Plano
	} else if plano.DataValidade.Before(time.Now()) {
		return errors.New("Data de validade, Plano de saude vencido.")
		// Verifica a validade do Contrato do Convênio
	} else if plano.Convenio.DataContratoConv.Before(time.Now()) {
		return errors.New("Não possível usar o Plano. Contratro do Convênio: " + plano.Convenio.NomeConv + " está vencido desde:" + plano.Convenio.DataContratoConv.Format("02/01/2006"))
	} else {
		return nil
	}
}
