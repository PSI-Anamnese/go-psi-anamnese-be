package repositories

type Patient struct {
	id string
}

var _ Entity = (*Patient)(nil)

func (p Patient) GetId() string {
	return p.id
}

type PatientRepository struct {
}

var _ Repository[string, Patient] = (*PatientRepository)(nil)

func NewPatientRepository() PatientRepository {
	return PatientRepository{}
}

func (p PatientRepository) Save(patient Patient) (Patient, error) {
	return Patient{}, nil
}

func (p PatientRepository) Update(patient Patient) (Patient, error) {
	return Patient{}, nil
}

func (p PatientRepository) Delete(id string) error {
	return nil
}

func (p PatientRepository) ListAll() ([]Patient, error) {
	return []Patient{}, nil
}

func (p PatientRepository) ListById(id string) (Patient, error) {
	return Patient{}, nil
}
