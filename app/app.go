package app

import (
	"Agenda/services/config"
	"Agenda/services/logger"
	"Agenda/services/router"
	"fmt"

	"go.uber.org/zap"
)

// Atributos da Aplicação
type Application struct {
	// Funcoes de Conexão ao Banco de Dados
	FuncoesDB *router.Funcs
	// Configuração do Ambiente
	Configuracao config.Config
}

// Executa a Aplicação
func (app *Application) Run() {
	fmt.Println("-- Executando a Aplicação --")

	// Inicializa o Logger
	logger.InicializaLogger(app.Configuracao)
	logger.Info("Logger inciado")

	// Incializa o Router
	if err := router.InitRouter(app.FuncoesDB, app.Configuracao); err != nil {
		logger.Error("Router falhou: ", err, zap.String("jorney", "app.go"))
	}
}
