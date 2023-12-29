package armazenamento

import (
	"Agenda/pkgs/common"
	"Agenda/pkgs/convenio"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

// CRUD Convenios
// Criar Convênio a ser utilizado na criação dos Planos de Saude.
func CriarConvenio(cv convenio.Convenios) (interface{}, error) {
	// Antes de qq coisa, verificar os dados do Convenio.
	err := convenio.VerificarConvenio(cv)
	if err != nil {
		return nil, err
	}
	// Definir o Banco e a Coleção de Dados
	Agendamentos = Cliente.Database(common.ConfigInicial.ArmazemDatabase).Collection("Convenios")
	// Inserir os Dados no contexto atual
	result, err := Agendamentos.InsertOne(ctx, cv)
	if err != nil {
		return nil, err
	}
	// Retornar o resultado
	return result.InsertedID, nil
}

// Ler/Retorna Convênios, retorna uma lista de Convênios ao passar uma String independete de Caixa/Baixa.
// A String "*" indica todos os Convênios.
func GetConvenios(nome string) ([]convenio.Convenios, error) {
	// Definir o Banco e a Coleção de Dados
	Agendamentos = Cliente.Database(common.ConfigInicial.ArmazemDatabase).Collection("Convenios")
	// Cria os filtros adequados de pesquisa no MongoDB
	var filter bson.M
	if nome == "*" {
		filter = bson.M{}
	} else {
		filter = bson.M{"plano": primitive.Regex{Pattern: nome, Options: "i"}}
	}
	// Alinha o cursor de busca
	cursor, err := Agendamentos.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	// Cria um Slice para receber os Docs apontados pelo cursor do Driver do Mongo
	var conv []convenio.Convenios
	err = cursor.All(ctx, &conv)
	if err != nil {
		return nil, err
	}
	return conv, nil
}

// Deleta os Convênios de acordo com o "Nome" passado como parâmetro.
// Se "todos" = "true", todos os Docs do filtro serão deletados.
func DeletarConvenio(nome string, todos bool) (*mongo.DeleteResult, error) {
	// Definir o Banco e a Coleção de Dados
	Convenios = Cliente.Database(common.ConfigInicial.ArmazemDatabase).Collection("Convenios")
	// Inserir os Dados no contexto atual
	var result *mongo.DeleteResult
	// Cria os filtros adequados de pesquisa no MongoDB
	filter := bson.M{"plano": primitive.Regex{Pattern: nome, Options: "i"}}
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

func AtualizarConvenio(nomeConv string, novoConv convenio.Convenios, todos bool) (*mongo.UpdateResult, error) {
	// Definir o Banco e a Coleção de Dados
	Convenios = Cliente.Database(common.ConfigInicial.ArmazemDatabase).Collection("Convenios")
	// Cria os filtros adequados de pesquisa no MongoDB
	filter := bson.M{"plano": primitive.Regex{Pattern: nomeConv, Options: "i"}}
	update := bson.M{"$set": novoConv}
	var result *mongo.UpdateResult
	var err error
	if todos {
		result, err = Convenios.UpdateMany(ctx, filter, update)
	} else {
		result, err = Convenios.UpdateOne(ctx, filter, update)
	}
	if err != nil {
		return nil, err
	} else {
		return result, nil
	}
}
