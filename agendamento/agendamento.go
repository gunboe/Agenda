package agendamento

import (
	"Agenda/agente"
	"Agenda/paciente"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Agendamento struct {
	ID             primitive.ObjectID `bson:"_id"`
	DataInicio     time.Time          `bson:"data_inicio"`
	Duracao        time.Duration      `bson:"duracao"`
	Atividade      string             `bson:"atividade"`
	AgenteExecutor agente.Agente      `bson:"agente_executor"`
	PacienteAtend  paciente.Paciente  `bson:"paciente_atendido"`
	Confirmado     bool               `bson:"confirmado"`
	MeioPagamento  string             `bson:"meio_pagamento"`
	Pago           bool               `bson:"pago"`
	Cancelado      bool               `bson:"cancelado"`
}
