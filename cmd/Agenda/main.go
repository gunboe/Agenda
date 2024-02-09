package main

import (
	"Agenda/app"
	convControllers "Agenda/controllers/convenio"
	pacControllers "Agenda/controllers/paciente"
	"Agenda/repository"
	mdb "Agenda/services/armazenamento/mongodb"
	"Agenda/services/config"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Iniciando Agenda
	var err error
	fmt.Println("-- Iniciando Agenda --")

	// Carrega as configurações
	var conf config.Config
	err = conf.CarregaConfig("config.ini")
	if err != nil {
		fmt.Println("Erro de Configuração: ", err)
		os.Exit(1)
	}
	fmt.Print("Carregando as Configurações do Armazenamento...")
	fmt.Println(conf.ArmazemDados)

	// Inicialização das Funções de Conexão do Banco de Dados
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
			ConvRepo: repository.ConvRepo{DB: db}}
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

	// Instancia a Aplicação passando as Conexões de Banco de Dados
	application := &app.Application{
		// Conexões de Banco de Dados
		ConvFunc:     convFunc,
		PacienteFunc: pacienteFunc,
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

	// Avisa que está pronto
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

func testes() {
	// TESTES
	// var err error

	// Inicialização de algumas variáveis pra teste da Estrutra de Dados
	// var d, _ = time.Parse("02/01/2006", "22/06/2023")
	// var d, _ = time.Parse("02/01/2006", "25/05/2025")
	// var d, _ = time.Parse("02/01/2006", "07/07/2027")
	// var dur, _ = time.ParseDuration("1h")

	// Inicializa Convênio e Plano
	// convTeste := models.Convenio{NomeConv: "Sul America", Endereco: "Rua das Nações,163", DataContratoConv: d1, Indisponivel: false}
	// criaConvenio(convTeste)

	//  Retorna Conv
	// conv := "sul"
	// convs := getConvenios(conv)
	// if convs == nil {
	// 	os.Exit(1)
	// }

	// fmt.Println("Verificando o plano:", planoTeste.Convenio.NomeConv, "nos convenios:", listaConvs)
	// err = models.VerificarPlano(planoTeste)
	// if err != nil {
	// 	fmt.Println("Erro:", err)
	// 	os.Exit(1)
	// }

	// Antes de criar o Plano de Pagamento deve-se, obter o Convenio cadastrado no Mongo
	//convTeste = getConvenios("sul")[0]
	//Cria um Plano CASSI
	// id, _ := primitive.ObjectIDFromHex("65998064fab6d835ca0f5a5e") //
	// planoX := models.PlanoPgto{ID: primitive.NewObjectID(), ConvenioId: id,
	// NrPlano: "00000-01", DataValidade: d, Inativo: false, Particular: false}
	// planoParticular := models.PlanoPgto{Particular: true}

	// // fmt.Println("Checa PalnoX:", models.ChecarPlanoPgto(planoX))
	// // fmt.Println("Checa PalnoParticular:", models.ChecarPlanoPgto(planoParticular))

	// Criadno Paciente com um PlanoPgto
	// var pacienteA = models.Paciente{ID: primitive.NewObjectID(), Nome: "Gunther boeckmann", CPF: "891552974-04",
	// 	NrCelular: 81999998888, Email: "biel@net.io", Endereco: "Rua Manoel H PEssoa, 552",
	// 	Bloqueado: false}
	// pacienteA.PlanosPgts = append(pacienteA.PlanosPgts, planoX)
	// pacienteA.PlanosPgts = append(pacienteA.PlanosPgts, planoParticular)

	// fmt.Println("Checa PacienteA:", models.ChecarPaciente(pacienteA))
	// fmt.Println(printJSON(pacienteA))
	// controllers.CriaPaciente(pacienteA)

	// var pacienteB = models.Paciente{ID: primitive.NewObjectID(), Nome: "Guisela Silva", CPF: "194630144-20",
	// 	NrCelular: 81999998888, Email: "guiga@net.io", Endereco: "SolMAr, 47",
	// 	Bloqueado: false}
	// pacienteB.PlanosPgts = append(pacienteB.PlanosPgts, planoParticular)

	// Para gravar e recuperar o Secret, é necessário criar uma função propria de persistencia
	// com MAP ao invés de Struct, visto que esse atributo é PRIVADO!!
	// err := pacienteA.SetSecret("SEGREDOBIEL")
	// if err != nil {
	// 	fmt.Println("Erro:", err)
	// }

	// fmt.Println(printJSON(pacienteB))
	// criaPaciente(pacienteB)

	// // Alterando Paciente
	// var d, _ = time.Parse("02/01/2006", "25/05/2025")
	// idpac, _ := primitive.ObjectIDFromHex("65a424731b8ddd87b0e3e926") // Pac: Gunther
	// // controllers.HabilitePacPorId(idpac, true)
	// pacienteA, err := controllers.GetPacientePorId(idpac)
	// if err != nil {
	// 	fmt.Println("Não encotrou CPF:", pacienteA.ID, " Saindo...")
	// 	os.Exit(1)
	// }
	// convenioId, _ := primitive.ObjectIDFromHex("65a3f94a14e94f9bca4d4115")
	// pID := primitive.NewObjectID()
	// pID, _ := primitive.ObjectIDFromHex("659de2c86c357efadababe38")
	// plano := models.PlanoPgto{ID: pID, ConvenioId: convenioId, NrPlano: "00000-1", DataValidade: d}
	// plano := models.PlanoPgto{Particular: true}
	// // Adicionando novo plano
	// pacienteA.PlanosPgts = append(pacienteA.PlanosPgts, plano)
	// // pacienteA.ID = idpac
	// // pacienteA.PlanosPgts[0] = plano
	// controllers.AtualizaPacPorId(idpac, pacienteA)
	// listaPaciente("gunt")

	// Deletetando Plano do Paciente
	// plano = models.PlanoPgto{ID: plano.ID}
	// controllers.DelPlanoPgtoPaciente(idpac, plano)

	// Inserindo Plano do Paciente
	// controllers.InsPlanoPgtoPaciente(idpac, plano)

	// Deletando Paciente
	// idpac, _ := primitive.ObjectIDFromHex("659b5cddfda4bbbe7de29781") // Pac: Gunther
	// deletaPacientePorId(idpac)

	// var ag = agente.Agente{Nome: "Elke", CPF: "001.038.719-32", NrCelular: 123456798, Especialidades: []string{"Endocrino", "Clinico"}}
	// err = ag.SetSecret("senha123")
	// err = agente.VerificarAgente(ag)
	// if err != nil {
	// 	fmt.Println(err, ag.Nome)
	// }
	// fmt.Println(ag)
	// fmt.Println(pacienteA)
	// var agenteExec = agente.Agente{Nome: "Dr. Zebalos", CPF: "12345679-01", NrCelular: 8199997777, Especialidades: []string{"Ortopedista", "Cirurgião"}}
	// var agendaBiel = agendamento.Agendamento{ID: primitive.NewObjectID(), DataInicio: d2, Duracao: dur, Atividade: "Consulta padrão", AgenteExecutor: agenteExec, PacienteAtend: pacienteA, Confirmado: true, MeioPagamento: "Dinheiro", Pago: false, Cancelado: false}

	// fmt.Println(agendaBiel)
	// s, err := pacienteA.GetSecret()
	// if err != nil {
	// 	fmt.Println("Erro:", err)
	// } else {
	// 	fmt.Println(s)
	// }

	// // Criar Convenio
	// nomeConv := "CASSI"
	// endConv := "Av Rosa e Silva,25"
	// dataConv, _ := time.Parse("02/01/2006", "04/04/2024") // Data deve conter zero!!
	// novoConv := models.Convenio{ID: primitive.NewObjectID(), NomeConv: nomeConv, Endereco: endConv, DataContratoConv: dataConv}
	// controllers.CriaConvenio(novoConv)

	// nomeConv = "UNIMED"
	// endConv = "Rua Solimões, 35"
	// dataConv, _ = time.Parse("02/01/2006", "03/03/2033") // Data deve conter zero!!
	// novoConv = models.Convenio{ID: primitive.NewObjectID(), NomeConv: nomeConv, Endereco: endConv, DataContratoConv: dataConv, Indisponivel: false}
	// controllers.CriaConvenio(novoConv)

	// nomeConv := "Sul America"
	// endConv := "Rua das Nações,163"
	// dataConv, _ := time.Parse("02/01/2006", "05/05/2055") // Data deve conter zero!!
	// novoConv := models.Convenio{ID: primitive.NewObjectID(), NomeConv: nomeConv, Endereco: endConv, DataContratoConv: dataConv}
	// convControllers.CriaConvenio(novoConv)

	// Listar Convenios
	// controllers.ListaConvenio("*")

	// // Teste DELETE Convenio
	// plano := "teste"
	// todos := true
	// // var todos bool
	// controllers.DeletaConveniosPorNome(plano, todos)
	// // listaConvenio("*")

	// filtroNomeConv := "CASSI"
	// todos := false
	// atualizaConv("PARticula", convTeste, todos)

	// var d, _ = time.Parse("02/01/2006", "22/06/2023")
	// var d, _ = time.Parse("02/01/2006", "25/05/2025")
	// var d, _ = time.Parse("02/01/2006", "07/07/2027")

	// Alterando Convenio
	// ConvAlterado := models.Convenio{}
	// NomeConv := "sul"
	// idConv := controllers.GetConveniosPorNome(NomeConv)[0].ID
	// ID, _ := primitive.ObjectIDFromHex("65a3f94a14e94f9bca4d4115") //Sul America
	// ConvAlterado = models.Convenio{ID: idConv, NomeConv: "Sul da America", DataContratoConv: d, Endereco: "Rua das Creolas, 467"}
	// fmt.Println("ConvAlterado antes:", ConvAlterado)

	// controllers.AtualizaConvPorId(idConv, ConvAlterado)
	// fmt.Println(idConv)
	// controllers.HabiliteConvPorId(idConv, true)

	// controllers.ListaConvenio(NomeConv)

	// atualizaConvPorNome("CASSI", ConvAlterado, false)
	// novoConvs := getConvenioPorId(ID)
	// fmt.Println(novoConvs)
	// novoConv.NomeConv = "CASSI"
	// novoConv.Endereco = "Av Rosa e Silva,9999"
	// novoConv = novoConvs[0]
	// // fmt.Println("novoConv antes :", novoConv)
	// id := novoConv.ID
	// fmt.Println("ID:", id)
	// // novoConv.NomeConv = "CASSI"
	// novoConv.Indisponivel = false

	// deletaConvenioPorId(ID)
	// fmt.Println("novoConv depois:", getConvenioPorId(id))

	// fmt.Println("Listando todoso os Convenios")
	// listaConvenio("*", "json")

}
