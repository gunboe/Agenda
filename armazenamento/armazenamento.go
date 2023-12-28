package armazenamento

import (
	"Agenda/agendamento"
	"Agenda/lib"
	"Agenda/planosaude"
	"context"
	"fmt"
	"os"
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

// Ambiente do MongoDB
var ctx = context.TODO()
var Agendamentos *mongo.Collection
var Pacientes *mongo.Collection
var Agentes *mongo.Collection
var Convenios *mongo.Collection
var Cliente *mongo.Client

func IniciarArmazenamento() {
	// Testes iniciais do programa e verificação de requisitos
	// Por exemplo:
	// - Testar no config.ini os dias da semana de 0 à 6
	// - Lançar os Erros como Logs

	// Carrega as configurações
	fmt.Print("Iniciando as Configurações do Armazenamento...")
	var conf lib.Config
	conf = lib.ConfigInicial
	fmt.Println(conf.ArmazemDados)

	// Conectar e testar o acesso ao Armazem de Dados
	// MongoDB
	if conf.ArmazemDados == "Mongo" {
		fmt.Print("Conectando ao MongoDB...")
		var err error
		Cliente, err = ConnectMongo(conf)
		if err != nil {
			fmt.Println("\nErro:", err)
			os.Exit(1)
		} else {
			fmt.Println(" Pingou!")
		}
	} else if conf.ArmazemDados == "Postgres" {
		fmt.Println("Banco Postgres ainda não implementado. Use o MongoDB! Saindo...")
		os.Exit(1)
	} else {
		fmt.Println("Nao é MongoDB")
		os.Exit(1)
	}
}

func ConnectMongo(conf lib.Config) (*mongo.Client, error) {
	// Conectando ao MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://" + conf.ArmazemHost + ":" + strconv.Itoa(conf.ArmazemPort) + "/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return client, err
	}
	// Testa Conexão
	err = client.Ping(ctx, nil)
	if err != nil {
		return client, err
	}
	return client, nil
}

// Criar as Coleções/Tabelas no Database "Agenda" do MongoDB
// Agendamentos = client.Database(conf.ArmazemDatabase).Collection("Agendamentos") C
// Pacientes = client.Database(conf.ArmazemDatabase).Collection("Pacientes")
// Agentes = client.Database(conf.ArmazemDatabase).Collection("Agentes")
// Convenios = client.Database(conf.ArmazemDatabase).Collection("Convenios") vCRUD

// CRUD Convenios
func GravarConvenio(cf lib.Config, cv planosaude.Convenios) (interface{}, error) {
	// Antes de qq coisa, verificar os dados do Convenio
	err := planosaude.VerificarConvenio(cv)
	if err != nil {
		return nil, err
	}
	// // Conectar ao Armazem de Dados
	// client, err := ConnectMongo(cf)
	// if err != nil {
	// 	return nil, err
	// }
	// Definir o Banco e a Coleção de Dados
	Agendamentos = Cliente.Database(lib.ConfigInicial.ArmazemDatabase).Collection("Convenios")
	// Inserir os Dados no contexto atual
	result, err := Agendamentos.InsertOne(ctx, cv)
	if err != nil {
		return nil, err
	}
	// Retornar o resultado
	return result.InsertedID, nil
}
func GetConvenios(nome string) ([]planosaude.Convenios, error) {
	// client, err := ConnectMongo(conf)
	// if err != nil {
	// 	return nil, err
	// }
	Agendamentos = Cliente.Database(lib.ConfigInicial.ArmazemDatabase).Collection("Convenios")

	filter := bson.M{}
	if nome == "*" {
		filter = bson.M{}
	} else {
		filter = bson.M{"plano": primitive.Regex{Pattern: nome, Options: "i"}}
	}
	cursor, err := Agendamentos.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var conv []planosaude.Convenios
	err = cursor.All(ctx, &conv)
	if err != nil {
		return nil, err
	}
	return conv, nil
}
func DeletarConvenio(conf lib.Config, nome string, todos bool) (*mongo.DeleteResult, error) {
	// Conectar ao Armazem de Dados
	client, err := ConnectMongo(conf)
	if err != nil {
		return nil, err
	}
	// Definir o Banco e a Coleção de Dados
	Convenios = client.Database(conf.ArmazemDatabase).Collection("Convenios")
	// Inserir os Dados no contexto atual
	var result *mongo.DeleteResult
	filter := bson.M{"plano": primitive.Regex{Pattern: nome, Options: "i"}}
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
func AtualizarConvenio(conf lib.Config, nomeConv string, novoConv planosaude.Convenios, todos bool) (*mongo.UpdateResult, error) {
	// Conectar ao Armazem de Dados
	client, err := ConnectMongo(conf)
	if err != nil {
		return nil, err
	}
	// Definir o Banco e a Coleção de Dados
	Convenios = client.Database(conf.ArmazemDatabase).Collection("Convenios")

	filter := bson.M{"plano": primitive.Regex{Pattern: nomeConv, Options: "i"}}
	var update bson.M
	update = bson.M{"$set": novoConv}
	var result *mongo.UpdateResult
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

// CRUD Agendamentos
func GravarAgendamento(conf lib.Config, ag agendamento.Agendamento) (interface{}, error) {
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

// Ambiente do PostGres
// <vazio por enquanto>
