package agendamento

import (
	"Agenda/agente"
	"Agenda/paciente"
	"time"
)

type Agendamento struct {
	DataInicio      time.Time
	Duracao         time.Duration
	Atividade       string
	AgenteExecutor  agente.Agente
	PacienteAtender paciente.Paciente
	Confirmado      bool
	MeioPagamento   string
	Pago            bool
}
