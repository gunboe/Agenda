package planopgto

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PlanoPgto struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	ConvenioId   primitive.ObjectID `bson:"convenio,omitempty"`
	NrPlano      string             `bson:"nr_plano,omitempty"`
	DataValidade time.Time          `bson:"data_validade_plano,omitempty"`
	Inativo      bool               `bson:"inativo"`    // Default: Inativo=FALSE(NIL)    -> Ativo
	Particular   bool               `bson:"particular"` // Default: Particular=FALSE(NIL) -> Privado
}

// Checar Planos em Convênios
// Caso o Plano de Pagamento a ser utilizado seja "Particular",
// os campos do Plano de Pagamento devem ficar em Vazios/Zerados
func VerificarPlano(plano PlanoPgto) error {
	// Verifica os campos
	if plano.Inativo {
		return errors.New("Plano não está ativo!!")
	}
	if plano.Particular {
		if !plano.ConvenioId.IsZero() || plano.NrPlano != "" || !plano.DataValidade.IsZero() {
			return errors.New("ConvenioID, NrPlano, DataValidade ou Ativo não está(ão) vazio(s).\nPara Plano Particular todos esses campos devem estar vazios/nulos.")
		}
	} else {
		if plano.NrPlano == "" || plano.DataValidade.IsZero() {
			return errors.New("Nr do Plano ou Data de validade do Plano de Pagamento está(ão) vazio(s).")
		}
		if plano.DataValidade.Before(time.Now()) {
			return errors.New("Data de validade. Plano de Pagamento está vencido.")
		}
	}
	return nil
}
