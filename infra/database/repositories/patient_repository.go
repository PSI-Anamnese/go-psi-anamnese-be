package repositories

type Patient struct {
	id string
}

func (p Patient) GetId() string {
	return p.id
}

type PatientRepository struct {
}

func NewPatientRepository() PatientRepository {
	return PatientRepository{}
}

func (p PatientRepository) Save(patient Patient) (Patient, error) {
	return Patient{}, nil
}

func (p PatientRepository) Update(patient Patient) (Patient, error) {
	return Patient{}, nil
}

func (p PatientRepository) Delete(patient Patient) error {
	return nil
}

func (p PatientRepository) ListAll() ([]Patient, error) {
	return []Patient{}, nil
}

func (p PatientRepository) ListById(id string) (Patient, error) {
	return Patient{}, nil
}
