package controllers

import "github.com/gin-gonic/gin"

type PatientController struct {
}

func NewPatientController() PatientController {
	return PatientController{}
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
