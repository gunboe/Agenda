package main

import (
	"Agenda/app"
	convControllers "Agenda/controllers/convenio"
	pacControllers "Agenda/controllers/paciente"
	mdb "Agenda/services/armazenamento/mongodb"
	"Agenda/services/config"
	"Agenda/services/repository"
	"Agenda/services/router"
	"fmt"
	"os"
	"strings"
)

func main() {
	///////////////////
	// Iniciando Agenda
	var err error
	fmt.Println("-- Iniciando Agenda --")

	// Inincializa o Serviço de Configurações
	var conf config.Config
	err = conf.CarregaConfig("config.ini")
	if err != nil {
		fmt.Println("Erro de Configuração: ", err)
		os.Exit(1)
	}
	fmt.Print("Carregando as Configurações do Armazenamento...")
	fmt.Println(conf.ArmazemDados)

	///////////////////////////////////////////////////////////////
	// Inicializa as Funções/Serviços de Conexão do Banco de Dados
	var convFunc *convControllers.ConvenioFunc
	var pacienteFunc *pacControllers.PacienteFunc

	// Implementa o Banco de Dados definido na configuração
	switch conf.ArmazemDados {
	case "MongoDB":
		db := &mdb.MongoDB{Configuracao: conf}
		err := db.TestaBanco()
		if err != nil {
			fmt.Println("Erro: Teste inicial de Banco de Dados MongoDB")
			fmt.Println("verifique o Mongo e as configurações de acesso.")
			os.Exit(1)
		}
		// Iniicializa os serviços/Funções de Banco de Dados para o Objetos a serem armazenados
		convFunc = &convControllers.ConvenioFunc{
			ConvRepo: repository.ConvRepo{DB: db}}
		pacienteFunc = &pacControllers.PacienteFunc{
			PacRepo:  repository.PacRepo{DB: db},
			ConvRepo: convFunc.ConvRepo}
	case "Postgres":
		fmt.Println("Banco Postgres ainda não implementado. Use o MongoDB! Saindo...")
		os.Exit(1)
	default:
		fmt.Println("Erro: Escolha uma opção de Banco de Dados na Configuração (config.ini)")
		os.Exit(1)
	}

	// Chama a apresentação Inicial
	// apresentacaoIni()

	// Monta uma conexão com o Banco
	if err = convFunc.ConvRepo.DB.Connect(); err != nil {
		// Lidar com o erro de conexão
		fmt.Println("Erro de Conexão com o Banco:", err)
		os.Exit(1)
	}
	defer convFunc.ConvRepo.DB.Close()

	/////////////////////////////////////////////////
	// Instancia a Aplicação com as Conexões de Banco
	application := &app.Application{
		// Conexões de Banco de Dados
		FuncoesDB: &router.Funcs{
			FuncsConv:     convFunc,
			FuncsPaciente: pacienteFunc,
		},
		// Configurações
		Configuracao: conf,
	}
	// Executa a Aplicação passando as configurações
	application.Run()

	// Fecha as Conexões do Banco
	err = convFunc.ConvRepo.DB.Close()
	if err != nil {
		fmt.Println("Erro: Banco de Dados não fechou!")
		os.Exit(1)
	}

	// Encerramento não esperado da aplicação
	fmt.Println("Ambiente provavelmente encerrado!\n")
}

// Execução de duas consultas para apresentção incial
func apresentacaoIni() {
	// var err error
	var conf config.Config
	conf.CarregaConfig("config.ini")
	var convFunc *convControllers.ConvenioFunc
	var pacienteFunc *pacControllers.PacienteFunc
	db := &mdb.MongoDB{Configuracao: conf}
	db.Connect()
	convFunc = &convControllers.ConvenioFunc{
		ConvRepo: repository.ConvRepo{DB: db}}
	pacienteFunc = &pacControllers.PacienteFunc{
		PacRepo:  repository.PacRepo{DB: db},
		ConvRepo: repository.ConvRepo{DB: db}}

	// Reconecta ao Banco
	fmt.Println("Utilizando o DataBase:", conf.ArmazemDatabase)

	// Obtem todos os Convênios
	todosConvs := convFunc.GetConveniosPorNome("*")
	// Listagem de Convenios
	var listaConvs string
	for _, v := range todosConvs {
		listaConvs += " \"" + v.NomeConv + "\""
	}
	listaConvs = strings.TrimSpace(listaConvs)
	fmt.Println("Lista de Todos os Convenios:", listaConvs)

	// Obtem Todos Pacientes
	todosPacientes := pacienteFunc.GetPacientesPorNome("*")
	// Listagem de Pacientes
	var listaPacs string
	for _, p := range todosPacientes {
		listaPacs += " \"" + p.Nome + "\""
	}
	listaPacs = strings.TrimSpace(listaPacs)
	fmt.Println("Lista de Todos os Pacientes:", listaPacs)
}
