package agendaliberada

import (
	"Agenda/pkgs/agente"
	"errors"
	"time"
)

type AgendaLiberada struct {
	Agente     agente.Agente
	DataInicio time.Time
	DataFinal  time.Time
	HoraInicio time.Duration
	HoraFinal  time.Duration
	Liberada   bool
}

// Regras de Negocio dos Agesntes e AgendaLiberada
// AgendaLiberada
func VerificarAgendaLiberada(al AgendaLiberada) error {
	err := agente.VerificarAgente(al.Agente)
	if err != nil {
		return err
	} else if al.DataFinal.Before(al.DataInicio) {
		return errors.New("Data Final não pode ser anterior a Data Inicial.")
	} else if al.DataFinal.Before(time.Now()) {
		return errors.New("Data Final não pode ser anterior ao dia atual.")
	} else {
		return nil
	}
}
