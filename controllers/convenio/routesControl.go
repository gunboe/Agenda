package convControllers

import (
	"Agenda/models"
	"Agenda/services/expandErro"
	"Agenda/services/validation"
	"fmt"

	"github.com/gin-gonic/gin"
)

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
	var convRequest models.Convenio
	var err error
	// Realiza o Marshal dos Campos da requição no Objeto
	err = c.ShouldBindJSON(&convRequest)
	if err != nil {
		// Existindo um erro, ele será enviado para validação do Convênio
		reqErro := validation.ValidaCamposReq(err)
		fmt.Println(err)
		c.JSON(reqErro.Code, reqErro)
		return
	}
	// Cria o Convenio se não houver erros legados(checagem de campo) ou Erros de Negocio
	err = CriaConvenio(convRequest)
	if err != nil {
		reqErro := expandErro.NewBadRequestError("Erros na regra de negócio: " + err.Error())
		c.JSON(reqErro.Code, reqErro)
		return
	}
}

// RC: Deleta um Convênio por ID
func DeleteConvById(c *gin.Context) {
	// pac, err := GetPacientePorId(id)
}

// RC: Atualiza um Convênio por Json
func UpdateConv(c *gin.Context) {
	// pac, err := GetPacientePorId(id)
}
