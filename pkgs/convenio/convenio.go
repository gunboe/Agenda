package convenio

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Convenio struct {
	ID               primitive.ObjectID `bson:"_id,omitempty""`
	NomeConv         string             `bson:"nomeconv,omitempty"`
	Endereco         string             `bson:"endereco,omitempty"`
	DataContratoConv time.Time          `bson:"data_contrato_conv,omitempty"`
	Disponivel       bool               `bson:"disponivel"`
}

// Checar convenios
func VerificarConvenio(conv Convenio) error {
	// Verificar se os campos do convênio qualquer não estão zerados
	if conv.NomeConv == "" || conv.Endereco == "" || conv.DataContratoConv.IsZero() {
		return errors.New("Nome, Endereço, Data do Contrato ou Disponibilidade do Convêncio vazia/zerada/falso.")
	} else if !conv.Disponivel {
		return errors.New("Convenio não está mais disponível.")
	} else if conv.DataContratoConv.Before(time.Now()) {
		return errors.New("Convênio está com a data de contrato vencida.")
	} else {
		return nil
	}
}
