package armazenamento

import (
	"Agenda/pkgs/common"
	"Agenda/pkgs/planopgto"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

// CRUD PlanoPgto
// Criar Plano de Pagamento a ser utilizado pelos Pacientes.
func CriarPlanoPgto(pp planopgto.PlanoPgto) (interface{}, error) {
	// Antes de qq coisa, verificar os dados do PlanoPgto estão em conformidade.
	err := planopgto.VerificarPlano(pp)
	if err != nil {
		return nil, err
	}
	// Ajustar o Documento do Objeto a ser armazenado,
	// visto que existe relacionamento com outros Objetos.
	// A não ser que o Mongo ofereça algo para armazenar docs relacionados
	// Para o Convenio e Paciente:
	// var novoconv convenio.Convenio
	// var novopac paciente.Paciente
	// var id primitive.ObjectID
	// id = pp.Convenio.ID
	// pp.Convenio = novoconv
	// pp.Convenio.ID = id
	// id = pp.Paciente.ID
	// pp.Paciente = novopac
	// pp.Paciente.ID = id

	// Definir o Banco e a Coleção de Dados
	PlanoPgto = Cliente.Database(common.ConfigInicial.ArmazemDatabase).Collection("PlanoPgto")
	// Inserir os Dados no contexto atual
	result, err := PlanoPgto.InsertOne(ctx, pp)
	if err != nil {
		return nil, err
	}
	// Retornar o resultado
	return result.InsertedID, nil
}

// Ler/Retorna Plano de Pagamentos, retorna uma lista de Plano de Pagamentos ao passar uma String independete de Caixa/Baixa.
// A String "*" indica todos os Plano de Pagamentos.
func GetPlanoPgto(nome string) ([]planopgto.PlanoPgto, error) {
	// Definir o Banco e a Coleção de Dados
	PlanoPgto = Cliente.Database(common.ConfigInicial.ArmazemDatabase).Collection("PlanoPgto")
	// Cria os filtros adequados de pesquisa no MongoDB
	var filter bson.M
	if nome == "*" {
		filter = bson.M{}
	} else {
		filter = bson.M{"convenio": primitive.Regex{Pattern: nome, Options: "i"}}
	}
	// Alinha o cursor de busca
	cursor, err := PlanoPgto.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	// Cria um Slice para receber os Docs apontados pelo cursor do Driver do Mongo
	var plano []planopgto.PlanoPgto
	err = cursor.All(ctx, &plano)
	if err != nil {
		return nil, err
	}
	return plano, nil
}

// Deleta os Plano de Pagamentos de acordo com o "Nome" passado como parâmetro.
// Se "todos" = "true", todos os Docs do filtro serão deletados.
func DeletarPlanoPgto(nome string, todos bool) (*mongo.DeleteResult, error) {
	// Definir o Banco e a Coleção de Dados
	PlanoPgto = Cliente.Database(common.ConfigInicial.ArmazemDatabase).Collection("PlanoPgto")
	// Inserir os Dados no contexto atual
	var result *mongo.DeleteResult
	// Cria os filtros adequados de pesquisa no MongoDB
	filter := bson.M{"plano": primitive.Regex{Pattern: nome, Options: "i"}}
	var err error
	if todos {
		// opts := options.Delete().SetHint(bson.M{"_id": 1})
		result, err = PlanoPgto.DeleteMany(ctx, filter)
	} else {
		result, err = PlanoPgto.DeleteOne(ctx, filter)
	}
	if err != nil {
		return nil, err
	}
	// Retornar o resultado
	return result, nil
}

func AtualizarPlanoPgto(nomeConv string, novoConv planopgto.PlanoPgto, todos bool) (*mongo.UpdateResult, error) {
	// Definir o Banco e a Coleção de Dados
	PlanoPgto = Cliente.Database(common.ConfigInicial.ArmazemDatabase).Collection("PlanoPgto")
	// Cria os filtros adequados de pesquisa no MongoDB
	filter := bson.M{"plano": primitive.Regex{Pattern: nomeConv, Options: "i"}}
	update := bson.M{"$set": novoConv}
	var result *mongo.UpdateResult
	var err error
	if todos {
		result, err = PlanoPgto.UpdateMany(ctx, filter, update)
	} else {
		result, err = PlanoPgto.UpdateOne(ctx, filter, update)
	}
	if err != nil {
		return nil, err
	} else {
		return result, nil
	}
}
