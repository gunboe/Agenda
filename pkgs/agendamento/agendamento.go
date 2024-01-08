package agendamento

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Agendamento struct {
	ID          primitive.ObjectID `bson:"_id"`
	DataInicio  time.Time          `bson:"data_inicio"`
	Duracao     time.Duration      `bson:"duracao"`
	Atividade   string             `bson:"atividade"`
	PacienteID  primitive.ObjectID `bson:"paciente"`
	AgenteID    primitive.ObjectID `bson:"agente"`
	Confirmado  bool               `bson:"confirmado"`
	PlanoPgtoID primitive.ObjectID `bson:"planopgto"`
	ValorPago   float32            `bson:"valor_pago"`
	Pago        bool               `bson:"pago"`
	Cancelado   bool               `bson:"cancelado"`
}
