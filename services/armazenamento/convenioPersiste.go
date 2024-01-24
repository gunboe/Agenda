// PERSISTÊNCIA: Convenios
package armazenamento

import (
	"Agenda/models"
	"Agenda/services/config"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

// (CREATE) Criar Convênios para serem utilizados nos Planos de Pagamentos.
func CreateConvenio(cv models.Convenio) (interface{}, error) {
	// Antes de qq coisa, verificar os dados do Convenio.
	err := models.ChecarConvenio(cv)
	if err != nil {
		return nil, err
	}
	// Definir o Banco e a Coleção de Dados
	Convenios = Cliente.Database(config.ConfigInicial.ArmazemDatabase).Collection("Convenios")
	// Inserir os Dados no contexto atual
	result, err := Convenios.InsertOne(ctx, cv)
	if err != nil {
		return nil, err
	}
	// Retornar o resultado
	return result.InsertedID, nil
}

// (READ) Ler/Retorna Lista de Convênios buscando por Nome independete de Caixa/Baixa.
// A String "*" indica todos os Convênios.
func GetConveniosByName(nome string) ([]models.Convenio, error) {
	// Definir o Banco e a Coleção de Dados
	Convenios = Cliente.Database(config.ConfigInicial.ArmazemDatabase).Collection("Convenios")
	// Cria os filtros adequados de pesquisa no MongoDB
	var filter bson.M
	if nome == "*" {
		filter = bson.M{}
	} else {
		filter = bson.M{"nomeconv": primitive.Regex{Pattern: nome, Options: "i"}}
	}
	// Alinha o cursor de busca
	cursor, err := Convenios.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	// Cria um Slice para receber os Docs apontados pelo cursor do Driver do Mongo
	var convs []models.Convenio
	err = cursor.All(ctx, &convs)
	if err != nil {
		return nil, err
	}
	return convs, nil
}

// (READ) Ler/Retorna Convênios por ID. Se não encontrar retorna um Erro e um Convenio com atributos zerados.
func GetConvenioById(id primitive.ObjectID) (models.Convenio, error) {
	// Definir o Banco e a Coleção de Dados
	Convenios = Cliente.Database(config.ConfigInicial.ArmazemDatabase).Collection("Convenios")
	// Cria os filtros adequados de pesquisa no MongoDB
	filter := bson.M{"_id": id}
	// Cria um Convênio
	var conv models.Convenio
	// Alinha o cursor de busca
	err := Convenios.FindOne(ctx, filter).Decode(&conv)
	if err != nil {
		return conv, err
	}
	return conv, nil
}

// (UPDATE) Atualiza um ou mais Convenios pelo Nome do Convênio. Se não encontrar um Convênio NÃO retorna erro.
// Ao passar a String "*" todos os registros filtrados serão alterados.
func UpdateConvenioByName(nome string, novoConv models.Convenio, todos bool) (*mongo.UpdateResult, error) {
	// Definir o Banco e a Coleção de Dados
	Convenios = Cliente.Database(config.ConfigInicial.ArmazemDatabase).Collection("Convenios")
	// Cria os filtros adequados de pesquisa no MongoDB
	var filter bson.M
	if nome == "*" {
		filter = bson.M{}
	} else {
		filter = bson.M{"nomeconv": primitive.Regex{Pattern: nome, Options: "i"}}
	}
	update := bson.M{"$set": novoConv}
	var result *mongo.UpdateResult
	var err error
	if todos {
		result, err = Convenios.UpdateMany(ctx, filter, update)
	} else {
		// fmt.Println("tentando atualizar:", update, " com o filtro:", filter)
		result, err = Convenios.UpdateOne(ctx, filter, update)
	}
	if err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

// (UPDATE) Atualiza os Compos de um Convenios pelo ID do Convênio. Se não encontrar um Convênio NÃO retorna erro.
func UpdateConvenioById(id primitive.ObjectID, novoConv models.Convenio) (*mongo.UpdateResult, error) {
	// Definir o Banco e a Coleção de Dados
	Convenios = Cliente.Database(config.ConfigInicial.ArmazemDatabase).Collection("Convenios")
	// Cria os filtros adequados de pesquisa no MongoDB
	update := bson.M{"$set": novoConv}
	var result *mongo.UpdateResult
	var err error
	result, err = Convenios.UpdateByID(ctx, id, update)
	if err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

// (UPDATE) Disponibilizar um Convênio por ID. Quando um Convênio está marcado como Indisponivel,
// ele não pode ser alterado nem utilizado em PlanosPgto. Se não encontrar um Convênio NÃO retorna erro.
func AllowConveioById(id primitive.ObjectID, b bool) (*mongo.UpdateResult, error) {
	// Definir o Banco e a Coleção de Dados
	Convenios = Cliente.Database(config.ConfigInicial.ArmazemDatabase).Collection("Convenios")
	// Define o valor a ser atualizado
	update := bson.M{"$set": bson.M{"indisponivel": !b}}
	var result *mongo.UpdateResult
	var err error
	result, err = Convenios.UpdateByID(ctx, id, update)
	if err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

// (DELETE) Deleta os Convênios de acordo com o "Nome" passado como parâmetro.
// Se "todos" = "true", todos os Docs do filtro serão deletados.
func DeleteConvenioByName(nome string, todos bool) (*mongo.DeleteResult, error) {
	// Definir o Banco e a Coleção de Dados
	Convenios = Cliente.Database(config.ConfigInicial.ArmazemDatabase).Collection("Convenios")
	// Inserir os Dados no contexto atual
	var result *mongo.DeleteResult
	// Cria os filtros adequados de pesquisa no MongoDB
	var filter bson.M
	if nome == "*" {
		filter = bson.M{}
	} else {
		filter = bson.M{"nomeconv": primitive.Regex{Pattern: nome, Options: "i"}}
	}
	var err error
	if todos {
		// opts := options.Delete().SetHint(bson.M{"_id": 1})
		result, err = Convenios.DeleteMany(ctx, filter)
	} else {
		result, err = Convenios.DeleteOne(ctx, filter)
	}
	if err != nil {
		return nil, err
	}
	// Retornar o resultado
	return result, nil
}

// Deleta os Convênios de acordo com o "Nome" passado como parâmetro.
func DeleteConvenioById(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	// Definir o Banco e a Coleção de Dados
	Convenios = Cliente.Database(config.ConfigInicial.ArmazemDatabase).Collection("Convenios")
	// Inserir os Dados no contexto atual
	var result *mongo.DeleteResult
	// Cria os filtros adequados de pesquisa no MongoDB
	filter := bson.M{"_id": id}
	var err error
	result, err = Convenios.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}
	// Retornar o resultado
	return result, nil
}
