package models

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PlanoPgto struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ConvenioId   primitive.ObjectID `bson:"convenio,omitempty" json:"convenio"`
	NrPlano      string             `bson:"nr_plano,omitempty" json:"nr_plano"`
	DataValidade time.Time          `bson:"data_validade_plano,omitempty" json:"data_validade_plano"`
	Inativo      bool               `bson:"inativo"    json:"inativo"` // Default: Inativo=FALSE(NIL)    -> Ativo
}

// (Check) Verifica se Plano de Pagamento está com os atributos corretos, mas NÃO verifica no Armazém.
// Caso o Plano de Pagamento seja "Particular", os campos  devem ficar em Vazios/Zerados.
func ChecarPlanoPgto(plano PlanoPgto) error {
	// Verifica os campos
	if plano.Inativo {
		return errors.New("Plano não está ativo!!")
	} else if plano.ID.IsZero() {
		return errors.New("Object ID (_id) do Plano está zerado.")
	} else if plano.ConvenioId.IsZero() {
		return errors.New("O Id do Convênio do Plano (ConvenioID) está zerado.")
	} else if plano.NrPlano == "" || plano.DataValidade.IsZero() {
		return errors.New("Nr do Plano ou Data de validade do Plano de Pagamento está(ão) vazio(s).")
	} else if plano.DataValidade.Before(time.Now()) {
		return errors.New("Data de validade. Plano de Pagamento está vencido.")
	}
	return nil
}
