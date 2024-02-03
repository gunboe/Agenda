package controllers

import (
	"Agenda/services/validation"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type respGeneric struct {
	ID         primitive.ObjectID `json:"id,omitempty"`
	Criado     bool               `json:"criado,omitempty"`
	Inserido   bool               `json:"inserido,omitempty"`
	Atualizado bool               `json:"atualizado,omitempty"` // Serve também para o atributo "Indisponivel"
	Deletado   bool               `json:"deletado,omitempty"`
}

//////////////////////////////////
// Funções de Controller em Geral
//////////////////////////////////

func AvaliarRequest(c *gin.Context, obj interface{}) error {
	var err error
	err = c.ShouldBindJSON(obj)
	if err != nil {
		// Existindo um erro, ele será enviado para validação do Paciente
		reqErro := validation.ValidaCamposReq(err)
		fmt.Println(err)
		c.JSON(reqErro.Code, reqErro)
	}
	return err
}

func Resposta(c *gin.Context, result interface{}) {
	var resp respGeneric
	resp.ID = result.(primitive.ObjectID)
	path := c.Request.URL.Path
	switch {
	case strings.Contains(path, "create"):
		resp.Criado = true
	case strings.Contains(path, "insPlan"):
		resp.Inserido = true
	case (strings.Contains(path, "update") || strings.Contains(path, "indisp") || strings.Contains(path, "bloqPac")):
		resp.Atualizado = true
	case strings.Contains(path, "delete"):
		resp.Deletado = true
	}
	c.JSON(http.StatusOK, resp)
}
