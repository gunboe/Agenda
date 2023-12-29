package main

import (
	"Agenda/pkgs/armazenamento"
	"Agenda/pkgs/convenio"
	"encoding/json"
	"fmt"
	"strings"
)

// Criar as Coleções/Tabelas e Funções para o Database "Agenda" do MongoDB
// Agendamentos	: C
// Pacientes	:
// Agentes		:
// Convenios	: vCRUD

// Converter struct para Json
func printJSON(input interface{}) string {
	s, _ := json.MarshalIndent(input, "", "\t")
	return string(s)
}

// CRUD Convenios
// Atualiza os Dados de um ou mais Convênio armazenado utilizando como parâmetro o Nome do Convênio("nome"),
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

// Cria convênio no armazém
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

// Lista Convenios de forma textual passando como parâmetro o "Nome" do convênio.
// Se o argumento "nome" = "*", retornará todos os convênios armazenados.
// A listagem é no formato Json(padrão) ou Bson.
func listaConvenio(nome string, formato ...string) {
	fmt.Println("Localizando Convênios...")
	var convs []convenio.Convenios
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

// Deleta um Convênio específico ou mais de acordo com o Nome do Convênio usado como parâmetro de busca.
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
