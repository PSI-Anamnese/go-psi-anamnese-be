package entities

import "github.com/psi-anamnese/psi-anamnese-be/infra/database/repositories"

var _ repositories.Entity = (*Patient)(nil)

type Patient struct {
	id string
}

func (p Patient) GetId() string {
	return p.id
}
