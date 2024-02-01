package controllers

import (
	"Agenda/services/validation"
	"fmt"

	"github.com/gin-gonic/gin"
)

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
