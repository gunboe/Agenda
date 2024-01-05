// Pacote de funções de alto nível (CRUD-Create,Read,list,Update,Delete)
// para acessar os dados armazenados nas Coleções/Tabelas do Database "Agenda"
//  no MongoDB.
// Agendamentos	: C
// Pacientes	:
// Agentes		:
// Convenios	: CRlUD
package main

import (
	"Agenda/pkgs/armazenamento"
	"Agenda/pkgs/common"
	"Agenda/pkgs/convenio"
	"Agenda/pkgs/paciente"
	"Agenda/pkgs/planopgto"
	"encoding/json"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Inicializa o ambiente
func inicializaAmbiente() {
	// Carrega as configurações
	var conf common.Config
	conf = common.ConfigInicial
	// Conecta ao Banco
	fmt.Println("Utilizando o DataBase:", conf.ArmazemDatabase)
	// Testa o Banco relacionando todos os Convênnios
	todosConvs := getConveniosPorNome("*")
	var listaConvs string
	for _, v := range todosConvs {
		listaConvs += " \"" + v.NomeConv + "\""
	}
	listaConvs = strings.TrimSpace(listaConvs)
	fmt.Println("Lista de Todos os Convenios:", listaConvs)
	// Avisa que está pronto
	fmt.Println("Ambiente pronto para uso!\n")
}

// Converter struct para Json
func printJSON(input interface{}) string {
	s, _ := json.MarshalIndent(input, "", "\t")
	return string(s)
}

// CRUD Convenios

// (CREATE) Cria convênio e salva no armazém
func criaConvenio(conv convenio.Convenio) {
	// Checa se já existe Convenio pelo Nome
	convs, err := armazenamento.GetConveniosByName(conv.NomeConv)
	if err != nil {
		fmt.Println("Erro:", err)
	}
	if convs == nil {
		result, err := armazenamento.CreateConvenio(conv)
		if err != nil {
			fmt.Println("Erro:", err)
		} else {
			fmt.Println("Convenio Criado e armazenado:", result)
		}
	} else {
		fmt.Println("Convênio:\"" + conv.NomeConv + "\" já existe com o mesmo nome.")
	}
}

// (READ) Retorna um Vetor de Convenios passando como parâmetro o "Nome" do convênio.
// Se o argumento "nome" = "*", retornará todos os convênios armazenados.
func getConveniosPorNome(conv string) []convenio.Convenio {
	convs, err := armazenamento.GetConveniosByName(conv)
	if err != nil {
		fmt.Println("Erro:", err)
		return nil
	}
	if convs == nil {
		fmt.Println("Erro: Convênio: " + conv + " não encontrado.")
		return nil
	}
	return convs
}

// (READ) Retorna um Convenio passando como parâmetro o "ID" do convênio.
func getConvenioPorId(id primitive.ObjectID) convenio.Convenio {
	conv, err := armazenamento.GetConvenioById(id)
	if err != nil {
		fmt.Println("Erro:", err)
		return convenio.Convenio{}
	}
	if id == primitive.NilObjectID {
		fmt.Println("Erro: Convênio: " + id.String() + " não encontrado.")
		return convenio.Convenio{}
	}
	return conv
}

// (list) Retorna Lista de Convenios no formato Json ou Bson passando como parâmetro o "Nome" do convênio.
// Se o argumento "nome" = "*", retornará todos os convênios armazenados.
func listaConvenio(nome string, formato ...string) {
	fmt.Println("Localizando Convênios...")
	convs, err := armazenamento.GetConveniosByName(nome)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		// Se houver "formato" e do tipo "bson", imprima neste.
		if len(formato) > 0 && strings.EqualFold(formato[0], "bson") {
			fmt.Println("lista de Convênios:\n", convs)
			// Caso contrário, use por padrão "Json"
		} else {
			fmt.Println("lista de Convênios:\n", printJSON(convs))
		}
	}
}

// (UPDATE) Atualiza os Dados de um ou mais Convênio armazenado utilizando como parâmetro o Nome do Convênio("nome"),
// o Struct do Novo Convênio("novoConv") e a opção de alterar Todos("todos") simultaneamente.
func atualizaConvPorNome(nome string, novoConv convenio.Convenio, todos bool) {
	// Checa se o Nome do Convenio está vazio
	if nome == "" {
		fmt.Println("Erro: Nome do Convênio nulo/vazio.")
	} else {
		// Executar um Merge de nos Obejetos para poder verificar o resultado
		// convMerge :=
		// // Verifica se o convênio é válido
		// err := convenio.VerificarConvenio(novoConv)
		var err error
		if err != nil {
			fmt.Println("Erro:", err)
		} else {
			// Atualiza os dados do Convênio
			result, err := armazenamento.UpdateConvenioByName(nome, novoConv, todos)
			if err != nil {
				fmt.Println("Erro:", err)
			} else {
				fmt.Println("Convenios encontrados:", result.MatchedCount)
				fmt.Println("Convenios atualizados:", result.ModifiedCount)
			}
		}
	}
}

// (UPDATE) Atualiza os Dados de um Convênio armazenado utilizando como parâmetro o ID Convênio,
func atualizaConvPorId(id primitive.ObjectID, novoConv convenio.Convenio) {
	// Checa se o ID do Convenio está vazio
	if id.IsZero() {
		fmt.Println("Erro: ID nulo/vazio.")
	} else {
		// Atualiza os dados do Convênio
		result, err := armazenamento.UpdateConvenioById(id, novoConv)
		if err != nil {
			fmt.Println("Erro:", err)
		} else if result.ModifiedCount > 0 {
			fmt.Println("Convenio atualizado:", id.String())
		} else {
			fmt.Println("Convênio:\"" + id.String() + "\" não foi alterado ou não existe no Armazém.")
		}
	}
}

// (DELETE) Deleta um Convênio específico ou mais de um utilizando o Nome do Convênio como parâmetro de busca.
// Para Deletar todos os Convênios da busca é possível utilizar o parâmetro Boleano "todos".
func deletaConveniosPorNome(nome string, todos bool) {
	// Checa se o Nome do Convenio está vazio
	if nome == "" {
		fmt.Println("Erro: Nome do Convênio nulo/vazio.")
	} else {
		result, err := armazenamento.DeleteConvenioByName(nome, todos)
		if err != nil {
			fmt.Println("Erro:", err)
			fmt.Println("Provavel que o Convênio:\"" + nome + "\" não exista no Armazém.")
		} else {
			fmt.Println("Convenios deletados:", result.DeletedCount)
		}
	}
}

// (DELETE) Deleta um Convênio específico utilizando o ID do Convênio como parâmetro de busca.
func deletaConvenioPorId(id primitive.ObjectID) {
	// Checa se o Nome do Convenio está vazio
	if id.IsZero() {
		fmt.Println("Erro: ID nulo/vazio.")
	} else {
		result, err := armazenamento.DeleteConvenioById(id)
		if err != nil {
			fmt.Println("Erro:", err)
			fmt.Println("Provavel que o Convênio:\"" + id.String() + "\" não exista no Armazém.")
		} else {
			fmt.Println("Convenios deletados:", result.DeletedCount)
		}
	}
}

// CRUD Paciente
func CriaPaciente(pac paciente.Paciente) {

}

// // CRUD PlanoPgto
// Verificar se o Plano de Pgto passado é válido
func VerificaPlanoPgto(plano planopgto.PlanoPgto) {
	err := planopgto.VerificarPlano(plano)
	if err != nil {
		fmt.Println("Erro:", err)
	}
}

// // (CREATE) Cria objeto PlanoPgto do tipo Abstrato
// func criaPlanoPgto(plano planopgto.PlanoPgto) {
// 	// Checa se o PlanoPgto está ok
// 	err := planopgto.VerificarPlano(plano)
// 	if err != nil {
// 		fmt.Println("Erro:", err)
// 	}
// 	// Checar se o Convênio está válido
// 	// Checa se o PAciente já possui um Plano nesse convênio com o mesmo núemro
// 	planos, err := armazenamento.GetPlanoPgto(plano.Convenio.NomeConv)
// 	if err != nil {
// 		fmt.Println("Erro:", err)
// 	}
// 	if planos == nil {
// 		result, err := armazenamento.CriarPlanoPgto(plano)
// 		if err != nil {
// 			fmt.Println("Erro:", err)
// 		} else {
// 			fmt.Println("Plano de Pagamento salvo:", result)
// 		}
// 	} else {
// 		fmt.Println("Plano de Pagamento:\"" + plano.Convenio.NomeConv + "\" já existe.")
// 	}
// }

// // (READ) Retorna um Vetor de Planos de Pagamento passando como parâmetro o "Convênio" do Plano.
// // Se o argumento "nome" = "*", retornará todos os convênios armazenados.
// func GetPlanoPgto(conv string) []planopgto.PlanoPgto {
// 	planos, err := armazenamento.GetPlanoPgto(conv)
// 	if err != nil {
// 		fmt.Println("Erro:", err)
// 		return nil
// 	}
// 	if planos == nil {
// 		fmt.Println("Erro: Plano de Pagamento: " + conv + " não encontrado.")
// 		return nil
// 	}
// 	return planos
// }

// // (READ) Retorna um Vetor de Planos de Pagamento passando um "Paciente".
// func GetPlanoPgtoByPaciente(pac paciente.Paciente) []planopgto.PlanoPgto {
// 	return nil
// }
