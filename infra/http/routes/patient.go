package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/psi-anamnese/psi-anamnese-be/infra/http/controllers"
)

type PatientRoutes struct {
	engine            *gin.Engine
	patientController controllers.PatientController
}

func NewPatientRoutes(engine *gin.Engine, controller controllers.PatientController) PatientRoutes {
	return PatientRoutes{
		engine:            engine,
		patientController: controller,
	}
}

func (p PatientRoutes) Setup() {
	p.engine.GET("/patients/", p.patientController.GetAll)
	p.engine.GET("/patients/:id", p.patientController.GetOne)
	p.engine.PUT("/patients/", p.patientController.Save)
	p.engine.POST("/patients/:id", p.patientController.Update)
	p.engine.DELETE("/patients/:id", p.patientController.Delete)
}
