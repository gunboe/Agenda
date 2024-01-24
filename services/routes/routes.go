package routes

import (
	convControllers "Agenda/controllers/convenio"
	pacControllers "Agenda/controllers/paciente"
	"Agenda/services/config"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	// Paciente
	r.GET("/getPacById/:pacId", pacControllers.FindPacById)
	r.GET("/getPacByCPF/:pacCPF", pacControllers.FindPacByCPF)
	r.POST("/createPac", pacControllers.CreatePac)
	r.PUT("/updatePacById/:pacId", pacControllers.UpdatePac)
	r.DELETE("/deletePacById/:pacId", pacControllers.DeletePacById)
	// Convenio
	r.POST("/createConv", convControllers.CreateConv)
}

func InicializaRouter() {
	// Iniciando o Roteador
	router := gin.Default()
	InitRoutes(&router.RouterGroup)
	err := router.Run(":" + fmt.Sprint(config.ConfigInicial.WebServerPort))
	if err != nil {
		log.Fatal("Roteador falhou:", err)
	}
}
