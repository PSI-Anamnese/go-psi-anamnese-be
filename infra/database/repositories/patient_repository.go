package repositories

import "github.com/psi-anamnese/psi-anamnese-be/infra/database/repositories/entities"

var _ Repository[string, entities.Patient] = (*PatientRepository)(nil)

type PatientRepository struct {
}

func NewPatientRepository() PatientRepository {
	return PatientRepository{}
}

func (p PatientRepository) Save(patient entities.Patient) (entities.Patient, error) {
	return entities.Patient{}, nil
}

func (p PatientRepository) Update(patient entities.Patient) (entities.Patient, error) {
	return entities.Patient{}, nil
}

func (p PatientRepository) Delete(id string) error {
	return nil
}

func (p PatientRepository) ListAll() ([]entities.Patient, error) {
	return []entities.Patient{}, nil
}

func (p PatientRepository) ListById(id string) (entities.Patient, error) {
	return entities.Patient{}, nil
}
