package models

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PlanoPgto struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ConvenioId   primitive.ObjectID `bson:"convenio,omitempty" json:"convenio" binding:"required"`
	NrPlano      string             `bson:"nr_plano,omitempty" json:"nr_plano" binding:"required"`
	DataValidade time.Time          `bson:"data_validade_plano,omitempty" json:"data_validade_plano"  binding:"required,datavencida"`
	Inativo      bool               `bson:"inativo" json:"inativo"` // Default: Inativo=FALSE(NIL)    -> Ativo
}

// (Check) Checa os Campos do Plano de Pagamento com os atributos corretos,
// mas NÃO verifica dados do Nergócio (no Armazém).
func ChecarPlanoPgto(plano PlanoPgto) error {
	// Verifica os campos
	if plano.ConvenioId.IsZero() {
		return errors.New("o id do convênio do plano (convenioid) está zerado")
	} else if plano.NrPlano == "" {
		return errors.New("nr do plano de pagamento está vazio")
	} else if plano.DataValidade.IsZero() {
		return errors.New("data de validade do plano de pagamento está vazio")
	}
	return nil
}
