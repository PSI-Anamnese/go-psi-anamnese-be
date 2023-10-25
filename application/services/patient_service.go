package services

import (
	"github.com/psi-anamnese/psi-anamnese-be/infra/database/repositories"
	"github.com/psi-anamnese/psi-anamnese-be/infra/database/repositories/entities"
)

type PatientService struct {
	repository repositories.Repository[string, entities.Patient]
}

func NewPatientService(repository repositories.Repository[string, entities.Patient]) PatientService {
	return PatientService{
		repository: repository,
	}
}
