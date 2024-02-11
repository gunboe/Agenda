package router

import (
	convControllers "Agenda/controllers/convenio"
	pacControllers "Agenda/controllers/paciente"
	"Agenda/services/config"
	"Agenda/services/logger"
	"fmt"

	"github.com/gin-gonic/gin"
)

// var r *Router

type Funcs struct {
	FuncsConv     *convControllers.ConvenioFunc
	FuncsPaciente *pacControllers.PacienteFunc
}

func InitRoutes(rg *gin.RouterGroup, fr *Funcs) {
	// Paciente
	rg.POST("/createPac", fr.FuncsPaciente.CreatePac)
	rg.POST("/insPlanoPac/:pacId", fr.FuncsPaciente.InserePlanoPac)
	rg.GET("/getPacById/:pacId", fr.FuncsPaciente.FindPacById)
	rg.GET("/getPacientes/:nome", fr.FuncsPaciente.FindPacientes)
	rg.GET("/getPacByCPF/:pacCPF", fr.FuncsPaciente.FindPacByCPF)
	rg.PUT("/updatePacById/:pacId", fr.FuncsPaciente.UpdatePac)
	rg.PATCH("/bloqPacById/:pacId", fr.FuncsPaciente.BloqPac)
	rg.DELETE("/deletePacById/:pacId", fr.FuncsPaciente.DeletePacById)
	rg.DELETE("/deletePlanoPac/:pacId/:planoId", fr.FuncsPaciente.DelPlanoPac)
	// Convenio
	rg.POST("/createConv", fr.FuncsConv.CreateConv)
	rg.GET("/getConvById/:convId", fr.FuncsConv.FindConvById)
	rg.GET("/getConvenios/:nome", fr.FuncsConv.FindConvenios)
	rg.PUT("/updateConvById/:convId", fr.FuncsConv.UpdateConv)
	rg.PATCH("/indispConvById/:convId", fr.FuncsConv.IndispConv)
	rg.DELETE("/deleteConvById/:convId", fr.FuncsConv.DeleteConvById)
}

// func InicializaRouter(cF *convControllers.ConvenioFunc) {
func InitRouter(fDB *Funcs, conf config.Config) error {
	// Iniciando o Roteador
	var r *gin.Engine
	r = gin.Default()
	rv1 := r.Group("/api/v1")
	// Inicializando API - daqui não roda mais nada
	InitRoutes(rv1, fDB)
	logger.Info("Router inciado")
	if err := r.Run(":" + fmt.Sprint(conf.ApiServerPort)); err != nil {
		return err
	}
	return nil
}
