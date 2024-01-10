package convenio

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Convenio struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	NomeConv         string             `bson:"nomeconv,omitempty"`
	Endereco         string             `bson:"endereco,omitempty"`
	DataContratoConv time.Time          `bson:"data_contrato_conv,omitempty"`
	Indisponivel     bool               `bson:"indisponivel"` // Default: Indisponível=FALSE(NIL) -> Disponível
}

// Checar convenios
func ChecarConvenio(conv Convenio) error {
	// Verificar se os campos do convênio qualquer não estão zerados
	if conv.ID.IsZero() {
		return errors.New("ID do Convênio está zerado.")
	} else if conv.NomeConv == "" {
		return errors.New("Nome do Convênio está vazio.")
	} else if conv.Endereco == "" {
		return errors.New("Endereço do Convêncio está vazia.")
	} else if conv.DataContratoConv.IsZero() {
		return errors.New("Data do Contrato do Convêncio está zerada.")
	} else if conv.Indisponivel {
		return errors.New("Convenio não está mais disponível.")
	} else if conv.DataContratoConv.Before(time.Now()) {
		return errors.New("Convênio está com a data de contrato vencida.")
	} else {
		return nil
	}
}
