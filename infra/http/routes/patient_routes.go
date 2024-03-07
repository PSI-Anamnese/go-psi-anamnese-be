package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/psi-anamnese/psi-anamnese-be/application/services"
	"github.com/psi-anamnese/psi-anamnese-be/infra/database/repositories"
	"github.com/psi-anamnese/psi-anamnese-be/infra/http/controllers"
)

var _ Route = (*PatientRoutes)(nil)

type PatientRoutes struct {
	engine            *gin.Engine
	patientController controllers.PatientController
}

func newPatientRoutes(engine *gin.Engine, controller controllers.PatientController) PatientRoutes {
	return PatientRoutes{
		engine:            engine,
		patientController: controller,
	}
}

func NewPatientRoutes(engine *gin.Engine) PatientRoutes {
	return newPatientRoutes(engine, controllers.NewPatientController(
		services.NewPatientService(
			repositories.NewPatientRepository(),
		)))
}

func (p PatientRoutes) Setup() {
	p.engine.GET("/patients/", p.patientController.GetAll)
	p.engine.GET("/patients/:id", p.patientController.GetOne)
	p.engine.PUT("/patients/", p.patientController.Save)
	p.engine.POST("/patients/:id", p.patientController.Update)
	p.engine.DELETE("/patients/:id", p.patientController.Delete)
}
