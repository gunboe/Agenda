package mdg

import (
	"Agenda/models"
	"Agenda/services/config"
)

// CRUD Agendamentos
func GravarAgendamento(conf config.Config, ag models.Agendamento) (interface{}, error) {
	// client, err := ConnectMongo(conf)
	// if err != nil {
	// 	return nil, err
	// }
	// Agendamentos = client.Database(conf.ArmazemDatabase).Collection("Agendamentos")

	// result, err := Agendamentos.InsertOne(ctx, ag)
	// if err != nil {
	// 	return nil, err
	// }

	//return result.InsertedID, nil
	return nil, nil
}
