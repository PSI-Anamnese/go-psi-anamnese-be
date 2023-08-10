package services

import "github.com/psi-anamnese/psi-anamnese-be/infra/database/repositories"

type PatientService struct {
	repository repositories.PatientRepository
}

func NewPatientService(repository repositories.PatientRepository) PatientService {
	return PatientService{
		repository: repository,
	}
}
