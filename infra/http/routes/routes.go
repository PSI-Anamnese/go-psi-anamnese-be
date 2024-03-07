package routes

import (
	"github.com/gin-gonic/gin"
)

type Routes []Route

type Route interface {
	Setup()
}

func NewRoutes(engine *gin.Engine) Routes {
	return Routes{
		NewPatientRoutes(engine),
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
