package pacControllers

import (
	"Agenda/common"
	"Agenda/models"
	"Agenda/services/expandErro"
	"Agenda/services/validation"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/////////////////////////////////
// Routes Control para Pacientes
/////////////////////////////////

// RC: Cria Paciente por Json
func CreatePac(c *gin.Context) {
	var pacRequest models.Paciente
	var err error
	// Realiza o Marshal dos Campos da requição no Objeto
	err = c.ShouldBindJSON(&pacRequest)
	if err != nil {
		// Existindo um erro, ele será enviado para validação do Paciente
		reqErro := validation.ValidaCamposReq(err)
		fmt.Println(err)
		c.JSON(reqErro.Code, reqErro)
		return
	}
	// Cria o Paciente se não houver erros legados(checagem de campo) ou Erros de Negocio
	result, err := CriaPaciente(pacRequest)
	if err != nil {
		reqErro := expandErro.NewBadRequestError("Erros na regra de negócio: " + err.Error())
		c.JSON(reqErro.Code, reqErro)
		return
	}
	c.JSON(http.StatusOK, result)
}

// RC: Retornar um objeto Json do Paciente por ID
func FindPacById(c *gin.Context) {
	// Checa o Id recebido
	id, err := primitive.ObjectIDFromHex(c.Param("pacId"))
	if err != nil {
		reqErro := expandErro.NewBadRequestError("Erro na pesquisa:(pacId) " + err.Error())
		c.JSON(reqErro.Code, reqErro)
		return
	}
	// Tenta realizar a busca
	pac, err := GetPacientePorId(id)
	if err != nil {
		reqErro := expandErro.NewNotFoundError("Erro na pesquisa: " + err.Error())
		c.JSON(reqErro.Code, reqErro)
		fmt.Println(reqErro.Mensagem)
		return
	}
	fmt.Println("paciente encontrado com id:", id)
	c.JSON(http.StatusOK, pac)
}

// RC: Retornar TODOS ("*") objetos Json do Paciente ou por Nome
func FindPacientes(c *gin.Context) {
	// Tenta realizar a busca
	// nome := c.Param("pac")
	pacs := ListaPaciente(c.Param("nome"), "bson")
	jsize := len(pacs.([]models.Paciente))
	if jsize == 0 {
		reqErro := expandErro.NewNotFoundError("erro na pesquisa: " + "paciente(s) não encontrado(s)")
		c.JSON(reqErro.Code, reqErro)
		fmt.Println(reqErro.Mensagem)
		return
	}
	// Converte o Tipo Interface no Tipo dos dados reais e calucula o tamanho do array
	fmt.Println("paciente(s) encontrado(s):", jsize)
	c.JSON(http.StatusOK, pacs)
}

// RC: Retornar um objeto Json do Paciente por CPF
func FindPacByCPF(c *gin.Context) {
	var err error
	cpf := c.Param("pacCPF")
	// Checa o CPF recebido
	if _, ok := common.CPFvalido(cpf); !ok {
		err = errors.New("CPF inválido")
		reqErro := expandErro.NewBadRequestError("Erro na pesquisa:(pacCPF) " + err.Error())
		c.JSON(reqErro.Code, reqErro)
		fmt.Println(reqErro)
		return
	}
	// Tenta realizar a busca
	pac, err := GetPacientePorCPF(cpf)
	if err != nil {
		reqErro := expandErro.NewNotFoundError("Erro na pesquisa: " + err.Error())
		c.JSON(reqErro.Code, reqErro)
		fmt.Println(reqErro.Mensagem)
		return
	}
	fmt.Println("Paciente encontrado com id:", cpf)
	c.JSON(http.StatusOK, pac)
}

// RC: Atualiza um Paciente por Json
func UpdatePac(c *gin.Context) {
	var pacRequest models.Paciente
	var reqErro *expandErro.Lasquera
	var err error
	// Checa ID
	id, err := primitive.ObjectIDFromHex(c.Param("pacId"))
	if err != nil {
		reqErro := expandErro.NewBadRequestError("erro na atualização do Paciente: (pacId) " + err.Error())
		c.JSON(reqErro.Code, reqErro)
		return
	}
	// Realiza o Marshal dos Campos da requição no Objeto
	err = c.ShouldBindJSON(&pacRequest)
	if err != nil {
		// Existindo um erro, ele será enviado para validação do Objeto
		reqErro := validation.ValidaCamposReq(err)
		fmt.Println(err)
		c.JSON(reqErro.Code, reqErro)
		return
	}
	// Atualiza o Objeto
	err = AtualizaPacPorId(id, pacRequest)
	if err != nil {
		reqErro = expandErro.NewBadRequestError("erro na atualização do Paciente (regras de negócio): " + err.Error())
		c.JSON(reqErro.Code, reqErro)
		return
	}
}

// RC: Deleta um Paciente por ID
func DeletePacById(c *gin.Context) {
	var err error
	// Checa ID
	id, err := primitive.ObjectIDFromHex(c.Param("pacId"))
	if err != nil {
		reqErro := expandErro.NewBadRequestError("Erro na deleção: (pacId) " + err.Error())
		c.JSON(reqErro.Code, reqErro)
		fmt.Println(reqErro.Mensagem)
		return
	}
	// Tenta realizar a deleção
	err = DeletaPacientePorId(id)
	if err != nil {
		reqErro := expandErro.NewNotFoundError("Erro na deleção: " + err.Error())
		c.JSON(reqErro.Code, reqErro)
		fmt.Println(reqErro.Mensagem)
		return
	}
}
