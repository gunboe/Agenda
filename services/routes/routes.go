package routes

import (
	"Agenda/controllers"
	config "Agenda/services"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getPacById/:pacId", controllers.FindPacById)
	r.GET("/getPacByCPF/:pacCPF", controllers.FindPacByCPF)
	r.POST("/createPac", controllers.CreatePac)
	r.PUT("/updatePacById/:pacId", controllers.UpdatePac)
	r.DELETE("/deletePacById/:pacId", controllers.DeletePacById)
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
