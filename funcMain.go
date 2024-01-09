package main

import (
	"Agenda/pkgs/common"
	"encoding/json"
	"fmt"
	"strings"
)

////////////////////////////
// Funçoes Iniciais do Main
////////////////////////////

// Inicializa o ambiente
func inicializaAmbiente() {
	// Carrega as configurações
	var conf common.Config
	conf = common.ConfigInicial
	// Conecta ao Banco
	fmt.Println("Utilizando o DataBase:", conf.ArmazemDatabase)
	// Testa o Banco relacionando todos os Convênnios e Pacientes
	todosConvs := getConveniosPorNome("*")
	todosPacientes := getPacientesPorNome("*")
	// Listagem de Convenios
	var listaConvs string
	for _, v := range todosConvs {
		listaConvs += " \"" + v.NomeConv + "\""
	}
	listaConvs = strings.TrimSpace(listaConvs)
	fmt.Println("Lista de Todos os Convenios:", listaConvs)
	// Listagem de Pacientes
	var listaPacs string
	for _, p := range todosPacientes {
		listaPacs += " \"" + p.Nome + "\""
	}
	listaPacs = strings.TrimSpace(listaPacs)
	fmt.Println("Lista de Todos os Pacientes:", listaPacs)

	// Avisa que está pronto
	fmt.Println("Ambiente pronto para uso!\n")
}

// Converter struct para Json
func printJSON(input interface{}) string {
	s, _ := json.MarshalIndent(input, "", "\t")
	return string(s)
}
