package convControllers

import (
	"Agenda/controllers"
	"Agenda/models"
	"Agenda/services/expandErro"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/////////////////////////////////
// Routes Control para Convênio
/////////////////////////////////

// Cria Convênio por Json
func CreateConv(c *gin.Context) {
	var convRequest models.Convenio
	var err error
	// Realiza o Marshal dos Campos da requição no Objeto
	if err = controllers.AvaliarRequest(c, &convRequest); err != nil {
		return
	}
	// Cria o Convenio se não houver erros legados(checagem de campo) ou Erros de Negocio
	result, err := CriaConvenio(convRequest)
	if err != nil {
		reqErro := expandErro.NewBadRequestError("Erros na criação do convênio (Regras de Negócio): " + err.Error())
		fmt.Println(reqErro)
		c.JSON(reqErro.Code, reqErro)
		return
	}
	c.JSON(http.StatusOK, result)
}

// Retornar um objeto Json do Convênio por ID
func FindConvById(c *gin.Context) {
	// Checa o Id recebido
	id, err := primitive.ObjectIDFromHex(c.Param("convId"))
	if err != nil {
		reqErro := expandErro.NewBadRequestError("Erro na pesquisa:(convId) " + err.Error())
		fmt.Println(reqErro)
		c.JSON(reqErro.Code, reqErro)
		return
	}
	// Tenta realizar a busca
	conv, err := GetConvenioPorId(id)
	if err != nil {
		reqErro := expandErro.NewNotFoundError("Erro na pesquisa: " + err.Error())
		fmt.Println(reqErro)
		c.JSON(reqErro.Code, reqErro)
		return
	}
	fmt.Println("Paciente encontrado com id:", id)
	c.JSON(http.StatusOK, conv)
}

// Retornar Todos ("*") objetos Json do Convênio ou por Nome
func FindConvenios(c *gin.Context) {
	// Tenta realizar a busca
	pacs := ListaConvenio(c.Param("nome"), "bson")
	jsize := len(pacs.([]models.Convenio))
	if jsize == 0 {
		reqErro := expandErro.NewNotFoundError("Erro na pesquisa: convênios não encontrados")
		fmt.Println(reqErro)
		c.JSON(reqErro.Code, reqErro)
		return
	}
	// Converte o Tipo Interface no Tipo dos dados reais e calucula o tamanho do array
	fmt.Println("Convênio(s) encontrado(s):", jsize)
	c.JSON(http.StatusOK, pacs)
}

// Atualiza um Convênio por Json
func UpdateConv(c *gin.Context) {
	var convRequest models.Convenio
	var reqErro *expandErro.Lasquera
	var err error
	// Checa ID
	id, err := primitive.ObjectIDFromHex(c.Param("convId"))
	if err != nil {
		reqErro := expandErro.NewBadRequestError("Erro na atualização do convênio: (convId) " + err.Error())
		fmt.Println(reqErro)
		c.JSON(reqErro.Code, reqErro)
		return
	}
	// Realiza o Marshal dos Campos da requição no Objeto
	if err = controllers.AvaliarRequest(c, &convRequest); err != nil {
		return
	}
	// Atualiza o Convênio
	err = AtualizaConvPorId(id, convRequest)
	if err != nil {
		reqErro = expandErro.NewBadRequestError("Erro na atualização do convênio (Regras de Negócio): " + err.Error())
		fmt.Println(reqErro)
		c.JSON(reqErro.Code, reqErro)
		return
	}
}

// RC: Atualiza um Convênio por Json
func IndispConv(c *gin.Context) {
	// var convRequest models.Convenio
	var reqErro *expandErro.Lasquera
	var err error
	// Checa ID
	convId, err := primitive.ObjectIDFromHex(c.Param("convId"))
	if err != nil {
		reqErro := expandErro.NewBadRequestError("Erro na disponibilização do convênio: (convId) " + err.Error())
		fmt.Println(reqErro)
		c.JSON(reqErro.Code, reqErro)
		return
	}
	// Obter o valor do atributo "disponivel"
	s := c.Query("indisponivel")
	b, err := strconv.ParseBool(s)
	if err != nil {
		reqErro := expandErro.NewBadRequestError("Erro na disponibilização do convênio: " + err.Error())
		fmt.Println(reqErro)
		c.JSON(reqErro.Code, reqErro)
		return
	}
	// Tenta realizar a alterção do atributo
	err = HabiliteConvPorId(convId, b)
	if err != nil {
		reqErro = expandErro.NewBadRequestError("Erro na disponibilização do convênio (Regras de negócio): " + err.Error())
		fmt.Println(reqErro)
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
		fmt.Println(reqErro)
		c.JSON(reqErro.Code, reqErro)
		return
	}
	// Tenta realizar a deleção
	err = DeletaConvenioPorId(id)
	if err != nil {
		reqErro := expandErro.NewNotFoundError("Erro na deleção: " + err.Error())
		fmt.Println(reqErro)
		c.JSON(reqErro.Code, reqErro)
		return
	}
}
