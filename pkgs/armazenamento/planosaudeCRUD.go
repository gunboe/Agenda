package armazenamento

import (
	"Agenda/pkgs/common"
	"Agenda/pkgs/planosaude"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

// CRUD PlanoSaude
// Criar Plano de Saude a ser utilizado pelos Pacientes.
func CriarPlanoSaude(cv planosaude.PlanoSaude) (interface{}, error) {
	// Antes de qq coisa, verificar os dados do PlanoSaude.
	err := planosaude.VerificarPlano(cv)
	if err != nil {
		return nil, err
	}
	// Definir o Banco e a Coleção de Dados
	PlanoSaude = Cliente.Database(common.ConfigInicial.ArmazemDatabase).Collection("PlanoSaude")
	// Inserir os Dados no contexto atual
	result, err := PlanoSaude.InsertOne(ctx, cv)
	if err != nil {
		return nil, err
	}
	// Retornar o resultado
	return result.InsertedID, nil
}

// Ler/Retorna Plano de Saudes, retorna uma lista de Plano de Saudes ao passar uma String independete de Caixa/Baixa.
// A String "*" indica todos os Plano de Saudes.
func GetPlanoSaude(nome string) ([]planosaude.PlanoSaude, error) {
	// Definir o Banco e a Coleção de Dados
	PlanoSaude = Cliente.Database(common.ConfigInicial.ArmazemDatabase).Collection("PlanoSaude")
	// Cria os filtros adequados de pesquisa no MongoDB
	var filter bson.M
	if nome == "*" {
		filter = bson.M{}
	} else {
		filter = bson.M{"convenio": primitive.Regex{Pattern: nome, Options: "i"}}
	}
	// Alinha o cursor de busca
	cursor, err := PlanoSaude.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	// Cria um Slice para receber os Docs apontados pelo cursor do Driver do Mongo
	var plano []planosaude.PlanoSaude
	err = cursor.All(ctx, &plano)
	if err != nil {
		return nil, err
	}
	return plano, nil
}

// Deleta os Plano de Saudes de acordo com o "Nome" passado como parâmetro.
// Se "todos" = "true", todos os Docs do filtro serão deletados.
func DeletarPlanoSaude(nome string, todos bool) (*mongo.DeleteResult, error) {
	// Definir o Banco e a Coleção de Dados
	PlanoSaude = Cliente.Database(common.ConfigInicial.ArmazemDatabase).Collection("PlanoSaude")
	// Inserir os Dados no contexto atual
	var result *mongo.DeleteResult
	// Cria os filtros adequados de pesquisa no MongoDB
	filter := bson.M{"plano": primitive.Regex{Pattern: nome, Options: "i"}}
	var err error
	if todos {
		// opts := options.Delete().SetHint(bson.M{"_id": 1})
		result, err = PlanoSaude.DeleteMany(ctx, filter)
	} else {
		result, err = PlanoSaude.DeleteOne(ctx, filter)
	}
	if err != nil {
		return nil, err
	}
	// Retornar o resultado
	return result, nil
}

func AtualizarPlanoSaude(nomeConv string, novoConv planosaude.PlanoSaude, todos bool) (*mongo.UpdateResult, error) {
	// Definir o Banco e a Coleção de Dados
	PlanoSaude = Cliente.Database(common.ConfigInicial.ArmazemDatabase).Collection("PlanoSaude")
	// Cria os filtros adequados de pesquisa no MongoDB
	filter := bson.M{"plano": primitive.Regex{Pattern: nomeConv, Options: "i"}}
	update := bson.M{"$set": novoConv}
	var result *mongo.UpdateResult
	var err error
	if todos {
		result, err = PlanoSaude.UpdateMany(ctx, filter, update)
	} else {
		result, err = PlanoSaude.UpdateOne(ctx, filter, update)
	}
	if err != nil {
		return nil, err
	} else {
		return result, nil
	}
}
