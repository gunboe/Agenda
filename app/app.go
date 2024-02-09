package app

import (
	convControllers "Agenda/controllers/convenio"
	pacControllers "Agenda/controllers/paciente"
	"Agenda/services/config"
	"Agenda/services/logger"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

// Atributos da Aplicação
type Application struct {
	ConvFunc     *convControllers.ConvenioFunc
	PacienteFunc *pacControllers.PacienteFunc
	Configuracao config.Config
}

func (app *Application) Run() {
	fmt.Println("-- Executando a Aplicação --")
	// Inicializa o Logger
	logger.InicializaLogger(app.Configuracao)
	logger.Info("Logger inciado")
	// Incializa o Router
	InicializaRouter(app)
}

func InitRoutes(r *gin.RouterGroup, app *Application) {
	// Paciente
	r.POST("/createPac", app.PacienteFunc.CreatePac)
	r.POST("/insPlanoPac/:pacId", app.PacienteFunc.InserePlanoPac)
	r.GET("/getPacById/:pacId", app.PacienteFunc.FindPacById)
	r.GET("/getPacientes/:nome", app.PacienteFunc.FindPacientes)
	r.GET("/getPacByCPF/:pacCPF", app.PacienteFunc.FindPacByCPF)
	r.PUT("/updatePacById/:pacId", app.PacienteFunc.UpdatePac)
	r.PATCH("/bloqPacById/:pacId", app.PacienteFunc.BloqPac)
	r.DELETE("/deletePacById/:pacId", app.PacienteFunc.DeletePacById)
	r.DELETE("/deletePlanoPac/:pacId/:planoId", app.PacienteFunc.DelPlanoPac)
	// Convenio
	r.POST("/createConv", app.ConvFunc.CreateConv)
	r.GET("/getConvById/:convId", app.ConvFunc.FindConvById)
	r.GET("/getConvenios/:nome", app.ConvFunc.FindConvenios)
	r.PUT("/updateConvById/:convId", app.ConvFunc.UpdateConv)
	r.PATCH("/indispConvById/:convId", app.ConvFunc.IndispConv)
	r.DELETE("/deleteConvById/:convId", app.ConvFunc.DeleteConvById)
}

// func InicializaRouter(cF *convControllers.ConvenioFunc) {
func InicializaRouter(app *Application) {
	// Iniciando o Roteador
	router := gin.Default()
	rv1 := router.Group("/api/v1")
	// Inicializando API - daqui não roda mais nada
	InitRoutes(rv1, app)
	err := router.Run(":" + fmt.Sprint(app.Configuracao.ApiServerPort))
	if err != nil {
		log.Fatal("Roteador falhou:", err)
	}
}
