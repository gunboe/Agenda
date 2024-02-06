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
func (convFunc *ConvenioFunc) CreateConv(c *gin.Context) {
	var err error
	var convRequest models.Convenio
	// Realiza o Marshal dos Campos da requição no Objeto
	if err = controllers.AvaliarRequest(c, &convRequest); err != nil {
		return
	}
	// Cria o Convenio se não houver erros legados(checagem de campo) ou Erros de Negocio
	result, err := convFunc.CriaConvenio(convRequest)
	if err != nil {
		reqErro := expandErro.NewBadRequestError("Erros na criação do convênio (Regras de Negócio): " + err.Error())
		c.JSON(reqErro.Code, reqErro)
		return
	}
	// Processa e envia resposta (Qdo é para criação é necessário obter o ID do resultado do Mongo)
	controllers.Resposta(c, result)
}

// Retornar um objeto Json do Convênio por ID
func (convFunc *ConvenioFunc) FindConvById(c *gin.Context) {
	// Checa o Id recebido
	id, err := primitive.ObjectIDFromHex(c.Param("convId"))
	if err != nil {
		reqErro := expandErro.NewBadRequestError("Erro na pesquisa:(convId) " + err.Error())
		c.JSON(reqErro.Code, reqErro)
		return
	}
	// Tenta realizar a busca
	conv, err := convFunc.GetConvenioPorId(id)
	if err != nil {
		reqErro := expandErro.NewNotFoundError("Erro na pesquisa: " + err.Error())
		c.JSON(reqErro.Code, reqErro)
		return
	}
	fmt.Println("Convenio encontrado com id:", id)
	c.JSON(http.StatusOK, conv)
}

// Retornar Todos ("*") objetos Json do Convênio ou por Nome
func (convFunc *ConvenioFunc) FindConvenios(c *gin.Context) {
	// Tenta realizar a busca
	convs := convFunc.ListaConvenio(c.Param("nome"), "bson")
	jsize := len(convs.([]models.Convenio))
	if jsize == 0 {
		reqErro := expandErro.NewNotFoundError("Erro na pesquisa: convênios não encontrados")
		c.JSON(reqErro.Code, reqErro)
		return
	}
	// Converte o Tipo Interface no Tipo dos dados reais e calucula o tamanho do array
	fmt.Println("Convênio(s) encontrado(s):", jsize)
	c.JSON(http.StatusOK, convs)
}

// Atualiza um Convênio por Json
func (convFunc *ConvenioFunc) UpdateConv(c *gin.Context) {
	var convRequest models.Convenio
	var reqErro *expandErro.Lasquera
	var err error
	// Checa ID
	id, err := primitive.ObjectIDFromHex(c.Param("convId"))
	if err != nil {
		reqErro := expandErro.NewBadRequestError("Erro na atualização do convênio: (convId) " + err.Error())
		c.JSON(reqErro.Code, reqErro)
		return
	}
	// Realiza o Marshal dos Campos da requição no Objeto
	if err = controllers.AvaliarRequest(c, &convRequest); err != nil {
		return
	}
	// Atualiza o Convênio
	err = convFunc.AtualizaConvPorId(id, convRequest)
	if err != nil {
		reqErro = expandErro.NewBadRequestError("Erro na atualização do convênio (Regras de Negócio): " + err.Error())
		c.JSON(reqErro.Code, reqErro)
		return
	}
	// Processa e envia resposta
	controllers.Resposta(c, id)
}

// (In)Disponibiliza Convênio
func (convFunc *ConvenioFunc) IndispConv(c *gin.Context) {

	// var convRequest models.Convenio
	var reqErro *expandErro.Lasquera
	var err error
	// Checa ID
	id, err := primitive.ObjectIDFromHex(c.Param("convId"))
	if err != nil {
		reqErro := expandErro.NewBadRequestError("Erro: ID convênio: " + err.Error())
		c.JSON(reqErro.Code, reqErro)
		return
	}
	// Obter o valor do atributo "disponivel"
	s := c.Query("indisponivel")
	b, err := strconv.ParseBool(s)
	if err != nil {
		if s == "" {
			reqErro = expandErro.NewBadRequestError("Erro: Query: indisponivel não encontrada")
		} else {
			reqErro = expandErro.NewBadRequestError("Erro: Query: indisponivel deve ser true ou false, recebido: " + s)
		}
		c.JSON(reqErro.Code, reqErro)
		return
	}
	// Tenta realizar a alterção do atributo
	err = convFunc.HabiliteConvPorId(id, b)
	if err != nil {
		reqErro = expandErro.NewBadRequestError("Erro na disponibilização do convênio (Regras de negócio): " + err.Error())
		c.JSON(reqErro.Code, reqErro)
		return
	}
	// Processa e envia resposta
	controllers.Resposta(c, id)
}

// Deleta um Convênio por ID
func (convFunc *ConvenioFunc) DeleteConvById(c *gin.Context) {

	// pac, err := GetPacientePorId(id)
	var err error
	var reqErro *expandErro.Lasquera
	// Checa ID
	id, err := primitive.ObjectIDFromHex(c.Param("convId"))
	if err != nil {
		reqErro = expandErro.NewBadRequestError("Erro na deleção: (convId) " + err.Error())
		c.JSON(reqErro.Code, reqErro)
		return
	} else {
		// Tenta realizar a deleção
		err = convFunc.DeletaConvenioPorId(id)
		if err != nil {
			reqErro = expandErro.NewNotFoundError("Erro na deleção: " + err.Error())
			c.JSON(reqErro.Code, reqErro)
			return
		}
	}
	// Processa e envia resposta
	controllers.Resposta(c, id)
}
