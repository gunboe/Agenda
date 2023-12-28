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
// func VerificarPlano(plano string, nr string, data time.Time, convenios []Convenios) (PlanoSaude, error) {
func VerificarPlano(plano PlanoSaude) error {

	// convenios, err := armazenamento.GetConvenios("*")
	// if err != nil {
	// 	return err
	// }
	// for _, conv := range convenios {
	// if strings.EqualFold(conv.NomeConv, plano.Convenio.NomeConv) {
	if plano.Convenio.NomeConv == "" || plano.NrPlano == "" || plano.DataValidade.IsZero() {
		return errors.New("Nome, Número ou Data de validade vazio.")
	} else if plano.DataValidade.Before(time.Now()) {
		return errors.New("Data de validade, Plano de saude vencido.")
		// } else if conv.DataContratoConv.Before(time.Now()) {
		// 	return errors.New("Não possível usar o Plano. Contratro do Convênio: " + conv.NomeConv + " está vencido.")
	} else {
		return nil
	}
	// }
	// }
	// return errors.New("Plano não conveniado.")
}
