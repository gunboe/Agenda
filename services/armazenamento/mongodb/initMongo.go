package mdg

import (
	"Agenda/services/config"
	"context"
	"fmt"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Variáveis Globais do Ambiente do MongoDB
var ctx = context.TODO()

// Implementação da interface Database para MongoDB
type MongoDB struct {
	// campos específicos do MongoDB, se necessário
	Client *mongo.Client
}

// Conectar ao MongoDB
func (m *MongoDB) Connect(conf config.Config) error {
	// lógica de conexão com MongoDB
	var err error
	url := "mongodb://" + conf.ArmazemHost + ":" + strconv.Itoa(conf.ArmazemPort) + "/"
	clientOptions := options.Client().ApplyURI(url)
	m.Client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}
	// Testa Conexão (Essa Ping deve ser retirada em Produção, por possivel lentião)
	err = m.Client.Ping(ctx, nil)
	if err != nil {
		return err
	}
	return nil
}

// Desconectar do MongoDB
func (m *MongoDB) Close() error {
	// lógica de fechamento de conexão
	defer m.Client.Disconnect(ctx)
	return nil // so para não ficar em erro
}

// Inicialização do serviço de armazenamento do MongoDB
func (m *MongoDB) TestaBanco(conf config.Config) error {
	// Conectar e testar o acesso ao Armazem de Dados
	fmt.Print("Conectando ao MongoDB...")
	err := m.Connect(conf)
	if err != nil {
		return err
	} else {
		fmt.Println(" Pingou!")
		m.Close()
		return nil
	}
}
