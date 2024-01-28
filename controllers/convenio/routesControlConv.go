package convControllers

import (
	"Agenda/models"
	"Agenda/services/expandErro"
	"Agenda/services/validation"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/////////////////////////////////
// Routes Control para Convênio
/////////////////////////////////

// RC: Retornar um objeto Json do Convênio por ID
func FindConvById(c *gin.Context) {
	// Checa o Id recebido
	id, err := primitive.ObjectIDFromHex(c.Param("convId"))
	if err != nil {
		reqErro := expandErro.NewBadRequestError("Erro na pesquisa:(convId) " + err.Error())
		c.JSON(reqErro.Code, reqErro)
		return
	}
	// Tenta realizar a busca
	conv, err := GetConvenioPorId(id)
	if err != nil {
		reqErro := expandErro.NewNotFoundError("Erro na pesquisa: " + err.Error())
		c.JSON(reqErro.Code, reqErro)
		fmt.Println(reqErro.Mensagem)
		return
	}
	fmt.Println("paciente encontrado com id:", id)
	c.JSON(http.StatusOK, conv)
}

// RC: Retornar Todos ("*") objetos Json do Convênio ou por Nome
func FindConvenios(c *gin.Context) {
	// Tenta realizar a busca
	pacs := ListaConvenio(c.Param("nome"), "bson")
	jsize := len(pacs.([]models.Convenio))
	if jsize == 0 {
		reqErro := expandErro.NewNotFoundError("erro na pesquisa: " + "convênio(s) não encontrado(s)")
		c.JSON(reqErro.Code, reqErro)
		fmt.Println(reqErro.Mensagem)
		return
	}
	// Converte o Tipo Interface no Tipo dos dados reais e calucula o tamanho do array
	fmt.Println("convênio(s) encontrado(s):", jsize)
	c.JSON(http.StatusOK, pacs)
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
	var err error
	// Checa ID
	id, err := primitive.ObjectIDFromHex(c.Param("convId"))
	if err != nil {
		reqErro := expandErro.NewBadRequestError("Erro na deleção: (convId) " + err.Error())
		c.JSON(reqErro.Code, reqErro)
		return
	}
	// Tenta realizar a deleção
	err = DeletaConvenioPorId(id)
	if err != nil {
		reqErro := expandErro.NewNotFoundError("Erro na deleção: " + err.Error())
		c.JSON(reqErro.Code, reqErro)
		return
	}
}

// RC: Atualiza um Convênio por Json
func UpdateConv(c *gin.Context) {
	// pac, err := GetPacientePorId(id)
}
