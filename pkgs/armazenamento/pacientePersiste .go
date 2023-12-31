package armazenamento

import (
	"Agenda/pkgs/common"
	"Agenda/pkgs/paciente"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

// PERSISTÊNCIA: Pacientes

// (CREATE) Criar Pacientes para serem utilizados nos Agendamentos
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

// (READ) Ler/Retorna Lista de Pacientes buscando por Nome independete de Caixa/Baixa.
// A String "*" indica todos os Pacientes.
func GetPacientesByName(nome string) ([]paciente.Paciente, error) {
	// Definir o Banco e a Coleção de Dados
	Pacientes = Cliente.Database(common.ConfigInicial.ArmazemDatabase).Collection("Pacientes")
	// Cria os filtros adequados de pesquisa no MongoDB
	var filter bson.M
	if nome == "*" {
		filter = bson.M{}
	} else {
		filter = bson.M{"nome": primitive.Regex{Pattern: nome, Options: "i"}}
	}
	// Alinha o cursor de busca
	cursor, err := Pacientes.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	// Cria um Slice para receber os Docs apontados pelo cursor do Driver do Mongo
	var pacs []paciente.Paciente
	err = cursor.All(ctx, &pacs)
	if err != nil {
		return nil, err
	}
	return pacs, nil
}

// (READ) Ler/Retorna Pacientes por ID
func GetPacienteById(id primitive.ObjectID) (paciente.Paciente, error) {
	// Definir o Banco e a Coleção de Dados
	Pacientes = Cliente.Database(common.ConfigInicial.ArmazemDatabase).Collection("Pacientes")
	// Cria os filtros adequados de pesquisa no MongoDB
	filter := bson.M{"_id": id}
	// Cria um Paciente
	var pac paciente.Paciente
	// Alinha o cursor de busca
	err := Pacientes.FindOne(ctx, filter).Decode(&pac)
	if err != nil {
		return pac, err
	}
	return pac, nil
}

// (READ) Ler/Retorna Pacientes por CPF
func GetPacienteByCPF(cpf string) (paciente.Paciente, error) {
	// Definir o Banco e a Coleção de Dados
	Pacientes = Cliente.Database(common.ConfigInicial.ArmazemDatabase).Collection("Pacientes")
	// Cria os filtros adequados de pesquisa no MongoDB
	filter := bson.M{"cpf": cpf}
	// Cria um Paciente
	var pac paciente.Paciente
	// Alinha o cursor de busca
	err := Pacientes.FindOne(ctx, filter).Decode(&pac)
	if err != nil {
		return pac, err
	}
	return pac, nil
}

// (UPDATE) Atualiza um ou mais Pacientes pelo Nome do Paciente. Se não encontrar um Paciente NÃO retorna erro.
// Ao passar a String "*" todos os registros filtrados serão alterados
func UpdatePacienteByName(nome string, novoPac paciente.Paciente, todos bool) (*mongo.UpdateResult, error) {
	// Definir o Banco e a Coleção de Dados
	Pacientes = Cliente.Database(common.ConfigInicial.ArmazemDatabase).Collection("Pacientes")
	// Cria os filtros adequados de pesquisa no MongoDB
	var filter bson.M
	if nome == "*" {
		filter = bson.M{}
	} else {
		filter = bson.M{"nome": primitive.Regex{Pattern: nome, Options: "i"}}
	}
	update := bson.M{"$set": novoPac}
	var result *mongo.UpdateResult
	var err error
	if todos {
		result, err = Pacientes.UpdateMany(ctx, filter, update)
	} else {
		// fmt.Println("tentando atualizar:", update, " com o filtro:", filter)
		result, err = Pacientes.UpdateOne(ctx, filter, update)
	}
	if err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

// (UPDATE) Atualiza os Compos de um Pacientes pelo ID do Convênio. Se não encontrar um Paciente NÃO retorna erro.
func UpdatePacienteById(id primitive.ObjectID, novoPac paciente.Paciente) (*mongo.UpdateResult, error) {
	// Definir o Banco e a Coleção de Dados
	Pacientes = Cliente.Database(common.ConfigInicial.ArmazemDatabase).Collection("Pacientes")
	// Cria os filtros adequados de pesquisa no MongoDB
	// filter := bson.M{"_id": id}
	update := bson.M{"$set": novoPac}
	var result *mongo.UpdateResult
	var err error
	result, err = Pacientes.UpdateByID(ctx, id, update)
	if err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

// (UPDATE) Besbloquear um Paciente por ID. Quando um Paciente está marcado como Bloqueado,
// ele não pode ser alterado nem utilizado em Agendamentos. Se não encontrar um Convênio NÃO retorna erro.
func AllowPacienteById(id primitive.ObjectID, b bool) (*mongo.UpdateResult, error) {
	// Definir o Banco e a Coleção de Dados
	Pacientes = Cliente.Database(common.ConfigInicial.ArmazemDatabase).Collection("Pacientes")
	// Define o valor a ser atualizado
	update := bson.M{"$set": bson.M{"bloqueado": !b}}
	var result *mongo.UpdateResult
	var err error
	result, err = Pacientes.UpdateByID(ctx, id, update)
	if err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

// (DELETE) Deleta os Pacientes de acordo com o "Nome" passado como parâmetro.
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

// Deleta os Pacientes de acordo com o "Nome" passado como parâmetro.
func DeletePacienteById(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	// Definir o Banco e a Coleção de Dados
	Pacientes = Cliente.Database(common.ConfigInicial.ArmazemDatabase).Collection("Pacientes")
	// Inserir os Dados no contexto atual
	var result *mongo.DeleteResult
	// Cria os filtros adequados de pesquisa no MongoDB
	filter := bson.M{"_id": id}
	var err error
	result, err = Pacientes.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}
	// Retornar o resultado
	return result, nil
}
