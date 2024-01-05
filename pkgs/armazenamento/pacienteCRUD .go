package armazenamento

import (
	"Agenda/pkgs/common"
	"Agenda/pkgs/paciente"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

// CRUD Pacientes

// Criar Pacientes para serem utilizados nos Agendamentos
func CreatePaciente(pac paciente.Paciente) (interface{}, error) {
	// Antes de qq coisa, verificar os dados do Paciente.
	err := paciente.VerificarPaciente(pac)
	if err != nil {
		return nil, err
	}
	// Definir o Banco e a Coleção de Dados
	Pacientes = Cliente.Database(common.ConfigInicial.ArmazemDatabase).Collection("Pacientes")
	// Inserir os Dados no contexto atual
	result, err := Pacientes.InsertOne(ctx, pac)
	if err != nil {
		return nil, err
	}
	// Retornar o resultado
	return result.InsertedID, nil
}

// Ler/Retorna Pacientes, retorna uma lista de Pacientes ao passar uma String independete de Caixa/Baixa.
// A String "*" indica todos os Pacientes.
func GetPacientesByName(nome string) ([]paciente.Paciente, error) {
	// Definir o Banco e a Coleção de Dados
	Pacientes = Cliente.Database(common.ConfigInicial.ArmazemDatabase).Collection("Pacientes")
	// Cria os filtros adequados de pesquisa no MongoDB
	var filter bson.M
	if nome == "*" {
		filter = bson.M{}
	} else {
		filter = bson.M{"nomeconv": primitive.Regex{Pattern: nome, Options: "i"}}
	}
	// Alinha o cursor de busca
	cursor, err := Pacientes.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	// Cria um Slice para receber os Docs apontados pelo cursor do Driver do Mongo
	var convs []paciente.Paciente
	err = cursor.All(ctx, &convs)
	if err != nil {
		return nil, err
	}
	return convs, nil
}

// Ler/Retorna Pacientes, retorna um Convênio por ID
func GetPacienteById(id primitive.ObjectID) (paciente.Paciente, error) {
	// Definir o Banco e a Coleção de Dados
	Pacientes = Cliente.Database(common.ConfigInicial.ArmazemDatabase).Collection("Pacientes")
	// Cria os filtros adequados de pesquisa no MongoDB
	var filter bson.M
	filter = bson.M{"_id": id}

	// Cria um Convênio
	var conv paciente.Paciente
	// Alinha o cursor de busca
	err := Pacientes.FindOne(ctx, filter).Decode(&conv)
	if err != nil {
		return conv, err
	}
	return conv, nil
}

// Deleta os Pacientes de acordo com o "Nome" passado como parâmetro.
// Se "todos" = "true", todos os Docs do filtro serão deletados.
func DeletePacienteByName(nome string, todos bool) (*mongo.DeleteResult, error) {
	// Definir o Banco e a Coleção de Dados
	Pacientes = Cliente.Database(common.ConfigInicial.ArmazemDatabase).Collection("Pacientes")
	// Inserir os Dados no contexto atual
	var result *mongo.DeleteResult
	// Cria os filtros adequados de pesquisa no MongoDB
	var filter bson.M
	if nome == "*" {
		filter = bson.M{}
	} else {
		filter = bson.M{"plano": primitive.Regex{Pattern: nome, Options: "i"}}
	}
	var err error
	if todos {
		// opts := options.Delete().SetHint(bson.M{"_id": 1})
		result, err = Pacientes.DeleteMany(ctx, filter)
	} else {
		result, err = Pacientes.DeleteOne(ctx, filter)
	}
	if err != nil {
		return nil, err
	}
	// Retornar o resultado
	return result, nil
}

func UpdatePacienteByName(nomeConv string, novoConv paciente.Paciente, todos bool) (*mongo.UpdateResult, error) {
	// Definir o Banco e a Coleção de Dados
	Pacientes = Cliente.Database(common.ConfigInicial.ArmazemDatabase).Collection("Pacientes")
	// Cria os filtros adequados de pesquisa no MongoDB
	filter := bson.M{"plano": primitive.Regex{Pattern: nomeConv, Options: "i"}}
	update := bson.M{"$set": novoConv}
	var result *mongo.UpdateResult
	var err error
	if todos {
		result, err = Pacientes.UpdateMany(ctx, filter, update)
	} else {
		result, err = Pacientes.UpdateOne(ctx, filter, update)
	}
	if err != nil {
		return nil, err
	} else {
		return result, nil
	}
}
