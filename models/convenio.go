package models

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Convenio struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	NomeConv         string             `bson:"nomeconv,omitempty" json:"nomeconv" binding:"required"`
	NrPrestador      string             `bson:"nrprestador,omitempty" json:"nrprestador" binding:"required"`
	Endereco         string             `bson:"endereco,omitempty" json:"endereco" binding:"required"`
	DataContratoConv time.Time          `bson:"data_contrato_conv,omitempty" json:"data_contrato_conv" binding:"required,datavencida"`
	// Indisponivel: Default = FALSE(NIL) -> Disponível
	Indisponivel bool `bson:"indisponivel" json:"indisponivel"`
}

type ConvenioResp struct {
	ID         primitive.ObjectID `json:"id,omitempty"`
	Criado     bool               `json:"criado,omitempty"`
	Inserido   bool               `json:"inserido,omitempty"`
	Atualizado bool               `json:"atualizado,omitempty"` // Serve também para o atributo "Indisponivel"
	Deletado   bool               `json:"deletado,omitempty"`
}

// Checar convenios
func ChecarConvenio(conv Convenio) error {
	// Verificar se os campos do convênio qualquer não estão zerados
	if conv.NomeConv == "" {
		return errors.New("Nome do Convênio está vazio")
	} else if conv.NrPrestador == "" {
		return errors.New("Nr Prestador do Convêncio está vazio")
	} else if conv.Endereco == "" {
		return errors.New("Endereço do Convêncio está vazia")
	} else if conv.DataContratoConv.IsZero() {
		return errors.New("Data do Contrato do Convêncio está zerada")
	} else if conv.Indisponivel {
		return errors.New("Convenio não está disponível")
	} else if conv.DataContratoConv.Before(time.Now()) {
		return errors.New("Convênio está com a data de contrato vencida")
	} else {
		return nil
	}
}
