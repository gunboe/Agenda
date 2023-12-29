package convenio

import (
	"errors"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Convenios struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	NomeConv         string             `bson:"plano,omitempty"`
	Endereco         string             `bson:"endereco,omitempty"`
	DataContratoConv time.Time          `bson:"data_contrato_conv,omitempty"`
	Disponivel       bool               `bson:"disponivel,omitempty"`
}

// Checar convenios
// Se não houver Convenio/Plano, o atendimento é Particular,
// os campos a serem utilizados são:
// 		NomeConv:         Particular
// 		Endereco:         <vazio>
// 		DataContratoConv: 0
// 		Disponivel:       true
func VerificarConvenio(conv Convenios) error {
	// Testa se o Convênio é Particular, os outros campos devem estar no padrão "Particular"
	if strings.EqualFold(conv.NomeConv, "Particular") {
		if conv.Endereco != "" || !conv.DataContratoConv.IsZero() || !conv.Disponivel {
			return errors.New("O Plano *Particular* deve estar com Endereço:nil, DataContratoConv:0 e Disponivel:true.")
		}
		return nil
	} else {
		// Verificar os campos
		if conv.NomeConv == "" || conv.Endereco == "" || conv.DataContratoConv.IsZero() || !conv.Disponivel {
			return errors.New("Nome, Número ou Data de validade vazia.")
		} else if !conv.Disponivel {
			return errors.New("Convenio não está mais disponível.")
		} else if conv.DataContratoConv.Before(time.Now()) {
			return errors.New("Convênio está com a data de contrato vencida.")
		} else {
			return nil
		}
	}
}
