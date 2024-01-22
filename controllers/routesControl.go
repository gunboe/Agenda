package controllers

import (
	"Agenda/models"
	"Agenda/services/expandErro"
	"Agenda/services/validation"
	"fmt"

	"github.com/gin-gonic/gin"
)

/////////////////////////////////
// Routes Control para Pacientes
/////////////////////////////////

// RC: Retornar um objeto Json do Paciente por ID
func FindPacById(c *gin.Context) {
	var pacRequest models.Paciente
	var err error
	err = c.ShouldBindJSON(&pacRequest)
	if err != nil {
		reqErro := expandErro.NewBadRequestError("Existe alguns campos errados: " + err.Error())
		c.JSON(reqErro.Code, reqErro)
		return
	}
	fmt.Println((pacRequest))
	// pac, err := GetPacientePorId(id)

}

// RC: Retornar um objeto Json do Paciente por CPF
func FindPacByCPF(c *gin.Context) {
	// pac, err := GetPacientePorId(id)
}

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
	err = CriaPaciente(pacRequest)
	if err != nil {
		reqErro := expandErro.NewBadRequestError("Erros na regra de negócio: " + err.Error())
		c.JSON(reqErro.Code, reqErro)
		return
	}
}

// RC: Deleta um Paciente por ID
func DeletePacById(c *gin.Context) {
	// pac, err := GetPacientePorId(id)
}

// RC: Atualiza um Paciente por Json
func UpdatePac(c *gin.Context) {
	// pac, err := GetPacientePorId(id)
}

/////////////////////////////////
// Routes Control para Convênio
/////////////////////////////////

// RC: Retornar um objeto Json do Convênio por ID
func FindConvById(c *gin.Context) {
	// pac, err := GetConvPorId(id)
}

// RC: Retornar um objeto Json do Convênio por Nome
func FindConvByName(c *gin.Context) {
	// pac, err := GetConvPorNome(id)
}

// RC: Cria Convênio por Json
func CreateConv(c *gin.Context) {
	var pac models.Paciente
	// pac = c.
	CriaPaciente(pac)
}

// RC: Deleta um Convênio por ID
func DeleteConvById(c *gin.Context) {
	// pac, err := GetPacientePorId(id)
}

// RC: Atualiza um Convênio por Json
func UpdateConv(c *gin.Context) {
	// pac, err := GetPacientePorId(id)
}
