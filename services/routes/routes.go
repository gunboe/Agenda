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
	r.POST("/createPac", pacControllers.CreatePac)
	r.GET("/getPacById/:pacId", pacControllers.FindPacById)
	r.GET("/getPacientes/:nome", pacControllers.FindPacientes)
	r.GET("/getPacByCPF/:pacCPF", pacControllers.FindPacByCPF)
	r.PUT("/updatePacById/:pacId", pacControllers.UpdatePac)
	r.DELETE("/deletePacById/:pacId", pacControllers.DeletePacById)
	// Convenio
	r.POST("/createConv", convControllers.CreateConv)
	r.GET("/getConvById/:convId", convControllers.FindConvById)
	r.GET("/getConvenios/:nome", convControllers.FindConvenios)
	r.PUT("/updateConvById/:convId", convControllers.UpdateConv)
	r.PATCH("/dispConvById/:convId", convControllers.DispConv)
	r.DELETE("/deleteConvById/:convId", convControllers.DeleteConvById)

}

func InicializaRouter() {
	// Iniciando o Roteador
	router := gin.Default()
	rv1 := router.Group("/api/v1")
	// InitRoutes(&router.RouterGroup)
	InitRoutes(rv1)
	err := router.Run(":" + fmt.Sprint(config.ConfigInicial.WebServerPort))
	if err != nil {
		log.Fatal("Roteador falhou:", err)
	}
}
