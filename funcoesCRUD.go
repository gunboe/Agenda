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
	"Agenda/pkgs/convenio"
	"Agenda/pkgs/paciente"
	"Agenda/pkgs/planosaude"
	"encoding/json"
	"fmt"
	"strings"
)

// Converter struct para Json
func printJSON(input interface{}) string {
	s, _ := json.MarshalIndent(input, "", "\t")
	return string(s)
}

// CRUD Convenios

// (CREATE) Cria convênio e salva no armazém
func criaConvenio(conv convenio.Convenios) {
	// Checa se já existe Convenio
	convs, err := armazenamento.GetConvenios(conv.NomeConv)
	if err != nil {
		fmt.Println("Erro:", err)
	}
	if convs == nil {
		result, err := armazenamento.CriarConvenio(conv)
		if err != nil {
			fmt.Println("Erro:", err)
		} else {
			fmt.Println("Convenio salvo:", result)
		}
	} else {
		fmt.Println("Convênio:\"" + conv.NomeConv + "\" já existe.")
	}
}

// (READ-funcoesCRUD.go) Retorna um Vetor de Convenios passando como parâmetro o "Nome" do convênio.
// Se o argumento "nome" = "*", retornará todos os convênios armazenados.
func getConvenios(conv string) []convenio.Convenios {
	convs, err := armazenamento.GetConvenios(conv)
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

// (list) Retorna Lista de Convenios no formato Json ou Bson passando como parâmetro o "Nome" do convênio.
// Se o argumento "nome" = "*", retornará todos os convênios armazenados.
func listaConvenio(nome string, formato ...string) {
	fmt.Println("Localizando Convênios...")
	convs, err := armazenamento.GetConvenios(nome)
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
func atualizaConv(nome string, novoConv convenio.Convenios, todos bool) {
	// Checa se já existe Convenio
	if nome == "" || novoConv.NomeConv == "" {
		fmt.Println("Erro: Insira um nome de convênio válido para atualizar.")
	} else {
		// Verifica se o convênio é válido
		err := convenio.VerificarConvenio(novoConv)
		if err != nil {
			fmt.Println("Erro:", err)
		} else {
			// Atualiza os dados do Convênio
			result, err := armazenamento.AtualizarConvenio(nome, novoConv, todos)
			if err != nil {
				fmt.Println("Erro:", err)
			} else {
				fmt.Println("Convenios atualizados:", result.ModifiedCount)
			}
		}
	}
}

// (DELETE) Deleta um Convênio específico ou mais de acordo com o Nome do Convênio usado como parâmetro de busca.
// Para Deletar todos os Convênios da busca é possível utilizar o parâmetro Boleano "todos".
func deletaConvenio(sconv string, todos bool) {
	// Checa se existe algum Convenio com o padrão de Nome passado
	convs, err := armazenamento.GetConvenios(sconv)
	if err != nil {
		fmt.Println(err)
	}
	if convs != nil {
		result, err := armazenamento.DeletarConvenio(sconv, todos)
		if err != nil {
			fmt.Println(err)
		} else {
			var p string
			for _, s := range convs {
				p += " " + s.NomeConv
			}
			fmt.Println("Convenios deletados:", result.DeletedCount, "("+strings.TrimSpace(p)+")")
		}
	} else {
		fmt.Println("Convênio:\"" + sconv + "\" não existe no Armazém.")
	}
}

// CRUD PlanoSaude

// (CREATE) Cria PlanoSaude e salva no armazém
func criaPlanoSaude(plano planosaude.PlanoSaude) {
	// Checa se já existe PlanoSaude
	planos, err := armazenamento.GetPlanoSaude(plano.Convenio.NomeConv)
	if err != nil {
		fmt.Println("Erro:", err)
	}
	if planos == nil {
		result, err := armazenamento.CriarPlanoSaude(plano)
		if err != nil {
			fmt.Println("Erro:", err)
		} else {
			fmt.Println("Convenio salvo:", result)
		}
	} else {
		fmt.Println("Convênio:\"" + plano.Convenio.NomeConv + "\" já existe.")
	}
}

// (READ) Retorna um Vetor de Planos de Saude passando como parâmetro o "Convênio" do Plano.
// Se o argumento "nome" = "*", retornará todos os convênios armazenados.
func GetPlanoSaude(conv string) []planosaude.PlanoSaude {
	planos, err := armazenamento.GetPlanoSaude(conv)
	if err != nil {
		fmt.Println("Erro:", err)
		return nil
	}
	if planos == nil {
		fmt.Println("Erro: Convênio: " + conv + " não encontrado.")
		return nil
	}
	return planos
}

// (READ) Retorna um Vetor de Planos de Saude passando um "Paciente".
func GetPlanoSaudeByPaciente(pac paciente.Paciente) []planosaude.PlanoSaude {
	return nil
}
