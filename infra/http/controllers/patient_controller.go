package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/psi-anamnese/psi-anamnese-be/application/services"
)

type PatientController struct {
	patientService services.PatientService
}

func NewPatientController(service services.PatientService) PatientController {
	return PatientController{
		patientService: service,
	}
}

func (p PatientController) GetOne(c *gin.Context) {
	c.JSON(200, gin.H{
		"id": "teste",
	})
}

func (p PatientController) GetAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"id": "teste",
	})
}

func (p PatientController) Save(c *gin.Context) {
	c.JSON(200, gin.H{
		"id": "teste",
	})
}

func (p PatientController) Update(c *gin.Context) {
	c.JSON(200, gin.H{
		"id": "teste",
	})
}

func (p PatientController) Delete(c *gin.Context) {
	c.JSON(200, gin.H{
		"id": "teste",
	})
}
