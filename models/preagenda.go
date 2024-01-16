package models

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PreAgenda struct {
	ID         primitive.ObjectID
	AgenteID   primitive.ObjectID
	DataInicio time.Time
	DataFinal  time.Time
	HoraInicio time.Duration
	HoraFinal  time.Duration
	Liberada   bool
}

// Regras de Negocio dos Agesntes e PreAgenda
// PreAgenda
func VerificarPreAgenda(pa PreAgenda) error {
	var err error
	// err := VerificarAgente(pa.AgenteID)
	if err != nil {
		return err
	} else if pa.DataFinal.Before(pa.DataInicio) {
		return errors.New("Data Final não pode ser anterior a Data Inicial.")
	} else if pa.DataFinal.Before(time.Now()) {
		return errors.New("Data Final não pode ser anterior ao dia atual.")
	} else {
		return nil
	}
}
