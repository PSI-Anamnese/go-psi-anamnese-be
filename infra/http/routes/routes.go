package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/psi-anamnese/psi-anamnese-be/application/services"
	"github.com/psi-anamnese/psi-anamnese-be/infra/database/repositories"
	"github.com/psi-anamnese/psi-anamnese-be/infra/http/controllers"
)

type Routes []Route

type Route interface {
	Setup()
}

func NewRoutes(engine *gin.Engine) Routes {
	return Routes{
		NewPatientRoutes(engine, controllers.NewPatientController(
			services.NewPatientService(
				repositories.NewPatientRepository(),
			))),
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
