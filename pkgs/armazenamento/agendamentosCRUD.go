package armazenamento

import (
	"Agenda/pkgs/agendamento"
	"Agenda/pkgs/common"
)

// CRUD Agendamentos
func GravarAgendamento(conf common.Config, ag agendamento.Agendamento) (interface{}, error) {
	client, err := ConnectMongo(conf)
	if err != nil {
		return nil, err
	}
	Agendamentos = client.Database(conf.ArmazemDatabase).Collection("Agendamentos")

	result, err := Agendamentos.InsertOne(ctx, ag)
	if err != nil {
		return nil, err
	}

	return result.InsertedID, nil
}
