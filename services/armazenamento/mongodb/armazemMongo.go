package armazenamento

import (
	"Agenda/services/config"
	"context"
	"fmt"
	"os"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Variáveis Globais do Ambiente do MongoDB
var ctx = context.TODO()
var Agendamentos *mongo.Collection
var Convenios *mongo.Collection
var Pacientes *mongo.Collection
var Agentes *mongo.Collection
var Cliente *mongo.Client

// Implementação da interface Database para MongoDB
type MongoDB struct {
	// campos específicos do MongoDB, se necessário
	Client *mongo.Client
}

// Conectar ao MongoDB
func (m *MongoDB) Connect() error {
	// lógica de conexão com MongoDB
	return nil // so para não ficar em erro
}

// Desconectar do MongoDB
func (m *MongoDB) Close() error {
	// lógica de fechamento de conexão
	return nil // so para não ficar em erro
}

// Inicialização do serviço de armazenamento do MongoDB
func init() {
	// Carrega as configurações
	fmt.Print("Iniciando as Configurações do Armazenamento...")
	conf := config.ConfigInicial
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

// Conectar ao armazenamento MongoDB
func ConnectMongo(conf config.Config) (*mongo.Client, error) {
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
